import React, { FunctionComponent, useState } from "react";
import { useDispatch } from "react-redux";
import { graphqlRequest } from "../src/clients/graphqlClient";
import {
  CreateUserMutation,
  CreateUserMutationVariables,
  LoginMutation,
  LoginMutationVariables,
  useCreateUserMutation,
  useLoginMutation,
} from "../src/generated/generated";
import { login } from "../store/user-slice";
import Button from "./Button";
import Input from "./Input";

enum Type {
  button = "button",
  submit = "submit",
  reset = "reset",
}

const Login: FunctionComponent<{
  modal: Function;
}> = ({ modal }) => {
  const dispatch = useDispatch();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [avatar, setAvatar] = useState("");
  const [email, setEmail] = useState("");
  const [bio, setBio] = useState("");
  const [newUser, setNewUser] = useState(false);

  const { mutate: loginUser, data: loginData } = useLoginMutation(
    graphqlRequest(),
    {
      onSuccess: (
        data: LoginMutation,
        variables: LoginMutationVariables,
        context: unknown
      ) => {
        dispatch(
          login({
            user: {
              id: data.login.user.id,
              username: data.login.user.username,
              bio: data.login.user.bio,
              avatar: data.login.user.avatar,
            },
            token: data.login.authToken.accessToken,
          })
        );
        return console.log("Login", data);
      },
    }
  );

  const { mutate: createUser, data: createData } = useCreateUserMutation(
    graphqlRequest(),
    {
      onSuccess: (
        data: CreateUserMutation,
        varaiables: CreateUserMutationVariables,
        context: unknown
      ) => {
        dispatch(
          login({
            user: {
              id: data.createUser.user.id,
              username: data.createUser.user.username,
              bio: data.createUser.user.bio,
              avatar: data.createUser.user.avatar,
            },
            token: data.createUser.authToken.accessToken,
          })
        );
        return console.log("Uxer Created", data);
      },
    }
  );

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    if (newUser) {
      createUser({ input: { username, password, avatar, email, bio } });
    } else {
      loginUser({ input: { username, password } });
    }
    modal(false);
  };

  return (
    <div className=" flex justify-center mt-7">
      <form
        className=" flex flex-col w-5/6 lg:w-1/3 space-y-3"
        onSubmit={handleSubmit}
      >
        <Input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e: any) => setUsername(e.target.value)}
          required
        />
        <Input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e: any) => setPassword(e.target.value)}
          required
        />
        {newUser ? (
          <>
            <Input
              type="text"
              placeholder="Avatar Url"
              value={avatar}
              onChange={(e: any) => setAvatar(e.target.value)}
              required
            />
            <Input
              type="email"
              placeholder="Email"
              value={email}
              onChange={(e: any) => setEmail(e.target.value)}
              required
            />
            <Input
              type="Bio"
              placeholder="Something about you"
              value={bio}
              onChange={(e: any) => setBio(e.target.value)}
              required
            />
          </>
        ) : null}

        <Button type={Type.submit}>
          {newUser ? "Create Account" : "Login"}
        </Button>
        <p
          onClick={() => setNewUser(!newUser)}
          className=" text-gray-400 text-sm cursor-pointer"
        >
          I {newUser ? "" : "don't"} have an account
        </p>
      </form>
    </div>
  );
};

export default Login;
