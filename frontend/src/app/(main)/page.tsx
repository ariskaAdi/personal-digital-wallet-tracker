"use client";
import { Button } from "@/components/ui/button";
import axios from "axios";
import Link from "next/link";
import { useState } from "react";

const apiUrl = process.env.NEXT_PUBLIC_API_URL;

type Idata = {
  id: number;
  name: string;
  email: string;
};

export default function Home() {
  const [data, setData] = useState<Idata[]>([]);
  const fetchData = async () => {
    try {
      const res = await axios.get(`${apiUrl}/user`);
      setData(res.data.data);
      console.log(res.data.data);
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <div className="h-screen flex justify-center items-center">
      <Button onClick={fetchData}>CHECK USER</Button>
      <div className="relative p-2 m-2 flex justify-center items-center">
        {data.length > 0 && (
          <ul>
            {data.map((item) => (
              <li key={item.id}>
                {item.id} {item.name} - {item.email}
              </li>
            ))}
          </ul>
        )}
      </div>
      <Link href="/auth/login">
        <Button>Login</Button>
      </Link>
    </div>
  );
}
