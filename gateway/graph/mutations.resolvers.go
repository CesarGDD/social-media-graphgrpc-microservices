package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/grpc-gateway/authUser"
	"cesargdd/grpc-gateway/graph/generated"
	"cesargdd/grpc-gateway/graph/model"
	"cesargdd/grpc-gateway/pb"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	res, err := u.GetUserByUsername(ctx, &pb.GetUserByUsernameRequest{Username: input.Username})
	if err != nil {
		return nil, errors.New("invalid username")
	}
	if !CheckPasswordHash(input.Password, res.GetUser().Password) {
		return nil, errors.New("invalid password")
	}
	token, err := authUser.GenerateToken(res.User.Username)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: &model.AuthToken{
			AccessToken: token,
		},
		User: &pb.User{
			Id:        res.User.Id,
			CreatedAt: res.User.CreatedAt,
			UpdatedAt: res.User.UpdatedAt,
			Username:  res.User.Username,
			Bio:       res.User.Bio,
			Email:     res.User.Email,
			Avatar:    res.User.Avatar,
		},
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.AuthResponse, error) {
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	inputUser := pb.User{
		Username: input.Username,
		Bio:      input.Bio,
		Email:    input.Email,
		Password: hashedPassword,
		Avatar:   input.Avatar,
	}

	res, err := u.CreateUser(ctx, &pb.CreateUserRequest{User: &inputUser})
	if err != nil {
		fmt.Println(err)
	}
	user := res.GetUser()
	token, err := authUser.GenerateToken(user.Username)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{AuthToken: &model.AuthToken{
		AccessToken: token,
	},
		User: &pb.User{
			Id:        user.Id,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Username:  user.Username,
			Bio:       user.Bio,
			Email:     user.Email,
			Avatar:    user.Avatar,
		}}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.EditUser) (*pb.User, error) {
	userCon := authUser.ForContext(ctx)
	if userCon == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := u.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:     int32(input.ID),
		Bio:    *input.Bio,
		Avatar: *input.Avatar,
	})
	if err != nil {
		fmt.Println(err)
	}
	user := res.User
	return &pb.User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Email:     user.Email,
		Password:  user.Password,
		Avatar:    user.Avatar,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input *int) (*pb.User, error) {
	userCon := authUser.ForContext(ctx)
	if userCon == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := u.DeleteUser(ctx, &pb.DeleteUserRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	user := res.User
	return &pb.User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Bio:       user.Bio,
		Email:     user.Email,
		Password:  user.Password,
		Avatar:    user.Avatar,
	}, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*pb.Post, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := p.CreatePost(ctx, &pb.CreatePostRequest{Post: &pb.Post{
		Url:     input.URL,
		Caption: *input.Caption,
		UserId:  int32(input.UserID),
	}})
	if err != nil {
		fmt.Println(err)
	}
	post := res.Post
	return &pb.Post{
		Id:        post.Id,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Url:       post.Url,
		Caption:   post.Caption,
		UserId:    post.UserId,
	}, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input *model.EditPost) (*pb.Post, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := p.UpdatePost(ctx, &pb.UpdatePostRequest{
		PostId:  int32(input.ID),
		Url:     *input.URL,
		Caption: *input.Caption,
	})
	if err != nil {
		fmt.Println(err)
	}
	post := res.Post
	return &pb.Post{
		Id:        post.Id,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Url:       post.Url,
		Caption:   post.Caption,
		UserId:    post.UserId,
	}, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *int) (*pb.Post, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := p.DeletePost(ctx, &pb.DeletePostRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	post := res.Post
	return &pb.Post{
		Id:        post.Id,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Url:       post.Url,
		Caption:   post.Caption,
		UserId:    post.UserId,
	}, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*pb.Comment, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := c.CreateComment(ctx, &pb.CreateCommentRequest{
		Comment: &pb.Comment{
			Contents: input.Contents,
			UserId:   int32(input.UserID),
			PostId:   int32(input.PostID),
		}})
	if err != nil {
		fmt.Println(err)
	}
	comment := res.Comment
	return &pb.Comment{
		Id:        comment.Id,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Contents:  comment.Contents,
		UserId:    comment.UserId,
		PostId:    comment.PostId,
	}, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, input *model.EditComment) (*pb.Comment, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := c.UpdateComment(ctx, &pb.UpdateCommentRequest{
		Id:       int32(input.ID),
		Contents: input.Contents,
	})
	if err != nil {
		fmt.Println(err)
	}
	comment := res.Comment
	return &pb.Comment{
		Id:        comment.Id,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Contents:  comment.Contents,
		UserId:    comment.UserId,
		PostId:    comment.PostId,
	}, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, input *int) (*pb.Comment, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := c.DeleteComment(ctx, &pb.DeleteCommentRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	comment := res.Comment
	return &pb.Comment{
		Id:        comment.Id,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Contents:  comment.Contents,
		UserId:    comment.UserId,
		PostId:    comment.PostId,
	}, nil
}

func (r *mutationResolver) CreatePostLike(ctx context.Context, input *model.NewPostLike) (*pb.PostLike, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := pl.CreatePostLike(ctx, &pb.CreatePostLikeRequest{
		PostLike: &pb.PostLike{
			UserId: int32(input.UserID),
			PostId: int32(input.PostID),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	postLike := res.PostLike
	return &pb.PostLike{
		Id:        postLike.Id,
		CreatedAt: postLike.CreatedAt,
		UserId:    postLike.UserId,
		PostId:    postLike.PostId,
	}, nil
}

func (r *mutationResolver) DeletePostLikeByID(ctx context.Context, input *int) (*pb.PostLike, error) {
	res, err := pl.DeletePostLikeById(ctx, &pb.DeletePostLikeByIdRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	postLike := res.PostLike
	return &pb.PostLike{
		Id:        postLike.Id,
		CreatedAt: postLike.CreatedAt,
		UserId:    postLike.UserId,
		PostId:    postLike.PostId,
	}, nil
}

func (r *mutationResolver) DeletePostLikeByUserID(ctx context.Context, input *int) (*pb.PostLike, error) {
	res, err := pl.DeletePostLikeByUserId(ctx, &pb.DeletePostLikeByUserIdRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	postLike := res.PostLike
	return &pb.PostLike{
		Id:        postLike.Id,
		CreatedAt: postLike.CreatedAt,
		UserId:    postLike.UserId,
		PostId:    postLike.PostId,
	}, nil
}

func (r *mutationResolver) CreateCommentLike(ctx context.Context, input *model.NewCommentLike) (*pb.CommentLike, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := cl.CreateCommentLike(ctx, &pb.CreateCommentLikeRequest{
		CommentLike: &pb.CommentLike{
			UserId:    int32(input.UserID),
			CommentId: int32(input.CommentID),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	comLike := res.CommentLike
	return &pb.CommentLike{
		Id:        comLike.Id,
		CreatedAt: comLike.CreatedAt,
		UserId:    comLike.UserId,
		CommentId: comLike.CommentId,
	}, nil
}

func (r *mutationResolver) DeleteCommentLikeByID(ctx context.Context, input *int) (*pb.CommentLike, error) {
	res, err := cl.DeleteCommentLikeById(ctx, &pb.DeleteCommentLikeByIdRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	comLike := res.CommentLike
	return &pb.CommentLike{
		Id:        comLike.Id,
		CreatedAt: comLike.CreatedAt,
		UserId:    comLike.UserId,
		CommentId: comLike.CommentId,
	}, nil
}

func (r *mutationResolver) DeleteCommentLikeByUserID(ctx context.Context, input *int) (*pb.CommentLike, error) {
	res, err := cl.DeleteCommentLikeByUserId(ctx, &pb.DeleteCommentLikeByUserIdRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	comLike := res.CommentLike
	return &pb.CommentLike{
		Id:        comLike.Id,
		CreatedAt: comLike.CreatedAt,
		UserId:    comLike.UserId,
		CommentId: comLike.CommentId,
	}, nil
}

func (r *mutationResolver) CreateHashtag(ctx context.Context, input *model.NewHashtag) (*pb.Hashtag, error) {
	res, err := h.CreateHashtag(ctx, &pb.CreateHashtagRequest{
		Hashtag: &pb.Hashtag{
			Title: input.Title,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	hashtag := res.Hashtag
	return &pb.Hashtag{
		Id:        hashtag.Id,
		CreatedAt: hashtag.CreatedAt,
		Title:     hashtag.Title,
	}, nil
}

func (r *mutationResolver) UpdateHashtag(ctx context.Context, input *model.EditHashtag) (*pb.Hashtag, error) {
	res, err := h.UpdateHashtag(ctx, &pb.UpdateHashtagRequest{
		Id:    int32(input.ID),
		Title: input.Title,
	})
	if err != nil {
		fmt.Println(err)
	}
	hashtag := res.Hashtag
	return &pb.Hashtag{
		Id:        hashtag.Id,
		CreatedAt: hashtag.CreatedAt,
		Title:     hashtag.Title,
	}, nil
}

func (r *mutationResolver) DeleteHashtag(ctx context.Context, input *int) (*pb.Hashtag, error) {
	res, err := h.DeleteHashtag(ctx, &pb.DeleteHashtagRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	hashtag := res.Hashtag
	return &pb.Hashtag{
		Id:        hashtag.Id,
		CreatedAt: hashtag.CreatedAt,
		Title:     hashtag.Title,
	}, nil
}

func (r *mutationResolver) CreateHashtagPost(ctx context.Context, input *model.NewHashtagPost) (*pb.HashtagPost, error) {
	res, err := hp.CreateHashtagPost(ctx, &pb.CreateHashtagPostRequest{
		HashtagPost: &pb.HashtagPost{
			HashtagId: int32(input.HashtagID),
			PostId:    int32(input.PostID),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	hasPost := res.HashtagPost
	return &pb.HashtagPost{
		Id:        hasPost.Id,
		HashtagId: hasPost.HashtagId,
		PostId:    hasPost.PostId,
	}, nil
}

func (r *mutationResolver) DeleteHashtagPost(ctx context.Context, input *int) (*pb.HashtagPost, error) {
	res, err := hp.DeleteHashtagPost(ctx, &pb.DeleteHashtagPostRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	hasPost := res.HashtagPost
	return &pb.HashtagPost{
		Id:        hasPost.Id,
		HashtagId: hasPost.HashtagId,
		PostId:    hasPost.PostId,
	}, nil
}

func (r *mutationResolver) CreateFollower(ctx context.Context, input *model.NewFollower) (*pb.Follower, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := f.CreateFollower(ctx, &pb.CreateFollowerRequest{
		Follower: &pb.Follower{
			LeaderId:   int32(input.LeaderID),
			FollowerId: int32(input.FollowerID),
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	follower := res.Follower
	return &pb.Follower{
		Id:         follower.Id,
		CreatedAt:  follower.CreatedAt,
		LeaderId:   follower.LeaderId,
		FollowerId: follower.FollowerId,
	}, nil
}

func (r *mutationResolver) DeleteFollower(ctx context.Context, input *int) (*pb.Follower, error) {
	user := authUser.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	res, err := f.DeleteFollower(ctx, &pb.DeleteFollowerRequest{Id: int32(*input)})
	if err != nil {
		fmt.Println(err)
	}
	follower := res.Follower
	return &pb.Follower{
		Id:         follower.Id,
		CreatedAt:  follower.CreatedAt,
		LeaderId:   follower.LeaderId,
		FollowerId: follower.FollowerId,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
