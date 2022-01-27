package main

import (
	"cesargdd/grpc-likes/likespb"
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
	likespb.PostLikesServiceServer
	likespb.CommentLikesServiceServer
}

var conn = pg.Connect()
var db = pg.New(conn)

// PostLikes

func (*server) CreatePostLike(ctx context.Context, req *likespb.CreatePostLikeRequest) (*likespb.CreatePostLikeResponse, error) {
	createPostLike, err := db.CreatePostLike(ctx, pg.CreatePostLikeParams{
		UserId:    req.GetPostLike().GetUserId(),
		PostId:    req.GetPostLike().GetPostId(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating postlike", err)
	}
	return &likespb.CreatePostLikeResponse{
		PostLike: &likespb.PostLike{
			Id:        createPostLike.Id,
			CreatedAt: createPostLike.CreatedAt,
			UserId:    createPostLike.UserId,
			PostId:    createPostLike.PostId,
		},
	}, nil
}
func (*server) GetPostLike(ctx context.Context, req *likespb.GetPostLikeRequest) (*likespb.GetPostLikeResponse, error) {
	postLike, err := db.GetPostLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("error getting postLike", err)
	}
	return &likespb.GetPostLikeResponse{
		PostLike: &likespb.PostLike{
			Id:        postLike.Id,
			CreatedAt: postLike.CreatedAt,
			UserId:    postLike.UserId,
			PostId:    postLike.PostId,
		},
	}, nil
}
func (*server) DeletePostLike(ctx context.Context, req *likespb.DeletePostLikeRequest) (*likespb.DeletePostLikeResponse, error) {
	delPostLike, err := db.DeletePostLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting postLike", err)
	}
	return &likespb.DeletePostLikeResponse{
		PostLike: &likespb.PostLike{
			Id:        delPostLike.Id,
			CreatedAt: delPostLike.CreatedAt,
			UserId:    delPostLike.UserId,
			PostId:    delPostLike.PostId,
		},
	}, nil
}
func (*server) ListPostLikes(ctx context.Context, req *likespb.ListPostLikesRequest) (*likespb.ListPostLikesResponse, error) {
	postLikes, err := db.ListPostLikes(ctx)
	if err != nil {
		fmt.Println("Error listing postLikes", err)
	}
	data := &likespb.ListPostLikesResponse{}
	copier.Copy(&data.PostLike, &postLikes)
	return &likespb.ListPostLikesResponse{
		PostLike: data.PostLike,
	}, nil
}

func (*server) ListPostLikesById(ctx context.Context, req *likespb.ListPostLikesByIdRequest) (*likespb.ListPostLikesByIdResponse, error) {
	postLikes, err := db.ListPostLikesById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing postLikes", err)
	}
	data := &likespb.ListPostLikesByIdResponse{}
	copier.Copy(&data.PostLike, &postLikes)
	return &likespb.ListPostLikesByIdResponse{
		PostLike: data.PostLike,
	}, nil
}

func (*server) ListPostLikesByPostId(ctx context.Context, req *likespb.ListPostLikesByPostIdRequest) (*likespb.ListPostLikesByPostIdResponse, error) {
	postLikes, err := db.ListPostLikesByPostId(ctx, req.GetPostId())
	if err != nil {
		fmt.Println("Error listing postLikes", err)
	}
	data := &likespb.ListPostLikesByPostIdResponse{}
	copier.Copy(&data.PostLike, &postLikes)
	return &likespb.ListPostLikesByPostIdResponse{
		PostLike: data.PostLike,
	}, nil
}

// CommentLikes

func (*server) CreateCommentLike(ctx context.Context, req *likespb.CreateCommentLikeRequest) (*likespb.CreateCommentLikeResponse, error) {
	createCommentLike, err := db.CreateCommentLike(ctx, pg.CreateCommentLikeParams{
		UserId:    req.GetCommentLike().GetUserId(),
		CommentId: req.GetCommentLike().GetCommentId(),
		CreatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Commentlike", err)
	}
	return &likespb.CreateCommentLikeResponse{
		CommentLike: &likespb.CommentLike{
			Id:        createCommentLike.Id,
			CreatedAt: createCommentLike.CreatedAt,
			UserId:    createCommentLike.UserId,
			CommentId: createCommentLike.CommentId,
		},
	}, nil
}
func (*server) GetCommentLike(ctx context.Context, req *likespb.GetCommentLikeRequest) (*likespb.GetCommentLikeResponse, error) {
	commentLike, err := db.GetCommentLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("error getting CommentLike", err)
	}
	return &likespb.GetCommentLikeResponse{
		CommentLike: &likespb.CommentLike{
			Id:        commentLike.Id,
			CreatedAt: commentLike.CreatedAt,
			UserId:    commentLike.UserId,
			CommentId: commentLike.CommentId,
		},
	}, nil
}
func (*server) DeleteCommentLike(ctx context.Context, req *likespb.DeleteCommentLikeRequest) (*likespb.DeleteCommentLikeResponse, error) {
	delCommentLike, err := db.DeleteCommentLike(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting CommentLike", err)
	}
	return &likespb.DeleteCommentLikeResponse{
		CommentLike: &likespb.CommentLike{
			Id:        delCommentLike.Id,
			CreatedAt: delCommentLike.CreatedAt,
			UserId:    delCommentLike.UserId,
			CommentId: delCommentLike.CommentId,
		},
	}, nil
}
func (*server) ListCommentLikes(ctx context.Context, req *likespb.ListCommentLikesRequest) (*likespb.ListCommentLikesResponse, error) {
	commentLikes, err := db.ListCommentLikes(ctx)
	if err != nil {
		fmt.Println("Error listing CommentLikes", err)
	}
	data := &likespb.ListCommentLikesResponse{}
	copier.Copy(&data.CommentLike, &commentLikes)
	return &likespb.ListCommentLikesResponse{
		CommentLike: data.CommentLike,
	}, nil
}

func (*server) ListCommentLikesById(ctx context.Context, req *likespb.ListCommentLikesByIdRequest) (*likespb.ListCommentLikesByIdResponse, error) {
	commentLikes, err := db.ListCommentLikesById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing CommentLikes", err)
	}
	data := &likespb.ListCommentLikesByIdResponse{}
	copier.Copy(&data.CommentLike, &commentLikes)
	return &likespb.ListCommentLikesByIdResponse{
		CommentLike: data.CommentLike,
	}, nil
}
func (*server) ListCommentLikesByCommentId(ctx context.Context, req *likespb.ListCommentLikesByCommentIdRequest) (*likespb.ListCommentLikesByCommentIdResponse, error) {
	commentLikes, err := db.ListCommentLikesByCommentId(ctx, req.GetCommentId())
	if err != nil {
		fmt.Println("Error listing CommentLikes", err)
	}
	data := &likespb.ListCommentLikesByCommentIdResponse{}
	copier.Copy(&data.CommentLike, &commentLikes)
	return &likespb.ListCommentLikesByCommentIdResponse{
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
	likespb.RegisterPostLikesServiceServer(s, &server{})
	likespb.RegisterCommentLikesServiceServer(s, &server{})

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
