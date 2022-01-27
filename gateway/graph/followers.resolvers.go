package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/grpc-gateway/graph/generated"
	"cesargdd/grpc-gateway/pb"
	"context"
	"fmt"
)

func (r *followerResolver) User(ctx context.Context, obj *pb.Follower) (*pb.User, error) {
	res, err := u.GetUser(ctx, &pb.GetUserRequest{Id: int32(obj.FollowerId)})
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

// Follower returns generated.FollowerResolver implementation.
func (r *Resolver) Follower() generated.FollowerResolver { return &followerResolver{r} }

type followerResolver struct{ *Resolver }
