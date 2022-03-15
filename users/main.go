package main

import (
	"cesargdd/users-grpc/pb"
	"cesargdd/users-grpc/pg"
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
	pb.UsersServiceServer
	pb.FollowersServiceServer
}

var conn = pg.Connect()
var db = pg.New(conn)

// Users

func (*server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("Create user request")
	userReq := req.GetUser()
	user, err := db.CreateUser(ctx, pg.CreateUserParams{
		Username:  userReq.GetUsername(),
		Bio:       userReq.GetBio(),
		Avatar:    userReq.GetAvatar(),
		Email:     userReq.GetEmail(),
		Password:  userReq.GetPassword(),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error Creating User", err)
	}
	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:        user.Id,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Username:  user.Username,
			Bio:       user.Bio,
			Email:     user.Email,
			Password:  user.Password,
			Avatar:    user.Avatar,
		},
	}, nil
}

func (*server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	fmt.Println("Get user request")
	getUser, err := db.GetUserById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error getting the user", err)
	}
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:        getUser.Id,
			CreatedAt: getUser.CreatedAt,
			UpdatedAt: getUser.UpdatedAt,
			Username:  getUser.Username,
			Bio:       getUser.Bio,
			Email:     getUser.Email,
			Password:  getUser.Password,
			Avatar:    getUser.Avatar,
		},
	}, nil
}
func (*server) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.GetUserByUsernameResponse, error) {
	fmt.Println("Get user request")
	getUser, err := db.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		fmt.Println("Error getting the user", err)
	}
	return &pb.GetUserByUsernameResponse{
		User: &pb.User{
			Id:        getUser.Id,
			CreatedAt: getUser.CreatedAt,
			UpdatedAt: getUser.UpdatedAt,
			Username:  getUser.Username,
			Bio:       getUser.Bio,
			Email:     getUser.Email,
			Password:  getUser.Password,
			Avatar:    getUser.Avatar,
		},
	}, nil
}
func (*server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	updateUser, err := db.UpdateUser(ctx, pg.UpdateUserParams{
		Id:        req.GetId(),
		Bio:       req.GetBio(),
		Avatar:    req.GetAvatar(),
		UpdatedAt: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("error updating user", err)
	}
	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:        updateUser.Id,
			CreatedAt: updateUser.CreatedAt,
			UpdatedAt: updateUser.UpdatedAt,
			Username:  updateUser.Username,
			Bio:       updateUser.Bio,
			Email:     updateUser.Email,
			Password:  updateUser.Password,
			Avatar:    updateUser.Avatar,
		},
	}, nil
}
func (*server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	deleteUser, err := db.DeleteUser(ctx, req.GetId())
	if err != nil {
		fmt.Println("error deleting user", err)
	}
	return &pb.DeleteUserResponse{
		User: &pb.User{
			Id:       deleteUser.Id,
			Username: deleteUser.Username,
			Email:    deleteUser.Email,
		},
	}, nil
}

func (*server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	fmt.Println("List user request")
	users, err := db.ListUsers(ctx)
	if err != nil {
		fmt.Println("error listing users", err)
	}
	data := &pb.ListUsersResponse{}
	copier.Copy(&data.User, &users)
	return &pb.ListUsersResponse{
		User: data.User,
	}, nil
}

func (*server) ListUsersById(ctx context.Context, req *pb.ListUsersByIdRequest) (*pb.ListUsersByIdResponse, error) {
	users, err := db.ListUsersById(ctx, req.GetId())
	if err != nil {
		fmt.Println("error listing users by id", err)
	}
	data := &pb.ListUsersByIdResponse{}
	copier.Copy(&data.User, &users)
	return &pb.ListUsersByIdResponse{
		User: data.User,
	}, nil
}

func (*server) CreateFollower(ctx context.Context, req *pb.CreateFollowerRequest) (*pb.CreateFollowerResponse, error) {
	createFollower, err := db.CreateFollower(ctx, pg.CreateFollowerParams{
		LeaderId:   req.GetFollower().GetLeaderId(),
		FollowerId: req.GetFollower().GetFollowerId(),
		CreatedAt:  time.Now().Unix(),
	})
	if err != nil {
		fmt.Println("Error creating Follower", err)
	}
	return &pb.CreateFollowerResponse{
		Follower: &pb.Follower{
			Id:         createFollower.Id,
			CreatedAt:  createFollower.CreatedAt,
			LeaderId:   createFollower.LeaderId,
			FollowerId: createFollower.FollowerId,
		},
	}, nil
}
func (*server) GetFollower(ctx context.Context, req *pb.GetFollowerRequest) (*pb.GetFollowerResponse, error) {
	follower, err := db.GetFollower(ctx, req.GetId())
	if err != nil {
		fmt.Println("error getting Follower", err)
	}
	return &pb.GetFollowerResponse{
		Follower: &pb.Follower{
			Id:         follower.Id,
			CreatedAt:  follower.CreatedAt,
			LeaderId:   follower.LeaderId,
			FollowerId: follower.FollowerId,
		},
	}, nil
}
func (*server) DeleteFollower(ctx context.Context, req *pb.DeleteFollowerRequest) (*pb.DeleteFollowerResponse, error) {
	delFollower, err := db.DeleteFollower(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error deleting Follower", err)
	}
	return &pb.DeleteFollowerResponse{
		Follower: &pb.Follower{
			Id:         delFollower.Id,
			CreatedAt:  delFollower.CreatedAt,
			LeaderId:   delFollower.LeaderId,
			FollowerId: delFollower.FollowerId,
		},
	}, nil
}
func (*server) ListFollowers(ctx context.Context, req *pb.ListFollowersRequest) (*pb.ListFollowersResponse, error) {
	Followers, err := db.ListFollowers(ctx)
	if err != nil {
		fmt.Println("Error listing Followers", err)
	}
	data := &pb.ListFollowersResponse{}
	copier.Copy(&data.Follower, &Followers)
	return &pb.ListFollowersResponse{
		Follower: data.Follower,
	}, nil
}
func (*server) ListFollowersById(ctx context.Context, req *pb.ListFollowersByIdRequest) (*pb.ListFollowersByIdResponse, error) {
	Followers, err := db.ListFollowersById(ctx, req.GetId())
	if err != nil {
		fmt.Println("Error listing Followers", err)
	}
	data := &pb.ListFollowersByIdResponse{}
	copier.Copy(&data.Follower, &Followers)
	return &pb.ListFollowersByIdResponse{
		Follower: data.Follower,
	}, nil
}

func (*server) ListFollowersByLeaderId(ctx context.Context, req *pb.ListFollowersByLeaderIdRequest) (*pb.ListFollowersByLeaderIdResponse, error) {
	Followers, err := db.ListFollowersByLeaderId(ctx, req.GetLeaderId())
	if err != nil {
		fmt.Println("Error listing Followers", err)
	}
	data := &pb.ListFollowersByLeaderIdResponse{}
	copier.Copy(&data.Follower, &Followers)
	return &pb.ListFollowersByLeaderIdResponse{
		Follower: data.Follower,
	}, nil
}

func main() {
	//if we crashed the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Conectring with postgres")
	pg.Connect()

	fmt.Println("Users service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterUsersServiceServer(s, &server{})
	pb.RegisterFollowersServiceServer(s, &server{})

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
