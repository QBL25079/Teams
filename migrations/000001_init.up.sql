CREATE SCHEMA teams;

CREATE TABLE IF NOT EXISTS teams.team (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    parent_id   INTEGER REFERENCES groups(id) ON DELETE CASCADE,
    created_at  TIMESTAMPTZ DEFAULT NOW(),
    updated_at  TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_groups_parent ON groups(parent_id);

CREATE TABLE IF NOT EXISTS teams.user (
    id          SERIAL PRIMARY KEY,
    first_name  VARCHAR(100) NOT NULL,
    last_name   VARCHAR(100) NOT NULL,
    birth_year  SMALLINT NOT NULL 
        CHECK (birth_year > 1900 AND birth_year <= EXTRACT(YEAR FROM CURRENT_DATE)),
    group_id    INTEGER NOT NULL REFERENCES groups(id) ON DELETE RESTRICT,
    created_at  TIMESTAMPTZ DEFAULT NOW(),
    updated_at  TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_people_group_id ON teams.user(group_id);
CREATE INDEX IF NOT EXISTS idx_people_last_name ON teams.user(last_name);