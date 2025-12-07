import dynamic from "next/dynamic";

const HomeDashboardPage = dynamic(
  () => import("@/features/dashboard/homeDashboard"),
  {
    ssr: true,
  }
);

const ServerHomeDashboardPage = () => {
  return <HomeDashboardPage />;
};

export default ServerHomeDashboardPage;
