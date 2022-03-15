package servers

import (
	"cesargdd/grpc-gateway/pb"
	"log"
	"os"

	"google.golang.org/grpc"
)

func UserServer() (pb.UsersServiceClient, pb.FollowersServiceClient) {
	opts := grpc.WithInsecure()
	ccu, err := grpc.Dial(os.Getenv("USERS_ADDRESS"), opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	u := pb.NewUsersServiceClient(ccu)
	f := pb.NewFollowersServiceClient(ccu)
	return u, f
}
func PostServer() pb.PostsServiceClient {
	opts := grpc.WithInsecure()
	ccp, err := grpc.Dial(os.Getenv("POSTS_ADDRESS"), opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	p := pb.NewPostsServiceClient(ccp)
	return p
}
func CommentServer() pb.CommentsServiceClient {
	opts := grpc.WithInsecure()
	ccc, err := grpc.Dial(os.Getenv("COMMENTS_ADDRESS"), opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	c := pb.NewCommentsServiceClient(ccc)
	return c
}
func LikeServer() (pb.PostLikesServiceClient, pb.CommentLikesServiceClient) {
	opts := grpc.WithInsecure()
	ccl, err := grpc.Dial(os.Getenv("LIKES_ADDRESS"), opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	pl := pb.NewPostLikesServiceClient(ccl)
	cl := pb.NewCommentLikesServiceClient(ccl)
	return pl, cl
}
func HashtagServer() (pb.HashtagsServiceClient, pb.HashtagPostsServiceClient) {
	opts := grpc.WithInsecure()
	cch, err := grpc.Dial(os.Getenv("HASHTAGS_ADDRESS"), opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	h := pb.NewHashtagsServiceClient(cch)
	hp := pb.NewHashtagPostsServiceClient(cch)
	return h, hp
}
