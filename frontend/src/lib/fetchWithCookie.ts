"use server";

import { cookies } from "next/headers";

export const fetchWithCookie = async (
  url: string,
  options: RequestInit = {}
) => {
  const cookieStore = await cookies();
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  return fetch(url, {
    ...options,
    headers: {
      ...(options.headers || {}),
      cookie: cookieHeader,
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
};
