"use client";
import { formatCurrency, formatDate } from "@/lib/utils";
// const transactions = [
//   {
//     icon: "🎵",
//     name: "Spotify Premium",
//     date: "23 December 2025, 02:45 AM",
//     amount: "+$8,200.00",
//     type: "received",
//   },
//   {
//     icon: "🎧",
//     name: "Spotify Premium",
//     date: "23 December 2025, 07:22 PM",
//     amount: "-$199.00",
//     type: "sent",
//   },
//   {
//     icon: "✈️",
//     name: "Transferwise - Received",
//     date: "21 December 2025, 10:30 AM",
//     amount: "+$1,200.00",
//     type: "received",
//   },
//   {
//     icon: "💳",
//     name: "H&M Payment",
//     date: "19 December 2025, 06:50 PM",
//     amount: "-$2,300.00",
//     type: "sent",
//   },
// ];

import axios from "axios";
import { BanknoteArrowDown, BanknoteArrowUp } from "lucide-react";
import { useEffect, useState } from "react";

export type ITx = {
  id: number;
  amount: number;
  notes: string;
  type: string;
  created_at: string;
};

export default function RecentTx() {
  const [tx, setTx] = useState<ITx[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setIsLoading(true);
        const result = await axios.get(
          `${process.env.NEXT_PUBLIC_API_URL}/tx`,
          { withCredentials: true }
        );
        setTx(result.data.data);
        console.log(result.data.data);
        setIsLoading(false);
      } catch (error) {
        console.log(error);
      } finally {
        setIsLoading(false);
      }
    };
    fetchData();
  }, []);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="bg-white dark:bg-slate-800 rounded-2xl p-6 shadow-sm border border-slate-200 dark:border-slate-700">
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-lg font-semibold text-slate-900 dark:text-white">
          Recent Transactions
        </h2>
        <button className="text-xs px-3 py-1 text-teal-600 hover:bg-teal-50 dark:hover:bg-teal-900/20 rounded transition-colors">
          Last 7 days
        </button>
      </div>

      <div className="space-y-4">
        {tx.map((tx, idx) => (
          <div
            key={idx}
            className="flex items-center justify-between py-3 border-b border-slate-100 dark:border-slate-700 last:border-0">
            <div className="flex items-center gap-4 flex-1">
              <div className="w-10 h-10 bg-slate-100 dark:bg-slate-700 rounded-full flex items-center justify-center text-lg">
                {tx.type === "INCOME" ? (
                  <BanknoteArrowUp />
                ) : (
                  <BanknoteArrowDown />
                )}
              </div>
              <div className="flex-1 min-w-0">
                <p className="text-sm font-medium text-slate-900 dark:text-white truncate">
                  {tx.notes}
                </p>
                <p className="text-xs text-slate-500 dark:text-slate-400 truncate">
                  {formatDate(tx.created_at)}
                </p>
              </div>
            </div>
            <div className="text-right ml-2">
              <p
                className={`text-sm font-semibold ${
                  tx.type === "INCOME"
                    ? "text-green-600"
                    : "text-red-600 dark:text-white"
                }`}>
                {tx.type === "INCOME" ? "+" : "-"} {formatCurrency(tx.amount)}
              </p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
