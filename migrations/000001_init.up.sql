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

-- Команды
INSERT INTO teams.team (name, parent_id) VALUES
    ('Executive Board', NULL),
    ('Engineering', 1),
    ('Marketing', 1),
    ('Sales', 1),
    ('HR', 1),
    ('Backend Team', 2),
    ('Frontend Team', 2),
    ('DevOps', 2),
    ('Digital Marketing', 3),
    ('Content Marketing', 3),
    ('Enterprise Sales', 4),
    ('SMB Sales', 4),
    ('Recruiting', 5),
    ('Employee Relations', 5);

-- Пользователи
INSERT INTO teams.user (first_name, last_name, birth_year, team_id) VALUES
    ('John', 'Smith', 1975, 1),
    ('Sarah', 'Johnson', 1978, 1),
    ('Michael', 'Brown', 1985, 2),
    ('Emily', 'Davis', 1988, 2),
    ('James', 'Wilson', 1983, 2),
    ('Jessica', 'Martinez', 1990, 3),
    ('David', 'Anderson', 1986, 3),
    ('Robert', 'Taylor', 1982, 4),
    ('Maria', 'Thomas', 1989, 4),
    ('William', 'Moore', 1991, 4),
    ('Jennifer', 'Jackson', 1984, 5),
    ('Charles', 'White', 1980, 5),
    ('Daniel', 'Harris', 1992, 6),
    ('Lisa', 'Martin', 1993, 6),
    ('Paul', 'Thompson', 1991, 7),
    ('Karen', 'Garcia', 1994, 7),
    ('Steven', 'Martinez', 1987, 8),
    ('Michelle', 'Robinson', 1990, 8),
    ('Matthew', 'Clark', 1992, 9),
    ('Amanda', 'Rodriguez', 1993, 9),
    ('Kevin', 'Lewis', 1985, 11),
    ('Laura', 'Lee', 1988, 11),
    ('Brian', 'Walker', 1990, 12),
    ('Nancy', 'Hall', 1992, 12);