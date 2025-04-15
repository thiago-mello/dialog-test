import { cookies } from "next/headers";
import { getIronSession } from "iron-session";
import { AppSidebar } from "@/components/app-sidebar";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import React, { JSX } from "react";
import { SessionData, sessionOptions } from "@/lib/session";
import { QueryClientWrapper } from "@/providers/query";
import { Toaster } from "@/components/ui/toaster";

/**
 * A wrapper component that provides the application's layout structure
 * Includes sidebar functionality and responsive layout adjustments based on authentication state
 *
 * @param {Object} props - Component props
 * @param {React.ReactNode} props.children - Child elements to be rendered within the wrapper
 * @returns {JSX.Element} Wrapped application content with sidebar and layout structure
 */
export default async function AppWrapper({
  children,
}: {
  children: React.ReactNode;
}): Promise<JSX.Element> {
  const session = await getIronSession<SessionData>(
    await cookies(),
    sessionOptions
  );
  const isLoggedIn = !!session.userId;

  return (
    <QueryClientWrapper>
      <SidebarProvider>
        {isLoggedIn && (
          <AppSidebar
            user={{
              id: session.userId!!,
              email: session.email!!,
              name: session.name!!,
            }}
          />
        )}

        <main className="w-full">
          {isLoggedIn && <SidebarTrigger />}
          <div
            className={`w-full ${isLoggedIn ? "max-w-screen-xl" : ""} mx-auto`}
          >
            {children}
          </div>
        </main>
        <Toaster />
      </SidebarProvider>
    </QueryClientWrapper>
  );
}
