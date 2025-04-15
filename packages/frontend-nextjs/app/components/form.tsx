"use client";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useState } from "react";
import { AlertCircle, Eye, EyeOff, ShieldAlert } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";
import { handleLogin } from "@/actions/login";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import Link from "next/link";

export default function LoginForm({ expired }: { expired?: boolean }) {
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isLoggingIn, setIsLoggingIn] = useState(false);
  const router = useRouter();

  const formSchema = z.object({
    email: z.string().email({ message: "O email deve ser válido" }),
    password: z
      .string()
      .min(8, { message: "A senha deve conter ao menos 8 caracteres" }),
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    setIsLoggingIn(true);
    const loginError = await handleLogin(values.email, values.password);
    if (!loginError) {
      setError(null);
      router.push("/timeline");
      return;
    }

    switch (loginError.type) {
      case "USER_NOT_FOUND":
        setError("O usuário ou senha estão incorretos");
        break;
      case "INTERNAL_ERROR":
        setError(
          "Parece que o servidor está fora do ar no momento. Sentimos muito :("
        );
        break;
    }
    setIsLoggingIn(false);
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input
                  placeholder="usuario@gmail.com"
                  type="email"
                  {...field}
                  className="w-full"
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
                <div className="relative">
                  <Input
                    placeholder="Digite sua senha"
                    type={showPassword ? "text" : "password"}
                    {...field}
                    className="w-full pr-10"
                  />
                  <button
                    type="button"
                    onClick={() => setShowPassword(!showPassword)}
                    className="absolute right-2 top-1/2 -translate-y-1/2 text-zinc-500 hover:text-zinc-700"
                  >
                    {showPassword ? (
                      <EyeOff className="h-5 w-5" />
                    ) : (
                      <Eye className="h-5 w-5" />
                    )}
                  </button>
                </div>
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <br />

        {error && (
          <Alert variant="destructive">
            <AlertCircle className="h-4 w-4" />
            <AlertTitle>Erro ao fazer login</AlertTitle>
            <AlertDescription>{error}</AlertDescription>
          </Alert>
        )}

        {expired && (
          <Alert>
            <ShieldAlert className="h-4 w-4" />
            <AlertTitle>Sua sessão expirou</AlertTitle>
            <AlertDescription>
              Sua sessão chegou ao fim. Faça login novamente para acessar a
              plataforma.
            </AlertDescription>
          </Alert>
        )}

        <Button type="submit" className="w-full" disabled={isLoggingIn}>
          Entrar
        </Button>

        <div className="text-center text-sm text-gray-600">
          Não tem uma conta?{" "}
          <Link href="/register" className="text-blue-600 hover:underline">
            Cadastre-se
          </Link>
        </div>
      </form>
    </Form>
  );
}
