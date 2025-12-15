"use client";
import CardComponent from "@/components/molecules/CardComponent";
import { formatCurrency, formatDate } from "@/lib/utils";

import { BanknoteArrowDown, BanknoteArrowUp } from "lucide-react";
import { ITx } from "./type";

type RecentTxClientProps = {
  tx: ITx[];
};

export default function RecentTxClient({ tx }: RecentTxClientProps) {
  return (
    <CardComponent
      title="Recent Transactions"
      titleClassName="text-lg font-semibold text-slate-900 dark:text-white"
      actions={
        <button className="text-xs px-3 py-1 text-teal-600 hover:bg-teal-50 dark:hover:bg-teal-900/20 rounded transition-colors">
          Last 7 days
        </button>
      }>
      <div className="space-y-4">
        {tx?.map((item, index) => (
          <div
            key={index}
            className="flex items-center justify-between py-3 border-b border-slate-100 dark:border-slate-700 last:border-0">
            <div className="flex items-center gap-4 flex-1">
              <div className="w-10 h-10 bg-slate-100 dark:bg-slate-700 rounded-full flex items-center justify-center text-lg">
                {item.type === "INCOME" ? (
                  <BanknoteArrowUp />
                ) : (
                  <BanknoteArrowDown />
                )}
              </div>
              <div className="flex-1 min-w-0">
                <p className="text-sm font-medium text-slate-900 dark:text-white truncate">
                  {item.notes}
                </p>
                <p className="text-xs text-slate-500 dark:text-slate-400 truncate">
                  {formatDate(item.created_at)}
                </p>
              </div>
            </div>
            <div className="text-right ml-2">
              <p
                className={`text-sm font-semibold ${
                  item.type === "INCOME"
                    ? "text-green-600"
                    : "text-red-600 dark:text-white"
                }`}>
                {item.type === "INCOME" ? "+" : "-"}{" "}
                {formatCurrency(item.amount)}
              </p>
            </div>
          </div>
        ))}
      </div>
    </CardComponent>
  );
}
