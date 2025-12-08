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
import axios from "axios";
import React, { useActionState, useEffect, useState } from "react";
import { incomeAction } from "./action";

type IData = {
  id: number;
  name: string;
  balance: number;
};

const IncomeCard = () => {
  const [wallets, setWallets] = useState<IData[]>([]);
  const [selectedWalletId, setSelectedWalletId] = useState<number | null>(null);
  const [selectedWalletName, setSelectedWalletName] = useState("");
  const [loading, setLoading] = useState(false);
  const [message, formAction, isPending] = useActionState(incomeAction, null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const response = await axios.get(
          `${process.env.NEXT_PUBLIC_API_URL}/wallet/all`,
          { withCredentials: true }
        );
        const data: IData[] = response.data.data;

        setWallets(data);
        // auto select first wallet
        if (data.length > 0) {
          setSelectedWalletId(data[0].id);
          setSelectedWalletName(data[0].name);
        }
        setLoading(false);
      } catch (error) {
        console.log(error);
      } finally {
      }
    };
    fetchData();
  }, []);

  if (loading) {
    return <p>loading...</p>;
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
            {wallets.map((wallet) => (
              <DropdownMenuItem
                key={wallet.id}
                onClick={() => {
                  setSelectedWalletId(wallet.id);
                  setSelectedWalletName(wallet.name);
                }}>
                {wallet.name}
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

export default IncomeCard;
