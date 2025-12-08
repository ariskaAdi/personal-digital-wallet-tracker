import dynamic from "next/dynamic";

const TransactionsPage = dynamic(
  () => import("@/features/dashboard/transactions"),
  {
    ssr: true,
  }
);

const ServerTransactionsPage = () => {
  return <TransactionsPage />;
};

export default ServerTransactionsPage;
