"use client";

import CardComponent from "@/components/molecules/CardComponent";
import RealTimeClock from "@/components/molecules/RealTimeClock";
import { formatCurrency } from "@/lib/utils";
import axios from "axios";
import { Send, TrendingUp, MoreVertical } from "lucide-react";
import { useEffect, useState } from "react";

export default function BalanceCard() {
  const [balance, setBalance] = useState<number>(0);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    const response = async () => {
      setIsLoading(true);
      try {
        const result = await axios.get(
          `${process.env.NEXT_PUBLIC_API_URL}/wallet/all`,
          { withCredentials: true }
        );
        const wallet = result.data.data;
        if (wallet.length > 0) {
          setBalance(wallet[0].balance);
        }
      } catch (error) {
        console.log(error);
      } finally {
        setIsLoading(false);
      }
    };
    response();
  }, []);

  if (isLoading) {
    return (
      <div className="bg-white dark:bg-slate-800 rounded-2xl p-6 shadow-sm border border-slate-200 dark:border-slate-700">
        <p className="text-slate-600 dark:text-slate-400">Loading...</p>
      </div>
    );
  }

  return (
    <CardComponent
      title="Your Total Balance"
      actions={
        <button className="p-2 hover:bg-slate-100 dark:hover:bg-slate-700 rounded-lg transition-colors">
          <MoreVertical size={18} className="text-slate-400" />
        </button>
      }>
      <div className="mb-6">
        <p className="text-4xl font-bold text-teal-600 mb-2">
          {formatCurrency(balance)}
        </p>
        <RealTimeClock />
      </div>

      {/* Action buttons */}
      <div className="grid grid-cols-3 gap-3">
        <button className="flex flex-col items-center gap-2 p-3 hover:bg-slate-50 dark:hover:bg-slate-700 rounded-lg transition-colors">
          <Send size={20} className="text-teal-600" />
          <span className="text-xs font-medium text-slate-600 dark:text-slate-400">
            Send
          </span>
        </button>
        <button className="flex flex-col items-center gap-2 p-3 hover:bg-slate-50 dark:hover:bg-slate-700 rounded-lg transition-colors">
          <TrendingUp size={20} className="text-blue-600" />
          <span className="text-xs font-medium text-slate-600 dark:text-slate-400">
            Invest
          </span>
        </button>
        <button className="flex flex-col items-center gap-2 p-3 hover:bg-slate-50 dark:hover:bg-slate-700 rounded-lg transition-colors">
          <div className="w-5 h-5 rounded-full border-2 border-amber-600" />
          <span className="text-xs font-medium text-slate-600 dark:text-slate-400">
            More
          </span>
        </button>
      </div>
    </CardComponent>
  );
}
