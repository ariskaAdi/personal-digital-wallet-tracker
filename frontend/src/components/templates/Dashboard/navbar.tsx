import Notification from "@/components/molecules/Notification";
import UserInfo from "@/components/molecules/UserInfo";
import React from "react";

const NavbarDahsboard = () => {
  return (
    <>
      <UserInfo />

      <div className="w-8 sm:hidden" />
      <Notification />
    </>
  );
};

export default NavbarDahsboard;
