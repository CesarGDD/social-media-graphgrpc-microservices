import React, { useState } from "react";
import { AiFillCloseCircle } from "react-icons/ai";
// @ts-ignore
import { useQueryClient } from "react-query";
import { useSelector } from "react-redux";
import { graphqlRequest } from "../src/clients/graphqlClient";
import {
  CreateCommentMutation,
  CreateCommentMutationVariables,
  Post,
  useCommentByPostIdQuery,
  useCreateCommentMutation,
  useGetUserByIdQuery,
  Comment
} from "../src/generated/generated";
import { selectUser } from "../store";
import Button from "./Button";
import CommentComponent from "./CommentComponent";
import Input from "./Input";
import PostsCard from "./PostsCard";
type DataComment = {
  commentsByPostId: Comment[];
};
enum Type {
  button = "button",
  submit = "submit",
  reset = "reset",
}

const ModalPost: React.FC<{
  modal: Function;
  post: Post;
}> = ({ modal, post }) => {
  const queryClient = useQueryClient();
  const currentUser = useSelector(selectUser)
  const [content, setContent] = useState("");
  const { data: currentUserData } = useGetUserByIdQuery(graphqlRequest(), {
    userId: currentUser!.id,
  });
  const { data: comments } = useCommentByPostIdQuery<DataComment, Error>(
    graphqlRequest(),
    { commentsByPostIdId: post!.id }
  );
  const { mutate } = useCreateCommentMutation(graphqlRequest(), {
    onSuccess: (
      data: CreateCommentMutation,
      variables: CreateCommentMutationVariables,
      context: unknown
    ) => {
      queryClient.invalidateQueries("CommentByPostId");
      setContent("")
      return console.log("create comment");
    },
  });
  let commentLength = comments?.commentsByPostId.length;
  return (
    <div className="w-full m-auto md:w-7/12 lg:w-3/5 pb-3">
      <AiFillCloseCircle
        className=" text-gray-500 text-4xl absolute top-1 right-1 cursor-pointer"
        onClick={() => modal(false)}
      />
      <PostsCard post={post} />
      <div className=" space-x-2 flex bg-gray-50 pb-3 px-3 align-middle">
        <Input
          type="text"
          onChange={(e: any) => setContent(e.target.value)}
          value={content}
        />
        <Button
          type={Type.button}
          onClick={() =>
            mutate({
              input: {
                contents: content,
                post_id: post.id,
                user_id: currentUserData!.user.id,
              },
            })
          }
        >
          Comment
        </Button>
      </div>
      {comments?.commentsByPostId?.map((commment) => (
        <CommentComponent comment={commment} key={commment.id} />
      ))}
    </div>
  );
};

export default ModalPost;
