import { Filter, Download } from "lucide-react";

const invoices = [
  {
    date: "21 DEC, 2024",
    invoice: "INV-2024-001",
    recipient: "John Doe",
    status: "Paid",
    amount: "$1,200.00",
  },
  {
    date: "20 DEC, 2024",
    invoice: "INV-2024-002",
    recipient: "Jane Smith",
    status: "Pending",
    amount: "$850.50",
  },
  {
    date: "19 DEC, 2024",
    invoice: "INV-2024-003",
    recipient: "Bob Wilson",
    status: "Paid",
    amount: "$2,350.00",
  },
  {
    date: "18 DEC, 2024",
    invoice: "INV-2024-004",
    recipient: "Alice Johnson",
    status: "Rejected",
    amount: "$450.25",
  },
];

export default function InvoiceTable() {
  return (
    <div className="bg-white dark:bg-slate-800 rounded-2xl p-6 shadow-sm border border-slate-200 dark:border-slate-700">
      <div className="flex items-center justify-between mb-6">
        <h2 className="text-lg font-semibold text-slate-900 dark:text-white">
          Invoice Activity
        </h2>
        <div className="flex items-center gap-2">
          <button className="p-2 hover:bg-slate-100 dark:hover:bg-slate-700 rounded-lg transition-colors">
            <Filter size={18} className="text-slate-400" />
          </button>
          <button className="p-2 hover:bg-slate-100 dark:hover:bg-slate-700 rounded-lg transition-colors">
            <Download size={18} className="text-slate-400" />
          </button>
        </div>
      </div>

      <div className="overflow-x-auto">
        <table className="w-full text-sm">
          <thead>
            <tr className="border-b border-slate-200 dark:border-slate-700">
              <th className="px-4 py-3 text-left font-semibold text-slate-600 dark:text-slate-400 text-xs uppercase">
                Date & Time
              </th>
              <th className="px-4 py-3 text-left font-semibold text-slate-600 dark:text-slate-400 text-xs uppercase">
                Invoice Number
              </th>
              <th className="px-4 py-3 text-left font-semibold text-slate-600 dark:text-slate-400 text-xs uppercase">
                Recipient
              </th>
              <th className="px-4 py-3 text-left font-semibold text-slate-600 dark:text-slate-400 text-xs uppercase">
                Status
              </th>
              <th className="px-4 py-3 text-right font-semibold text-slate-600 dark:text-slate-400 text-xs uppercase">
                Amount
              </th>
            </tr>
          </thead>
          <tbody>
            {invoices.map((invoice, idx) => (
              <tr
                key={idx}
                className="border-b border-slate-100 dark:border-slate-700/50 hover:bg-slate-50 dark:hover:bg-slate-700/30 transition-colors">
                <td className="px-4 py-4 text-slate-600 dark:text-slate-300">
                  {invoice.date}
                </td>
                <td className="px-4 py-4 text-slate-900 dark:text-white font-medium">
                  {invoice.invoice}
                </td>
                <td className="px-4 py-4 text-slate-600 dark:text-slate-300">
                  {invoice.recipient}
                </td>
                <td className="px-4 py-4">
                  <span
                    className={`px-3 py-1 rounded-full text-xs font-medium ${
                      invoice.status === "Paid"
                        ? "bg-green-100 dark:bg-green-900/30 text-green-700 dark:text-green-400"
                        : invoice.status === "Pending"
                        ? "bg-yellow-100 dark:bg-yellow-900/30 text-yellow-700 dark:text-yellow-400"
                        : "bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-400"
                    }`}>
                    {invoice.status}
                  </span>
                </td>
                <td className="px-4 py-4 text-right text-slate-900 dark:text-white font-medium">
                  {invoice.amount}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
