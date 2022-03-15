package main

import (
	"cesargdd/grpc-likes/pb"
	"cesargdd/grpc-likes/pg"
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
	pb.PostLikesServiceServer
	pb.CommentLikesServiceServer
}

var conn = pg.Connect()
var db = pg.New(conn)

// PostLikes

func (*server) CreatePostLike(ctx context.Context, req *pb.CreatePostLikeRequest) (*pb.CreatePostLikeResponse, error) {
	createPostLike, err := db.CreatePostLike(ctx, pg.CreatePostLikeParams{
		UserId:    req.GetPostLike().GetUserId(),
		PostId:    req.GetPostLike().GetPostId(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating postlike", err)
	}
	return &pb.CreatePostLikeResponse{
		PostLike: &pb.PostLike{
			Id:        createPostLike.Id,
			CreatedAt: createPostLike.CreatedAt,
			UserId:    createPostLike.UserId,
			PostId:    createPostLike.PostId,
		},
	}, nil
}
func (*server) GetPostLike(ctx context.Context, req *pb.GetPostLikeRequest) (*pb.GetPostLikeResponse, error) {
	postLike, err := db.GetPostLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("error getting postLike", err)
	}
	return &pb.GetPostLikeResponse{
		PostLike: &pb.PostLike{
			Id:        postLike.Id,
			CreatedAt: postLike.CreatedAt,
			UserId:    postLike.UserId,
			PostId:    postLike.PostId,
		},
	}, nil
}
func (*server) DeletePostLikeById(ctx context.Context, req *pb.DeletePostLikeByIdRequest) (*pb.DeletePostLikeByIdResponse, error) {
	delPostLike, err := db.DeletePostLikeById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting postLike", err)
	}
	return &pb.DeletePostLikeByIdResponse{
		PostLike: &pb.PostLike{
			Id:        delPostLike.Id,
			CreatedAt: delPostLike.CreatedAt,
			UserId:    delPostLike.UserId,
			PostId:    delPostLike.PostId,
		},
	}, nil
}
func (*server) DeletePostLikeByUserId(ctx context.Context, req *pb.DeletePostLikeByUserIdRequest) (*pb.DeletePostLikeByUserIdResponse, error) {
	delPostLike, err := db.DeletePostLikeByUserId(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting postLike", err)
	}
	return &pb.DeletePostLikeByUserIdResponse{
		PostLike: &pb.PostLike{
			Id:        delPostLike.Id,
			CreatedAt: delPostLike.CreatedAt,
			UserId:    delPostLike.UserId,
			PostId:    delPostLike.PostId,
		},
	}, nil
}
func (*server) ListPostLikes(ctx context.Context, req *pb.ListPostLikesRequest) (*pb.ListPostLikesResponse, error) {
	postLikes, err := db.ListPostLikes(ctx)
	if err != nil {
		fmt.Println("Error listing postLikes", err)
	}
	data := &pb.ListPostLikesResponse{}
	copier.Copy(&data.PostLike, &postLikes)
	return &pb.ListPostLikesResponse{
		PostLike: data.PostLike,
	}, nil
}

func (*server) ListPostLikesById(ctx context.Context, req *pb.ListPostLikesByIdRequest) (*pb.ListPostLikesByIdResponse, error) {
	postLikes, err := db.ListPostLikesById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing postLikes", err)
	}
	data := &pb.ListPostLikesByIdResponse{}
	copier.Copy(&data.PostLike, &postLikes)
	return &pb.ListPostLikesByIdResponse{
		PostLike: data.PostLike,
	}, nil
}

func (*server) ListPostLikesByPostId(ctx context.Context, req *pb.ListPostLikesByPostIdRequest) (*pb.ListPostLikesByPostIdResponse, error) {
	postLikes, err := db.ListPostLikesByPostId(ctx, req.GetPostId())
	if err != nil {
		fmt.Println("Error listing postLikes", err)
	}
	data := &pb.ListPostLikesByPostIdResponse{}
	copier.Copy(&data.PostLike, &postLikes)
	return &pb.ListPostLikesByPostIdResponse{
		PostLike: data.PostLike,
	}, nil
}

// CommentLikes

func (*server) CreateCommentLike(ctx context.Context, req *pb.CreateCommentLikeRequest) (*pb.CreateCommentLikeResponse, error) {
	createCommentLike, err := db.CreateCommentLike(ctx, pg.CreateCommentLikeParams{
		UserId:    req.GetCommentLike().GetUserId(),
		CommentId: req.GetCommentLike().GetCommentId(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Commentlike", err)
	}
	return &pb.CreateCommentLikeResponse{
		CommentLike: &pb.CommentLike{
			Id:        createCommentLike.Id,
			CreatedAt: createCommentLike.CreatedAt,
			UserId:    createCommentLike.UserId,
			CommentId: createCommentLike.CommentId,
		},
	}, nil
}
func (*server) GetCommentLike(ctx context.Context, req *pb.GetCommentLikeRequest) (*pb.GetCommentLikeResponse, error) {
	commentLike, err := db.GetCommentLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("error getting CommentLike", err)
	}
	return &pb.GetCommentLikeResponse{
		CommentLike: &pb.CommentLike{
			Id:        commentLike.Id,
			CreatedAt: commentLike.CreatedAt,
			UserId:    commentLike.UserId,
			CommentId: commentLike.CommentId,
		},
	}, nil
}
func (*server) DeleteCommentLikeById(ctx context.Context, req *pb.DeleteCommentLikeByIdRequest) (*pb.DeleteCommentLikeByIdResponse, error) {
	delCommentLike, err := db.DeleteCommentLikeById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting CommentLike", err)
	}
	return &pb.DeleteCommentLikeByIdResponse{
		CommentLike: &pb.CommentLike{
			Id:        delCommentLike.Id,
			CreatedAt: delCommentLike.CreatedAt,
			UserId:    delCommentLike.UserId,
			CommentId: delCommentLike.CommentId,
		},
	}, nil
}
func (*server) DeleteCommentLikeByUserId(ctx context.Context, req *pb.DeleteCommentLikeByUserIdRequest) (*pb.DeleteCommentLikeByUserIdResponse, error) {
	delCommentLike, err := db.DeleteCommentLikeByUserId(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting CommentLike", err)
	}
	return &pb.DeleteCommentLikeByUserIdResponse{
		CommentLike: &pb.CommentLike{
			Id:        delCommentLike.Id,
			CreatedAt: delCommentLike.CreatedAt,
			UserId:    delCommentLike.UserId,
			CommentId: delCommentLike.CommentId,
		},
	}, nil
}
func (*server) ListCommentLikes(ctx context.Context, req *pb.ListCommentLikesRequest) (*pb.ListCommentLikesResponse, error) {
	commentLikes, err := db.ListCommentLikes(ctx)
	if err != nil {
		fmt.Println("Error listing CommentLikes", err)
	}
	data := &pb.ListCommentLikesResponse{}
	copier.Copy(&data.CommentLike, &commentLikes)
	return &pb.ListCommentLikesResponse{
		CommentLike: data.CommentLike,
	}, nil
}

func (*server) ListCommentLikesById(ctx context.Context, req *pb.ListCommentLikesByIdRequest) (*pb.ListCommentLikesByIdResponse, error) {
	commentLikes, err := db.ListCommentLikesById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing CommentLikes", err)
	}
	data := &pb.ListCommentLikesByIdResponse{}
	copier.Copy(&data.CommentLike, &commentLikes)
	return &pb.ListCommentLikesByIdResponse{
		CommentLike: data.CommentLike,
	}, nil
}
func (*server) ListCommentLikesByCommentId(ctx context.Context, req *pb.ListCommentLikesByCommentIdRequest) (*pb.ListCommentLikesByCommentIdResponse, error) {
	commentLikes, err := db.ListCommentLikesByCommentId(ctx, req.GetCommentId())
	if err != nil {
		fmt.Println("Error listing CommentLikes", err)
	}
	data := &pb.ListCommentLikesByCommentIdResponse{}
	copier.Copy(&data.CommentLike, &commentLikes)
	return &pb.ListCommentLikesByCommentIdResponse{
		CommentLike: data.CommentLike,
	}, nil
}

func main() {
	//if we crashed the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Conectring with postgres")
	pg.Connect()

	fmt.Println("Likes service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50054")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterPostLikesServiceServer(s, &server{})
	pb.RegisterCommentLikesServiceServer(s, &server{})

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
