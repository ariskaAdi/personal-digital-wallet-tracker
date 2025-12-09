"use server";

import { fetchWithCookie } from "@/lib/fetchWithCookie";
import { transactionShema } from "@/lib/validation/tx";

// logic api form
const transactionAction = async (
  endpoint: "/tx/income" | "/tx/expense",
  prevState: unknown,
  formData: FormData
) => {
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

  try {
    await fetchWithCookie(`${process.env.NEXT_PUBLIC_API_URL}${endpoint}`, {
      method: "POST",
      body: JSON.stringify({
        wallet_id,
        amount,
        notes,
      }),
    });

    return {
      success: true,
      message: "Transaction success",
    };
  } catch (error) {
    console.log(error);
    return {
      success: false,
      message: "Transaction failed",
    };
  }
};

export const incomeAction = transactionAction.bind(null, "/tx/income");
export const expenseAction = transactionAction.bind(null, "/tx/expense");

// call api for tanstack strategy
export const fetchWallets = async () => {
  try {
    const res = await fetchWithCookie(
      `${process.env.NEXT_PUBLIC_API_URL}/wallet/all`,
      {
        method: "GET",
      }
    );
    const data = await res.json();
    return data.data;
  } catch (error) {
    console.log(error);
  }
};
