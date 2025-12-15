import BalanceCard from "@/components/organisms/BalanceCard";
import DebitCard from "@/components/organisms/DebitCard";
import ExpensesChart from "@/components/organisms/ExpenseChart";
import InvoiceTable from "@/components/organisms/InvoiceTable";
import RecentTx from "@/components/organisms/RecentTx";
export default function HomeDashboard() {
  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      {/* Scrollable content */}
      <div className="flex-1 overflow-auto">
        <div className="p-4 sm:p-6 space-y-6">
          {/* Cards grid */}
          <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
            {/* Debit card - spans 2 columns on desktop */}
            <div className="lg:col-span-1">
              <DebitCard />
            </div>

            {/* Balance card */}
            <div className="lg:col-span-1">
              <BalanceCard />
            </div>

            {/* Expenses chart - spans full width or 1 column */}
            <div className="lg:col-span-1">
              <ExpensesChart />
            </div>
          </div>

          {/* Bottom section */}
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            {/* Recent transactions */}
            <div>
              <RecentTx />
            </div>

            {/* Invoice activity table */}
            <div className="w-full lg:w-auto order-2 lg:order-1">
              <InvoiceTable />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
