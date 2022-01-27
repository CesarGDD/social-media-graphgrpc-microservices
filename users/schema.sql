CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	created_at BIGINT NOT NULL,
	updated_at BIGINT NOT NULL,
	username VARCHAR(30) UNIQUE NOT NULL,
	bio VARCHAR(400) NOT NULL DEFAULT '\',
	avatar VARCHAR(200) NOT NULL DEFAULT '\',
	email VARCHAR(40) NOT NULL,
	password VARCHAR(50) NOT NULL
);


CREATE TABLE followers (
	id SERIAL PRIMARY KEY,
	created_at BIGINT NOT NULL,
	leader_id INTEGER NOT NULL,
	follower_id INTEGER NOT NULL
);