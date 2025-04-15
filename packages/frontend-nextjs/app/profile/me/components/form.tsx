"use client";

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import {
  isEmailAvaliable,
  updateMyUser,
  UserPublicData,
} from "@/actions/api/users/registration";
import { useToast } from "@/hooks/use-toast";
import { useRouter } from "next/navigation";

// Zod Schema
const formSchema = z
  .object({
    name: z
      .string()
      .max(255, "Máximo de 255 caracteres")
      .nonempty("Nome é obrigatório"),
    email: z.string().email("E-mail inválido").nonempty("E-mail é obrigatório"),
    bio: z.string().max(400, "Máximo de 400 caracteres").optional(),
    password: z
      .string()
      .min(8, "A senha deve ter no mínimo 8 caracteres")
      .optional(),
    password_confirm: z.string().optional(),
  })
  .refine(
    (data) => {
      if (data.password || data.password_confirm) {
        return data.password === data.password_confirm;
      }
      return true;
    },
    {
      message: "As senhas não coincidem",
      path: ["password_confirm"],
    }
  );

type FormData = z.infer<typeof formSchema>;

export function ProfileForm({ user }: { user: UserPublicData }) {
  const [emailAvalilable, setEmailAvalilable] = useState<boolean | null>(null);
  const { toast } = useToast();
  const router = useRouter();

  const formDefaultValues: FormData = {
    name: user.name,
    email: user.email,
    bio: user.bio,
    password: undefined,
    password_confirm: undefined,
  };
  const form = useForm<FormData>({
    resolver: zodResolver(formSchema),
    defaultValues: formDefaultValues,
  });

  const onSubmit = async (data: FormData) => {
    if (emailAvalilable === false) {
      form.setError("email", {
        type: "manual",
        message: "Este e-mail já está em uso",
      });
      return;
    }

    const passwordChanged = !!data.password;
    await updateMyUser({ ...data });
    toast({
      title: "Informações atualizadas com sucesso",
      description: `As mudanças no seu perfil já estão visíveis para outros usuários${
        passwordChanged ? " e a sua senha foi atualizada." : "."
      }`,
    });
    router.refresh();
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-10">
        <div className="space-y-6 border p-6 rounded-lg shadow-sm">
          <h2 className="text-lg font-semibold mb-2">Informações Pessoais</h2>

          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Nome</FormLabel>
                <FormControl>
                  <Input placeholder="Seu nome" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>E-mail</FormLabel>
                <FormControl>
                  <Input
                    placeholder="email@exemplo.com"
                    {...field}
                    onBlur={async (e) => {
                      field.onBlur(); // Keep RHF behavior

                      const email = e.target.value;
                      if (!email || email === formDefaultValues.email) {
                        return;
                      }

                      const availableEmail = await checkEmailAvailability(
                        email
                      );
                      if (!availableEmail) {
                        form.setError("email", {
                          type: "",
                          message: "Este e-mail já está em uso",
                        });
                        setEmailAvalilable(false);
                        return;
                      }
                      setEmailAvalilable(true);
                    }}
                    onChange={(e) => {
                      field.onChange(e);
                      form.clearErrors("email");
                      setEmailAvalilable(null); // Reset state
                    }}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="bio"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Biografia</FormLabel>
                <FormControl>
                  <Textarea
                    placeholder="Conte um pouco sobre você"
                    className="resize-none"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="space-y-6 border p-6 rounded-lg shadow-sm">
          <h2 className="text-lg font-semibold mb-2">Alterar Senha</h2>

          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Nova Senha</FormLabel>
                <FormControl>
                  <Input
                    type="password"
                    placeholder="Deixe em branco para não alterar"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="password_confirm"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Confirmar Nova Senha</FormLabel>
                <FormControl>
                  <Input
                    type="password"
                    placeholder="Confirme a nova senha"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="flex flex-row justify-end">
          <Button
            type="reset"
            onClick={() => {
              form.reset({ ...formDefaultValues });
              setEmailAvalilable(true);
            }}
            variant="secondary"
          >
            Cancelar Mudanças
          </Button>
          <Button
            type="submit"
            className="ml-3"
            disabled={emailAvalilable === false}
          >
            {form.formState.isSubmitting ? (
              <span className="animate-spin rounded-full h-4 w-4 border border-white border-t-transparent mr-2"></span>
            ) : (
              "Salvar"
            )}
          </Button>
        </div>
      </form>
    </Form>
  );

  async function checkEmailAvailability(email: string): Promise<Boolean> {
    if (email && email !== formDefaultValues.email) {
      return await isEmailAvaliable(email);
    }

    return true;
  }
}
