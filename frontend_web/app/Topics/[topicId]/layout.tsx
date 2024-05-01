import { Inter } from 'next/font/google';

import TopHeader from '@/components/shared/top_header/top_header';

import { TopicSideBar } from './components/sideBar';

const inter = Inter({ subsets: ["latin"] });


export default function TopicDetailLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
        <div className=" ">
          
        <div className='sticky top-0 z-50 h-[100px]'><TopHeader></TopHeader>
        </div>
        <div className='flex'>
        <div className='flex-none w-[26%]'><TopicSideBar/></div>
        <div className='grow'>
        
          { children }
        </div>
        </div>
        </div>
  );
}
