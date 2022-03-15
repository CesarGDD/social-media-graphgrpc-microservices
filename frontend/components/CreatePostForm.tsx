import React, { FunctionComponent, useState } from "react";
import { useQueryClient } from "react-query";
import { useSelector } from "react-redux";
import { graphqlRequest } from "../src/clients/graphqlClient";
import { CreatePostMutation, CreatePostMutationVariables, Post, useCreatePostMutation, useGetUserByIdQuery } from "../src/generated/generated";
import { selectUser } from "../store";
import Button from "./Button";
import Input from "./Input";

enum Type {
  button = "button",
  submit = "submit",
  reset = "reset",
}

const CreatePostForm:FunctionComponent<{
  modal: Function
}> = ({modal}) => {
  const currentUser = useSelector(selectUser)
  const {data: currentUserData} = useGetUserByIdQuery(graphqlRequest(), {userId: currentUser!.id})
  const queryClient = useQueryClient()
  const [imageUrl, setImageUrl] = useState("");
  const [caption, setCaption] = useState("");
  const {mutate} = useCreatePostMutation(graphqlRequest(), {onSuccess: (data:CreatePostMutation, variables: CreatePostMutationVariables, context: unknown) => {
    queryClient.invalidateQueries<Post[]>('GetPosts')
    return console.log('create post', data)
  }})

  const handleSubmit = (event:React.FormEvent) => {
    event.preventDefault()
    mutate({input: {url: imageUrl, caption: caption, user_id: currentUserData!.user.id}})
    modal(false)
  }

  return (
    <div className=" flex justify-center">
      <form className=" flex flex-col w-5/6 lg:w-1/3 space-y-3" onSubmit={handleSubmit}>
      <h1 className="" >New Post!</h1>
        <Input
          type="text"
          placeholder="Image Url"
          value={imageUrl}
          onChange={(e:any) => setImageUrl(e.target.value)}
          required
        />
        <Input
          type="text"
          placeholder="Caption"
          value={caption}
          onChange={(e:any) => setCaption(e.target.value)}
          required
        />
        <Button type={Type.submit}>Create</Button>
      </form>
    </div>
  );
};

export default CreatePostForm;
