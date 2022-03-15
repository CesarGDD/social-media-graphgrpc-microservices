package main

import (
	"cesargdd/posts-grpc/pb"
	"cesargdd/posts-grpc/pg"
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
	pb.PostsServiceServer
}

var conn = pg.Connect()
var db = pg.New(conn)

// Posts

func (*server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	fmt.Println("Create post request")
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
	return &pb.CreatePostResponse{
		Post: &pb.Post{
			Id:        createPost.Id,
			CreatedAt: createPost.CreatedAt,
			UpdatedAt: createPost.UpdatedAt,
			Url:       createPost.Url,
			Caption:   createPost.Caption,
			UserId:    createPost.UserId,
		},
	}, nil
}
func (*server) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, err := db.GetPost(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting Post", err)
	}
	return &pb.GetPostResponse{
		Post: &pb.Post{
			Id:        post.Id,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Url:       post.Url,
			Caption:   post.Caption,
			UserId:    post.UserId,
		},
	}, nil
}
func (*server) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	updatePost, err := db.UpdatePost(ctx, pg.UpdatePostParams{
		Id:        req.GetPostId(),
		Url:       req.GetUrl(),
		Caption:   req.GetCaption(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error updating Post")
	}
	return &pb.UpdatePostResponse{
		Post: &pb.Post{
			Id:        updatePost.Id,
			CreatedAt: updatePost.CreatedAt,
			UpdatedAt: updatePost.UpdatedAt,
			Url:       updatePost.Url,
			Caption:   updatePost.Caption,
			UserId:    updatePost.UserId,
		},
	}, nil
}
func (*server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	delPost, err := db.DeletePost(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Post")
	}
	return &pb.DeletePostResponse{
		Post: &pb.Post{
			Id:        delPost.Id,
			CreatedAt: delPost.CreatedAt,
			UpdatedAt: delPost.UpdatedAt,
			Url:       delPost.Url,
			Caption:   delPost.Caption,
			UserId:    delPost.UserId,
		},
	}, nil
}
func (*server) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	fmt.Println("List posts request")
	listPost, err := db.ListPosts(ctx)
	if err != nil {
		fmt.Println("Error getting Posts")
	}
	data := &pb.ListPostsResponse{}
	copier.Copy(&data.Post, &listPost)
	return &pb.ListPostsResponse{
		Post: data.Post,
	}, nil
}

func (*server) ListPostsById(ctx context.Context, req *pb.ListPostsByIdRequest) (*pb.ListPostsByIdResponse, error) {
	listPost, err := db.ListPostsById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting Posts")
	}
	data := &pb.ListPostsByIdResponse{}
	copier.Copy(&data.Post, &listPost)
	return &pb.ListPostsByIdResponse{
		Post: data.Post,
	}, nil
}

func (*server) ListPostsByUserId(ctx context.Context, req *pb.ListPostsByUserIdRequest) (*pb.ListPostsByUserIdResponse, error) {
	listPost, err := db.ListPostsByUserId(ctx, req.GetUserId())
	if err != nil {
		fmt.Println("Error getting Posts")
	}
	data := &pb.ListPostsByUserIdResponse{}
	copier.Copy(&data.Post, &listPost)
	return &pb.ListPostsByUserIdResponse{
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
	pb.RegisterPostsServiceServer(s, &server{})

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
