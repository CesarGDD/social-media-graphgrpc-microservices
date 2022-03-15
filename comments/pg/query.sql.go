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
ORDER BY id DESC
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
