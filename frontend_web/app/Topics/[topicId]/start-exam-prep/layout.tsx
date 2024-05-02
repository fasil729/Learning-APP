'use client'
import { Inter } from "next/font/google";
import TopHeader from "@/components/shared/top_header/top_header";

const inter = Inter({ subsets: ["latin"] });


export default function ExamPrepLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className=" ">
          
      <div className='sticky top-0 z-50 h-[100px]'><TopHeader></TopHeader>
      </div>
      <div className='flex'>
      <div className='grow'>
      
        { children }
      </div>
      </div>
    </div>
  );
}
