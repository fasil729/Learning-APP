"use client"
import React from 'react'
import { useRouter } from 'next/navigation'
import { useSelector } from 'react-redux'
import { RootState } from '@/lib/store'
import QuizComponent from '@/components/quiz/quiz'
import ChapterContent from '@/components/chapter-content'


const StartExamPrep = () => {
  const router=useRouter();
  
  return (<>
    <ChapterContent/>
  </>
  )
}

export default StartExamPrep