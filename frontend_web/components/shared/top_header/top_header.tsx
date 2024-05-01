'use client';

export default function TopHeader() {
return <nav className="bg-white h-[6.2rem] flex justify-between items-center px-14 gap-20 ">                
                <div className="flex ">
                    <a href="#" className="block  h-[80px] w-[80px] py-2 px-3 text-black text-[20px]  hover:bg-gray-100 rounded "><img src='/profile1.png' alt='logo' /></a>
                    </div>
                <div  >
                   <div className="   flex flex-col ">
                    <ul className="font-medium flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                        {/* <li>
                            <a href="#" className="block py-2 px-3 text-white bg-blue-700 rounded md:bg-transparent md:text-blue-700 md:p-0 dark:text-white md:dark:text-blue-500" aria-current="page">Home</a>
                        </li> */}
                        <li >
                            <a href="/Topics" className="block  py-2 px-3 text-black text-[20px]  hover:bg-gray-100 rounded active:bg-gray-200">Topics</a>
                        </li>
                        <li >
                            <a href="#" className="block  py-2 px-3 text-black text-[20px]  hover:bg-gray-100 rounded ">About</a>
                        </li>
                        <li >
                            <a href="#" className="block  py-2 px-3 text-black text-[20px]  hover:bg-gray-100 rounded ">countact</a>
                        </li>
                        {/* <li>
                            <a href="#" className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">Exam Preparation</a>
                        </li> */}
                        {/* <li>
                            <a href="#" className="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">Contact</a>
                        </li> */}
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