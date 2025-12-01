import { UserInfoProps } from "@/components/molecules/UserInfo/UserInfo";
import { apiUrl } from "@/constant/api_url";
import axios from "axios";

export const getMeUserService = async (): Promise<{
  success: boolean;
  data: UserInfoProps;
} | null> => {
  try {
    const res = await axios.get(`${apiUrl}/user/me`, { withCredentials: true });
    console.log(res.data);
    return res.data;
  } catch (error) {
    console.error("Error getMeUserService:", error);
    return null;
  }
};
