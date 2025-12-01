import dynamic from "next/dynamic";

const RegisterPage = dynamic(() => import("@/features/auth/register"), {
  ssr: true,
});

const ServerLoginPage = async () => {
  return <RegisterPage />;
};

export default ServerLoginPage;
