import React, { FunctionComponent, useState } from "react";
import PostsCard from "./PostsCard";
// @ts-ignore
import Modal from "react-modal";
import {
  Comment,
  CreateCommentMutation,
  CreateCommentMutationVariables,
  Post,
  useCommentByPostIdQuery,
  useCreateCommentMutation,
  useGetUserByIdQuery,
} from "../src/generated/generated";
import { graphqlRequest } from "../src/clients/graphqlClient";
import { useQueryClient } from "react-query";
import ModalPost from "./ModalPost";

Modal.setAppElement("body");

enum Type {
  button = "button",
  submit = "submit",
  reset = "reset",
}

type DataComment = {
  commentsByPostId: Comment[];
};

const PostComponent: FunctionComponent<{ post: Post }> = ({ post }) => {
  const queryClient = useQueryClient()
  const [modal, setModal] = useState(false);
  const { data: comments } = useCommentByPostIdQuery<DataComment, Error>(
    graphqlRequest(),
    { commentsByPostIdId: post.id }
  );
  const { mutate } = useCreateCommentMutation(graphqlRequest(), {
    onSuccess: (
      data: CreateCommentMutation,
      variables: CreateCommentMutationVariables,
      context: unknown
    ) => {
      queryClient.invalidateQueries('CommentByPostId')
      return console.log("create comment");
    },
  });
  let commentLength = comments?.commentsByPostId.length;


  return (
    <div className=" bg-gray-50 pb-4">
      <PostsCard post={post} modalOpen={()=> setModal(true)} />
      <div className="mx-4 text-sm mt-2">
        <p
          onClick={() => setModal(!modal)}
          className=" text-gray-500 cursor-pointer"
        >
          {commentLength} Comments
        </p>
      </div>
      <Modal isOpen={modal}>
       <ModalPost modal={(mod: boolean) => setModal(mod)} post={post} />
      </Modal>
    </div>
  );
};

export default PostComponent;
