import React, { FunctionComponent } from "react";

enum Type {
  button = "button",
  submit = "submit",
  reset = "reset",
}

const Button: FunctionComponent<{
  type: Type;
  onClick?: () => void
}> = ({ children, type, onClick }) => {
  return (
    <button
      className=" p-1 bg-green-600 text-white font-bold rounded-md hover:bg-green-700"
      type={type}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

export default Button;
