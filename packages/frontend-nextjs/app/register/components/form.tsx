"use client";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import {
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from "@/components/ui/form";
import Link from "next/link";
import { useState } from "react";
import {
  createUser,
  isEmailAvaliable,
  User,
} from "@/actions/api/users/registration";
import { useToast } from "@/hooks/use-toast";
import { useRouter } from "next/navigation";

const formSchema = z
  .object({
    name: z
      .string()
      .max(255, "O nome deve ter no máximo 255 caracteres")
      .nonempty("Nome é obrigatório"),
    email: z.string().email("E-mail inválido").nonempty("E-mail é obrigatório"),
    bio: z
      .string()
      .max(400, "A biografia deve ter no máximo 400 caracteres")
      .optional(),
    password: z
      .string()
      .min(8, "A senha deve ter no mínimo 8 caracteres")
      .nonempty("Senha é obrigatória"),
    password_confirm: z.string().nonempty("Confirmação de senha é obrigatória"),
  })
  .refine((data) => data.password === data.password_confirm, {
    message: "As senhas não coincidem",
    path: ["password_confirm"],
  });

type FormData = z.infer<typeof formSchema>;

export default function RegistrationForm() {
  const [availableEmail, setAvaliableEmail] = useState<boolean | null>(null);
  const { toast } = useToast();
  const router = useRouter();

  const form = useForm<FormData>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      email: "",
      bio: "",
      password: "",
      password_confirm: "",
    },
  });

  async function onSubmit(data: FormData) {
    if (availableEmail === false) {
      form.setError("email", {
        type: "",
        message: "Este e-mail já está em uso",
      });
      return;
    }

    const user: User = {
      name: data.name,
      email: data.email,
      bio: data.bio,
      password: data.password,
      password_confirm: data.password_confirm,
    };

    const message = await createUser(user);
    if (message) {
      toast({
        title: "Erro ao criar o usuário",
        description: message.error,
        variant: "destructive",
      });
      return;
    }

    toast({
      title: "Usuário registrado com sucesso",
      description: "Você será redirecionado para a página de login",
    });
    router.push("/");
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Nome</FormLabel>
              <FormControl>
                <Input placeholder="Seu nome completo" {...field} />
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
                    if (!email) {
                      return;
                    }

                    const availableEmail = await checkEmailAvailability(email);
                    if (!availableEmail) {
                      form.setError("email", {
                        type: "",
                        message: "Este e-mail já está em uso",
                      });
                      setAvaliableEmail(false);
                      return;
                    }
                    setAvaliableEmail(true);
                  }}
                  onChange={(e) => {
                    field.onChange(e);
                    form.clearErrors("email");
                    setAvaliableEmail(null); // Reset state
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

        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Senha</FormLabel>
              <FormControl>
                <Input type="password" placeholder="Sua senha" {...field} />
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
              <FormLabel>Confirmar Senha</FormLabel>
              <FormControl>
                <Input
                  type="password"
                  placeholder="Confirme sua senha"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />

        <div className="flex flex-row justify-end">
          <Button variant="secondary" asChild>
            <Link href="/">Voltar</Link>
          </Button>
          <Button
            type="submit"
            className="ml-3"
            disabled={availableEmail === false}
          >
            {form.formState.isSubmitting ? (
              <span className="animate-spin rounded-full h-4 w-4 border border-white border-t-transparent mr-2"></span>
            ) : (
              "Registrar"
            )}
          </Button>
        </div>
      </form>
    </Form>
  );

  async function checkEmailAvailability(email: string): Promise<Boolean> {
    if (email) {
      return await isEmailAvaliable(email);
    }

    return true;
  }
}
