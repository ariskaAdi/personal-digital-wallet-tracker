"use client";
import Notification from "@/components/molecules/Notification";
import UserInfo from "@/components/molecules/UserInfo";
import { cn } from "@/lib/utils";
import {
  BanknoteArrowDown,
  BanknoteArrowUp,
  LayoutDashboard,
  Menu,
  User,
  X,
} from "lucide-react";
import Link from "next/link";
import React, { useState } from "react";

interface DashboardLayoutProps {
  children: React.ReactNode;
}

const sideBarMenu = [
  { icon: LayoutDashboard, label: "Dashboard", link: "/profile" },
  { icon: BanknoteArrowUp, label: "Income", link: "/profile" },
  { icon: BanknoteArrowDown, label: "Expense", link: "/profile" },
];

const userMenu = [{ icon: User, label: "Profile", link: "/profile" }];

const DashboardLayout = ({ children }: DashboardLayoutProps) => {
  const [sidebarOpen, setSidebarOpen] = useState(false);
  return (
    <div className="flex h-screen bg-background">
      {/* Overlay mobile */}
      {sidebarOpen && (
        <div
          className="fixed inset-0 bg-black/50 z-40 lg:hidden"
          onClick={() => setSidebarOpen(false)}
        />
      )}

      {/* Sidebar */}
      <div
        className={cn(
          "fixed inset-y-0 left-0 z-50 w-64 bg-sidebar transform transition-transform duration-300 ease-in-out lg:translate-x-0 lg:static lg:inset-0",
          sidebarOpen ? "translate-x-0" : "-translate-x-full"
        )}>
        <div className="flex flex-col h-full">
          {/* User Info */}
          <div className="flex items-center justify-between p-6 border-b border-sidebar-border">
            <Link href="/">
              <div className="flex items-center gap-2">
                <div className="w-8 h-8 bg-black rounded flex items-center justify-center">
                  {/* for avatar */}
                  <span className="text-white text-sm font-bold">A</span>
                </div>
                <div>
                  <h2 className="text-sidebar-foreground font-bold">
                    Name App
                  </h2>
                  <p className="text-sidebar-foreground/60 text-xs">
                    Personal Digital Wallet
                  </p>
                </div>
              </div>
            </Link>
            <button
              onClick={() => setSidebarOpen(false)}
              className="lg:hidden text-sidebar-foreground">
              <X className="w-5 h-5" />
            </button>
          </div>

          {/* Main Menu */}
          <div className="flex-1 px-4 py-6">
            <div className="mb-6">
              <h3 className="text-sidebar-foreground/60 text-xs font-medium uppercase tracking-wider mb-3">
                MAIN MENU
              </h3>
              <nav className="space-y-1">
                {sideBarMenu.map((item, index) => {
                  //   const isActive = pathname === item.link;
                  return (
                    <Link
                      key={index}
                      href={item.link}
                      className={cn(
                        "flex items-center justify-between px-3 py-2 rounded-lg text-sm transition-colors"
                      )}>
                      <div className="flex items-center gap-3">
                        <item.icon className="w-5 h-5" />
                        <span>{item.label}</span>
                      </div>
                    </Link>
                  );
                })}
              </nav>
            </div>
            {/* Setting */}
            <div>
              <h3 className="text-sidebar-foreground/60 text-xs font-medium uppercase tracking-wider mb-3">
                SETTING ACCOUNT
              </h3>
              <nav className="space-y-1">
                {/* setting */}
                {userMenu.map((item, index) => {
                  //   const isActive = pathname === item.link;
                  return (
                    <Link
                      key={index}
                      href={item.link}
                      className={cn(
                        "flex items-center justify-between px-3 py-2 rounded-lg text-sm transition-colors"
                      )}>
                      <div className="flex items-center gap-3">
                        <item.icon className="w-5 h-5" />
                        <span>{item.label}</span>
                      </div>
                    </Link>
                  );
                })}
                <div className="mt-8">{/* <ButtonLogout /> */}</div>
              </nav>
            </div>
          </div>
        </div>
      </div>

      {/* Main content */}
      <div className="flex-2 flex flex-col overflow-hidden">
        {/* Mobile header */}
        <div className="flex items-center justify-between p-4 border-b border-border bg-card">
          <button
            onClick={() => setSidebarOpen(true)}
            className="text-foreground lg:hidden">
            <Menu className="w-6 h-6" />
          </button>
          <UserInfo />

          <div className="w-8 sm:hidden" />
          <Notification />
        </div>

        {/* Page content */}
        <main className="flex-1 overflow-auto p-6">{children}</main>
      </div>
    </div>
  );
};

export default DashboardLayout;
