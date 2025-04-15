"use client";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  BookOpenCheck,
  FileSpreadsheet,
  LogOut,
  ThumbsUp,
  UserPen,
} from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from "./ui/sidebar";
import { handleLogout } from "@/actions/logout";
import { JSX } from "react";

interface User {
  id: string;
  name: string;
  email: string;
}
export function AppSidebar({ user }: { user: User }): JSX.Element {
  const { open } = useSidebar();
  const pathname = usePathname();
  const router = useRouter();

  return (
    <Sidebar collapsible="icon">
      <SidebarHeader className="w-full flex justify-center items-center">
        <Image
          src="/images/logo.svg"
          alt="dialog"
          height={open ? 50 : 20}
          width={open ? 75 : 28}
        />
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Feed</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              <SidebarMenuItem key="LinhaTempo">
                <SidebarMenuButton
                  asChild
                  className={getMenuItemStyle(pathname, "/timeline")}
                >
                  <Link href="/timeline">
                    <ThumbsUp />
                    <span>Linha do Tempo</span>
                  </Link>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        <SidebarGroup>
          <SidebarGroupLabel>Perfil</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              <SidebarMenuItem key="User">
                <SidebarMenuButton
                  asChild
                  className={getMenuItemStyle(pathname, "/profile/me")}
                >
                  <Link href="/profile/me">
                    <UserPen />
                    <span>Editar Perfil</span>
                  </Link>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <SidebarMenuButton
                  size="lg"
                  className="text-sidebar-accent-foreground"
                >
                  <div className="grid flex-1 text-left text-sm leading-tight">
                    <span className="truncate font-semibold">{user.name}</span>
                    <span className="truncate text-xs">{user.email}</span>
                  </div>
                </SidebarMenuButton>
              </DropdownMenuTrigger>
              <DropdownMenuContent
                className="w-[--radix-dropdown-menu-trigger-width] min-w-56 rounded-lg"
                side="top"
                align="center"
                sideOffset={4}
              >
                <DropdownMenuGroup>
                  <DropdownMenuItem
                    onClick={async () => {
                      await handleLogout();
                      router.replace("/");
                      router.refresh();
                    }}
                  >
                    <LogOut />
                    Sair
                  </DropdownMenuItem>
                </DropdownMenuGroup>
              </DropdownMenuContent>
            </DropdownMenu>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  );
}

function getMenuItemStyle(pathname: string, targetPath: string): string {
  const selectedStyle = "bg-zinc-200 hover:bg-zinc-200";

  return pathname?.match(targetPath) ? selectedStyle : "";
}
