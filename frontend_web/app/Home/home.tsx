"use client";

import Link from 'next/link';

export default function Home() {
  return (
    <main className="w-screen h-screen relative">
      <div
        className="flex items-center  w-full h-full  bg-cover bg-center"
        style={{ backgroundImage: "url(/images2.jpeg)"   }}
      >
        <div className="pl-20 md:pl-40 pb-56 md:pb-20 flex flex-col  gap-5 z-[10] max-w-[750px]">
          <h1 className="text-[50px] text-white font-semibold">
            Well Come To Brilliant Wep App 
          </h1>
          <h3 className="text-white text-[25px] hidden md:block">
          Explore the World of Learning. Start your journey today and
                unlock endless possibilities for growth and development.          </h3>
          <div className="flex-col md:flex-row hidden md:flex gap-5">
            <Link
              href="/login"
              className="rounded-[20px] group relative bg-blue-700 hover:bg-blue-400 px-5 py-3 text-lg text-white max-w-[200px]"
            >
              Login
            </Link>
            <Link
              href="/register"
              className="rounded-[20px] group relative bg-transparent px-5 border border-white py-3 text-lg text-white max-w-[200px]"
            >
              <div className="absolute rounded-[20px] z-[1] bg-white inset-0 opacity-0 group-hver:opacity-20" />
              Sign Up
            </Link>
            
          </div>
        </div>
      </div>

      
    

    
    </main>
  );
}
