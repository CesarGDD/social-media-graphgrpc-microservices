package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/grpc-gateway/graph/generated"
	"cesargdd/grpc-gateway/pb"
	"context"
	"fmt"
)

func (r *userResolver) Posts(ctx context.Context, obj *pb.User) ([]*pb.Post, error) {
	res, err := p.ListPostsById(ctx, &pb.ListPostsByIdRequest{Id: obj.Id})
	if err != nil {
		fmt.Println(err)
	}
	return res.Post, nil
}

func (r *userResolver) Followers(ctx context.Context, obj *pb.User) ([]*pb.Follower, error) {
	res, err := f.ListFollowersByLeaderId(ctx, &pb.ListFollowersByLeaderIdRequest{LeaderId: obj.Id})
	if err != nil {
		fmt.Println(err)
	}
	return res.Follower, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
