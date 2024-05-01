'use client'
import { Inter } from "next/font/google";
import TopHeader from "@/components/shared/top_header/top_header";
import { ExperimentSideBar } from "./components/experimentSideBar";

const inter = Inter({ subsets: ["latin"] });


export default function TopicDetailLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
        <div>
            <ExperimentSideBar/>
            <div className="ml-[18%]">
                <TopHeader />
                {children}
            </div>
        </div>
  );
}
