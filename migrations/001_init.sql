-- Copyright (C) 2026 Gleb Obitotsky ... AGPLv3

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    telegram_id TEXT UNIQUE NOT NULL,
    phone       TEXT DEFAULT '',
    metadata    JSONB NOT NULL DEFAULT '{}'
);

CREATE TABLE branches (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name       TEXT NOT NULL,
    address    TEXT,
    metadata   JSONB NOT NULL DEFAULT '{}'
);

CREATE TABLE slots (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    branch_id  UUID REFERENCES branches(id) ON DELETE CASCADE,
    start_time TIMESTAMPTZ NOT NULL,
    end_time   TIMESTAMPTZ NOT NULL,
    is_booked  BOOLEAN NOT NULL DEFAULT FALSE,
    metadata   JSONB NOT NULL DEFAULT '{}'
);

CREATE TABLE bookings (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slot_id    UUID REFERENCES slots(id) ON DELETE CASCADE,
    user_id    UUID REFERENCES users(id) ON DELETE CASCADE,
    booked_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    status     TEXT NOT NULL DEFAULT 'confirmed',
    metadata   JSONB NOT NULL DEFAULT '{}'
);

CREATE INDEX idx_slots_branch_booked ON slots(branch_id, is_booked);