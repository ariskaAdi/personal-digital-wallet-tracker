"use client";

import CardComponent from "@/components/molecules/CardComponent";
import InputForm from "@/components/molecules/InputForm";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import React, { useActionState, useEffect, useState } from "react";
import {
  fetchWallets,
  incomeAction,
} from "../../../features/dashboard/transactions/action";
import { useQuery } from "@tanstack/react-query";
import { toast } from "sonner";

type IData = {
  id: number;
  name: string;
  balance: number;
};

const IncomeCardTanstack = ({
  initialWallets = [],
}: {
  initialWallets?: IData[];
}) => {
  // store data in state
  const [selectedWalletId, setSelectedWalletId] = useState(
    initialWallets[0]?.id ?? null
  );
  const [selectedWalletName, setSelectedWalletName] =
    useState(initialWallets[0]?.name) ?? "";

  // form action react-19
  const [message, formAction, isPending] = useActionState(incomeAction, null);

  //   showing toast
  useEffect(() => {
    if (!message) return;

    if (message.success) {
      toast.success(message.message);
    } else {
      toast.error(message.message);
    }
  }, [message]);

  //   tanstack query
  const { data: wallets } = useQuery({
    queryKey: ["wallets"],
    queryFn: fetchWallets,
    initialData: initialWallets,
  });

  if (!wallets || wallets.length === 0) {
    return <p>No wallets found</p>;
  }

  return (
    <CardComponent
      title="Add Income"
      actions={
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline">
              {selectedWalletName || "Select Wallet"}
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-56">
            {wallets.map((w: IData) => (
              <DropdownMenuItem
                key={w.id}
                onClick={() => {
                  setSelectedWalletId(w.id);
                  setSelectedWalletName(w.name);
                }}>
                {w.name}
              </DropdownMenuItem>
            ))}
          </DropdownMenuContent>
        </DropdownMenu>
      }>
      {/* ERROR MESSSAGE */}
      {message?.success === false && (
        <p className="mb-3 text-red-500 text-center">{message?.message}</p>
      )}
      <form action={formAction} className="flex flex-col gap-4">
        <input type="hidden" name="wallet_id" value={selectedWalletId ?? ""} />
        <InputForm
          id="amount"
          name="amount"
          type="number"
          placeholder="ex : 10000, ">
          amount
        </InputForm>
        <InputForm
          id="notes"
          name="notes"
          type="text"
          placeholder="ex : salary, bonus ">
          notes
        </InputForm>
        <Button type="submit" className="bg-teal-700">
          {isPending ? "loading..." : "Submit"}
        </Button>
      </form>
    </CardComponent>
  );
};

export default IncomeCardTanstack;
