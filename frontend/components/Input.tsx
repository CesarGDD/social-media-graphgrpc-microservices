import React, { FunctionComponent } from "react";

const Input: FunctionComponent<{
  type: string;
  placeholder?: string;
  value: string;
  onChange: any;
  required?: boolean
}> = ({ type, placeholder, value, onChange, required }) => {
  return (
    <>
      <input
        className="p-1 border-2 rounded-md w-full"
        type={type}
        placeholder={placeholder}
        value={value}
        onChange={onChange}
        required={required}
      />
    </>
  );
};

export default Input;
