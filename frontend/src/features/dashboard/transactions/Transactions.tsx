import ExpenseCard from "@/components/organisms/ExpenseCard";
import IncomeCard from "@/components/organisms/IncomeCard";
import RecentTx from "@/components/organisms/RecentTx";
import React from "react";

const Transactions = () => {
  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      {/* Scrollable content */}
      <div className="flex-1 overflow-auto">
        <div className="p-4 sm:p-6 space-y-6">
          {/* Bottom section */}
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            {/* Recent transactions */}
            <div>
              <IncomeCard />
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
