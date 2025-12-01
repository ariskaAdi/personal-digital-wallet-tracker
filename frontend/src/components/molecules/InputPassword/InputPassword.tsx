"use client";
import Input from "@/components/atoms/Input";
import Label from "@/components/atoms/Label";
import { Eye, EyeOff } from "lucide-react";
import React, { useState } from "react";

// type for input
type InputProps = React.ComponentProps<"input">;

//type for label and input
interface InputPasswordProps extends React.ComponentPropsWithoutRef<"div"> {
  id: InputProps["id"];
  type: InputProps["type"];
  name?: InputProps["name"];
  placeholder?: InputProps["placeholder"];
}

const InputPassword = ({
  className,
  id,
  type,
  name,
  placeholder,
  children,
  ...props
}: InputPasswordProps) => {
  const [showPassword, setShowPassword] = useState(false);
  const defaultClassName = "flex-col";
  const finalClassName = className ? className : defaultClassName;

  const inputProps = { id, type, name, placeholder, ...props };
  return (
    <div className={finalClassName}>
      <Label htmlFor={id}>{children}</Label>
      <div className="relative">
        <Input {...inputProps} type={showPassword ? "text" : "password"} />
        <button
          type="button"
          className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500"
          onClick={() => setShowPassword((prev) => !prev)}>
          {showPassword ? (
            <EyeOff className="w-5 h-5" />
          ) : (
            <Eye className="w-5 h-5" />
          )}
        </button>
      </div>
    </div>
  );
};

export default InputPassword;
