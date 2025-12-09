import ExpenseCard from "@/components/organisms/ExpenseCard";
import IncomeCardTanstack from "@/components/organisms/IncomeCard";
import RecentTx from "@/components/organisms/RecentTx";
import { fetchWallets } from "@/features/dashboard/transactions/action";

const Transactions = async () => {
  const wallets = await fetchWallets();
  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      {/* Scrollable content */}
      <div className="flex-1 overflow-auto">
        <div className="p-4 sm:p-6 space-y-6">
          {/* Bottom section */}
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            {/* Recent transactions */}
            <div>
              <IncomeCardTanstack initialWallets={wallets} />
            </div>
            {/* Invoice activity table - hidden on mobile, shown on tablet+ */}
            <div>
              <ExpenseCard />
            </div>
          </div>
          <div>
            <RecentTx />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Transactions;
