"use client";

import { PieChart, Pie, Cell, ResponsiveContainer } from "recharts";
import { MoreVertical } from "lucide-react";

const data = [
  { name: "Spent", value: 85.5, color: "#0d9488" },
  { name: "Remaining", value: 14.5, color: "#d1d5db" },
];

export default function ExpensesChart() {
  return (
    <div className="bg-white dark:bg-slate-800 rounded-2xl p-6 shadow-sm border border-slate-200 dark:border-slate-700">
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-sm font-semibold text-slate-600 dark:text-slate-400">
          Expenses Instead
        </h2>
        <button className="p-2 hover:bg-slate-100 dark:hover:bg-slate-700 rounded-lg transition-colors">
          <MoreVertical size={18} className="text-slate-400" />
        </button>
      </div>

      <div className="flex flex-col items-center">
        <div className="w-32 h-32 mb-4">
          <ResponsiveContainer width="100%" height="100%">
            <PieChart>
              <Pie
                data={data}
                cx="50%"
                cy="50%"
                innerRadius={50}
                outerRadius={65}
                dataKey="value"
                startAngle={90}
                endAngle={-270}>
                {data.map((entry, index) => (
                  <Cell key={`cell-${index}`} fill={entry.color} />
                ))}
              </Pie>
            </PieChart>
          </ResponsiveContainer>
        </div>

        <p className="text-2xl font-bold text-slate-900 dark:text-white">
          85.5%
        </p>
        <p className="text-sm text-slate-600 dark:text-slate-400 mb-4">
          Normal Level
        </p>

        <div className="w-full pt-4 border-t border-slate-200 dark:border-slate-700">
          <p className="text-xs text-slate-500 dark:text-slate-400">
            Total Exp:
          </p>
          <p className="text-lg font-bold text-teal-600">$1,820.80</p>
        </div>
      </div>
    </div>
  );
}
