"use client"
import React from 'react'
import { useRouter } from 'next/navigation'
import QuizContent from './quiz-content'
import QuizPrompt from './quiz-prompt'
import { useSelector } from 'react-redux'
import { RootState } from '@/lib/store'
import QuizComponent from '@/components/quiz/quiz'


const StartQuizPage = () => {
  const router=useRouter();
  const quizData=useSelector((state:RootState)=>state.quiz.quizData)
  return (<>
  {quizData.length?<QuizComponent/>:<div className=' min-h-screen py-20 flex flex-col justify-center items-center justify-center items-center'>

  <div className="w-full md:w-11/12 lg:w-10/12 xl:w-8/12 ">
  <QuizPrompt/>
</div>
<div className="w-full md:w-11/12 lg:w-10/12 xl:w-8/12 ">
<QuizContent/>
</div>

  </div>}
  </>
  )
}

export default StartQuizPage