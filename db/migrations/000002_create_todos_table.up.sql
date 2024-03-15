CREATE TABLE IF NOT EXISTS todos (
    id serial PRIMARY KEY,
    activity_group_id int4 NOT NULL,
    title varchar(50) NOT NULL,
    is_active bool NOT NULL DEFAULT true,
    priority varchar(50) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp,
    CONSTRAINT fk_activity_group
    FOREIGN KEY (activity_group_id)
    REFERENCES activities (id)
);