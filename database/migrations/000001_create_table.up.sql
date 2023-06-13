BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users
(
    id                   uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name           text,
    last_name            text,
    dob                  text,
    phone                text,
    email                text,
    created_at           timestamp with time zone default now(),
    updated_at           timestamp with time zone,
    archived_at          timestamp with time zone
);

COMMIT ;

