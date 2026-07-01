CREATE SCHEMA IF NOT EXISTS teams;

CREATE TABLE IF NOT EXISTS teams.team (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    parent_id   INTEGER REFERENCES teams.team(id) ON DELETE SET NULL,
    created_at  TIMESTAMPTZ DEFAULT NOW(),
    updated_at  TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_team_parent ON teams.team(parent_id);

CREATE TABLE IF NOT EXISTS teams.user (
    id          SERIAL PRIMARY KEY,
    first_name  VARCHAR(100) NOT NULL,
    last_name   VARCHAR(100) NOT NULL,
    birth_year  SMALLINT NOT NULL 
        CHECK (birth_year > 1900 AND birth_year <= EXTRACT(YEAR FROM CURRENT_DATE)),
    team_id     INTEGER NULL REFERENCES teams.team(id) ON DELETE SET NULL,
    created_at  TIMESTAMPTZ DEFAULT NOW(),
    updated_at  TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_user_team_id ON teams.user(team_id);
CREATE INDEX IF NOT EXISTS idx_user_last_name ON teams.user(last_name);