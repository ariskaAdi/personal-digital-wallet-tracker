"use client";
import { useEffect, useState } from "react";

export default function RealTimeClock() {
  const [now, setNow] = useState(new Date());

  useEffect(() => {
    const timer = setInterval(() => {
      setNow(new Date());
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  return (
    <p className="text-xs text-slate-500 dark:text-slate-400">
      {now.toLocaleDateString("en-US", {
        year: "numeric",
        month: "long",
        day: "2-digit",
      })}{" "}
      •{" "}
      {now.toLocaleTimeString("en-US", {
        hour: "2-digit",
        minute: "2-digit",
        hour12: true,
      })}
    </p>
  );
}
