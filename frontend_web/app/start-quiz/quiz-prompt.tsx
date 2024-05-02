/* eslint-disable react-hooks/rules-of-hooks */
"use client"

import { useDispatch, useSelector } from "react-redux";

import { useEffect, useState } from "react";

import { setQuizData, useGenerateQuiz } from "@/lib/features/quiz/quizSlice";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useQuizGenerateMutation } from "@/lib/features/quiz/quizApi";
import { RootState } from "@/lib/store";




const QuizPrompt = () => {
  const [prompt, setPrompt] = useState("")
  const [quizNo, setQuizNo] = useState(5)
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(useGenerateQuiz({
      prompt: prompt,
      number_of_quizzes: quizNo
    }))
  }, [dispatch, prompt, quizNo]);



  const topis = useSelector((state: RootState) => state.quiz.topics)
  const [isClose, setClose] = useState<any>(topis?.length);
  useEffect(() => {
    setClose(topis?.length)
  }, [topis])
  const quizdata = useSelector((state: RootState) => state.quiz)


  const [QuizPrompt, { isError, isSuccess, data, isLoading, error }] = useQuizGenerateMutation();

  console.log("quizData", data);
  console.log("quizError", error);

  const onGenerateQuiz = async () => {
    await QuizPrompt(quizdata);
  }
  useEffect(() => {
    if (data) {
      dispatch(setQuizData({
        quizData: data
      }))
    }
  }, [data])
  return (<>
    <div className=" relative bg-white shadow">
      <div className="  w-full  m-2 p-4 space-y-4">
        <Input className=" p-1 " type="number" placeholder="Add Quiz.No"
          onChange={(event) => setQuizNo(Number(event.target.value))} />

        <Input className=" p-1 " type="text" placeholder="write quiz prompts..."
          onChange={(event) => setPrompt(event.target.value)} />

        <div className="p-2 flex justify-end">
          <Button className=" col-span-2 bg-blue-500 hover:bg-blue-600" onClick={onGenerateQuiz}
            disabled={isLoading}>{isLoading ? "Loading..." : "Generate"}</Button>
        </div>

      </div>
    </div>
  </>);
}

export default QuizPrompt;