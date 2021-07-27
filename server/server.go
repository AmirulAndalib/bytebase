package server

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "embed"

	"github.com/bytebase/bytebase/api"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	scas "github.com/qiangmzsx/string-adapter/v2"
	"go.uber.org/zap"
)

type Server struct {
	TaskScheduler *TaskScheduler
	SchemaSyncer  *SchemaSyncer

	CacheService api.CacheService

	SettingService         api.SettingService
	PrincipalService       api.PrincipalService
	MemberService          api.MemberService
	ProjectService         api.ProjectService
	ProjectMemberService   api.ProjectMemberService
	EnvironmentService     api.EnvironmentService
	InstanceService        api.InstanceService
	DatabaseService        api.DatabaseService
	TableService           api.TableService
	ColumnService          api.ColumnService
	IndexService           api.IndexService
	DataSourceService      api.DataSourceService
	BackupService          api.BackupService
	IssueService           api.IssueService
	IssueSubscriberService api.IssueSubscriberService
	PipelineService        api.PipelineService
	StageService           api.StageService
	TaskService            api.TaskService
	ActivityService        api.ActivityService
	BookmarkService        api.BookmarkService
	VCSService             api.VCSService
	RepositoryService      api.RepositoryService

	e *echo.Echo

	l         *zap.Logger
	version   string
	mode      string
	host      string
	port      int
	startedTs int64
	secret    string
	readonly  bool
	demo      bool
	plan      api.PlanType
}

//go:embed acl_casbin_model.conf
var casbinModel string

//go:embed acl_casbin_policy_owner.csv
var casbinOwnerPolicy string

//go:embed acl_casbin_policy_dba.csv
var casbinDBAPolicy string

//go:embed acl_casbin_policy_developer.csv
var casbinDeveloperPolicy string

func NewServer(logger *zap.Logger, version string, host string, port int, mode string, secret string, readonly bool, demo bool, debug bool) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	embedFrontend(logger, e)

	s := &Server{
		l:            logger,
		CacheService: NewCacheService(logger),
		e:            e,
		version:      version,
		mode:         mode,
		host:         host,
		port:         port,
		startedTs:    time.Now().Unix(),
		secret:       secret,
		readonly:     readonly,
		demo:         demo,
		plan:         api.TEAM,
	}

	if !readonly {
		scheduler := NewTaskScheduler(logger, s)
		defaultExecutor := NewDefaultTaskExecutor(logger)
		createDBExecutor := NewDatabaseCreateTaskExecutor(logger)
		sqlExecutor := NewSchemaUpdateTaskExecutor(logger)
		scheduler.Register(string(api.TaskGeneral), defaultExecutor)
		scheduler.Register(string(api.TaskDatabaseCreate), createDBExecutor)
		scheduler.Register(string(api.TaskDatabaseSchemaUpdate), sqlExecutor)
		s.TaskScheduler = scheduler

		schemaSyncer := NewSchemaSyncer(logger, s)
		s.SchemaSyncer = schemaSyncer
	}

	// Middleware
	if mode == "dev" || debug {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Skipper: func(c echo.Context) bool {
				return !strings.HasPrefix(c.Path(), "/api") && !strings.HasPrefix(c.Path(), "/hook")
			},
			Format: `{"time":"${time_rfc3339}",` +
				`"method":"${method}","uri":"${uri}",` +
				`"status":${status},"error":"${error}"}` + "\n",
		}))
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return RecoverMiddleware(logger, next)
	})

	webhookGroup := e.Group("/hook")
	s.registerWebhookRoutes(webhookGroup)

	apiGroup := e.Group("/api")

	apiGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return JWTMiddleware(logger, s.PrincipalService, next, mode, secret)
	})

	m, err := model.NewModelFromString(casbinModel)
	if err != nil {
		e.Logger.Fatal(err)
	}
	sa := scas.NewAdapter(strings.Join([]string{casbinOwnerPolicy, casbinDBAPolicy, casbinDeveloperPolicy}, "\n"))
	ce, err := casbin.NewEnforcer(m, sa)
	if err != nil {
		e.Logger.Fatal(err)
	}
	apiGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return ACLMiddleware(logger, s, ce, next, readonly)
	})
	s.registerSettingRoutes(apiGroup)
	s.registerActuatorRoutes(apiGroup)
	s.registerAuthRoutes(apiGroup)
	s.registerPrincipalRoutes(apiGroup)
	s.registerMemberRoutes(apiGroup)
	s.registerProjectRoutes(apiGroup)
	s.registerProjectMemberRoutes(apiGroup)
	s.registerEnvironmentRoutes(apiGroup)
	s.registerInstanceRoutes(apiGroup)
	s.registerDatabaseRoutes(apiGroup)
	s.registerIssueRoutes(apiGroup)
	s.registerIssueSubscriberRoutes(apiGroup)
	s.registerTaskRoutes(apiGroup)
	s.registerActivityRoutes(apiGroup)
	s.registerBookmarkRoutes(apiGroup)
	s.registerSqlRoutes(apiGroup)
	s.registerVCSRoutes(apiGroup)
	s.registerPlanRoutes(apiGroup)

	allRoutes, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Fatal(err)
	}

	logger.Debug(fmt.Sprintf("All registered routes: %v", string(allRoutes)))

	return s
}

func (server *Server) Run() error {
	if !server.readonly {
		if err := server.TaskScheduler.Run(); err != nil {
			return err
		}

		if err := server.SchemaSyncer.Run(); err != nil {
			return err
		}
	}

	// Sleep for 1 sec to make sure port is released between runs.
	time.Sleep(time.Duration(1) * time.Second)

	return server.e.Start(fmt.Sprintf(":%d", server.port))
}

func (server *Server) Shutdown(ctx context.Context) {
	if err := server.e.Shutdown(ctx); err != nil {
		server.e.Logger.Fatal(err)
	}
}
