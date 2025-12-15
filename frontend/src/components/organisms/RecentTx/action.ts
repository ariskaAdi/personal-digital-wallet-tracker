import { fetchWithCookie } from "@/lib/fetchWithCookie";

export const fetchBalance = async () => {
  try {
    const res = await fetchWithCookie(`${process.env.NEXT_PUBLIC_API_URL}/tx`, {
      method: "GET",
      cache: "no-store",
    });
    const data = await res.json();
    console.log(data.data);
    return data.data ?? [];
  } catch (error) {
    console.log(error);
    throw error;
  }
};

// export default async function RecentTxServer() {
//   const tx = await fetchBalance();
//   return (
//     <ComponentType>
//       <RecentTxClient tx={tx} />

//   )
// }
