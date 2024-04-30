import type { Metadata } from "next";
import { Inter } from "next/font/google";
import AppProvider from "@/app/provider";
import { TopicSideBar } from "./components/sideBar";

const inter = Inter({ subsets: ["latin"] });


export default function TopicDetailLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <AppProvider>
      <body className={inter.className + " " + "body-color"}>
        <div className="flex">
        <TopicSideBar/>
        <div className='ml-[21%]'>
          { children }
        </div>
        </div>
        </body>
      </AppProvider>
    </html>
  );
}
