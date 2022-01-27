package main

import (
	"cesargdd/grpc-hashtags/hashtagspb"
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
	hashtagspb.HashtagsServiceServer
	hashtagspb.HashtagPostsServiceServer
}

var conn = pg.Connect()
var db = pg.New(conn)

// Hashtag

func (*server) CreateHashtag(ctx context.Context, req *hashtagspb.CreateHashtagRequest) (*hashtagspb.CreateHashtagResponse, error) {
	createHashtag, err := db.CreateHashtag(ctx, pg.CreateHashtagParams{
		Title:     req.GetHashtag().GetTitle(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Hashtag", err)
	}
	return &hashtagspb.CreateHashtagResponse{
		Hashtag: &hashtagspb.Hashtag{
			Id:        createHashtag.Id,
			CreatedAt: createHashtag.CreatedAt,
			Title:     createHashtag.Title,
		},
	}, nil
}
func (*server) GetHashtagByTitle(ctx context.Context, req *hashtagspb.GetHashtagByTitleRequest) (*hashtagspb.GetHashtagByTitleResponse, error) {
	hashtag, err := db.GetHashtagByTitle(ctx, req.GetTitle())
	if err != nil {
		fmt.Println("Can not get Hashtag", err)
	}
	return &hashtagspb.GetHashtagByTitleResponse{
		Hashtag: &hashtagspb.Hashtag{
			Id:        hashtag.Id,
			CreatedAt: hashtag.CreatedAt,
			Title:     hashtag.Title,
		},
	}, nil
}
func (*server) GetHashtag(ctx context.Context, req *hashtagspb.GetHashtagRequest) (*hashtagspb.GetHashtagResponse, error) {
	hashtag, err := db.GetHashtagById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Can not get Hashtag", err)
	}
	return &hashtagspb.GetHashtagResponse{
		Hashtag: &hashtagspb.Hashtag{
			Id:        hashtag.Id,
			CreatedAt: hashtag.CreatedAt,
			Title:     hashtag.Title,
		},
	}, nil
}
func (*server) UpdateHashtag(ctx context.Context, req *hashtagspb.UpdateHashtagRequest) (*hashtagspb.UpdateHashtagResponse, error) {
	updateHashtag, err := db.UpdateHashtag(ctx, pg.UpdateHashtagParams{
		Id:    req.GetId(),
		Title: req.GetTitle(),
	})
	if err != nil {
		fmt.Println("Can not update Hashtag")
	}
	return &hashtagspb.UpdateHashtagResponse{
		Hashtag: &hashtagspb.Hashtag{
			Id:        updateHashtag.Id,
			CreatedAt: updateHashtag.CreatedAt,
			Title:     updateHashtag.Title,
		},
	}, nil
}
func (*server) DeleteHashtag(ctx context.Context, req *hashtagspb.DeleteHashtagRequest) (*hashtagspb.DeleteHashtagResponse, error) {
	delHashtag, err := db.DeleteHashtag(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Hashtag", err)
	}
	return &hashtagspb.DeleteHashtagResponse{
		Hashtag: &hashtagspb.Hashtag{
			Id:        delHashtag.Id,
			CreatedAt: delHashtag.CreatedAt,
			Title:     delHashtag.Title,
		},
	}, nil
}
func (*server) ListHashtags(ctx context.Context, req *hashtagspb.ListHashtagsRequest) (*hashtagspb.ListHashtagsResponse, error) {
	hashtags, err := db.ListHashtags(ctx)
	if err != nil {
		fmt.Println("Error listing Hashtags", err)
	}
	data := &hashtagspb.ListHashtagsResponse{}
	copier.Copy(&data.Hashtag, &hashtags)
	return &hashtagspb.ListHashtagsResponse{
		Hashtag: data.Hashtag,
	}, nil
}
func (*server) ListHashtagsById(ctx context.Context, req *hashtagspb.ListHashtagsByIdRequest) (*hashtagspb.ListHashtagsByIdResponse, error) {
	hashtags, err := db.ListHashtagsById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing Hashtags", err)
	}
	data := &hashtagspb.ListHashtagsByIdResponse{}
	copier.Copy(&data.Hashtag, &hashtags)
	return &hashtagspb.ListHashtagsByIdResponse{
		Hashtag: data.Hashtag,
	}, nil
}

// HashtagPost
func (*server) CreateHashtagPost(ctx context.Context, req *hashtagspb.CreateHashtagPostRequest) (*hashtagspb.CreateHashtagPostResponse, error) {
	createHasPost, err := db.CreateHashtagPost(ctx, pg.CreateHashtagPostParams{
		HashtagId: req.GetHashtagPost().GetHashtagId(),
		PostId:    req.GetHashtagPost().GetPostId(),
	})
	if err != nil {
		fmt.Println("Error creating hashtagPost", err)
	}
	return &hashtagspb.CreateHashtagPostResponse{
		HashtagPost: &hashtagspb.HashtagPost{
			Id:        createHasPost.Id,
			HashtagId: createHasPost.HashtagId,
			PostId:    createHasPost.PostId,
		},
	}, nil
}
func (*server) GetHashtagPost(ctx context.Context, req *hashtagspb.GetHashtagPostRequest) (*hashtagspb.GetHashtagPostResponse, error) {
	hashPost, err := db.GetHashtagPostById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting hashtagPost", err)
	}
	return &hashtagspb.GetHashtagPostResponse{
		HashtagPost: &hashtagspb.HashtagPost{
			Id:        hashPost.Id,
			HashtagId: hashPost.HashtagId,
			PostId:    hashPost.PostId,
		},
	}, nil
}
func (*server) DeleteHashtagPost(ctx context.Context, req *hashtagspb.DeleteHashtagPostRequest) (*hashtagspb.DeleteHashtagPostResponse, error) {
	delHasPost, err := db.DeleteHashtagPost(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting hashtagPost", err)
	}
	return &hashtagspb.DeleteHashtagPostResponse{
		HashtagPost: &hashtagspb.HashtagPost{
			Id:        delHasPost.Id,
			HashtagId: delHasPost.HashtagId,
			PostId:    delHasPost.PostId,
		},
	}, nil
}
func (*server) ListHashtagPosts(ctx context.Context, req *hashtagspb.ListHashtagPostsRequest) (*hashtagspb.ListHashtagPostsResponse, error) {
	hasPosts, err := db.ListHashtagsPost(ctx)
	if err != nil {
		fmt.Println("Error Listing hashtagsPost", err)
	}
	data := &hashtagspb.ListHashtagPostsResponse{}
	copier.Copy(&data.HashtagPost, &hasPosts)
	return &hashtagspb.ListHashtagPostsResponse{
		HashtagPost: data.HashtagPost,
	}, nil
}
func (*server) ListHashtagPostsById(ctx context.Context, req *hashtagspb.ListHashtagPostsByIdRequest) (*hashtagspb.ListHashtagPostsByIdResponse, error) {
	hasPosts, err := db.ListHashtagsPost(ctx)
	if err != nil {
		fmt.Println("Error Listing hashtagsPost", err)
	}
	data := &hashtagspb.ListHashtagPostsByIdResponse{}
	copier.Copy(&data.HashtagPost, &hasPosts)
	return &hashtagspb.ListHashtagPostsByIdResponse{
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
	hashtagspb.RegisterHashtagsServiceServer(s, &server{})
	hashtagspb.RegisterHashtagPostsServiceServer(s, &server{})

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
