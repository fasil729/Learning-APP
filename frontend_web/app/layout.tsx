import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import AppProvider from "@/app/provider";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Learning App",
  description: " learning app with Ai integration",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <AppProvider>
      <body className={inter.className + " " + "body-color"}>{children}</body>
      </AppProvider>
    </html>
  );
}
