BEGIN;

CREATE TABLE IF NOT EXISTS upload
(
    id                     uuid primary key default gen_random_uuid(),
    bucket_name            text,
    path                   text,
    user_id                uuid references users(id),
    created_at             timestamp with time zone default now(),
    updated_at             timestamp with time zone,
    archived_at            timestamp with time zone
);
COMMIT ;
