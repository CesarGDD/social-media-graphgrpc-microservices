import { GraphQLClient } from 'graphql-request';
import { RequestInit } from 'graphql-request/dist/types.dom';
import { useMutation, UseMutationOptions, useQuery, UseQueryOptions } from 'react-query';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };

function fetcher<TData, TVariables>(client: GraphQLClient, query: string, variables?: TVariables, headers?: RequestInit['headers']) {
  return async (): Promise<TData> => client.request<TData, TVariables>(query, variables, headers);
}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type AuthResponse = {
  __typename?: 'AuthResponse';
  authToken: AuthToken;
  user: User;
};

export type AuthToken = {
  __typename?: 'AuthToken';
  accessToken: Scalars['String'];
};

export type Comment = {
  __typename?: 'Comment';
  contents: Scalars['String'];
  created_at: Scalars['Int'];
  id: Scalars['Int'];
  likes?: Maybe<Array<CommentLike>>;
  post_id: Scalars['Int'];
  updated_at: Scalars['Int'];
  user: User;
  user_id: Scalars['Int'];
};

export type CommentLike = {
  __typename?: 'CommentLike';
  commentId: Scalars['Int'];
  created_at: Scalars['Int'];
  id: Scalars['Int'];
  user: User;
  user_id: Scalars['Int'];
};

export type EditComment = {
  contents: Scalars['String'];
  id: Scalars['Int'];
};

export type EditHashtag = {
  id: Scalars['Int'];
  title: Scalars['String'];
};

export type EditPost = {
  caption?: InputMaybe<Scalars['String']>;
  id: Scalars['Int'];
  url?: InputMaybe<Scalars['String']>;
};

export type EditUser = {
  avatar?: InputMaybe<Scalars['String']>;
  bio?: InputMaybe<Scalars['String']>;
  id: Scalars['Int'];
};

export type Follower = {
  __typename?: 'Follower';
  created_at: Scalars['Int'];
  follower_id: Scalars['Int'];
  id: Scalars['Int'];
  leader_id: Scalars['Int'];
  user: User;
};

export type Hashtag = {
  __typename?: 'Hashtag';
  created_at: Scalars['Int'];
  id: Scalars['Int'];
  title: Scalars['String'];
};

export type HashtagPost = {
  __typename?: 'HashtagPost';
  hashtag: Hashtag;
  hashtag_id: Scalars['Int'];
  id: Scalars['Int'];
  post_id: Scalars['Int'];
  posts?: Maybe<Array<Post>>;
};

export type LoginInput = {
  password: Scalars['String'];
  username: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createComment: Comment;
  createCommentLike: CommentLike;
  createFollower: Follower;
  createHashtag: Hashtag;
  createHashtagPost: HashtagPost;
  createPost: Post;
  createPostLike: PostLike;
  createUser: AuthResponse;
  deleteComment?: Maybe<Comment>;
  deleteCommentLikeById?: Maybe<CommentLike>;
  deleteCommentLikeByUserId?: Maybe<CommentLike>;
  deleteFollower?: Maybe<Follower>;
  deleteHashtag?: Maybe<Hashtag>;
  deleteHashtagPost?: Maybe<HashtagPost>;
  deletePost?: Maybe<Post>;
  deletePostLikeById?: Maybe<PostLike>;
  deletePostLikeByUserId?: Maybe<PostLike>;
  deleteUser?: Maybe<User>;
  login: AuthResponse;
  updateComment: Comment;
  updateHashtag: Hashtag;
  updatePost: Post;
  updateUser: User;
};


export type MutationCreateCommentArgs = {
  input?: InputMaybe<NewComment>;
};


export type MutationCreateCommentLikeArgs = {
  input?: InputMaybe<NewCommentLike>;
};


export type MutationCreateFollowerArgs = {
  input?: InputMaybe<NewFollower>;
};


export type MutationCreateHashtagArgs = {
  input?: InputMaybe<NewHashtag>;
};


export type MutationCreateHashtagPostArgs = {
  input?: InputMaybe<NewHashtagPost>;
};


export type MutationCreatePostArgs = {
  input?: InputMaybe<NewPost>;
};


export type MutationCreatePostLikeArgs = {
  input?: InputMaybe<NewPostLike>;
};


export type MutationCreateUserArgs = {
  input?: InputMaybe<NewUser>;
};


export type MutationDeleteCommentArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationDeleteCommentLikeByIdArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationDeleteCommentLikeByUserIdArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationDeleteFollowerArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationDeleteHashtagArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationDeleteHashtagPostArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationDeletePostArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationDeletePostLikeByIdArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationDeletePostLikeByUserIdArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationDeleteUserArgs = {
  input?: InputMaybe<Scalars['Int']>;
};


export type MutationLoginArgs = {
  input: LoginInput;
};


export type MutationUpdateCommentArgs = {
  input?: InputMaybe<EditComment>;
};


export type MutationUpdateHashtagArgs = {
  input?: InputMaybe<EditHashtag>;
};


export type MutationUpdatePostArgs = {
  input?: InputMaybe<EditPost>;
};


export type MutationUpdateUserArgs = {
  input?: InputMaybe<EditUser>;
};

export type NewComment = {
  contents: Scalars['String'];
  post_id: Scalars['Int'];
  user_id: Scalars['Int'];
};

export type NewCommentLike = {
  commentId: Scalars['Int'];
  user_id: Scalars['Int'];
};

export type NewFollower = {
  follower_id: Scalars['Int'];
  leader_id: Scalars['Int'];
};

export type NewHashtag = {
  title: Scalars['String'];
};

export type NewHashtagPost = {
  hashtag_id: Scalars['Int'];
  post_id: Scalars['Int'];
};

export type NewPost = {
  caption?: InputMaybe<Scalars['String']>;
  url: Scalars['String'];
  user_id: Scalars['Int'];
};

export type NewPostLike = {
  post_id: Scalars['Int'];
  user_id: Scalars['Int'];
};

export type NewUser = {
  avatar: Scalars['String'];
  bio: Scalars['String'];
  email: Scalars['String'];
  password: Scalars['String'];
  username: Scalars['String'];
};

export type Post = {
  __typename?: 'Post';
  caption?: Maybe<Scalars['String']>;
  comments?: Maybe<Array<Comment>>;
  created_at: Scalars['Int'];
  id: Scalars['Int'];
  likes?: Maybe<Array<PostLike>>;
  updated_at: Scalars['Int'];
  url: Scalars['String'];
  user: User;
  user_id: Scalars['Int'];
};

export type PostLike = {
  __typename?: 'PostLike';
  created_at: Scalars['Int'];
  id: Scalars['Int'];
  post_id: Scalars['Int'];
  user: User;
  user_id: Scalars['Int'];
};

export type Query = {
  __typename?: 'Query';
  comment: Comment;
  commentLike: CommentLike;
  commentLikeByCommentId: Array<CommentLike>;
  commentLikes: Array<CommentLike>;
  comments: Array<Comment>;
  commentsByPostId: Array<Comment>;
  follower: Follower;
  followers: Array<Follower>;
  hashtag: Hashtag;
  hashtagByTitle: Hashtag;
  hashtagPost: HashtagPost;
  hashtags: Array<Hashtag>;
  hashtagsPost: Array<HashtagPost>;
  post: Post;
  postLike: PostLike;
  postLikeByPostId: Array<PostLike>;
  postLikes: Array<PostLike>;
  posts: Array<Post>;
  user: User;
  userByUsername: User;
  users: Array<User>;
};


export type QueryCommentArgs = {
  id: Scalars['Int'];
};


export type QueryCommentLikeArgs = {
  id: Scalars['Int'];
};


export type QueryCommentLikeByCommentIdArgs = {
  comment_id: Scalars['Int'];
};


export type QueryCommentsByPostIdArgs = {
  id: Scalars['Int'];
};


export type QueryFollowerArgs = {
  id: Scalars['Int'];
};


export type QueryHashtagArgs = {
  id: Scalars['Int'];
};


export type QueryHashtagByTitleArgs = {
  title: Scalars['String'];
};


export type QueryHashtagPostArgs = {
  id: Scalars['Int'];
};


export type QueryPostArgs = {
  id: Scalars['Int'];
};


export type QueryPostLikeArgs = {
  id: Scalars['Int'];
};


export type QueryPostLikeByPostIdArgs = {
  post_id: Scalars['Int'];
};


export type QueryUserArgs = {
  id: Scalars['Int'];
};


export type QueryUserByUsernameArgs = {
  username: Scalars['String'];
};

export type User = {
  __typename?: 'User';
  avatar?: Maybe<Scalars['String']>;
  bio?: Maybe<Scalars['String']>;
  created_at?: Scalars['Int'];
  email?: Maybe<Scalars['String']>;
  followers?: Array<Follower>;
  id: Scalars['Int'];
  password?: Scalars['String'];
  posts?: Maybe<Array<Post>>;
  updated_at?: Scalars['Int'];
  username: Scalars['String'];
};

export type LoginMutationVariables = Exact<{
  input: LoginInput;
}>;


export type LoginMutation = { __typename?: 'Mutation', login: { __typename?: 'AuthResponse', authToken: { __typename?: 'AuthToken', accessToken: string }, user: { __typename?: 'User', id: number, username: string, bio?: string | null, avatar?: string | null } } };

export type CommentByPostIdQueryVariables = Exact<{
  commentsByPostIdId: Scalars['Int'];
}>;


export type CommentByPostIdQuery = { __typename?: 'Query', commentsByPostId: Array<{ __typename?: 'Comment', id: number, created_at: number, updated_at: number, contents: string, user_id: number, post_id: number, user: { __typename?: 'User', id: number, username: string, avatar?: string | null } }> };

export type CreateCommentMutationVariables = Exact<{
  input?: InputMaybe<NewComment>;
}>;


export type CreateCommentMutation = { __typename?: 'Mutation', createComment: { __typename?: 'Comment', id: number, created_at: number, updated_at: number, contents: string, user_id: number, post_id: number } };

export type UpdateCommentMutationVariables = Exact<{
  input?: InputMaybe<EditComment>;
}>;


export type UpdateCommentMutation = { __typename?: 'Mutation', updateComment: { __typename?: 'Comment', id: number, created_at: number, updated_at: number, contents: string, user_id: number, post_id: number } };

export type DeleteCommentMutationVariables = Exact<{
  input?: InputMaybe<Scalars['Int']>;
}>;


export type DeleteCommentMutation = { __typename?: 'Mutation', deleteComment?: { __typename?: 'Comment', id: number, contents: string } | null };

export type CreateFolloweMutationVariables = Exact<{
  input?: InputMaybe<NewFollower>;
}>;


export type CreateFolloweMutation = { __typename?: 'Mutation', createFollower: { __typename?: 'Follower', id: number, created_at: number, leader_id: number, follower_id: number } };

export type DeleteFollowerMutationVariables = Exact<{
  input?: InputMaybe<Scalars['Int']>;
}>;


export type DeleteFollowerMutation = { __typename?: 'Mutation', deleteFollower?: { __typename?: 'Follower', id: number, leader_id: number, follower_id: number, created_at: number } | null };

export type HashtagByTitleQueryVariables = Exact<{
  title: Scalars['String'];
}>;


export type HashtagByTitleQuery = { __typename?: 'Query', hashtagByTitle: { __typename?: 'Hashtag', id: number, created_at: number, title: string } };

export type HashTagsPostsQueryVariables = Exact<{ [key: string]: never; }>;


export type HashTagsPostsQuery = { __typename?: 'Query', hashtagsPost: Array<{ __typename?: 'HashtagPost', id: number, hashtag_id: number, post_id: number, posts?: Array<{ __typename?: 'Post', id: number, url: string, caption?: string | null }> | null }> };

export type CreateHastagMutationVariables = Exact<{
  input?: InputMaybe<NewHashtag>;
}>;


export type CreateHastagMutation = { __typename?: 'Mutation', createHashtag: { __typename?: 'Hashtag', id: number, title: string } };

export type MutationMutationVariables = Exact<{
  input?: InputMaybe<NewHashtagPost>;
}>;


export type MutationMutation = { __typename?: 'Mutation', createHashtagPost: { __typename?: 'HashtagPost', id: number, hashtag_id: number, post_id: number } };

export type PostLikeQueryVariables = Exact<{
  postId: Scalars['Int'];
}>;


export type PostLikeQuery = { __typename?: 'Query', postLikeByPostId: Array<{ __typename?: 'PostLike', id: number, user_id: number, post_id: number, user: { __typename?: 'User', username: string, avatar?: string | null } }> };

export type CommentLikeQueryVariables = Exact<{
  commentId: Scalars['Int'];
}>;


export type CommentLikeQuery = { __typename?: 'Query', commentLikeByCommentId: Array<{ __typename?: 'CommentLike', id: number, created_at: number, user_id: number, commentId: number, user: { __typename?: 'User', id: number, username: string, avatar?: string | null } }> };

export type CreatePostLikeMutationVariables = Exact<{
  input?: InputMaybe<NewPostLike>;
}>;


export type CreatePostLikeMutation = { __typename?: 'Mutation', createPostLike: { __typename?: 'PostLike', id: number, created_at: number, user_id: number, post_id: number, user: { __typename?: 'User', id: number, username: string, avatar?: string | null } } };

export type DeletePostLikeMutationVariables = Exact<{
  input?: InputMaybe<Scalars['Int']>;
}>;


export type DeletePostLikeMutation = { __typename?: 'Mutation', deletePostLikeById?: { __typename?: 'PostLike', id: number, created_at: number, user_id: number, post_id: number } | null };

export type CreateCommentLikeMutationVariables = Exact<{
  input?: InputMaybe<NewCommentLike>;
}>;


export type CreateCommentLikeMutation = { __typename?: 'Mutation', createCommentLike: { __typename?: 'CommentLike', id: number, created_at: number, user_id: number, commentId: number, user: { __typename?: 'User', id: number, username: string, avatar?: string | null } } };

export type DeleteCommentLikeMutationVariables = Exact<{
  input?: InputMaybe<Scalars['Int']>;
}>;


export type DeleteCommentLikeMutation = { __typename?: 'Mutation', deleteCommentLikeById?: { __typename?: 'CommentLike', id: number, created_at: number, user_id: number, commentId: number, user: { __typename?: 'User', id: number, username: string, avatar?: string | null } } | null };

export type GetPostsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetPostsQuery = { __typename?: 'Query', posts: Array<{ __typename?: 'Post', id: number, created_at: number, updated_at: number, url: string, caption?: string | null, user_id: number, user: { __typename?: 'User', id: number, username: string, avatar?: string | null } }> };

export type GetPostByIdQueryVariables = Exact<{
  postId: Scalars['Int'];
}>;


export type GetPostByIdQuery = { __typename?: 'Query', post: { __typename?: 'Post', id: number, created_at: number, url: string, caption?: string | null, updated_at: number, user_id: number, user: { __typename?: 'User', id: number, username: string, avatar?: string | null } } };

export type CreatePostMutationVariables = Exact<{
  input?: InputMaybe<NewPost>;
}>;


export type CreatePostMutation = { __typename?: 'Mutation', createPost: { __typename?: 'Post', id: number, created_at: number, updated_at: number, url: string, caption?: string | null, user_id: number } };

export type UpdatePostMutationVariables = Exact<{
  input?: InputMaybe<EditPost>;
}>;


export type UpdatePostMutation = { __typename?: 'Mutation', updatePost: { __typename?: 'Post', id: number, created_at: number, updated_at: number, url: string, caption?: string | null, user_id: number } };

export type DeletePostMutationVariables = Exact<{
  input?: InputMaybe<Scalars['Int']>;
}>;


export type DeletePostMutation = { __typename?: 'Mutation', deletePost?: { __typename?: 'Post', id: number, url: string, caption?: string | null } | null };

export type GetUserByIdQueryVariables = Exact<{
  userId: Scalars['Int'];
}>;


export type GetUserByIdQuery = { __typename?: 'Query', user: { __typename?: 'User', id: number, created_at: number, updated_at: number, username: string, bio?: string | null, avatar?: string | null, email?: string | null } };

export type GetUserByUsernameQueryVariables = Exact<{
  username: Scalars['String'];
}>;


export type GetUserByUsernameQuery = { __typename?: 'Query', userByUsername: { __typename?: 'User', id: number, created_at: number, updated_at: number, username: string, bio?: string | null, avatar?: string | null, email?: string | null } };

export type GetUserWithPostsByIdQueryVariables = Exact<{
  userId?: Scalars['Int'];
}>;


export type GetUserWithPostsByIdQuery = { __typename?: 'Query', user: { __typename?: 'User', id: number, created_at: number, updated_at: number, username: string, bio?: string | null, avatar?: string | null, email?: string | null, posts?: Array<{ __typename?: 'Post', id: number, created_at: number, updated_at: number, url: string, caption?: string | null, user_id: number, user: { __typename?: 'User', id: number, username: string, avatar?: string | null } }> | null, followers: Array<{ __typename?: 'Follower', id: number, follower_id: number, user: { __typename?: 'User', username: string, avatar?: string | null, id: number } }> } };

export type CreateUserMutationVariables = Exact<{
  input?: InputMaybe<NewUser>;
}>;


export type CreateUserMutation = { __typename?: 'Mutation', createUser: { __typename?: 'AuthResponse', authToken: { __typename?: 'AuthToken', accessToken: string }, user: { __typename?: 'User', id: number, username: string, bio?: string | null, avatar?: string | null } } };

export type UpdateUserMutationVariables = Exact<{
  input?: InputMaybe<EditUser>;
}>;


export type UpdateUserMutation = { __typename?: 'Mutation', updateUser: { __typename?: 'User', id: number, created_at: number, updated_at: number, username: string, bio?: string | null, avatar?: string | null, email?: string | null } };

export type DeleteUserMutationVariables = Exact<{
  input?: InputMaybe<Scalars['Int']>;
}>;


export type DeleteUserMutation = { __typename?: 'Mutation', deleteUser?: { __typename?: 'User', id: number, username: string, avatar?: string | null } | null };


export const LoginDocument = `
    mutation Login($input: LoginInput!) {
  login(input: $input) {
    authToken {
      accessToken
    }
    user {
      id
      username
      bio
      avatar
    }
  }
}
    `;
export const useLoginMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<LoginMutation, TError, LoginMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<LoginMutation, TError, LoginMutationVariables, TContext>(
      ['Login'],
      (variables?: LoginMutationVariables) => fetcher<LoginMutation, LoginMutationVariables>(client, LoginDocument, variables, headers)(),
      options
    );
export const CommentByPostIdDocument = `
    query CommentByPostId($commentsByPostIdId: Int!) {
  commentsByPostId(id: $commentsByPostIdId) {
    id
    created_at
    updated_at
    contents
    user_id
    post_id
    user {
      id
      username
      avatar
    }
  }
}
    `;
export const useCommentByPostIdQuery = <
      TData = CommentByPostIdQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables: CommentByPostIdQueryVariables,
      options?: UseQueryOptions<CommentByPostIdQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<CommentByPostIdQuery, TError, TData>(
      ['CommentByPostId', variables],
      fetcher<CommentByPostIdQuery, CommentByPostIdQueryVariables>(client, CommentByPostIdDocument, variables, headers),
      options
    );
export const CreateCommentDocument = `
    mutation CreateComment($input: NewComment) {
  createComment(input: $input) {
    id
    created_at
    updated_at
    contents
    user_id
    post_id
  }
}
    `;
export const useCreateCommentMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<CreateCommentMutation, TError, CreateCommentMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<CreateCommentMutation, TError, CreateCommentMutationVariables, TContext>(
      ['CreateComment'],
      (variables?: CreateCommentMutationVariables) => fetcher<CreateCommentMutation, CreateCommentMutationVariables>(client, CreateCommentDocument, variables, headers)(),
      options
    );
export const UpdateCommentDocument = `
    mutation UpdateComment($input: EditComment) {
  updateComment(input: $input) {
    id
    created_at
    updated_at
    contents
    user_id
    post_id
  }
}
    `;
export const useUpdateCommentMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<UpdateCommentMutation, TError, UpdateCommentMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<UpdateCommentMutation, TError, UpdateCommentMutationVariables, TContext>(
      ['UpdateComment'],
      (variables?: UpdateCommentMutationVariables) => fetcher<UpdateCommentMutation, UpdateCommentMutationVariables>(client, UpdateCommentDocument, variables, headers)(),
      options
    );
export const DeleteCommentDocument = `
    mutation DeleteComment($input: Int) {
  deleteComment(input: $input) {
    id
    contents
  }
}
    `;
export const useDeleteCommentMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<DeleteCommentMutation, TError, DeleteCommentMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<DeleteCommentMutation, TError, DeleteCommentMutationVariables, TContext>(
      ['DeleteComment'],
      (variables?: DeleteCommentMutationVariables) => fetcher<DeleteCommentMutation, DeleteCommentMutationVariables>(client, DeleteCommentDocument, variables, headers)(),
      options
    );
export const CreateFolloweDocument = `
    mutation CreateFollowe($input: NewFollower) {
  createFollower(input: $input) {
    id
    created_at
    leader_id
    follower_id
  }
}
    `;
export const useCreateFolloweMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<CreateFolloweMutation, TError, CreateFolloweMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<CreateFolloweMutation, TError, CreateFolloweMutationVariables, TContext>(
      ['CreateFollowe'],
      (variables?: CreateFolloweMutationVariables) => fetcher<CreateFolloweMutation, CreateFolloweMutationVariables>(client, CreateFolloweDocument, variables, headers)(),
      options
    );
export const DeleteFollowerDocument = `
    mutation DeleteFollower($input: Int) {
  deleteFollower(input: $input) {
    id
    leader_id
    follower_id
    created_at
  }
}
    `;
export const useDeleteFollowerMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<DeleteFollowerMutation, TError, DeleteFollowerMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<DeleteFollowerMutation, TError, DeleteFollowerMutationVariables, TContext>(
      ['DeleteFollower'],
      (variables?: DeleteFollowerMutationVariables) => fetcher<DeleteFollowerMutation, DeleteFollowerMutationVariables>(client, DeleteFollowerDocument, variables, headers)(),
      options
    );
export const HashtagByTitleDocument = `
    query HashtagByTitle($title: String!) {
  hashtagByTitle(title: $title) {
    id
    created_at
    title
  }
}
    `;
export const useHashtagByTitleQuery = <
      TData = HashtagByTitleQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables: HashtagByTitleQueryVariables,
      options?: UseQueryOptions<HashtagByTitleQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<HashtagByTitleQuery, TError, TData>(
      ['HashtagByTitle', variables],
      fetcher<HashtagByTitleQuery, HashtagByTitleQueryVariables>(client, HashtagByTitleDocument, variables, headers),
      options
    );
export const HashTagsPostsDocument = `
    query HashTagsPosts {
  hashtagsPost {
    id
    hashtag_id
    post_id
    posts {
      id
      url
      caption
    }
  }
}
    `;
export const useHashTagsPostsQuery = <
      TData = HashTagsPostsQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables?: HashTagsPostsQueryVariables,
      options?: UseQueryOptions<HashTagsPostsQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<HashTagsPostsQuery, TError, TData>(
      variables === undefined ? ['HashTagsPosts'] : ['HashTagsPosts', variables],
      fetcher<HashTagsPostsQuery, HashTagsPostsQueryVariables>(client, HashTagsPostsDocument, variables, headers),
      options
    );
export const CreateHastagDocument = `
    mutation CreateHastag($input: NewHashtag) {
  createHashtag(input: $input) {
    id
    title
  }
}
    `;
export const useCreateHastagMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<CreateHastagMutation, TError, CreateHastagMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<CreateHastagMutation, TError, CreateHastagMutationVariables, TContext>(
      ['CreateHastag'],
      (variables?: CreateHastagMutationVariables) => fetcher<CreateHastagMutation, CreateHastagMutationVariables>(client, CreateHastagDocument, variables, headers)(),
      options
    );
export const MutationDocument = `
    mutation Mutation($input: NewHashtagPost) {
  createHashtagPost(input: $input) {
    id
    hashtag_id
    post_id
  }
}
    `;
export const useMutationMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<MutationMutation, TError, MutationMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<MutationMutation, TError, MutationMutationVariables, TContext>(
      ['Mutation'],
      (variables?: MutationMutationVariables) => fetcher<MutationMutation, MutationMutationVariables>(client, MutationDocument, variables, headers)(),
      options
    );
export const PostLikeDocument = `
    query PostLike($postId: Int!) {
  postLikeByPostId(post_id: $postId) {
    id
    user_id
    post_id
    user {
      username
      avatar
    }
  }
}
    `;
export const usePostLikeQuery = <
      TData = PostLikeQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables: PostLikeQueryVariables,
      options?: UseQueryOptions<PostLikeQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<PostLikeQuery, TError, TData>(
      ['PostLike', variables],
      fetcher<PostLikeQuery, PostLikeQueryVariables>(client, PostLikeDocument, variables, headers),
      options
    );
export const CommentLikeDocument = `
    query CommentLike($commentId: Int!) {
  commentLikeByCommentId(comment_id: $commentId) {
    id
    created_at
    user_id
    commentId
    user {
      id
      username
      avatar
    }
  }
}
    `;
export const useCommentLikeQuery = <
      TData = CommentLikeQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables: CommentLikeQueryVariables,
      options?: UseQueryOptions<CommentLikeQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<CommentLikeQuery, TError, TData>(
      ['CommentLike', variables],
      fetcher<CommentLikeQuery, CommentLikeQueryVariables>(client, CommentLikeDocument, variables, headers),
      options
    );
export const CreatePostLikeDocument = `
    mutation CreatePostLike($input: NewPostLike) {
  createPostLike(input: $input) {
    id
    created_at
    user_id
    post_id
    user {
      id
      username
      avatar
    }
  }
}
    `;
export const useCreatePostLikeMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<CreatePostLikeMutation, TError, CreatePostLikeMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<CreatePostLikeMutation, TError, CreatePostLikeMutationVariables, TContext>(
      ['CreatePostLike'],
      (variables?: CreatePostLikeMutationVariables) => fetcher<CreatePostLikeMutation, CreatePostLikeMutationVariables>(client, CreatePostLikeDocument, variables, headers)(),
      options
    );
export const DeletePostLikeDocument = `
    mutation DeletePostLike($input: Int) {
  deletePostLikeById(input: $input) {
    id
    created_at
    user_id
    post_id
  }
}
    `;
export const useDeletePostLikeMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<DeletePostLikeMutation, TError, DeletePostLikeMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<DeletePostLikeMutation, TError, DeletePostLikeMutationVariables, TContext>(
      ['DeletePostLike'],
      (variables?: DeletePostLikeMutationVariables) => fetcher<DeletePostLikeMutation, DeletePostLikeMutationVariables>(client, DeletePostLikeDocument, variables, headers)(),
      options
    );
export const CreateCommentLikeDocument = `
    mutation CreateCommentLike($input: NewCommentLike) {
  createCommentLike(input: $input) {
    id
    created_at
    user_id
    commentId
    user {
      id
      username
      avatar
    }
  }
}
    `;
export const useCreateCommentLikeMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<CreateCommentLikeMutation, TError, CreateCommentLikeMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<CreateCommentLikeMutation, TError, CreateCommentLikeMutationVariables, TContext>(
      ['CreateCommentLike'],
      (variables?: CreateCommentLikeMutationVariables) => fetcher<CreateCommentLikeMutation, CreateCommentLikeMutationVariables>(client, CreateCommentLikeDocument, variables, headers)(),
      options
    );
export const DeleteCommentLikeDocument = `
    mutation DeleteCommentLike($input: Int) {
  deleteCommentLikeById(input: $input) {
    id
    created_at
    user_id
    commentId
    user {
      id
      username
      avatar
    }
  }
}
    `;
export const useDeleteCommentLikeMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<DeleteCommentLikeMutation, TError, DeleteCommentLikeMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<DeleteCommentLikeMutation, TError, DeleteCommentLikeMutationVariables, TContext>(
      ['DeleteCommentLike'],
      (variables?: DeleteCommentLikeMutationVariables) => fetcher<DeleteCommentLikeMutation, DeleteCommentLikeMutationVariables>(client, DeleteCommentLikeDocument, variables, headers)(),
      options
    );
export const GetPostsDocument = `
    query GetPosts {
  posts {
    id
    created_at
    updated_at
    url
    caption
    user_id
    user {
      id
      username
      avatar
    }
  }
}
    `;
export const useGetPostsQuery = <
      TData = GetPostsQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables?: GetPostsQueryVariables,
      options?: UseQueryOptions<GetPostsQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<GetPostsQuery, TError, TData>(
      variables === undefined ? ['GetPosts'] : ['GetPosts', variables],
      fetcher<GetPostsQuery, GetPostsQueryVariables>(client, GetPostsDocument, variables, headers),
      options
    );
export const GetPostByIdDocument = `
    query GetPostById($postId: Int!) {
  post(id: $postId) {
    id
    created_at
    url
    caption
    updated_at
    user_id
    user {
      id
      username
      avatar
    }
  }
}
    `;
export const useGetPostByIdQuery = <
      TData = GetPostByIdQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables: GetPostByIdQueryVariables,
      options?: UseQueryOptions<GetPostByIdQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<GetPostByIdQuery, TError, TData>(
      ['GetPostById', variables],
      fetcher<GetPostByIdQuery, GetPostByIdQueryVariables>(client, GetPostByIdDocument, variables, headers),
      options
    );
export const CreatePostDocument = `
    mutation CreatePost($input: NewPost) {
  createPost(input: $input) {
    id
    created_at
    updated_at
    url
    caption
    user_id
  }
}
    `;
export const useCreatePostMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<CreatePostMutation, TError, CreatePostMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<CreatePostMutation, TError, CreatePostMutationVariables, TContext>(
      ['CreatePost'],
      (variables?: CreatePostMutationVariables) => fetcher<CreatePostMutation, CreatePostMutationVariables>(client, CreatePostDocument, variables, headers)(),
      options
    );
export const UpdatePostDocument = `
    mutation UpdatePost($input: EditPost) {
  updatePost(input: $input) {
    id
    created_at
    updated_at
    url
    caption
    user_id
  }
}
    `;
export const useUpdatePostMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<UpdatePostMutation, TError, UpdatePostMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<UpdatePostMutation, TError, UpdatePostMutationVariables, TContext>(
      ['UpdatePost'],
      (variables?: UpdatePostMutationVariables) => fetcher<UpdatePostMutation, UpdatePostMutationVariables>(client, UpdatePostDocument, variables, headers)(),
      options
    );
export const DeletePostDocument = `
    mutation DeletePost($input: Int) {
  deletePost(input: $input) {
    id
    url
    caption
  }
}
    `;
export const useDeletePostMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<DeletePostMutation, TError, DeletePostMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<DeletePostMutation, TError, DeletePostMutationVariables, TContext>(
      ['DeletePost'],
      (variables?: DeletePostMutationVariables) => fetcher<DeletePostMutation, DeletePostMutationVariables>(client, DeletePostDocument, variables, headers)(),
      options
    );
export const GetUserByIdDocument = `
    query GetUserById($userId: Int!) {
  user(id: $userId) {
    id
    created_at
    updated_at
    username
    bio
    avatar
    email
  }
}
    `;
export const useGetUserByIdQuery = <
      TData = GetUserByIdQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables: GetUserByIdQueryVariables,
      options?: UseQueryOptions<GetUserByIdQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<GetUserByIdQuery, TError, TData>(
      ['GetUserById', variables],
      fetcher<GetUserByIdQuery, GetUserByIdQueryVariables>(client, GetUserByIdDocument, variables, headers),
      options
    );
export const GetUserByUsernameDocument = `
    query GetUserByUsername($username: String!) {
  userByUsername(username: $username) {
    id
    created_at
    updated_at
    username
    bio
    avatar
    email
  }
}
    `;
export const useGetUserByUsernameQuery = <
      TData = GetUserByUsernameQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables: GetUserByUsernameQueryVariables,
      options?: UseQueryOptions<GetUserByUsernameQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<GetUserByUsernameQuery, TError, TData>(
      ['GetUserByUsername', variables],
      fetcher<GetUserByUsernameQuery, GetUserByUsernameQueryVariables>(client, GetUserByUsernameDocument, variables, headers),
      options
    );
export const GetUserWithPostsByIdDocument = `
    query GetUserWithPostsById($userId: Int!) {
  user(id: $userId) {
    id
    created_at
    updated_at
    username
    bio
    avatar
    email
    posts {
      id
      created_at
      updated_at
      url
      caption
      user_id
      user {
        id
        username
        avatar
      }
    }
    followers {
      id
      follower_id
      user {
        username
        avatar
        id
      }
    }
  }
}
    `;
export const useGetUserWithPostsByIdQuery = <
      TData = GetUserWithPostsByIdQuery,
      TError = unknown
    >(
      client: GraphQLClient,
      variables: GetUserWithPostsByIdQueryVariables,
      options?: UseQueryOptions<GetUserWithPostsByIdQuery, TError, TData>,
      headers?: RequestInit['headers']
    ) =>
    useQuery<GetUserWithPostsByIdQuery, TError, TData>(
      ['GetUserWithPostsById', variables],
      fetcher<GetUserWithPostsByIdQuery, GetUserWithPostsByIdQueryVariables>(client, GetUserWithPostsByIdDocument, variables, headers),
      options
    );
export const CreateUserDocument = `
    mutation CreateUser($input: NewUser) {
  createUser(input: $input) {
    authToken {
      accessToken
    }
    user {
      id
      username
      bio
      avatar
    }
  }
}
    `;
export const useCreateUserMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<CreateUserMutation, TError, CreateUserMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<CreateUserMutation, TError, CreateUserMutationVariables, TContext>(
      ['CreateUser'],
      (variables?: CreateUserMutationVariables) => fetcher<CreateUserMutation, CreateUserMutationVariables>(client, CreateUserDocument, variables, headers)(),
      options
    );
export const UpdateUserDocument = `
    mutation UpdateUser($input: EditUser) {
  updateUser(input: $input) {
    id
    created_at
    updated_at
    username
    bio
    avatar
    email
  }
}
    `;
export const useUpdateUserMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<UpdateUserMutation, TError, UpdateUserMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<UpdateUserMutation, TError, UpdateUserMutationVariables, TContext>(
      ['UpdateUser'],
      (variables?: UpdateUserMutationVariables) => fetcher<UpdateUserMutation, UpdateUserMutationVariables>(client, UpdateUserDocument, variables, headers)(),
      options
    );
export const DeleteUserDocument = `
    mutation DeleteUser($input: Int) {
  deleteUser(input: $input) {
    id
    username
    avatar
  }
}
    `;
export const useDeleteUserMutation = <
      TError = unknown,
      TContext = unknown
    >(
      client: GraphQLClient,
      options?: UseMutationOptions<DeleteUserMutation, TError, DeleteUserMutationVariables, TContext>,
      headers?: RequestInit['headers']
    ) =>
    useMutation<DeleteUserMutation, TError, DeleteUserMutationVariables, TContext>(
      ['DeleteUser'],
      (variables?: DeleteUserMutationVariables) => fetcher<DeleteUserMutation, DeleteUserMutationVariables>(client, DeleteUserDocument, variables, headers)(),
      options
    );