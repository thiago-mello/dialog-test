import type { Metadata } from "next";
import "./globals.css";
import AppWrapper from "./wrapper";

export const metadata: Metadata = {
  title: "Dialog",
  description: "Desafio Técnico",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="antialiased">
        <AppWrapper>{children}</AppWrapper>
      </body>
    </html>
  );
}
