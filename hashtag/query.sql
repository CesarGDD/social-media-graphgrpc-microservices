-- name: GetHashtagByTitle :one
SELECT * FROM hashtags
WHERE title = $1;

-- name: GetHashtagById :one
SELECT * FROM hashtags
WHERE id = $1;

-- name: ListHashtags :many
SELECT * FROM hashtags
ORDER BY id;

-- name: ListHashtagsById :many
SELECT * FROM hashtags
WHERE id = $1;

-- name: CreateHashtag :one
INSERT INTO hashtags (title, created_at)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateHashtag :one
UPDATE hashtags
SET title = $2
WHERE id = $1
RETURNING *;

-- name: DeleteHashtag :one
DELETE FROM hashtags
WHERE id = $1
RETURNING *;

-- name: GetHashtagPostById :one
SELECT * FROM hashtag_post
WHERE id = $1;

-- name: ListHashtagsPost :many
SELECT * FROM hashtag_post
ORDER BY id;

-- name: ListHashtagsPostById :many
SELECT * FROM hashtag_post
WHERE id = $1;

-- name: CreateHashtagPost :one
INSERT INTO hashtag_post (hashtag_id, post_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteHashtagPost :one
DELETE FROM hashtag_post
WHERE id = $1
RETURNING *;