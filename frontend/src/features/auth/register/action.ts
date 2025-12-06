"use server";

import { apiUrl } from "@/constant/api_url";
import { registerSchema } from "@/lib/validation/auth";
import axios from "axios";
import { redirect } from "next/navigation";

export async function RegisterAction(prevState: unknown, formData: FormData) {
  // GET VALUE form
  const newUser = {
    name: formData.get("name"),
    email: formData.get("email"),
    password: formData.get("password"),
  };

  // VALIDATION ZOD
  const parsed = registerSchema.safeParse(newUser);
  if (!parsed.success) {
    return {
      success: false,
      message: parsed.error.issues[0].message,
    };
  }

  // CALL API
  try {
    const result = await axios.post(`${apiUrl}/auth/register`, newUser, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    console.log(result.data);
  } catch (error) {
    console.log(error);
    const msg = "Error RegisterHandler";
    return { success: false, message: msg };
  }

  redirect("/auth/login");
}
