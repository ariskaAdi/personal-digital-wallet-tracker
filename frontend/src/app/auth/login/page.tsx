import dynamic from "next/dynamic";

const LoginPage = dynamic(() => import("@/features/auth/login"), {
  ssr: true,
});

const ServerLoginPage = async () => {
  return <LoginPage />;
};

export default ServerLoginPage;
