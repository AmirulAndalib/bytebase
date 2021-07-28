-- Activity for issue 13001
INSERT INTO
    activity (
        id,
        creator_id,
        updater_id,
        container_id,
        `type`
    )
VALUES
    (
        14001,
        1,
        1,
        13001,
        'bb.issue.create'
    );

INSERT INTO
    activity (
        id,
        creator_id,
        updater_id,
        container_id,
        `type`,
        `comment`
    )
VALUES
    (
        14002,
        101,
        101,
        13001,
        'bb.issue.comment.create',
        'Welcome!'
    );

INSERT INTO
    activity (
        id,
        creator_id,
        updater_id,
        container_id,
        `type`,
        `comment`
    )
VALUES
    (
        14003,
        102,
        102,
        13001,
        'bb.issue.comment.create',
        'Let''s rock!'
    );

INSERT INTO
    activity (
        id,
        creator_id,
        updater_id,
        container_id,
        `type`,
        `comment`
    )
VALUES
    (
        14004,
        103,
        103,
        13001,
        'bb.issue.comment.create',
        'Go fish!'
    );

-- Activity for issue 13002
INSERT INTO
    activity (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`
    )
VALUES
    (
        14005,
        1,
        1624873710,
        1,
        1624873710,
        13002,
        'bb.issue.create'
    );

INSERT INTO
    activity (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        14006,
        1,
        1624873710,
        1,
        1624873710,
        13002,
        'bb.pipeline.task.status.update',
        '',
        '{"taskId":11002,"oldStatus":"PENDING","newStatus":"RUNNING"}'
    );

INSERT INTO
    activity (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        14007,
        1,
        1624873710,
        1,
        1624873710,
        13002,
        'bb.pipeline.task.status.update',
        'Established baseline version %s for database ''shop''',
        '{"taskId":11002,"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

INSERT INTO
    activity (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        14008,
        1,
        1624873710,
        1,
        1624873710,
        13002,
        'bb.pipeline.task.status.update',
        '',
        '{"taskId":11003,"oldStatus":"PENDING","newStatus":"RUNNING"}'
    );

INSERT INTO
    activity (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        14009,
        1,
        1624873710,
        1,
        1624873710,
        13002,
        'bb.pipeline.task.status.update',
        'Established baseline version %s for database ''shop''',
        '{"taskId":11003,"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

-- Activity for issue 13003
INSERT INTO
    activity (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`
    )
VALUES
    (
        14010,
        103,
        1624873710,
        103,
        1624873710,
        13003,
        'bb.issue.create'
    );

-- Activity for failed task_run 12001
INSERT INTO
    activity (
        id,
        creator_id,
        updater_id,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        14011,
        1,
        1,
        13003,
        'bb.pipeline.task.status.update',
        'table "tbl1" already exists',
        '{"taskId":11006,"oldStatus":"RUNNING","newStatus":"FAILED"}'
    );

INSERT INTO
    activity (
        id,
        creator_id,
        updater_id,
        container_id,
        `type`,
        payload
    )
VALUES
    (
        14012,
        102,
        102,
        13003,
        'bb.issue.status.update',
        '{"oldStatus":"OPEN","newStatus":"CANCELED"}'
    );

-- Activity for issue 13004
INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624865387,
        1,
        1624865387,
        13004,
        'bb.issue.create',
        '',
        ''
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624865388,
        1,
        1624865388,
        13004,
        'bb.pipeline.task.status.update',
        '',
        '{"taskId":11008,"oldStatus":"PENDING","newStatus":"RUNNING"}'
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624865388,
        1,
        1624865388,
        13004,
        'bb.pipeline.task.status.update',
        'Established baseline version 202106280000 for database ''blog''',
        '{"taskId":11008,"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624865388,
        1,
        1624865388,
        13004,
        'bb.issue.status.update',
        '',
        '{"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

-- Activity for issue 13005
INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624866790,
        1,
        1624866790,
        13005,
        'bb.issue.create',
        '',
        ''
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624866791,
        1,
        1624866791,
        13005,
        'bb.pipeline.task.status.update',
        '',
        '{"taskId":11009,"oldStatus":"PENDING","newStatus":"RUNNING"}'
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624866791,
        1,
        1624866791,
        13005,
        'bb.pipeline.task.status.update',
        'Established baseline version 202106280000 for database ''blog''',
        '{"taskId":11009,"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624866791,
        1,
        1624866791,
        13005,
        'bb.issue.status.update',
        '',
        '{"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

-- Activity for issue 13006
INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624868407,
        1,
        1624868407,
        13006,
        'bb.issue.create',
        '',
        ''
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624868408,
        1,
        1624868408,
        13006,
        'bb.pipeline.task.status.update',
        '',
        '{"taskId":11010,"oldStatus":"PENDING","newStatus":"RUNNING"}'
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624868408,
        1,
        1624868408,
        13006,
        'bb.pipeline.task.status.update',
        'Established baseline version 202106280000 for database ''blog''',
        '{"taskId":11010,"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624868408,
        1,
        1624868408,
        13006,
        'bb.issue.status.update',
        '',
        '{"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

-- Activity for issue 13007
INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624868680,
        1,
        1624868680,
        13007,
        'bb.issue.create',
        '',
        ''
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624868681,
        1,
        1624868681,
        13007,
        'bb.pipeline.task.status.update',
        '',
        '{"taskId":11011,"oldStatus":"PENDING","newStatus":"RUNNING"}'
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624868681,
        1,
        1624868681,
        13007,
        'bb.pipeline.task.status.update',
        'Established baseline version 202106280000 for database ''blog''',
        '{"taskId":11011,"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624868681,
        1,
        1624868681,
        13007,
        'bb.issue.status.update',
        '',
        '{"oldStatus":"RUNNING","newStatus":"DONE"}'
    );

-- Activity for issue 13008
INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624869944,
        1,
        1624869944,
        13008,
        'bb.issue.create',
        '',
        ''
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624869945,
        1,
        1624869945,
        13008,
        'bb.pipeline.task.status.update',
        '',
        '{"taskId":11012,"oldStatus":"PENDING","newStatus":"RUNNING"}'
    );

INSERT INTO
    activity (
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        container_id,
        `type`,
        COMMENT,
        payload
    )
VALUES
    (
        1,
        1624869945,
        1,
        1624869945,
        13008,
        'bb.pipeline.task.status.update',
        'database ''blog'' has already applied version 202106280100',
        '{"taskId":11012,"oldStatus":"RUNNING","newStatus":"FAILED"}'
    );