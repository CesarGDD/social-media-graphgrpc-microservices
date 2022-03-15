import Image from "next/image";
import React, { FunctionComponent } from "react";
import { BsEmojiHeartEyes, BsFillTrashFill } from "react-icons/bs";
import { useQueryClient } from "react-query";
import { useSelector } from "react-redux";
import { graphqlRequest } from "../src/clients/graphqlClient";
import {
  Comment,
  CreateCommentLikeMutation,
  CreateCommentLikeMutationVariables,
  DeleteCommentLikeMutation,
  DeleteCommentMutation,
  DeleteCommentMutationVariables,
  DeletePostLikeMutationVariables,
  useCommentLikeQuery,
  useCreateCommentLikeMutation,
  useDeleteCommentLikeMutation,
  useDeleteCommentMutation,
  useGetUserByIdQuery,
} from "../src/generated/generated";
import { selectUser } from "../store";

const CommentComponent: FunctionComponent<{
  comment: Comment;
}> = ({ comment: { contents, id, post_id, user_id, user } }) => {
  const queryClient = useQueryClient();
  const currentUser = useSelector(selectUser)
  const { data: currentUserData } = useGetUserByIdQuery(graphqlRequest(), {
    userId: currentUser!.id,
  });
  const { mutate: addLike } = useCreateCommentLikeMutation<Error>(
    graphqlRequest(),
    {
      onSuccess: (
        data: CreateCommentLikeMutation,
        variables: CreateCommentLikeMutationVariables,
        context: unknown
      ) => {
        queryClient.invalidateQueries("CommentLike");
        return console.log("create like", data);
      },
    }
  );
  const { mutate: delLike } = useDeleteCommentLikeMutation<Error>(
    graphqlRequest(),
    {
      onSuccess: (
        data: DeleteCommentLikeMutation,
        variables: DeletePostLikeMutationVariables,
        context: unknown
      ) => {
        queryClient.invalidateQueries("CommentLike");
        console.log("delete like", data);
      },
    }
  );
  const { mutate: delComment } = useDeleteCommentMutation(graphqlRequest(), {
    onSuccess: (
      data: DeleteCommentMutation,
      variables: DeleteCommentMutationVariables,
      context: unknown
    ) => {
      queryClient.invalidateQueries("CommentByPostId");
      return console.log("delete comment", data);
    },
  });
  const { data: comLikes } = useCommentLikeQuery(graphqlRequest(), {
    commentId: id,
  });

  const like = comLikes?.commentLikeByCommentId?.find(
    (obj) => obj.user_id === currentUserData?.user.id
  );

  const likeHandler = () => {
    if (like) {
      delLike({ input: like.id });
      console.log("deleteLike");
    } else {
      addLike({ input: { commentId: id, user_id: currentUserData!.user.id } });
      console.log("addLike");
    }
  };

  return (
    <div className=" flex justify-between px-4 pb-1 bg-gray-50">
      <div className="flex ">
        <Image
          className=" rounded-full"
          src={user.avatar!}
          width={25}
          height={25}
        />
        <p className=" ml-1 text-gray-600">
          <span className=" text-black font-bold mr-1">{user.username}</span>
          {contents}
        </p>
      </div>
      <div className=" flex  space-x-4">
        {currentUserData!.user.id === user_id ? (
          <BsFillTrashFill
            className=" cursor-pointer text-xl font-bold"
            onClick={() => delComment({ input: id })}
          />
        ) : null}
        <BsEmojiHeartEyes
          onClick={likeHandler}
          className={` text-xl ${
            // @ts-ignore
            like ? "text-red-500 font-bold" : null
          }  cursor-pointer`}
        />
      </div>
    </div>
  );
};

export default CommentComponent;
