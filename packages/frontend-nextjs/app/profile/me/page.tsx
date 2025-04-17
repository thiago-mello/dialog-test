import { getMyUser } from "@/actions/api/users/registration";
import { ProfileForm } from "./components/form";
import { notFound } from "next/navigation";
import DeleteAccountButton from "./components/delete-button";

export default async function ProfilePage() {
  const user = await getMyUser();
  if (!user) {
    notFound();
  }

  return (
    <div className="max-w-2xl mx-auto py-10">
      <div className="flex flex-row justify-between items-center mb-6">
        <h1 className="text-2xl font-bold ">Editar Perfil</h1>
        <DeleteAccountButton />
      </div>
      <ProfileForm user={user} />
    </div>
  );
}
