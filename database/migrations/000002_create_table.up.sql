
ALTER TABLE IF EXISTS users ADD COLUMN IF NOT EXISTS password text;

CREATE TABLE IF NOT EXISTS links
(
    id                   uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title                text,
    address              text,
    user_id              uuid references users(id),
    created_at           timestamp with time zone default now(),
    updated_at           timestamp with time zone,
    archived_at          timestamp with time zone
);