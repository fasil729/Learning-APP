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
        <div className="flex">
          
        <TopicSideBar/>
        <div className='ml-[21%]'>
        <TopHeader></TopHeader>
          { children }
        </div>
        </div>
  );
}
