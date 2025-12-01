"use client";
import InputForm from "@/components/molecules/InputForm";
import InputPassword from "@/components/molecules/InputPassword";
import { Button } from "@/components/ui/button";

import Link from "next/link";
import { useActionState } from "react";
import { loginAction } from "./action";
import AuthLayout from "@/components/templates/AuthLayout";

const Login = () => {
  const [message, formAction, isPending] = useActionState(loginAction, null);
  return (
    <AuthLayout
      title="Login"
      footer={
        <>
          Don’t have an account?{" "}
          <Link
            href="/auth/register"
            className="text-blue-500 hover:underline ml-1">
            Sign up
          </Link>
        </>
      }>
      {/* ERROR MESSSAGE */}
      {message?.success === false && (
        <p className="mb-3 text-red-500 text-center">{message?.message}</p>
      )}
      <form className="space-y-4" action={formAction}>
        {/* Email */}
        <InputForm
          id="email"
          name="email"
          placeholder="Enter your email"
          type="email">
          Email
        </InputForm>

        {/* Password */}
        <InputPassword
          id="password"
          name="password"
          placeholder="Enter your password"
          type="password">
          Password
        </InputPassword>

        {/* Submit */}
        <Button type="submit" className="w-full cursor-pointer">
          {isPending ? "Loading..." : "Login"}
        </Button>
      </form>

      {/* Divider */}
      <div className="flex items-center gap-2 my-4">
        <div className="flex-1 h-px bg-gray-300"></div>
        <span className="text-sm text-gray-500">OR</span>
        <div className="flex-1 h-px bg-gray-300"></div>
      </div>

      {/* <LoginGoogle /> */}
    </AuthLayout>
  );
};

export default Login;
