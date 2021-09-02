package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/bytebase/bytebase/api"
	"github.com/bytebase/bytebase/common"
	"go.uber.org/zap"
)

var (
	_ api.RepositoryService = (*RepositoryService)(nil)
)

// RepositoryService represents a service for managing repository.
type RepositoryService struct {
	l  *zap.Logger
	db *DB

	projectService api.ProjectService
}

// NewRepositoryService returns a new instance of RepositoryService.
func NewRepositoryService(logger *zap.Logger, db *DB, projectService api.ProjectService) *RepositoryService {
	return &RepositoryService{l: logger, db: db, projectService: projectService}
}

// CreateRepository creates a new repository.
func (s *RepositoryService) CreateRepository(ctx context.Context, create *api.RepositoryCreate) (*api.Repository, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	repository, err := s.createRepository(ctx, tx, create)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}

	return repository, nil
}

// FindRepositoryList retrieves a list of repositorys based on find.
func (s *RepositoryService) FindRepositoryList(ctx context.Context, find *api.RepositoryFind) ([]*api.Repository, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	list, err := findRepositoryList(ctx, tx, find)
	if err != nil {
		return []*api.Repository{}, err
	}

	return list, nil
}

// FindRepository retrieves a single repository based on find.
// Returns ENOTFOUND if no matching record.
// Returns ECONFLICT if finding more than 1 matching records.
func (s *RepositoryService) FindRepository(ctx context.Context, find *api.RepositoryFind) (*api.Repository, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	list, err := findRepositoryList(ctx, tx, find)
	if err != nil {
		return nil, err
	} else if len(list) == 0 {
		return nil, &common.Error{Code: common.ENOTFOUND, Message: fmt.Sprintf("repository not found: %+v", find)}
	} else if len(list) > 1 {
		return nil, &common.Error{Code: common.ECONFLICT, Message: fmt.Sprintf("found %d repositories with filter %+v, expect 1", len(list), find)}
	}
	return list[0], nil
}

// PatchRepository updates an existing repository by ID.
// Returns ENOTFOUND if repository does not exist.
func (s *RepositoryService) PatchRepository(ctx context.Context, patch *api.RepositoryPatch) (*api.Repository, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	repository, err := patchRepository(ctx, tx, patch)
	if err != nil {
		return nil, FormatError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}

	return repository, nil
}

// DeleteRepository deletes an existing repository by ID.
// Returns ENOTFOUND if repository does not exist.
func (s *RepositoryService) DeleteRepository(ctx context.Context, delete *api.RepositoryDelete) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return FormatError(err)
	}
	defer tx.Rollback()

	err = s.deleteRepository(ctx, tx, delete)
	if err != nil {
		return FormatError(err)
	}

	if err := tx.Commit(); err != nil {
		return FormatError(err)
	}

	return nil
}

// createRepository creates a new repository.
func (s *RepositoryService) createRepository(ctx context.Context, tx *Tx, create *api.RepositoryCreate) (*api.Repository, error) {
	// Updates the project workflow_type to "VCS"
	workflowType := api.VCS_WORKFLOW
	projectPatch := api.ProjectPatch{
		ID:           create.ProjectId,
		UpdaterId:    create.CreatorId,
		WorkflowType: &workflowType,
	}
	if _, err := s.projectService.PatchProjectTx(ctx, tx.Tx, &projectPatch); err != nil {
		return nil, err
	}

	// Insert row into database.
	row, err := tx.QueryContext(ctx, `
		INSERT INTO repository (
		    creator_id,
		    updater_id,
			vcs_id,
			project_id,
			name,
			full_path,
			web_url,
			base_directory,
			branch_filter,
			external_id,
			external_webhook_id,
			webhook_url_host,
			webhook_endpoint_id,
			webhook_secret_token,
			access_token,
			expires_ts,
			refresh_token
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, creator_id, created_ts, updater_id, updated_ts, vcs_id, project_id, name, full_path, web_url, base_directory, branch_filter, external_id, external_webhook_id, webhook_url_host, webhook_endpoint_id, webhook_secret_token, access_token, expires_ts, refresh_token
	`,
		create.CreatorId,
		create.CreatorId,
		create.VCSId,
		create.ProjectId,
		create.Name,
		create.FullPath,
		create.WebURL,
		create.BaseDirectory,
		create.BranchFilter,
		create.ExternalId,
		create.ExternalWebhookId,
		create.WebhookURLHost,
		create.WebhookEndpointId,
		create.WebhookSecretToken,
		create.AccessToken,
		create.ExpiresTs,
		create.RefreshToken,
	)

	if err != nil {
		return nil, FormatError(err)
	}
	defer row.Close()

	row.Next()
	var repository api.Repository
	if err := row.Scan(
		&repository.ID,
		&repository.CreatorId,
		&repository.CreatedTs,
		&repository.UpdaterId,
		&repository.UpdatedTs,
		&repository.VCSId,
		&repository.ProjectId,
		&repository.Name,
		&repository.FullPath,
		&repository.WebURL,
		&repository.BaseDirectory,
		&repository.BranchFilter,
		&repository.ExternalId,
		&repository.ExternalWebhookId,
		&repository.WebhookURLHost,
		&repository.WebhookEndpointId,
		&repository.WebhookSecretToken,
		&repository.AccessToken,
		&repository.ExpiresTs,
		&repository.RefreshToken,
	); err != nil {
		return nil, FormatError(err)
	}

	return &repository, nil
}

func findRepositoryList(ctx context.Context, tx *Tx, find *api.RepositoryFind) (_ []*api.Repository, err error) {
	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := find.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	if v := find.VCSId; v != nil {
		where, args = append(where, "vcs_id = ?"), append(args, *v)
	}
	if v := find.ProjectId; v != nil {
		where, args = append(where, "project_id = ?"), append(args, *v)
	}
	if v := find.WebhookEndpointId; v != nil {
		where, args = append(where, "webhook_endpoint_id = ?"), append(args, *v)
	}

	rows, err := tx.QueryContext(ctx, `
		SELECT 
		    id,
		    creator_id,
		    created_ts,
		    updater_id,
		    updated_ts,
			vcs_id,
			project_id,
			name,
			full_path,
			web_url,
			base_directory,
			branch_filter,
			external_id,
			external_webhook_id,
			webhook_url_host,
			webhook_endpoint_id,
			webhook_secret_token,
			access_token,
			expires_ts,
			refresh_token
		FROM repository
		WHERE `+strings.Join(where, " AND "),
		args...,
	)
	if err != nil {
		return nil, FormatError(err)
	}
	defer rows.Close()

	// Iterate over result set and deserialize rows into list.
	list := make([]*api.Repository, 0)
	for rows.Next() {
		var repository api.Repository
		if err := rows.Scan(
			&repository.ID,
			&repository.CreatorId,
			&repository.CreatedTs,
			&repository.UpdaterId,
			&repository.UpdatedTs,
			&repository.VCSId,
			&repository.ProjectId,
			&repository.Name,
			&repository.FullPath,
			&repository.WebURL,
			&repository.BaseDirectory,
			&repository.BranchFilter,
			&repository.ExternalId,
			&repository.ExternalWebhookId,
			&repository.WebhookURLHost,
			&repository.WebhookEndpointId,
			&repository.WebhookSecretToken,
			&repository.AccessToken,
			&repository.ExpiresTs,
			&repository.RefreshToken,
		); err != nil {
			return nil, FormatError(err)
		}

		list = append(list, &repository)
	}
	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}

	return list, nil
}

// patchRepository updates a repository by ID. Returns the new state of the repository after update.
func patchRepository(ctx context.Context, tx *Tx, patch *api.RepositoryPatch) (*api.Repository, error) {
	// Build UPDATE clause.
	set, args := []string{"updater_id = ?"}, []interface{}{patch.UpdaterId}
	if v := patch.BaseDirectory; v != nil {
		set, args = append(set, "base_directory = ?"), append(args, *v)
	}
	if v := patch.BranchFilter; v != nil {
		set, args = append(set, "branch_filter = ?"), append(args, *v)
	}

	args = append(args, patch.ID)

	// Execute update query with RETURNING.
	row, err := tx.QueryContext(ctx, `
		UPDATE repository
		SET `+strings.Join(set, ", ")+`
		WHERE id = ?
		RETURNING id, creator_id, created_ts, updater_id, updated_ts, vcs_id, project_id, name, full_path, web_url, base_directory, branch_filter, external_id, external_webhook_id, webhook_url_host, webhook_endpoint_id, webhook_secret_token, access_token, expires_ts, refresh_token
	`,
		args...,
	)
	if err != nil {
		return nil, FormatError(err)
	}
	defer row.Close()

	if row.Next() {
		var repository api.Repository
		if err := row.Scan(
			&repository.ID,
			&repository.CreatorId,
			&repository.CreatedTs,
			&repository.UpdaterId,
			&repository.UpdatedTs,
			&repository.VCSId,
			&repository.ProjectId,
			&repository.Name,
			&repository.FullPath,
			&repository.WebURL,
			&repository.BaseDirectory,
			&repository.BranchFilter,
			&repository.ExternalId,
			&repository.ExternalWebhookId,
			&repository.WebhookURLHost,
			&repository.WebhookEndpointId,
			&repository.WebhookSecretToken,
			&repository.AccessToken,
			&repository.ExpiresTs,
			&repository.RefreshToken,
		); err != nil {
			return nil, FormatError(err)
		}

		return &repository, nil
	}

	return nil, &common.Error{Code: common.ENOTFOUND, Message: fmt.Sprintf("repository ID not found: %d", patch.ID)}
}

// deleteRepository permanently deletes a repository by ID.
func (s *RepositoryService) deleteRepository(ctx context.Context, tx *Tx, delete *api.RepositoryDelete) error {
	// Updates the project workflow_type to "UI"
	workflowType := api.UI_WORKFLOW
	projectPatch := api.ProjectPatch{
		ID:           delete.ProjectId,
		UpdaterId:    delete.DeleterId,
		WorkflowType: &workflowType,
	}
	if _, err := s.projectService.PatchProjectTx(ctx, tx.Tx, &projectPatch); err != nil {
		return err
	}

	// Remove row from database.
	result, err := tx.ExecContext(ctx, `DELETE FROM repository WHERE project_id = ?`, delete.ProjectId)
	if err != nil {
		return FormatError(err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return &common.Error{Code: common.ENOTFOUND, Message: fmt.Sprintf("repository not found for project ID: %d", delete.ProjectId)}
	}

	return nil
}
