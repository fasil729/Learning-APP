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
        <div className="flex ">
          <div className='flex-none w-[26%]'> 
        <TopicSideBar/>
        </div>
        <div className='grow'>
        <TopHeader></TopHeader>
          { children }
        </div>
        </div>
  );
}
