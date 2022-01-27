package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/grpc-gateway/graph/generated"
	"cesargdd/grpc-gateway/pb"
	"context"
	"fmt"
)

func (r *postResolver) Comments(ctx context.Context, obj *pb.Post) ([]*pb.Comment, error) {
	res, err := c.ListCommentsById(ctx, &pb.ListCommentsByIdRequest{Id: obj.UserId})
	if err != nil {
		fmt.Println(err)
	}
	return res.Comment, nil
}

func (r *postResolver) Likes(ctx context.Context, obj *pb.Post) ([]*pb.PostLike, error) {
	res, err := pl.ListPostLikesByPostId(ctx, &pb.ListPostLikesByPostIdRequest{PostId: obj.Id})
	if err != nil {
		fmt.Println(err)
	}
	return res.PostLike, nil
}

func (r *postResolver) User(ctx context.Context, obj *pb.Post) (*pb.User, error) {
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

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
