"use client"
import React, { useEffect, useState } from 'react'
import { LessonList } from './lessonList'
import { Button } from '../../../../components/ui/button';
import QuizGenerate from './quiz.generate'
import { useSelector } from 'react-redux';

interface Props {
  isLoading: boolean,
  errors: any,
  data: any
}
const Content = (props: Props) => {

  const [progress, setProgress] = useState(13)
  const {isSuccess, isLoading, errors, data} = useSelector((state: any) => state.subjects.topic);
  console.log(data);

 
  return (<div className='px-[100px] py-1 space-y-1'>
{/* 
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
    </div> */}
    <div className="">
    <QuizGenerate/>
    <div className='mb-5 text-[25px] font-semibold'>Content</div>
     <LessonList isLoading={props.isLoading} data={props.data} errors={props.data}/>
    </div>

  </div>
  )
}

export default Content