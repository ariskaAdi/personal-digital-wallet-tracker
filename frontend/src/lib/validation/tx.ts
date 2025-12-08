import z from "zod";

export const transactionShema = z.object({
  amount: z.number().min(1000, "Amount must be at least 1000"),
  notes: z.string().min(3, "Notes must be at least 3 characters"),
});

export type Transaction = z.infer<typeof transactionShema>;
