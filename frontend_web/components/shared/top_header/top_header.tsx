'use client';

import Link from "next/link";

export default function TopHeader() {
return <nav className="sticky bg-white h-[6.2rem] flex justify-between items-center px-14 gap-20 ">                
                <div className="flex ">
                    <a href="#" className="block  h-[80px] w-[80px] py-2 px-3 text-black text-[20px]  hover:bg-gray-100 rounded "><img src='/profile11.jpg'  className="h-[60px] w-[60px] rounded-full" alt='logo' /></a>
                    </div>
                <div  >
                   <div className="   flex flex-col ">
                    <ul className="font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                        <li >
                            <Link href="/topics" className="block  py-2 px-3 text-black text-[20px]  hover:bg-gray-100 rounded active:bg-gray-200">Topics</Link>
                        </li>
                        <li >
                            <Link href="#" className="block  py-2 px-3 text-black text-[20px]  hover:bg-gray-100 rounded ">About</Link>
                        </li>
                        <li >
                            <Link href="#" className="block  py-2 px-3 text-black text-[20px]  hover:bg-gray-100 rounded ">countact</Link>
                        </li>
                        <li>
                        <div className=" gap-9">
                     <div className="flex gap-4">
                        <div>
                             <img src="/profile1.png" alt="user" className="w-10 h-10 rounded-full"/>
                         </div>
                         <div>
                             <h3>Fasika</h3>
                         </div>
                     </div>
                 </div>
                        </li>
                    </ul>
                    </div>
                
                
                 </div>
            </nav>
    }