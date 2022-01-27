package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/grpc-gateway/graph/generated"
	"cesargdd/grpc-gateway/pb"
	"context"
	"fmt"
)

func (r *hashtagPostResolver) Posts(ctx context.Context, obj *pb.HashtagPost) ([]*pb.Post, error) {
	res, err := p.ListPostsById(ctx, &pb.ListPostsByIdRequest{Id: obj.PostId})
	if err != nil {
		fmt.Println(err)
	}
	return res.Post, nil
}

func (r *hashtagPostResolver) Hashtag(ctx context.Context, obj *pb.HashtagPost) (*pb.Hashtag, error) {
	res, err := h.GetHashtag(ctx, &pb.GetHashtagRequest{Id: obj.HashtagId})
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

// HashtagPost returns generated.HashtagPostResolver implementation.
func (r *Resolver) HashtagPost() generated.HashtagPostResolver { return &hashtagPostResolver{r} }

type hashtagPostResolver struct{ *Resolver }
