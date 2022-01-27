CREATE TABLE hashtags (
	id SERIAL PRIMARY KEY,
	created_at BIGINT NOT NULL,
	title VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE hashtag_post (
	id SERIAL PRIMARY KEY,
	hashtag_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	UNIQUE(hashtag_id, post_id)
);