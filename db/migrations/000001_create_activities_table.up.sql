CREATE TABLE IF NOT EXISTS activities (
    id serial PRIMARY KEY,
    email varchar(50) NOT NULL,
    title varchar(50) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp
);