"use client"
import React, { useEffect, useState } from 'react'
import { LessonList } from './lessonList'
import { Button } from '../../../../components/ui/button';
import QuizGenerate from './quiz.generate'
const Content = () => {

  const [progress, setProgress] = useState(13)

 
  return (<div className='p-4 md:p-10 xl:p-20 space-y-10'>

    <div className="bg-[#F8F8FB] rounded-[15px] p-6 space-y-2">
      <p className='text-sm text-gray-500'>Course 1 of 6</p>
      <h4 className='text-black font-medium'>Course 1 - Introduction</h4>
      <div className="py-4 flex items-center gap-4 relative w-full ">
      <div className="bg-[#FFEFE0] w-full h-[6px] rounded-full overflow-hidden ">
        <div 
         style={{width:`${progress}%` }}
        className="bg-[#FF7800] h-full ">

        </div>
      </div>
      <Button  className='bg-[#4C6FFF] rounded-[5px] text-white hover:text-black'>Go to Unit 3</Button>
      </div>
      <p className='text-gray-800 text-sm'>2 completed units of 11 available</p>
    </div>
    <div className="">
    <QuizGenerate/>
     <LessonList/>
    </div>

  </div>
  )
}

export default Content