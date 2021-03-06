// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package pg

import (
	"context"
)

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

const listPosts = `-- name: ListPosts :many
SELECT id, created_at, updated_at, url, caption, user_id FROM posts
ORDER BY id DESC
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
