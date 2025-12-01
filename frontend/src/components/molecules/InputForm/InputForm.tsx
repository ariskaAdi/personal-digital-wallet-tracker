import Input from "@/components/atoms/Input";
import Label from "@/components/atoms/Label";
import React from "react";

// type for input component
type InputProps = React.ComponentProps<"input">;

// type for label and input
interface InputFormProps extends React.ComponentPropsWithoutRef<"div"> {
  id: InputProps["id"];
  type: InputProps["type"];
  name?: InputProps["name"];
  placeholder?: InputProps["placeholder"];
}

const InputForm = ({
  className,
  id,
  type,
  name,
  placeholder,
  children,
  ...props
}: InputFormProps) => {
  const defaultClassName = "flex flex-col gap-2";
  const finalClassName = className ? className : defaultClassName;

  const inputProps = { id, type, name, placeholder, ...props };
  return (
    <div className={finalClassName}>
      <Label htmlFor={id}>{children}</Label>
      <Input {...inputProps} />
    </div>
  );
};

export default InputForm;
