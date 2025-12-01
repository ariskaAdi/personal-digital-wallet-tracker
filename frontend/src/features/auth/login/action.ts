"use server";

import { apiUrl } from "@/constant/api_url";
import { loginShema } from "@/lib/validation/auth";
import axios from "axios";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export async function loginAction(prevState: unknown, formData: FormData) {
  // GET VALUE form
  const user = {
    email: formData.get("email"),
    password: formData.get("password"),
  };

  // VALIDATION ZOD
  const parsed = loginShema.safeParse(user);
  if (!parsed.success) {
    return {
      success: false,
      message: parsed.error.issues[0].message,
    };
  }

  // CALL API
  try {
    const result = await axios.post(`${apiUrl}/auth/login`, user, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (!result.data.success) {
      return {
        success: false,
        message: result.data.message,
      };
    }

    // SAVE TOKEN TO COOKIE
    (await cookies()).set("token", result.data.data.token, {
      httpOnly: true,
      secure: true,
    });
  } catch (error) {
    console.log("LOGIN ERROR", error);

    const msg =
      // error?.response?.data?.message ||
      // error?.message ||
      "Something went wrong";

    return {
      success: false,
      message: msg,
    };
  }

  // REDIRECT TO DASHBOARD

  redirect("/dashboard");
}
