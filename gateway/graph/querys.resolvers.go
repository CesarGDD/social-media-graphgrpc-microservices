package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/grpc-gateway/graph/generated"
	"cesargdd/grpc-gateway/pb"
	"cesargdd/grpc-gateway/servers"
	"context"
	"fmt"
)

var u, f = servers.UserServer()
var p = servers.PostServer()
var c = servers.CommentServer()
var pl, cl = servers.LikesServer()
var h, hp = servers.HashtagServer()

func (r *queryResolver) User(ctx context.Context, id int) (*pb.User, error) {
	res, err := u.GetUser(ctx, &pb.GetUserRequest{Id: int32(id)})
	if err != nil {
		fmt.Println(err)
	}
	user := res.User
	return &pb.User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.CreatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Email:     user.Email,
		Password:  user.Password,
		Avatar:    user.Avatar,
	}, nil
}

func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*pb.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*pb.User, error) {
	res, err := u.ListUsers(ctx, &pb.ListUsersRequest{})
	if err != nil {
		fmt.Println(err)
	}
	return res.User, nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*pb.Post, error) {
	res, err := p.GetPost(ctx, &pb.GetPostRequest{Id: int32(id)})
	if err != nil {
		fmt.Println(err)
	}
	post := res.Post
	return &pb.Post{
		Id:        post.Id,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Url:       post.Url,
		Caption:   post.Caption,
		UserId:    post.UserId,
	}, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*pb.Post, error) {
	res, err := p.ListPosts(ctx, &pb.ListPostsRequest{})
	if err != nil {
		fmt.Println(err)
	}
	return res.Post, nil
}

func (r *queryResolver) Comment(ctx context.Context, id int) (*pb.Comment, error) {
	res, err := c.GetComment(ctx, &pb.GetCommentRequest{Id: int32(id)})
	if err != nil {
		fmt.Println(err)
	}
	comment := res.Comment
	return &pb.Comment{
		Id:        comment.Id,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Contents:  comment.Contents,
		UserId:    comment.UserId,
		PostId:    comment.PostId,
	}, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*pb.Comment, error) {
	res, err := c.ListComments(ctx, &pb.ListCommentsRequest{})
	if err != nil {
		fmt.Println(err)
	}
	return res.Comment, nil
}

func (r *queryResolver) PostLike(ctx context.Context, id int) (*pb.PostLike, error) {
	res, err := pl.GetPostLike(ctx, &pb.GetPostLikeRequest{Id: int32(id)})
	if err != nil {
		fmt.Println(err)
	}
	postLike := res.PostLike
	return &pb.PostLike{
		Id:        postLike.Id,
		CreatedAt: postLike.CreatedAt,
		UserId:    postLike.UserId,
		PostId:    postLike.PostId,
	}, nil
}

func (r *queryResolver) PostLikes(ctx context.Context) ([]*pb.PostLike, error) {
	res, err := pl.ListPostLikes(ctx, &pb.ListPostLikesRequest{})
	if err != nil {
		fmt.Println(err)
	}
	return res.PostLike, nil
}

func (r *queryResolver) CommentLike(ctx context.Context, id int) (*pb.CommentLike, error) {
	res, err := cl.GetCommentLike(ctx, &pb.GetCommentLikeRequest{Id: int32(id)})
	if err != nil {
		fmt.Println(err)
	}
	commentLike := res.CommentLike
	return &pb.CommentLike{
		Id:        commentLike.Id,
		CreatedAt: commentLike.CreatedAt,
		UserId:    commentLike.UserId,
		CommentId: commentLike.CommentId,
	}, nil
}

func (r *queryResolver) CommentLikes(ctx context.Context) ([]*pb.CommentLike, error) {
	res, err := cl.ListCommentLikes(ctx, &pb.ListCommentLikesRequest{})
	if err != nil {
		fmt.Println(err)
	}
	return res.CommentLike, nil
}

func (r *queryResolver) HashtagByTitle(ctx context.Context, title string) (*pb.Hashtag, error) {
	res, err := h.GetHashtagByTitle(ctx, &pb.GetHashtagByTitleRequest{Title: title})
	if err != nil {
		fmt.Println(err)
	}
	hashtag := res.Hashtag
	return &pb.Hashtag{
		Id:        hashtag.Id,
		CreatedAt: hashtag.CreatedAt,
		Title:     hashtag.Title,
	}, nil
}

func (r *queryResolver) Hashtag(ctx context.Context, id int) (*pb.Hashtag, error) {
	res, err := h.GetHashtag(ctx, &pb.GetHashtagRequest{Id: int32(id)})
	if err != nil {
		fmt.Println(err)
	}
	hashtag := res.Hashtag
	return &pb.Hashtag{
		Id:        hashtag.Id,
		CreatedAt: hashtag.CreatedAt,
		Title:     hashtag.Title,
	}, nil
}

func (r *queryResolver) Hashtags(ctx context.Context) ([]*pb.Hashtag, error) {
	res, err := h.ListHashtags(ctx, &pb.ListHashtagsRequest{})
	if err != nil {
		fmt.Println(err)
	}
	return res.Hashtag, nil
}

func (r *queryResolver) HashtagPost(ctx context.Context, id int) (*pb.HashtagPost, error) {
	res, err := hp.GetHashtagPost(ctx, &pb.GetHashtagPostRequest{Id: int32(id)})
	if err != nil {
		fmt.Println(err)
	}
	hasPost := res.HashtagPost
	return &pb.HashtagPost{
		Id:        hasPost.Id,
		HashtagId: hasPost.HashtagId,
		PostId:    hasPost.PostId,
	}, nil
}

func (r *queryResolver) HashtagsPost(ctx context.Context) ([]*pb.HashtagPost, error) {
	res, err := hp.ListHashtagPosts(ctx, &pb.ListHashtagPostsRequest{})
	if err != nil {
		fmt.Println(err)
	}
	return res.HashtagPost, nil
}

func (r *queryResolver) Follower(ctx context.Context, id int) (*pb.Follower, error) {
	res, err := f.GetFollower(ctx, &pb.GetFollowerRequest{Id: int32(id)})
	if err != nil {
		fmt.Println(err)
	}
	follower := res.Follower
	return &pb.Follower{
		Id:         follower.Id,
		CreatedAt:  follower.CreatedAt,
		LeaderId:   follower.LeaderId,
		FollowerId: follower.FollowerId,
	}, nil
}

func (r *queryResolver) Followers(ctx context.Context) ([]*pb.Follower, error) {
	res, err := f.ListFollowers(ctx, &pb.ListFollowersRequest{})
	if err != nil {
		fmt.Println(err)
	}
	return res.Follower, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
