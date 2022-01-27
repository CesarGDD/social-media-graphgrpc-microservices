package main

import (
	"cesargdd/posts-grpc/pg"
	"cesargdd/posts-grpc/postspb"
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
	postspb.PostsServiceServer
}

var conn = pg.Connect()
var db = pg.New(conn)

// Posts

func (*server) CreatePost(ctx context.Context, req *postspb.CreatePostRequest) (*postspb.CreatePostResponse, error) {
	postReq := req.GetPost()
	createPost, err := db.CreatePost(ctx, pg.CreatePostParams{
		Caption:   postReq.GetCaption(),
		Url:       postReq.GetUrl(),
		UserId:    postReq.GetUserId(),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("error creating Post", err)
	}
	return &postspb.CreatePostResponse{
		Post: &postspb.Post{
			Id:        createPost.Id,
			CreatedAt: createPost.CreatedAt,
			UpdatedAt: createPost.UpdatedAt,
			Url:       createPost.Url,
			Caption:   createPost.Caption,
			UserId:    createPost.UserId,
		},
	}, nil
}
func (*server) GetPost(ctx context.Context, req *postspb.GetPostRequest) (*postspb.GetPostResponse, error) {
	post, err := db.GetPost(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting Post", err)
	}
	return &postspb.GetPostResponse{
		Post: &postspb.Post{
			Id:        post.Id,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Url:       post.Url,
			Caption:   post.Caption,
			UserId:    post.UserId,
		},
	}, nil
}
func (*server) UpdatePost(ctx context.Context, req *postspb.UpdatePostRequest) (*postspb.UpdatePostResponse, error) {
	updatePost, err := db.UpdatePost(ctx, pg.UpdatePostParams{
		Id:        req.GetPostId(),
		Url:       req.GetUrl(),
		Caption:   req.GetCaption(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error updating Post")
	}
	return &postspb.UpdatePostResponse{
		Post: &postspb.Post{
			Id:        updatePost.Id,
			CreatedAt: updatePost.CreatedAt,
			UpdatedAt: updatePost.UpdatedAt,
			Url:       updatePost.Url,
			Caption:   updatePost.Caption,
			UserId:    updatePost.UserId,
		},
	}, nil
}
func (*server) DeletePost(ctx context.Context, req *postspb.DeletePostRequest) (*postspb.DeletePostResponse, error) {
	delPost, err := db.DeletePost(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Post")
	}
	return &postspb.DeletePostResponse{
		Post: &postspb.Post{
			Id:        delPost.Id,
			CreatedAt: delPost.CreatedAt,
			UpdatedAt: delPost.UpdatedAt,
			Url:       delPost.Url,
			Caption:   delPost.Caption,
			UserId:    delPost.UserId,
		},
	}, nil
}
func (*server) ListPosts(ctx context.Context, req *postspb.ListPostsRequest) (*postspb.ListPostsResponse, error) {
	listPost, err := db.ListPosts(ctx)
	if err != nil {
		fmt.Println("Error getting Posts")
	}
	data := &postspb.ListPostsResponse{}
	copier.Copy(&data.Post, &listPost)
	return &postspb.ListPostsResponse{
		Post: data.Post,
	}, nil
}

func (*server) ListPostsById(ctx context.Context, req *postspb.ListPostsByIdRequest) (*postspb.ListPostsByIdResponse, error) {
	listPost, err := db.ListPostsById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting Posts")
	}
	data := &postspb.ListPostsByIdResponse{}
	copier.Copy(&data.Post, &listPost)
	return &postspb.ListPostsByIdResponse{
		Post: data.Post,
	}, nil
}

func (*server) ListPostsByUserId(ctx context.Context, req *postspb.ListPostsByUserIdRequest) (*postspb.ListPostsByUserIdResponse, error) {
	listPost, err := db.ListPostsByUserId(ctx, req.GetUserId())
	if err != nil {
		fmt.Println("Error getting Posts")
	}
	data := &postspb.ListPostsByUserIdResponse{}
	copier.Copy(&data.Post, &listPost)
	return &postspb.ListPostsByUserIdResponse{
		Post: data.Post,
	}, nil
}

func main() {
	//if we crashed the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Conectring with postgres")
	pg.Connect()

	fmt.Println("Posts service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	postspb.RegisterPostsServiceServer(s, &server{})

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
