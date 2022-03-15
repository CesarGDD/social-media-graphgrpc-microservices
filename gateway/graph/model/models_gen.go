// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"cesargdd/grpc-gateway/pb"
)

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *pb.User   `json:"user"`
}

type AuthToken struct {
	AccessToken string `json:"accessToken"`
}

type EditComment struct {
	ID       int    `json:"id"`
	Contents string `json:"contents"`
}

type EditHashtag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type EditPost struct {
	ID      int     `json:"id"`
	URL     *string `json:"url"`
	Caption *string `json:"caption"`
}

type EditUser struct {
	ID     int     `json:"id"`
	Bio    *string `json:"bio"`
	Avatar *string `json:"avatar"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewComment struct {
	Contents string `json:"contents"`
	UserID   int    `json:"user_id"`
	PostID   int    `json:"post_id"`
}

type NewCommentLike struct {
	UserID    int `json:"user_id"`
	CommentID int `json:"commentId"`
}

type NewFollower struct {
	LeaderID   int `json:"leader_id"`
	FollowerID int `json:"follower_id"`
}

type NewHashtag struct {
	Title string `json:"title"`
}

type NewHashtagPost struct {
	HashtagID int `json:"hashtag_id"`
	PostID    int `json:"post_id"`
}

type NewPost struct {
	URL     string  `json:"url"`
	Caption *string `json:"caption"`
	UserID  int     `json:"user_id"`
}

type NewPostLike struct {
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}

type NewUser struct {
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
