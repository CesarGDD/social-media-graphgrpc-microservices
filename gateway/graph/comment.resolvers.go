package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/grpc-gateway/graph/generated"
	"cesargdd/grpc-gateway/pb"
	"context"
	"fmt"
)

func (r *commentResolver) Likes(ctx context.Context, obj *pb.Comment) ([]*pb.CommentLike, error) {
	res, err := cl.ListCommentLikesByCommentId(ctx, &pb.ListCommentLikesByCommentIdRequest{CommentId: obj.Id})
	if err != nil {
		fmt.Println(err)
	}
	return res.CommentLike, nil
}

func (r *commentResolver) User(ctx context.Context, obj *pb.Comment) (*pb.User, error) {
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

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
