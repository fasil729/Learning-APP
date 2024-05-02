"use client"
import React, { useEffect, useState } from 'react'
import { LessonList } from './lessonList';
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
    <div className="">
    <QuizGenerate/>
    <div className='mb-5 text-[25px] font-semibold'>Content</div>
     <LessonList isLoading={props.isLoading} data={props.data} errors={props.data}/>
    </div>

  </div>
  )
}

export default Content