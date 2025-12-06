import { getMeUserService } from "@/services/user.service";
import { useEffect, useState } from "react";

export type UserInfoProps = {
  name: string;
  email: string;
};

const UserInfo = () => {
  const [dataUser, setDataUser] = useState<UserInfoProps | null>(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setLoading(true);

    const fetchData = async () => {
      try {
        const res = await getMeUserService();
        setDataUser(res && res.data);
      } catch (error) {
        console.log(error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <p>Loading user...</p>;
  return (
    <div className="flex justify-evenly gap-2">
      {dataUser && (
        <div>
          <p className="text-sm font-bold">{dataUser.name}</p>
          <p className="text-xs">{dataUser.email}</p>
        </div>
      )}
    </div>
  );
};

export default UserInfo;
