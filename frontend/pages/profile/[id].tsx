import Image from "next/image";
import React, { useEffect, useState } from "react";
// @ts-ignore
import Modal from "react-modal";
import { graphqlRequest } from "../../src/clients/graphqlClient";
import {
  CreateFolloweMutation,
  CreateFolloweMutationVariables,
  DeleteFollowerMutation,
  DeleteFollowerMutationVariables,
  useCreateFolloweMutation,
  useDeleteFollowerMutation,
  useGetUserWithPostsByIdQuery,
  User,
} from "../../src/generated/generated";
import { useRouter } from "next/router";
import ModalPost from "../../components/ModalPost";
import { useQueryClient } from "react-query";
import { useSelector } from "react-redux";
import { selectUser } from "../../store";
type UserData = {
  user: User;
};

const Profile = () => {
  const [modalOpen, setModalOpen] = useState(false);
  const router = useRouter();
  const { id } = router.query;
  const queryClient = useQueryClient();
  const currentUser = useSelector(selectUser);

  const {
    data: user,
    isLoading,
    isSuccess,
  } = useGetUserWithPostsByIdQuery<UserData>(
    graphqlRequest(),
    {
      userId: Number(id),
    });

  const {
    data: currentUserData,
    isLoading: loading,
    error,
  } = useGetUserWithPostsByIdQuery(graphqlRequest(), {
    userId: currentUser?.id,
  });
  console.log(error)

  const { mutate: follow } = useCreateFolloweMutation<Error>(graphqlRequest(), {
    onSuccess: (
      data: CreateFolloweMutation,
      variables: CreateFolloweMutationVariables,
      context: unknown
    ) => {
      queryClient.invalidateQueries("GetUserWithPostsById");
      return console.log("following", data);
    },
  });
  const { mutate: unfollow } = useDeleteFollowerMutation<Error>(
    graphqlRequest(),
    {
      onSuccess: (
        data: DeleteFollowerMutation,
        variables: DeleteFollowerMutationVariables,
        context: unknown
      ) => {
        queryClient.invalidateQueries("GetUserWithPostsById");
        return console.log("following", data);
      },
    }
  );

  const follower = user?.user.followers?.find(
    (obj) => obj.follower_id === currentUserData!.user?.id
  );

  const followHandler = () => {
    if (follower) {
      unfollow({ input: follower.id });
    } else {
      follow({
        input: {
          follower_id: currentUserData!.user.id,
          leader_id: user!.user.id,
        },
      });
    }
  };

  if (isLoading) return <p>Loading....</p>;
  if (error) return <p>Permision denegated</p>;

  return (
    <div>
      <div className=" border-b-2">
        <div className=" flex justify-around border-t-2 pt-3 items-center">
          {/* image */}
          <div className="">
            {isSuccess && (
              <Image
                loader={ ()=> user!.user.avatar!}
                className=" rounded-full"
                src={user!.user.avatar!}
                width={100}
                height={100}
              />
            )}
          </div>
          {/* info */}
          <div>
            <div className=" space-y-2">
              <h1 className=" text-xl px-4"> {user?.user.username} </h1>
              {currentUserData?.user.id === user?.user.id ? (
                <button className=" border-2 px-5">Edit profile</button>
              ) : (
                <button className=" border-2 px-5" onClick={followHandler}>
                  {follower ? "Unfollow" : "Follow"}
                </button>
              )}
            </div>
          </div>
          <div className="">
            <div>
              <p>
                {user?.user.posts?.length}{" "}
                <span className=" text-gray-500">posts</span>
              </p>
              <p>
                {user?.user.followers?.length}{" "}
                <span className=" text-gray-500">followers</span>
              </p>
              <p></p>
            </div>
            <div></div>
          </div>
        </div>
        <div className=" flex mt-3 px-4 justify-center pb-3">
          <h2> {user?.user.bio} </h2>
        </div>
      </div>
      <div className="flex flex-wrap mt-4 justify-center">
        {user!.user.posts?.map((post) => (
          <div className=" px-1" key={post.id}>
            <Image
              loader={()=> post.url}
              className=" rounded-md cursor-pointer"
              src={post.url}
              width={180}
              height={180}
              onClick={() => setModalOpen(true)}
            />
            <Modal isOpen={modalOpen}>
              <ModalPost
                post={post}
                modal={(mod: boolean) => setModalOpen(mod)}
              />
            </Modal>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Profile;
