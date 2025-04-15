import { getMyUser } from "@/actions/api/users/registration";
import { ProfileForm } from "./components/form";
import { notFound } from "next/navigation";

export default async function ProfilePage() {
  const user = await getMyUser();
  if (!user) {
    notFound();
  }

  return (
    <div className="max-w-2xl mx-auto py-10">
      <h1 className="text-2xl font-bold mb-6">Editar Perfil</h1>
      <ProfileForm user={user} />
    </div>
  );
}
