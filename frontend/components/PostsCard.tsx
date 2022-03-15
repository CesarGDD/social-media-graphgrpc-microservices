import Image from "next/image";
import React, { FunctionComponent } from "react";
import { AiOutlineComment } from "react-icons/ai";
import { BsEmojiHeartEyes } from "react-icons/bs";
import {
  CreatePostLikeMutation,
  CreatePostLikeMutationVariables,
  DeletePostLikeMutation,
  DeletePostLikeMutationVariables,
  Post,
  PostLike,
  useCreatePostLikeMutation,
  useDeletePostLikeMutation,
  useGetUserByIdQuery,
  usePostLikeQuery,
} from "../src/generated/generated";
import { graphqlRequest } from "../src/clients/graphqlClient";
import { useQueryClient } from "react-query";
import EditDeletePost from "./EditDeletePost";
import { useRouter } from "next/router";
import { useSelector } from "react-redux";
import { selectUser } from "../store";


const PostsCard: FunctionComponent<{
  post: Post;
  modalOpen?: () => void;
}> = ({
  post: { id, created_at, updated_at, url, user, user_id, caption },
  modalOpen,
}) => {
  const currentUser = useSelector(selectUser)
  const queryClient = useQueryClient();
  const { data: currentUserData } = useGetUserByIdQuery(graphqlRequest(), {
    userId: currentUser!.id,
  });

  const route = useRouter();

  const { data: postLike } = usePostLikeQuery(graphqlRequest(), { postId: id });

  const { mutate: addLike } = useCreatePostLikeMutation<Error>(graphqlRequest(), {
    onSuccess: (
      data: CreatePostLikeMutation,
      variables: CreatePostLikeMutationVariables,
      context: unknown
    ) => {
      queryClient.invalidateQueries<PostLike>("PostLike");
      return console.log("create like", data);
    },
  });

  const { mutate: deleteLike } = useDeletePostLikeMutation<Error>(
    graphqlRequest(),
    {
      onSuccess: (
        data: DeletePostLikeMutation,
        variables: DeletePostLikeMutationVariables,
        context: unknown
      ) => {
        queryClient.invalidateQueries<PostLike>("PostLike");
        console.log("delete like", data);
      },
    }
  );

  const like = postLike?.postLikeByPostId.find(
    (obj) => obj.user_id === currentUserData?.user?.id
  );

  const likeHandler = () => {
    if (like) {
      deleteLike({ input: like.id });
    } else {
      addLike({ input: { post_id: id, user_id: currentUserData!.user.id } });
    }
  };

  return (
    <div className=" mt-4 py-3 bg-gray-50 col-span-12 ">
      {/* User Info */}
      <div className=" flex justify-between ml-2">
        <div className=" flex items-center">
          <Image
            onClick={() => route.push(`/profile/${user_id}`)}
            className=" rounded-full border-5 border-gray-500 cursor-pointer"
            src={user.avatar!}
            width={50}
            height={50}
          />
          <p className=" ml-2"> {user.username} </p>
        </div>
        {currentUserData?.user.id === user_id ? (
          <EditDeletePost postId={id} />
        ) : null}
      </div>
      {/* Post */}
      <div className="flex flex-col mt-4 border-t">
        <Image src={url} width={600} height={750} />
        <div className=" flex mx-2 items-center p-2 space-x-6">
          {/* Like Create comment */}
          <div onClick={likeHandler}>
            <BsEmojiHeartEyes
              className={` text-3xl ${
                // @ts-ignore
                like ? "text-red-500 font-bold" : null
              }  cursor-pointer`}
            />
          </div>
          {/* <BsEmojiHeartEyesFill /> */}
          <AiOutlineComment
            className=" text-4xl cursor-pointer"
            onClick={modalOpen}
          />
        </div>
        <div className=" mx-4 space-y-2">
          <div>
            <p className=" text-sm cursor-pointer">
              Liked By SomePeople and ...
            </p>
          </div>
          <div>
            <p> {caption} </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default PostsCard;
