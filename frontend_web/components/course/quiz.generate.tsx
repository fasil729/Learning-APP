/* eslint-disable react-hooks/rules-of-hooks */
"use client"

import { useDispatch, useSelector } from "react-redux";
import { RooState } from "@/lib/store";
import { useEffect, useState } from "react";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { useGenerateQuiz } from "@/lib/features/quiz/quizSlice";
import { useQuizGenerateMutation } from "@/lib/features/quiz/quizApi";


const QuizGenerate = () => {
  const[prompt,setPrompt] = useState("")
  const[quizNo,setQuizNo] = useState(5)
const dispatch=useDispatch();

useEffect(() => {
  dispatch(useGenerateQuiz({
    prompt:prompt,
    number_of_quizzes:quizNo
  }))
},[dispatch, prompt, quizNo]);


 
    const topis=useSelector((state:RooState)=>state.quiz.topics)
  const [isClose,setClose]=useState<any>(topis?.length);
useEffect(()=>{
  setClose(topis?.length)
},[topis])
const quizdata =useSelector((state:RooState)=>state.quiz)
  
 
  const [quizGenerate,{isError,isSuccess,data,isLoading,error}]=useQuizGenerateMutation();

  console.log("quizData",data);
  console.log("quizError",error);

  const onGenerateQuiz=async() => {
    await quizGenerate(quizdata);
  }
  
  return ( <>

{isClose?  
<div className=" relative">
<div className="bg-sky-50 shadow-lg border rounded-lg  fixed w-full lg:w-1/2  top-[30%] m-2 p-4 space-y-2">
  <Input type="number" placeholder="Add Quiz.No" onChange={(event)=>setQuizNo(Number(event.target.value))}/>
 <div className="p-4 w-full grid grid-cols-12 gap-2">
 <Input className=" col-span-10" type="text" placeholder="write quiz prompts..." onChange={(event)=>setPrompt(event.target.value)}/>
 <Button className=" col-span-2" onClick={onGenerateQuiz} disabled={isLoading}>{isLoading?"Loading...":"Generate"}</Button>
 </div>
</div>
</div>:""
}
</>);
}
 
export default QuizGenerate;