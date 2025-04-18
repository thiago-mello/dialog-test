import { Card, CardTitle } from "@/components/ui/card";
import RegistrationForm from "./components/form";

export default function RegistrationPage() {
  return (
    <div className="flex flex-col w-full min-h-screen items-center justify-center bg-zinc-100 p-4">
      <Card className="w-full max-w-xl p-5">
        <CardTitle className="text-2xl font-bold mb-6">Cadastro</CardTitle>

        <RegistrationForm />
      </Card>
    </div>
  );
}
