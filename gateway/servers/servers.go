package servers

import (
	"cesargdd/grpc-gateway/pb"
	"log"

	"google.golang.org/grpc"
)

func UserServer() (pb.UsersServiceClient, pb.FollowersServiceClient) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("users:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	u := pb.NewUsersServiceClient(cc)
	f := pb.NewFollowersServiceClient(cc)
	return u, f
}
func PostServer() pb.PostsServiceClient {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("posts:50052", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	p := pb.NewPostsServiceClient(cc)
	return p
}
func CommentServer() pb.CommentsServiceClient {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("comments:50053", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	c := pb.NewCommentsServiceClient(cc)
	return c
}
func LikesServer() (pb.PostLikesServiceClient, pb.CommentLikesServiceClient) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("likes:50054", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	pl := pb.NewPostLikesServiceClient(cc)
	cl := pb.NewCommentLikesServiceClient(cc)
	return pl, cl
}
func HashtagServer() (pb.HashtagsServiceClient, pb.HashtagPostsServiceClient) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("hashtags:50055", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	h := pb.NewHashtagsServiceClient(cc)
	hp := pb.NewHashtagPostsServiceClient(cc)
	return h, hp
}