-- Rename jwt_token to token in session table
ALTER TABLE sessions RENAME COLUMN jwt_token TO token;

-- Rename state to name in states table
ALTER TABLE states RENAME COLUMN state TO name;