// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package pg

import (
	"context"
)

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
