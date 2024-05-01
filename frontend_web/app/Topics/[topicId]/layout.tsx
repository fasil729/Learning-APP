'use client'
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import AppProvider from "@/app/provider";
import { TopicSideBar } from "./components/sideBar";
import TopHeader from "@/components/shared/top_header/top_header";

const inter = Inter({ subsets: ["latin"] });


export default function TopicDetailLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
        <div className="flex">
        <TopicSideBar/>
        <div className='ml-[18%]'>
        <TopHeader/>
        <div className="py-10 px-10">
          { children }
        </div>
        </div>
        </div>
  );
}
