import BalanceCard from "@/components/molecules/BalanceCard";
import DebitCard from "@/components/molecules/DebitCard";
import ExpensesChart from "@/components/molecules/ExpenseChart";
import InvoiceTable from "@/components/molecules/InvoiceTable";
import RecentTx from "@/components/molecules/RecentTx";

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

            {/* Invoice activity table - hidden on mobile, shown on tablet+ */}
            <div className="hidden lg:block">
              <InvoiceTable />
            </div>
          </div>

          {/* Invoice table full width on mobile/tablet */}
          <div className="lg:hidden">
            <InvoiceTable />
          </div>
        </div>
      </div>
    </div>
  );
}
