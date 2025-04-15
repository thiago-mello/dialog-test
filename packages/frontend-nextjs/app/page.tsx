import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import Image from "next/image";
import LoginForm from "./components/form";
import { getIronSession } from "iron-session";
import { cookies } from "next/headers";
import { SessionData, sessionOptions } from "@/lib/session";
import { redirect } from "next/navigation";

export default async function Home({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}) {
  const session = await getIronSession<SessionData>(
    await cookies(),
    sessionOptions
  );
  if (session.userId) {
    redirect("/timeline");
    return;
  }

  const sessionExpired = (await searchParams).expired === "true";

  return (
    <div className="flex flex-col min-h-screen items-center justify-center bg-zinc-100 p-4">
      <div className="max-w-md mx-auto w-full mb-2">
        <Image src={"images/logo.svg"} alt="dialog" height={45} width={105} />
      </div>
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle className="text-2xl">Login</CardTitle>
          <CardDescription>
            Digite seu e-mail e senha para acessar a plataforma
          </CardDescription>
        </CardHeader>
        <CardContent>
          <LoginForm expired={sessionExpired} />
        </CardContent>
      </Card>
    </div>
  );
}
