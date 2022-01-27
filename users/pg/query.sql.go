// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package pg

import (
	"context"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (contents, user_id, post_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, contents, user_id, post_id
`

type CreateCommentParams struct {
	Contents  string
	UserId    int32
	PostId    int32
	CreatedAt int64
	UpdatedAt int64
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment,
		arg.Contents,
		arg.UserId,
		arg.PostId,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Comment
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Contents,
		&i.UserId,
		&i.PostId,
	)
	return i, err
}

const createCommentLike = `-- name: CreateCommentLike :one
INSERT INTO comment_likes (user_id, comment_id, created_at)
VALUES ($1, $2, $3)
RETURNING id, created_at, user_id, comment_id
`

type CreateCommentLikeParams struct {
	UserId    int32
	CommentId int32
	CreatedAt int64
}

func (q *Queries) CreateCommentLike(ctx context.Context, arg CreateCommentLikeParams) (CommentLike, error) {
	row := q.db.QueryRowContext(ctx, createCommentLike, arg.UserId, arg.CommentId, arg.CreatedAt)
	var i CommentLike
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UserId,
		&i.CommentId,
	)
	return i, err
}

const createFollower = `-- name: CreateFollower :one
INSERT INTO followers (leader_id, follower_id, created_at)
VALUES ($1, $2, $3)
RETURNING id, created_at, leader_id, follower_id
`

type CreateFollowerParams struct {
	LeaderId   int32
	FollowerId int32
	CreatedAt  int64
}

func (q *Queries) CreateFollower(ctx context.Context, arg CreateFollowerParams) (Follower, error) {
	row := q.db.QueryRowContext(ctx, createFollower, arg.LeaderId, arg.FollowerId, arg.CreatedAt)
	var i Follower
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.LeaderId,
		&i.FollowerId,
	)
	return i, err
}

const createHashtag = `-- name: CreateHashtag :one
INSERT INTO hashtags (title, created_at)
VALUES ($1, $2)
RETURNING id, created_at, title
`

type CreateHashtagParams struct {
	Title     string
	CreatedAt int64
}

func (q *Queries) CreateHashtag(ctx context.Context, arg CreateHashtagParams) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, createHashtag, arg.Title, arg.CreatedAt)
	var i Hashtag
	err := row.Scan(&i.Id, &i.CreatedAt, &i.Title)
	return i, err
}

const createHashtagPost = `-- name: CreateHashtagPost :one
INSERT INTO hashtag_post (hashtag_id, post_id)
VALUES ($1, $2)
RETURNING id, hashtag_id, post_id
`

type CreateHashtagPostParams struct {
	HashtagId int32
	PostId    int32
}

func (q *Queries) CreateHashtagPost(ctx context.Context, arg CreateHashtagPostParams) (HashtagPost, error) {
	row := q.db.QueryRowContext(ctx, createHashtagPost, arg.HashtagId, arg.PostId)
	var i HashtagPost
	err := row.Scan(&i.Id, &i.HashtagId, &i.PostId)
	return i, err
}

const createPost = `-- name: CreatePost :one
INSERT INTO posts (url, caption, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, url, caption, user_id
`

type CreatePostParams struct {
	Url       string
	Caption   string
	UserId    int32
	CreatedAt int64
	UpdatedAt int64
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.Url,
		arg.Caption,
		arg.UserId,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Post
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Url,
		&i.Caption,
		&i.UserId,
	)
	return i, err
}

const createPostLike = `-- name: CreatePostLike :one
INSERT INTO post_likes (user_id, post_id, created_at)
VALUES ($1, $2, $3)
RETURNING id, created_at, user_id, post_id
`

type CreatePostLikeParams struct {
	UserId    int32
	PostId    int32
	CreatedAt int64
}

func (q *Queries) CreatePostLike(ctx context.Context, arg CreatePostLikeParams) (PostLike, error) {
	row := q.db.QueryRowContext(ctx, createPostLike, arg.UserId, arg.PostId, arg.CreatedAt)
	var i PostLike
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UserId,
		&i.PostId,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, bio, avatar, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at, username, bio, avatar, email, password
`

type CreateUserParams struct {
	Username  string
	Bio       string
	Avatar    string
	Email     string
	Password  string
	CreatedAt int64
	UpdatedAt int64
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Bio,
		arg.Avatar,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Bio,
		&i.Avatar,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1
RETURNING id, created_at, updated_at, contents, user_id, post_id
`

func (q *Queries) DeleteComment(ctx context.Context, id int32) (Comment, error) {
	row := q.db.QueryRowContext(ctx, deleteComment, id)
	var i Comment
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Contents,
		&i.UserId,
		&i.PostId,
	)
	return i, err
}

const deleteCommentLike = `-- name: DeleteCommentLike :one
DELETE FROM comment_likes
WHERE id = $1
RETURNING id, created_at, user_id, comment_id
`

func (q *Queries) DeleteCommentLike(ctx context.Context, id int32) (CommentLike, error) {
	row := q.db.QueryRowContext(ctx, deleteCommentLike, id)
	var i CommentLike
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UserId,
		&i.CommentId,
	)
	return i, err
}

const deleteFollower = `-- name: DeleteFollower :one
DELETE FROM followers
WHERE id = $1
RETURNING id, created_at, leader_id, follower_id
`

func (q *Queries) DeleteFollower(ctx context.Context, id int32) (Follower, error) {
	row := q.db.QueryRowContext(ctx, deleteFollower, id)
	var i Follower
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.LeaderId,
		&i.FollowerId,
	)
	return i, err
}

const deleteHashtag = `-- name: DeleteHashtag :one
DELETE FROM hashtags
WHERE id = $1
RETURNING id, created_at, title
`

func (q *Queries) DeleteHashtag(ctx context.Context, id int32) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, deleteHashtag, id)
	var i Hashtag
	err := row.Scan(&i.Id, &i.CreatedAt, &i.Title)
	return i, err
}

const deleteHashtagPost = `-- name: DeleteHashtagPost :one
DELETE FROM hashtag_post
WHERE id = $1
RETURNING id, hashtag_id, post_id
`

func (q *Queries) DeleteHashtagPost(ctx context.Context, id int32) (HashtagPost, error) {
	row := q.db.QueryRowContext(ctx, deleteHashtagPost, id)
	var i HashtagPost
	err := row.Scan(&i.Id, &i.HashtagId, &i.PostId)
	return i, err
}

const deletePost = `-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1
RETURNING id, created_at, updated_at, url, caption, user_id
`

func (q *Queries) DeletePost(ctx context.Context, id int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, deletePost, id)
	var i Post
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Url,
		&i.Caption,
		&i.UserId,
	)
	return i, err
}

const deletePostLike = `-- name: DeletePostLike :one
DELETE FROM post_likes
WHERE id = $1
RETURNING id, created_at, user_id, post_id
`

func (q *Queries) DeletePostLike(ctx context.Context, id int32) (PostLike, error) {
	row := q.db.QueryRowContext(ctx, deletePostLike, id)
	var i PostLike
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UserId,
		&i.PostId,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id, created_at, updated_at, username, bio, avatar, email, password
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Bio,
		&i.Avatar,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getComment = `-- name: GetComment :one
SELECT id, created_at, updated_at, contents, user_id, post_id FROM comments
WHERE id = $1
`

func (q *Queries) GetComment(ctx context.Context, id int32) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getComment, id)
	var i Comment
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Contents,
		&i.UserId,
		&i.PostId,
	)
	return i, err
}

const getCommentLike = `-- name: GetCommentLike :one
SELECT id, created_at, user_id, comment_id FROM comment_likes
WHERE id = $1
`

func (q *Queries) GetCommentLike(ctx context.Context, id int32) (CommentLike, error) {
	row := q.db.QueryRowContext(ctx, getCommentLike, id)
	var i CommentLike
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UserId,
		&i.CommentId,
	)
	return i, err
}

const getFollower = `-- name: GetFollower :one
SELECT id, created_at, leader_id, follower_id FROM followers
WHERE id = $1
`

func (q *Queries) GetFollower(ctx context.Context, id int32) (Follower, error) {
	row := q.db.QueryRowContext(ctx, getFollower, id)
	var i Follower
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.LeaderId,
		&i.FollowerId,
	)
	return i, err
}

const getHashtagById = `-- name: GetHashtagById :one
SELECT id, created_at, title FROM hashtags
WHERE id = $1
`

func (q *Queries) GetHashtagById(ctx context.Context, id int32) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, getHashtagById, id)
	var i Hashtag
	err := row.Scan(&i.Id, &i.CreatedAt, &i.Title)
	return i, err
}

const getHashtagByTitle = `-- name: GetHashtagByTitle :one
SELECT id, created_at, title FROM hashtags
WHERE title = $1
`

func (q *Queries) GetHashtagByTitle(ctx context.Context, title string) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, getHashtagByTitle, title)
	var i Hashtag
	err := row.Scan(&i.Id, &i.CreatedAt, &i.Title)
	return i, err
}

const getHashtagPostById = `-- name: GetHashtagPostById :one
SELECT id, hashtag_id, post_id FROM hashtag_post
WHERE id = $1
`

func (q *Queries) GetHashtagPostById(ctx context.Context, id int32) (HashtagPost, error) {
	row := q.db.QueryRowContext(ctx, getHashtagPostById, id)
	var i HashtagPost
	err := row.Scan(&i.Id, &i.HashtagId, &i.PostId)
	return i, err
}

const getPost = `-- name: GetPost :one
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
WHERE id = $1
`

func (q *Queries) GetPost(ctx context.Context, id int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Url,
		&i.Caption,
		&i.UserId,
	)
	return i, err
}

const getPostLike = `-- name: GetPostLike :one
SELECT id, created_at, user_id, post_id FROM post_likes
WHERE id = $1
`

func (q *Queries) GetPostLike(ctx context.Context, id int32) (PostLike, error) {
	row := q.db.QueryRowContext(ctx, getPostLike, id)
	var i PostLike
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UserId,
		&i.PostId,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, created_at, updated_at, username, bio, avatar, email, password FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Bio,
		&i.Avatar,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, created_at, updated_at, username, bio, avatar, email, password FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Bio,
		&i.Avatar,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const listCommentLikes = `-- name: ListCommentLikes :many
SELECT id, created_at, user_id, comment_id FROM comment_likes
ORDER BY id
`

func (q *Queries) ListCommentLikes(ctx context.Context) ([]CommentLike, error) {
	rows, err := q.db.QueryContext(ctx, listCommentLikes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CommentLike
	for rows.Next() {
		var i CommentLike
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UserId,
			&i.CommentId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentLikesByCommentId = `-- name: ListCommentLikesByCommentId :many
SELECT id, created_at, user_id, comment_id FROM comment_likes
WHERE comment_id = $1
`

func (q *Queries) ListCommentLikesByCommentId(ctx context.Context, commentId int32) ([]CommentLike, error) {
	rows, err := q.db.QueryContext(ctx, listCommentLikesByCommentId, commentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CommentLike
	for rows.Next() {
		var i CommentLike
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UserId,
			&i.CommentId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentLikesById = `-- name: ListCommentLikesById :many
SELECT id, created_at, user_id, comment_id FROM comment_likes
WHERE id = $1
`

func (q *Queries) ListCommentLikesById(ctx context.Context, id int32) ([]CommentLike, error) {
	rows, err := q.db.QueryContext(ctx, listCommentLikesById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CommentLike
	for rows.Next() {
		var i CommentLike
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UserId,
			&i.CommentId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listComments = `-- name: ListComments :many
SELECT id, created_at, updated_at, contents, user_id, post_id FROM comments
ORDER BY id
`

func (q *Queries) ListComments(ctx context.Context) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listComments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Contents,
			&i.UserId,
			&i.PostId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsById = `-- name: ListCommentsById :many
SELECT id, created_at, updated_at, contents, user_id, post_id FROM comments
WHERE id= $1
`

func (q *Queries) ListCommentsById(ctx context.Context, id int32) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listCommentsById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Contents,
			&i.UserId,
			&i.PostId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsByPostId = `-- name: ListCommentsByPostId :many
SELECT id, created_at, updated_at, contents, user_id, post_id FROM comments
WHERE post_id= $1
`

func (q *Queries) ListCommentsByPostId(ctx context.Context, postId int32) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listCommentsByPostId, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Contents,
			&i.UserId,
			&i.PostId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFollowers = `-- name: ListFollowers :many
SELECT id, created_at, leader_id, follower_id FROM followers
ORDER BY id
`

func (q *Queries) ListFollowers(ctx context.Context) ([]Follower, error) {
	rows, err := q.db.QueryContext(ctx, listFollowers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follower
	for rows.Next() {
		var i Follower
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.LeaderId,
			&i.FollowerId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFollowersById = `-- name: ListFollowersById :many
SELECT id, created_at, leader_id, follower_id FROM followers
WHERE id = $1
`

func (q *Queries) ListFollowersById(ctx context.Context, id int32) ([]Follower, error) {
	rows, err := q.db.QueryContext(ctx, listFollowersById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follower
	for rows.Next() {
		var i Follower
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.LeaderId,
			&i.FollowerId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFollowersByLeaderId = `-- name: ListFollowersByLeaderId :many
SELECT id, created_at, leader_id, follower_id FROM followers
WHERE leader_id = $1
`

func (q *Queries) ListFollowersByLeaderId(ctx context.Context, leaderId int32) ([]Follower, error) {
	rows, err := q.db.QueryContext(ctx, listFollowersByLeaderId, leaderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follower
	for rows.Next() {
		var i Follower
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.LeaderId,
			&i.FollowerId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listHashtags = `-- name: ListHashtags :many
SELECT id, created_at, title FROM hashtags
ORDER BY id
`

func (q *Queries) ListHashtags(ctx context.Context) ([]Hashtag, error) {
	rows, err := q.db.QueryContext(ctx, listHashtags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Hashtag
	for rows.Next() {
		var i Hashtag
		if err := rows.Scan(&i.Id, &i.CreatedAt, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listHashtagsById = `-- name: ListHashtagsById :many
SELECT id, created_at, title FROM hashtags
WHERE id = $1
`

func (q *Queries) ListHashtagsById(ctx context.Context, id int32) ([]Hashtag, error) {
	rows, err := q.db.QueryContext(ctx, listHashtagsById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Hashtag
	for rows.Next() {
		var i Hashtag
		if err := rows.Scan(&i.Id, &i.CreatedAt, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listHashtagsPost = `-- name: ListHashtagsPost :many
SELECT id, hashtag_id, post_id FROM hashtag_post
ORDER BY id
`

func (q *Queries) ListHashtagsPost(ctx context.Context) ([]HashtagPost, error) {
	rows, err := q.db.QueryContext(ctx, listHashtagsPost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []HashtagPost
	for rows.Next() {
		var i HashtagPost
		if err := rows.Scan(&i.Id, &i.HashtagId, &i.PostId); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listHashtagsPostById = `-- name: ListHashtagsPostById :many
SELECT id, hashtag_id, post_id FROM hashtag_post
WHERE id = $1
`

func (q *Queries) ListHashtagsPostById(ctx context.Context, id int32) ([]HashtagPost, error) {
	rows, err := q.db.QueryContext(ctx, listHashtagsPostById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []HashtagPost
	for rows.Next() {
		var i HashtagPost
		if err := rows.Scan(&i.Id, &i.HashtagId, &i.PostId); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPostLikes = `-- name: ListPostLikes :many
SELECT id, created_at, user_id, post_id FROM post_likes
ORDER BY id
`

func (q *Queries) ListPostLikes(ctx context.Context) ([]PostLike, error) {
	rows, err := q.db.QueryContext(ctx, listPostLikes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PostLike
	for rows.Next() {
		var i PostLike
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UserId,
			&i.PostId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPostLikesById = `-- name: ListPostLikesById :many
SELECT id, created_at, user_id, post_id FROM post_likes
WHERE id= $1
`

func (q *Queries) ListPostLikesById(ctx context.Context, id int32) ([]PostLike, error) {
	rows, err := q.db.QueryContext(ctx, listPostLikesById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PostLike
	for rows.Next() {
		var i PostLike
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UserId,
			&i.PostId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPostLikesByPostId = `-- name: ListPostLikesByPostId :many
SELECT id, created_at, user_id, post_id FROM post_likes
WHERE post_id= $1
`

func (q *Queries) ListPostLikesByPostId(ctx context.Context, postId int32) ([]PostLike, error) {
	rows, err := q.db.QueryContext(ctx, listPostLikesByPostId, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PostLike
	for rows.Next() {
		var i PostLike
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UserId,
			&i.PostId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPosts = `-- name: ListPosts :many
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
ORDER BY id
`

func (q *Queries) ListPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Url,
			&i.Caption,
			&i.UserId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPostsById = `-- name: ListPostsById :many
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
WHERE id= $1
`

func (q *Queries) ListPostsById(ctx context.Context, id int32) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPostsById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Url,
			&i.Caption,
			&i.UserId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPostsByUserId = `-- name: ListPostsByUserId :many
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
WHERE user_id= $1
`

func (q *Queries) ListPostsByUserId(ctx context.Context, userId int32) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPostsByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Url,
			&i.Caption,
			&i.UserId,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, created_at, updated_at, username, bio, avatar, email, password FROM users
ORDER BY id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Username,
			&i.Bio,
			&i.Avatar,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersById = `-- name: ListUsersById :many
SELECT id, created_at, updated_at, username, bio, avatar, email, password FROM users
WHERE id = $1
`

func (q *Queries) ListUsersById(ctx context.Context, id int32) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Id,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Username,
			&i.Bio,
			&i.Avatar,
			&i.Email,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComment = `-- name: UpdateComment :one
UPDATE comments
SET contents = $2, updated_at = $3
WHERE id = $1
RETURNING id, created_at, updated_at, contents, user_id, post_id
`

type UpdateCommentParams struct {
	Id        int32
	Contents  string
	UpdatedAt int64
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, updateComment, arg.Id, arg.Contents, arg.UpdatedAt)
	var i Comment
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Contents,
		&i.UserId,
		&i.PostId,
	)
	return i, err
}

const updateHashtag = `-- name: UpdateHashtag :one
UPDATE hashtags
SET title = $2
WHERE id = $1
RETURNING id, created_at, title
`

type UpdateHashtagParams struct {
	Id    int32
	Title string
}

func (q *Queries) UpdateHashtag(ctx context.Context, arg UpdateHashtagParams) (Hashtag, error) {
	row := q.db.QueryRowContext(ctx, updateHashtag, arg.Id, arg.Title)
	var i Hashtag
	err := row.Scan(&i.Id, &i.CreatedAt, &i.Title)
	return i, err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET url = $2, caption = $3, updated_at = $4
WHERE id = $1
RETURNING id, created_at, updated_at, url, caption, user_id
`

type UpdatePostParams struct {
	Id        int32
	Url       string
	Caption   string
	UpdatedAt int64
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.Id,
		arg.Url,
		arg.Caption,
		arg.UpdatedAt,
	)
	var i Post
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Url,
		&i.Caption,
		&i.UserId,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET bio = $2, avatar = $3, updated_at = $4
WHERE id = $1
RETURNING id, created_at, updated_at, username, bio, avatar, email, password
`

type UpdateUserParams struct {
	Id        int32
	Bio       string
	Avatar    string
	UpdatedAt int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Id,
		arg.Bio,
		arg.Avatar,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.Id,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.Bio,
		&i.Avatar,
		&i.Email,
		&i.Password,
	)
	return i, err
}