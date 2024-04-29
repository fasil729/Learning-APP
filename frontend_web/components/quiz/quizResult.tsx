import React from 'react'


interface QuizResultProps{
  score: number;
  quizLength: number;
}
const QuizResult:React.FC<QuizResultProps> = ({score,quizLength}) => {
  return (<div className=' w-full md:w-8/12 xl:w-1/2 grid grid-cols-1 md:grid-cols-2  border-2 p-4 rounded-[10px]'>
  <div className="">
    <div className="w-32 flex justify-center items-center h-32 rounded-full shadow-lg border-2">
       <p className='text-2xl font-semibold'>{score}/{quizLength}</p>
    </div>
  </div>
  
  <div className="">
    <p> 
      <span>Test:</span>
      <span className=' font-medium'>General Knowlege Quiz</span>
    </p>
    <p> 
      <span>Test:</span>
      <span className=' font-medium'>General Knowlege Quiz</span>
    </p>
    <p> 
      <span>Test:</span>
      <span className=' font-medium'>General Knowlege Quiz</span>
    </p>
    <p> 
      <span>Test:</span>
      <span className=' font-medium'>General Knowlege Quiz</span>
    </p>
    <p> 
      <span>Test:</span>
      <span className=' font-medium'>General Knowlege Quiz</span>
    </p>
  </div>
  </div>
  )
}

export default QuizResult