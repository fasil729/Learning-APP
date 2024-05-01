import React from 'react';

interface QuizResultProps {
  score: number;
  quizLength: number;
}

const QuizResult: React.FC<QuizResultProps> = ({ score, quizLength }) => {
  const percent = (score / quizLength) * 100;

  return (
    <div className='w-full md:w-8/12 xl:w-1/2 grid  grid-cols-2 gap-10 border-2 p-4 rounded-[10px]'>
      <div className={`
      w-32 flex 
      justify-center
       items-center
       h-32 rounded-full
        shadow-md border-2
        ${percent>=75&&'border-green-500 text-green-500'}

        ${percent < 65&&'border-red-500 text-red-500'}
        `}>
        <p className='text-4xl font-semibold'>{score}/{quizLength}</p>
      </div>

      {percent >= 85 ? (
        <div className={`text-4xl text-green-500 h-full items-center w-full flex justify-center font-bold leading-10`}>
          <h1>Excellent!!</h1>
        </div>
      ) : null}

      {75 <= percent && percent < 85 ? (
        <div className={`text-4xl text-green-500 w-full flex justify-center font-bold leading-10`}>
          <h1>Very Good!!</h1>
        </div>
      ) : null}

      {65 <= percent && percent < 75 ? (
        <div className={`text-4xl text-gray-500 w-full flex justify-center font-bold leading-10`}>
          <h1>More Enough!!</h1>
        </div>
      ) : null}

      {50 <= percent && percent < 65 ? (
        <div className={`text-4xl text-gray-500 w-full flex justify-center font-bold leading-10`}>
          <h1>Enough!!</h1>
        </div>
      ) : null}

      {35 <= percent && percent < 50 ? (
        <div className={`text-4xl text-red-500 w-full flex justify-center font-bold leading-10`}>
          <h1>Not Enough!!</h1>
        </div>
      ) : null}

      {percent < 35 ? (
        <div className={`text-4xl text-red-500 w-full flex justify-center font-bold leading-10`}>
          <h1>Poor!!</h1>
        </div>
      ) : null}
    </div>
  );
};

export default QuizResult;
