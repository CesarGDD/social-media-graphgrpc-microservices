package main

import (
	"cesargdd/grpc-hashtags/pb"
	"cesargdd/grpc-hashtags/pg"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.HashtagsServiceServer
	pb.HashtagPostsServiceServer
}

var conn = pg.Connect()
var db = pg.New(conn)

// Hashtag

func (*server) CreateHashtag(ctx context.Context, req *pb.CreateHashtagRequest) (*pb.CreateHashtagResponse, error) {
	createHashtag, err := db.CreateHashtag(ctx, pg.CreateHashtagParams{
		Title:     req.GetHashtag().GetTitle(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Hashtag", err)
	}
	return &pb.CreateHashtagResponse{
		Hashtag: &pb.Hashtag{
			Id:        createHashtag.Id,
			CreatedAt: createHashtag.CreatedAt,
			Title:     createHashtag.Title,
		},
	}, nil
}
func (*server) GetHashtagByTitle(ctx context.Context, req *pb.GetHashtagByTitleRequest) (*pb.GetHashtagByTitleResponse, error) {
	hashtag, err := db.GetHashtagByTitle(ctx, req.GetTitle())
	if err != nil {
		fmt.Println("Can not get Hashtag", err)
	}
	return &pb.GetHashtagByTitleResponse{
		Hashtag: &pb.Hashtag{
			Id:        hashtag.Id,
			CreatedAt: hashtag.CreatedAt,
			Title:     hashtag.Title,
		},
	}, nil
}
func (*server) GetHashtag(ctx context.Context, req *pb.GetHashtagRequest) (*pb.GetHashtagResponse, error) {
	hashtag, err := db.GetHashtagById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Can not get Hashtag", err)
	}
	return &pb.GetHashtagResponse{
		Hashtag: &pb.Hashtag{
			Id:        hashtag.Id,
			CreatedAt: hashtag.CreatedAt,
			Title:     hashtag.Title,
		},
	}, nil
}
func (*server) UpdateHashtag(ctx context.Context, req *pb.UpdateHashtagRequest) (*pb.UpdateHashtagResponse, error) {
	updateHashtag, err := db.UpdateHashtag(ctx, pg.UpdateHashtagParams{
		Id:    req.GetId(),
		Title: req.GetTitle(),
	})
	if err != nil {
		fmt.Println("Can not update Hashtag")
	}
	return &pb.UpdateHashtagResponse{
		Hashtag: &pb.Hashtag{
			Id:        updateHashtag.Id,
			CreatedAt: updateHashtag.CreatedAt,
			Title:     updateHashtag.Title,
		},
	}, nil
}
func (*server) DeleteHashtag(ctx context.Context, req *pb.DeleteHashtagRequest) (*pb.DeleteHashtagResponse, error) {
	delHashtag, err := db.DeleteHashtag(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Hashtag", err)
	}
	return &pb.DeleteHashtagResponse{
		Hashtag: &pb.Hashtag{
			Id:        delHashtag.Id,
			CreatedAt: delHashtag.CreatedAt,
			Title:     delHashtag.Title,
		},
	}, nil
}
func (*server) ListHashtags(ctx context.Context, req *pb.ListHashtagsRequest) (*pb.ListHashtagsResponse, error) {
	hashtags, err := db.ListHashtags(ctx)
	if err != nil {
		fmt.Println("Error listing Hashtags", err)
	}
	data := &pb.ListHashtagsResponse{}
	copier.Copy(&data.Hashtag, &hashtags)
	return &pb.ListHashtagsResponse{
		Hashtag: data.Hashtag,
	}, nil
}
func (*server) ListHashtagsById(ctx context.Context, req *pb.ListHashtagsByIdRequest) (*pb.ListHashtagsByIdResponse, error) {
	hashtags, err := db.ListHashtagsById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing Hashtags", err)
	}
	data := &pb.ListHashtagsByIdResponse{}
	copier.Copy(&data.Hashtag, &hashtags)
	return &pb.ListHashtagsByIdResponse{
		Hashtag: data.Hashtag,
	}, nil
}

// HashtagPost
func (*server) CreateHashtagPost(ctx context.Context, req *pb.CreateHashtagPostRequest) (*pb.CreateHashtagPostResponse, error) {
	createHasPost, err := db.CreateHashtagPost(ctx, pg.CreateHashtagPostParams{
		HashtagId: req.GetHashtagPost().GetHashtagId(),
		PostId:    req.GetHashtagPost().GetPostId(),
	})
	if err != nil {
		fmt.Println("Error creating hashtagPost", err)
	}
	return &pb.CreateHashtagPostResponse{
		HashtagPost: &pb.HashtagPost{
			Id:        createHasPost.Id,
			HashtagId: createHasPost.HashtagId,
			PostId:    createHasPost.PostId,
		},
	}, nil
}
func (*server) GetHashtagPost(ctx context.Context, req *pb.GetHashtagPostRequest) (*pb.GetHashtagPostResponse, error) {
	hashPost, err := db.GetHashtagPostById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting hashtagPost", err)
	}
	return &pb.GetHashtagPostResponse{
		HashtagPost: &pb.HashtagPost{
			Id:        hashPost.Id,
			HashtagId: hashPost.HashtagId,
			PostId:    hashPost.PostId,
		},
	}, nil
}
func (*server) DeleteHashtagPost(ctx context.Context, req *pb.DeleteHashtagPostRequest) (*pb.DeleteHashtagPostResponse, error) {
	delHasPost, err := db.DeleteHashtagPost(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting hashtagPost", err)
	}
	return &pb.DeleteHashtagPostResponse{
		HashtagPost: &pb.HashtagPost{
			Id:        delHasPost.Id,
			HashtagId: delHasPost.HashtagId,
			PostId:    delHasPost.PostId,
		},
	}, nil
}
func (*server) ListHashtagPosts(ctx context.Context, req *pb.ListHashtagPostsRequest) (*pb.ListHashtagPostsResponse, error) {
	hasPosts, err := db.ListHashtagsPost(ctx)
	if err != nil {
		fmt.Println("Error Listing hashtagsPost", err)
	}
	data := &pb.ListHashtagPostsResponse{}
	copier.Copy(&data.HashtagPost, &hasPosts)
	return &pb.ListHashtagPostsResponse{
		HashtagPost: data.HashtagPost,
	}, nil
}
func (*server) ListHashtagPostsById(ctx context.Context, req *pb.ListHashtagPostsByIdRequest) (*pb.ListHashtagPostsByIdResponse, error) {
	hasPosts, err := db.ListHashtagsPost(ctx)
	if err != nil {
		fmt.Println("Error Listing hashtagsPost", err)
	}
	data := &pb.ListHashtagPostsByIdResponse{}
	copier.Copy(&data.HashtagPost, &hasPosts)
	return &pb.ListHashtagPostsByIdResponse{
		HashtagPost: data.HashtagPost,
	}, nil
}

func main() {
	//if we crashed the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Conectring with postgres")
	pg.Connect()

	fmt.Println("Hashtag service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50055")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterHashtagsServiceServer(s, &server{})
	pb.RegisterHashtagPostsServiceServer(s, &server{})

	//Register reflection service on gRPC server
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	//Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	//BLock until a signal is received
	<-ch

	fmt.Println("Closing Conection Connection")
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of program")
}
