package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/grpc-gateway/graph/generated"
	"cesargdd/grpc-gateway/pb"
	"context"
	"fmt"
)

func (r *commentLikeResolver) User(ctx context.Context, obj *pb.CommentLike) (*pb.User, error) {
	res, err := u.GetUser(ctx, &pb.GetUserRequest{Id: obj.UserId})
	if err != nil {
		fmt.Println(err)
	}
	user := res.User
	return &pb.User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Email:     user.Email,
		Password:  user.Password,
		Avatar:    user.Avatar,
	}, nil
}

func (r *postLikeResolver) User(ctx context.Context, obj *pb.PostLike) (*pb.User, error) {
	res, err := u.GetUser(ctx, &pb.GetUserRequest{Id: obj.UserId})
	if err != nil {
		fmt.Println(err)
	}
	user := res.User
	return &pb.User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Email:     user.Email,
		Password:  user.Password,
		Avatar:    user.Avatar,
	}, nil
}

// CommentLike returns generated.CommentLikeResolver implementation.
func (r *Resolver) CommentLike() generated.CommentLikeResolver { return &commentLikeResolver{r} }

// PostLike returns generated.PostLikeResolver implementation.
func (r *Resolver) PostLike() generated.PostLikeResolver { return &postLikeResolver{r} }

type commentLikeResolver struct{ *Resolver }
type postLikeResolver struct{ *Resolver }
