import React, { useState } from "react";
import { BsFileEarmarkPlusFill, BsPersonLinesFill } from "react-icons/bs";
import { FaSearch } from "react-icons/fa";
import { AiFillCloseCircle, AiFillHome } from "react-icons/ai";
import { RiLogoutBoxFill } from "react-icons/ri";
import { useRouter } from "next/router";
// @ts-ignore
import Modal from "react-modal";
import CreatePostForm from "./CreatePostForm";
import { useGetUserByIdQuery } from "../src/generated/generated";
import { graphqlRequest } from "../src/clients/graphqlClient";
import { selectUser } from "../store";
import { useDispatch, useSelector } from "react-redux";
import { logout } from "../store/user-slice";


const Navbar = () => {
  console.log("gateway", process.env.NEXT_PUBLIC_API_URL)
  console.log("port", process.env.NEXTGATEWAY_PORT)


  const dispatch = useDispatch()
  const route = useRouter();
  const [modalOpen, setModalOpen] = useState(false);
  const currentUser = useSelector(selectUser);
  const userId = currentUser?.id
  const { data: currentUserData } = useGetUserByIdQuery(graphqlRequest(), {userId: userId!}, {enabled: !!userId});

  const logoutUser = () => {
     route.push("/")
     dispatch(logout())
  }
  return (
    <div>
      <div className=" grid grid-flow-col py-5 bg-gray-50">
        <div className=" mx-4">
          <h1 className=" text-xl">Social Media</h1>
        </div>
        <div className=" flex justify-between mx-4 ml-16">
          <div className=" cursor-pointer">
            <div onClick={() => route.push("/")}>
              <AiFillHome className=" text-2xl" />
            </div>
          </div>
          <div className=" cursor-pointer">
            <div onClick={() => setModalOpen(true)}>
              <BsFileEarmarkPlusFill className=" text-2xl" />
            </div>
          </div>
          <div className=" cursor-pointer">
            <div onClick={logoutUser}>
              <RiLogoutBoxFill className=" text-2xl" />
            </div>
          </div>
          <div
            onClick={() => route.push(`/profile/${currentUserData?.user.id}`)}
            className=" cursor-pointer"
          >
            <div>
              <BsPersonLinesFill className=" text-2xl" />
            </div>
          </div>
        </div>
      </div>
      <div>
        <Modal isOpen={modalOpen}>
        <AiFillCloseCircle
        className=" text-gray-500 text-4xl absolute top-1 right-1 cursor-pointer"
        onClick={() => setModalOpen(false)}
      />
          <CreatePostForm modal={(mod:boolean) => setModalOpen(mod)} />
        </Modal>
      </div>
    </div>
  );
};

export default Navbar;
