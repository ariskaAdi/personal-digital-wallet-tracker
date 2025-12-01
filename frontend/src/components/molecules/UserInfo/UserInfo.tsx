"use client";
import { getMeUserService } from "@/services/user.service";
import React, { useState } from "react";

export type UserInfoProps = {
  name: string;
  email: string;
};

const UserInfo = () => {
  const [user, setUser] = useState<UserInfoProps | null>(null);

  return (
    <div className="flex justify-evenly gap-2">
      {user && (
        <div>
          <p>{user.name}</p>
          <p>{user.email}</p>
        </div>
      )}
      <button
        onClick={async () => {
          const res = await getMeUserService();
          if (res?.success) {
            setUser(res.data);
          } else {
            setUser(null);
          }
        }}>
        CHECK
      </button>
    </div>
  );
};

export default UserInfo;
