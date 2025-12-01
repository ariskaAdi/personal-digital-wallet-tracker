"use client";
import InputForm from "@/components/molecules/InputForm";
import InputPassword from "@/components/molecules/InputPassword";
import AuthLayout from "@/components/templates/AuthLayout";
import { Button } from "@/components/ui/button";

import Link from "next/link";
import React from "react";

const Register = () => {
  const formAction = (formData: FormData) => {
    const newUser = {
      name: formData.get("name"),
      email: formData.get("email"),
      password: formData.get("password"),
    };

    console.log(newUser);
  };

  return (
    <AuthLayout
      title="Register"
      footer={
        <>
          Already have an account?{" "}
          <Link
            href="/auth/login"
            className="text-blue-500 hover:underline ml-1">
            Login
          </Link>
        </>
      }>
      <form action={formAction} className="space-y-4">
        {/* Username */}
        <InputForm
          id="name"
          name="name"
          placeholder="Enter your username"
          type="text">
          Username
        </InputForm>
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
        <Button type="submit" className="w-full">
          Register
        </Button>
      </form>
    </AuthLayout>
  );
};

export default Register;
