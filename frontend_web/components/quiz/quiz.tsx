"use client"

import React, { useEffect, useState } from 'react'
import { GiCheckMark } from "react-icons/gi";
import { Button } from '../ui/button';
import QuizResult from './quizResult';
import { useSelector } from 'react-redux';
import { RootState } from '@/lib/store';
import { GiCancel } from "react-icons/gi";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '../ui/accordion';
const QuizComponent = () => {
  const quizdata=useSelector((state:RootState)=>state.quiz.quizData)
  const [selected,setSelected] = useState<number>(0)
  const [selectedQuestion,setSelectedQuestion] = useState<any>(quizdata[0])
  useEffect(()=>{
    setSelectedQuestion(quizdata[selected])
  },[selected])


  const onNextButton=()=>{
    if(quizdata.length-1>selected){
      setSelected((prev)=>prev+=1)
    }
    else{
      setSelected(0)
    }
  }

  const onPrevButton=()=>{
    if(0<selected){
      setSelected((prev)=>prev-=1)
    }
    else{
      setSelected(quizdata.length-1)
    }
  }
 
  const onSelectNumber=(data:number)=>{
    setSelected(data)
  }

  
  const [isSelectedAll, setSelectedAll] = useState(false);
  const [score, setScore] = useState(0);
  const [selectedChoices, setSelectedChoices] = useState<any>({});
  const [answeredQuestions, setAnsweredQuestions] = useState<Set<number>>(new Set());

  
  const onSelectionChanged = (
    questionId: number,
    choiceIndex: string,
    is_answer: boolean
  ) => {
    setAnsweredQuestions((prev)=>prev.add(questionId))
  
    const updatedChoices = {
      ...selectedChoices,
      [questionId]: {
        
        choiceIndex,
        is_answer: is_answer,
      },
    };
    setSelectedChoices(updatedChoices);

    // Calculate score
    const newScore = Object.values(updatedChoices).reduce(
      (acc: number, choice: any) => (choice.is_answer ? acc + 1 : acc),
      0
    );
    setScore(newScore);
  };





  const onSubmit = () => {
    setSelectedAll(true);
  };

  

  return (<div className=" min-h-screen w-full flex xl:p-20 gap-2 lg:justify-center items-center bg-white p-4">

  <div className="w-full  lg:w-8/12  flex  flex-col gap-10 items-center justify-center px-10">

  {isSelectedAll? <div className="w-full space-y-10">
<QuizResult score={score} quizLength={quizdata?.length}/>
  
  <div className=" space-y-4  border-2 rounded-[10px] p-4 md:p-8 w-full">
   {quizdata?.map((quiz:any,index)=>{
    return <div className="" key={index}>

    
   <div className="flex w-full justify-between ">
     <div className="space-y-4">
       <h4 className='text-[#153462] text-lg font-semibold'>Question {index+1}</h4>
       <p>{quiz?.question}</p>
     </div>
    
  
   </div>

   <div className="p-2 mt-10">
     <p className="text-[#BDBDBD] py-2">
     <span className=''>OPTIONS :</span> <span className='text-sm'>Select any one of the following</span>
     </p>
     {quiz.options.map((choose:any,ind:number)=>{
       return <div key={ind} className={`flex 
       gap-3 items-center p-2
       m-2

       ${
        (
          selectedChoices[index + 1] &&
          choose.is_answer &&
          "bg-green-200 dark:bg-green-700") ||
        (isSelectedAll && choose.is_answer && "bg-green-200 dark:bg-green-700")
      } ${
        selectedChoices[index + 1]?.choiceIndex ===
          ind.toString() &&
        
        selectedChoices[index + 1]?.is_answer
          ? "bg-green-200 dark:bg-green-700"
          : 
            selectedChoices[index + 1]?.choiceIndex ===
              ind.toString()
          ? "bg-red-200 dark:bg-red-700"
          : ""
      }
       `}>
          <button 
          disabled={
            (selectedChoices[index + 1]) ||
            isSelectedAll
          }
          onClick={() => {
           onSelectionChanged(
             index+ 1,
             ind.toString(),
             choose.is_answer
           );
         }}
          className={`h-5 w-5 border border-[#54BAB9] 
          
          font-semibold 
          rounded-full
          ${
            selectedChoices[index + 1]?.choiceIndex ===
            ind.toString() &&
          
          selectedChoices[index + 1]?.is_answer
             ? "bg-[#54BAB9] text-white "
             :selectedChoices[index + 1]?.choiceIndex ===
             ind.toString()? "bg- transparent":""
         }
          `}>
           {
              selectedChoices[index + 1]?.choiceIndex ===
              ind.toString() &&
            
            selectedChoices[index + 1]?.is_answer?<GiCheckMark size={16}/>: selectedChoices[index + 1]?.choiceIndex ===
            ind.toString()?<GiCancel className='text-red-700' size={16}/>:""
           }
           
          </button>

          <p>{choose.text}</p></div>



     })}



{(selectedChoices[index + 1]) ||
  isSelectedAll ? (
    <Accordion type="single" collapsible className="w-full">
      <AccordionItem value="item-1" className="border-none">
        <AccordionTrigger className="hover:no-underline justify-start">
          
          <p className="text-[16px]">
          Explanation
          </p>
        </AccordionTrigger>
        <AccordionContent>
          <div className="px-2 py-4 bg-slate-100 dark:bg-gray-700">
          <p className="text-[14px]">{quiz.explanation}</p>
          </div>
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  ) : (
    ""
  )}
   </div>
    </div>
   })}
         
   
     
 
</div></div>: <div className=" space-y-4  border-2 rounded-[10px] p-4 md:p-8 w-full">
   
    
   <div className="flex w-full justify-between ">
     <div className="space-y-4">
       <h4 className='text-[#153462] text-lg font-semibold'>Question {selected+1}</h4>
       <p>{selectedQuestion?.question}</p>
     </div>
   
   </div>

   <div className="p-2 mt-10">
     <p className="text-[#BDBDBD] py-2">
     <span className=''>OPTIONS :</span> <span className='text-sm'>Select any one of the following</span>
     </p>
     <div className="">
      
     {selectedQuestion?.options.map((choose:any,ind:number)=>{
       return <div key={ind} className="flex 
       gap-3 items-center py-1">
          <button 
          onClick={() => {
           onSelectionChanged(
             selected+ 1,
             ind.toString(),
             choose.is_answer
           );
         }}
          className={`h-5 w-5 border border-[#54BAB9] 
          
          font-semibold 
          rounded-full
          ${
           selectedChoices[selected + 1]?.choiceIndex ===
           ind.toString()
             ? "bg-[#54BAB9] text-white "
             : ""
         }
          `}>
           {
             selectedChoices[selected + 1]?.choiceIndex ===
             ind.toString()?<GiCheckMark size={16}/>:""
           }
           
          </button>

          <p>{choose.text}</p></div>
     })}
     </div>
   </div>
         
   
     
 
</div>}

   <div className="flex gap-10">
    <Button 
    disabled={selected>0?false:true}
    onClick={onPrevButton} 
    className={`px-5
     py-2
     border
     border-black 
    rounded-[5px] 
    ${selected>0?'bg-black text-white hover:bg-black hover:bg-opacity-95 hover:text-white ':' cursor-not-allowed bg-opacity-50'}
    transition 
     duration-300  
     `}>
      Previous</Button>

    <Button 
     disabled={selected<quizdata.length-1?false:true}
    onClick={
      onNextButton
    }
    className={`px-5
    py-2
    border
    border-black 
   rounded-[5px] 
   ${selected<quizdata.length-1?'bg-black text-white hover:bg-black hover:bg-opacity-95 hover:text-white ':' cursor-not-allowed bg-opacity-50'}
   transition 
    duration-300  
    `}>
      Save&Next</Button>

      <Button 
    
    onClick={
    onSubmit
    }
    className={`px-5
    py-2
    border
   
   rounded-[5px] 

   transition 
    duration-300  
    `}>
      Submit</Button>
   </div>
   
  </div>

 {!isSelectedAll? <div className="hidden lg:block lg:w-4/12 ">

<div className="flex flex-col w-full justify-center items-center gap-4 border-2 rounded-[10px] p-4">
  <h4 className='text-[#153462] text-lg font-semibold'>Question Bank</h4>
  <div className="grid grid-cols-5 gap-6">
    {quizdata.map((quiz,index)=>{
      return <button 
      onClick={()=>onSelectNumber(index)}
      key={index} className={`text-lg font-medium p-1 
      rounded-[5px] text-center 
      border 
      bg-opacity-15
     
      ${answeredQuestions.has(index+1)?'bg-[#2FD790] text-[#2FD790] ':'bg-[#F0EEED] text-black'}
      `}>{index+1}</button>
    })}
  </div>
  <div className="p-4 w-full flex justify-end border-b-2">
    <p><span>Answered:</span> <span className='font-bold'>6/10</span></p>
  </div>

  <div className="grid grid-cols-2 gap-10">
 <div className="flex gap-6 items-center">
 <button 
     className='text-lg font-medium p-1 
      rounded-[5px] text-center text-[#2FD790] 
      border bg-[#2FD790] 
      bg-opacity-15 w-8'>{answeredQuestions.size}</button>
      <p>Answered</p>
 </div>

      <div className="flex gap-6 items-center">
      <button 
     className='text-lg font-medium p-1 
      rounded-[5px] text-center text-black
      border bg-[#F0EEED] 
      bg-opacity-15 w-8'>{quizdata.length-answeredQuestions.size}</button>
      <p>Not Answered</p>
      </div>
  </div>

  <div className="w-full flex justify-center mt-10 ">
    <Button className='shadow-md border py-2 px-4 rounded-[5px]'>Exit Quizes</Button>
  </div>

</div>

</div>:""}

</div>
  )
}

export default QuizComponent

