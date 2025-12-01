import z from "zod";

export const loginShema = z.object({
  email: z.string().email("Email is required"),
  password: z.string().min(5, "Password must be at least 5 characters"),
});

export type LoginSchema = z.infer<typeof loginShema>;
