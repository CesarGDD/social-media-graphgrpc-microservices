import type { NextPage } from "next";
import PostComponent from "../components/PostComponet";
import { Post, useGetPostsQuery } from "../src/generated/generated";
import { graphqlRequest } from "../src/clients/graphqlClient";
import { useSelector } from "react-redux";
import { selectUser } from "../store";
// @ts-ignore
import Modal from "react-modal";
import { useEffect, useState } from "react";
import Login from "../components/Login";

type PostData = {
  posts: Post[];
};

const Home: NextPage = () => {
  const [modal, setModal] = useState(true);
  const { data, isLoading, error } = useGetPostsQuery<PostData, Error>(
    graphqlRequest(),
    {}
  );
  const currentUser = useSelector(selectUser);
  useEffect(() => {
    if (!currentUser) {
      setModal(true);
    }
  }, [currentUser]);
  console.log(currentUser);
  if (isLoading) return <p>Loading...</p>;
  if (error) return <p> {error.message} </p>;

  return (
    <div className=" w-full m-auto md:w-7/12 lg:w-2/5">
      {!currentUser && modal === false ? <p>Loading...</p> : null}
      {!currentUser ? (
        <Modal isOpen={modal}>
          <Login modal={(mod: boolean) => setModal(mod)} />
        </Modal>
      ) : (
        data?.posts?.map((post) => <PostComponent post={post} key={post.id} />)
      )}
    </div>
  );
};

export default Home;
