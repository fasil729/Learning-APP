"use client"
import { Button } from '@/components/ui/button'
import React from 'react'
import Quiz from "..//../public/quiz.jpg"
import Image from 'next/image'
import { useRouter } from 'next/navigation'

const StartQuizPage = () => {
  const router=useRouter();
  return (<div className=' h-screen w-full flex justify-center items-center'>

    <div className="flex  flex-col gap-4 justify-center items-center w-max">
      <div className=" min-w-max  overflow-hidden">
<Image height={400} width={500} src={Quiz} alt='quiz banner' className=' object-cover'/>
      </div>

      <h1 className=' font-semibold text-xl'>Get Started With Your Quizes</h1>
      <Button 
      onClick={()=>{
        router.push("/start-quiz/quiz");
      }}
      className='w-full 
      px-3 py-2 
      bg-[#153462] 
      text-lg
       text-white
       hover:bg-[#153462] 
       hover:bg-opacity-85 
       rounded-[10px]
       transition
       duration-300
       '>Start</Button>
    </div>
  </div>
  )
}

export default StartQuizPage