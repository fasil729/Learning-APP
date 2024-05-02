"use client"
import React from 'react'
import { useRouter } from 'next/navigation'
import { useSelector } from 'react-redux'
import { RootState } from '@/lib/store'
import QuizComponent from '@/components/quiz/quiz'
import ChapterContent from '@/components/chapter-content'


const StartExamPrep = () => {
  const router=useRouter();
  
  return (<div className='w-full md:w-11/12 lg:w-10/12 xl:w-8/12 m-auto'>
    <ChapterContent/>
  </div>
  )
}

export default StartExamPrep