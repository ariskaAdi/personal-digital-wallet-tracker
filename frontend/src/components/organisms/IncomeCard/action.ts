"use server";

import { transactionShema } from "@/lib/validation/tx";
import { cookies } from "next/headers";

export async function incomeAction(prevState: unknown, formData: FormData) {
  // get cookie
  const cookieStore = await cookies();

  const cookieHeader = cookieStore
    .getAll()
    .map((cookie) => `${cookie.name}=${cookie.value}`)
    .join("; ");
  // get value

  const wallet_id = Number(formData.get("wallet_id"));
  const amount = Number(formData.get("amount"));
  const notes = formData.get("notes");

  // validation zod
  const parsed = transactionShema.safeParse({ amount, notes });
  if (!parsed.success) {
    return {
      success: false,
      message: parsed.error.issues[0].message,
    };
  }

  // call api

  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/tx/income`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        cookie: cookieHeader,
      },
      credentials: "include",

      body: JSON.stringify({
        wallet_id,
        amount,
        notes,
      }),
    });
    console.log(res);
  } catch (error) {
    console.log(error);
  }
}
