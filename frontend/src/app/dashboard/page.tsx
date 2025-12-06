import BalanceCard from "@/components/organisms/tempalatev0/balance-card";
import DebitCard from "@/components/organisms/tempalatev0/debit-card";
import ExpensesChart from "@/components/organisms/tempalatev0/expenses-chart";
import InvoiceTable from "@/components/organisms/tempalatev0/invoice-table";
import RecentTransactions from "@/components/organisms/tempalatev0/recent-transactions";

export default function Home() {
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
              <RecentTransactions />
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
