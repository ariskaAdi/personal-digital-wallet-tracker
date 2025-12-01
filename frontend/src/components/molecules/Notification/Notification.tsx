import { Bell, MessageCircle } from "lucide-react";
import React from "react";

const Notification = () => {
  return (
    <div className="flex justify-between gap-4">
      <Bell size={20} className="text-foreground" />
      <MessageCircle size={20} className="text-foreground" />
    </div>
  );
};

export default Notification;
