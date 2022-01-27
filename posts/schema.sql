CREATE TABLE IF NOT EXISTS posts (
	id SERIAL PRIMARY KEY,
	created_at BIGINT NOT NULL,
	updated_at BIGINT NOT NULL,
	url VARCHAR(200) NOT NULL,
	caption VARCHAR(240) NOT NULL,
	user_id INTEGER NOT NULL
);
