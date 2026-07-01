DROP INDEX IF EXISTS teams.idx_user_team_id;
DROP INDEX IF EXISTS teams.idx_user_last_name;
DROP TABLE IF EXISTS teams.user;
DROP INDEX IF EXISTS teams.idx_team_parent;
DROP TABLE IF EXISTS teams.team;
DROP SCHEMA IF EXISTS teams;