-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: ListUsersById :many
SELECT * FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, bio, avatar, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET bio = $2, avatar = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;



-- name: GetFollower :one
SELECT * FROM followers
WHERE id = $1;

-- name: ListFollowers :many
SELECT * FROM followers
ORDER BY id;

-- name: ListFollowersById :many
SELECT * FROM followers
WHERE id = $1;

-- name: ListFollowersByLeaderId :many
SELECT * FROM followers
WHERE leader_id = $1;

-- name: CreateFollower :one
INSERT INTO followers (leader_id, follower_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteFollower :one
DELETE FROM followers
WHERE id = $1
RETURNING *;