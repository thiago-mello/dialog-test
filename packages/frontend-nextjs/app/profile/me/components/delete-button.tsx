"use client";
import { Button } from "@/components/ui/button";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { useState } from "react";
import { Spinner } from "@/components/ui/spinner";
import { deleteMyUser } from "@/actions/api/users/registration";
import { handleLogout } from "@/actions/logout";
import { useRouter } from "next/navigation";
import { useToast } from "@/hooks/use-toast";

export default function DeleteAccountButton() {
  const [isDeleting, setIsDeleting] = useState<boolean>(false);
  const router = useRouter();
  const { toast } = useToast();

  async function handleDelete() {
    setIsDeleting(true);

    const response = await deleteMyUser();
    // if user was deleted successfully
    if (response.status === 200) {
      await handleLogout();
      toast({
        title: "Sua conta foi excluída com sucesso",
        description: "É uma pena vê-lo ir...",
      });
      setIsDeleting(false);
      router.replace("/");
    }

    setIsDeleting(false);
  }

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button variant="destructive">Excluir Conta</Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Deseja excluir sua conta?</AlertDialogTitle>
          <AlertDialogDescription>
            Tem certeza que deseja excluir sua conta? Todos os seus posts e
            reações serão removidos definitivamente. Essa ação não pode ser
            desfeita.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Cancelar</AlertDialogCancel>
          <AlertDialogAction
            className="w-20"
            onClick={handleDelete}
            disabled={isDeleting}
          >
            {isDeleting ? <Spinner /> : "Excluir"}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
