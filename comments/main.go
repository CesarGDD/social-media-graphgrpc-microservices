package main

import (
	"cesargdd/grpc-comments/pb"
	"cesargdd/grpc-comments/pg"
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
	pb.CommentsServiceServer
}

var conn = pg.Connect()
var db = pg.New(conn)

// Comments

func (*server) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	createComment, err := db.CreateComment(ctx, pg.CreateCommentParams{
		Contents:  req.GetComment().GetContents(),
		UserId:    req.GetComment().GetUserId(),
		PostId:    req.GetComment().GetPostId(),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Comment", err)
	}
	return &pb.CreateCommentResponse{
		Comment: &pb.Comment{
			Id:        createComment.Id,
			CreatedAt: createComment.CreatedAt,
			UpdatedAt: createComment.UpdatedAt,
			Contents:  createComment.Contents,
			UserId:    createComment.UserId,
			PostId:    createComment.PostId,
		},
	}, nil
}
func (*server) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	comment, err := db.GetComment(ctx, req.GetId())
	if err != nil {
		fmt.Println("Can not get comment", err)
	}
	return &pb.GetCommentResponse{
		Comment: &pb.Comment{
			Id:        comment.Id,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
			Contents:  comment.Contents,
			UserId:    comment.UserId,
			PostId:    comment.UserId,
		},
	}, nil
}
func (*server) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	updateComment, err := db.UpdateComment(ctx, pg.UpdateCommentParams{
		Id:        req.GetId(),
		Contents:  req.GetContents(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Can not update comment")
	}
	return &pb.UpdateCommentResponse{
		Comment: &pb.Comment{
			Id:        updateComment.Id,
			CreatedAt: updateComment.CreatedAt,
			UpdatedAt: updateComment.UpdatedAt,
			Contents:  updateComment.Contents,
			UserId:    updateComment.UserId,
			PostId:    updateComment.UserId,
		},
	}, nil
}
func (*server) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	delComment, err := db.DeleteComment(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Comment", err)
	}
	return &pb.DeleteCommentResponse{
		Comment: &pb.Comment{
			Id:        delComment.Id,
			CreatedAt: delComment.CreatedAt,
			UpdatedAt: delComment.UpdatedAt,
			Contents:  delComment.Contents,
			UserId:    delComment.UserId,
			PostId:    delComment.UserId,
		},
	}, nil
}
func (*server) ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	comments, err := db.ListComments(ctx)
	if err != nil {
		fmt.Println("Error listing comments", err)
	}
	data := &pb.ListCommentsResponse{}
	copier.Copy(&data.Comment, &comments)
	return &pb.ListCommentsResponse{
		Comment: data.Comment,
	}, nil
}

func (*server) ListCommentsById(ctx context.Context, req *pb.ListCommentsByIdRequest) (*pb.ListCommentsByIdResponse, error) {
	comments, err := db.ListCommentsById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing comments", err)
	}
	data := &pb.ListCommentsByIdResponse{}
	copier.Copy(&data.Comment, &comments)
	return &pb.ListCommentsByIdResponse{
		Comment: data.Comment,
	}, nil
}
func (*server) ListCommentsByPostId(ctx context.Context, req *pb.ListCommentsByPostIdRequest) (*pb.ListCommentsByPostIdResponse, error) {
	comments, err := db.ListCommentsByPostId(ctx, req.GetPostId())
	if err != nil {
		fmt.Println("Error listing comments", err)
	}
	data := &pb.ListCommentsByPostIdResponse{}
	copier.Copy(&data.Comment, &comments)
	return &pb.ListCommentsByPostIdResponse{
		Comment: data.Comment,
	}, nil
}

func main() {
	//if we crashed the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Conectring with postgres")
	pg.Connect()

	fmt.Println("Comment service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterCommentsServiceServer(s, &server{})

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
