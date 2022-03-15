CREATE TABLE post_likes (
	id SERIAL PRIMARY KEY,
	created_at BIGINT NOT NULL,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL
);

CREATE TABLE comment_likes (
	id SERIAL PRIMARY KEY,
	created_at BIGINT NOT NULL,
	user_id INTEGER NOT NULL,
	comment_id INTEGER NOT NULL
);