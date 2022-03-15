import React, { FunctionComponent } from "react";
import { BsFillTrashFill } from "react-icons/bs";
import { FaEdit } from "react-icons/fa";
import { useQueryClient } from "react-query";
import { graphqlRequest } from "../src/clients/graphqlClient";
import { DeletePostMutation, DeletePostMutationVariables, useDeletePostMutation } from "../src/generated/generated";

const EditDeletePost:FunctionComponent<{
    postId: number
}> = ({postId}) => {
    const queryClient = useQueryClient()
    const {mutate: delPost} = useDeletePostMutation<Error>(graphqlRequest(), {
        onSuccess:(data: DeletePostMutation, variables: DeletePostMutationVariables, context: unknown) => {
            queryClient.invalidateQueries("GetPosts")
            return alert('Post deleted',)
        }
    });
  return (
    <div className=" flex items-center text-xl space-x-6 mr-2">
      <BsFillTrashFill className=" cursor-pointer" onClick={()=> delPost({input:postId})} />
      <FaEdit className=" cursor-pointer" />
    </div>
  );
};

export default EditDeletePost;
