/* eslint-disable react-hooks/rules-of-hooks */
"use client";

import { useDispatch, useSelector } from "react-redux";
import { useEffect, useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { RootState } from "@/lib/store";

const ExamPrompt = () => {
  const [prompt, setPrompt] = useState("");
  const [expectedStudyHour, setExpectedStudyHour] = useState(5); // Changed quizNo to expectedStudyHour
  const dispatch = useDispatch();

  const topics = useSelector((state: RootState) => state.exam.topics);
  const [isClose, setClose] = useState<any>(topics?.length);
  useEffect(() => {
    setClose(topics?.length);
  }, [topics]);

  const examData = useSelector((state: RootState) => state.exam);

  const [ExamPrompt, { isLoading, data, error }] = useExamGenerateMutation(); // Changed QuizPrompt to ExamPrompt

  console.log("examData", data);
  console.log("examError", error);

  const onGenerateExam = async () => {
    await ExamPrompt({ prompt, expectedStudyHour }); // Pass prompt and expectedStudyHour to ExamPrompt
  };

  useEffect(() => {
    if (data) {
      dispatch(setExamData({ examData: data }));
    }
  }, [data]);

  return (
    <>
      <div className=" relative bg-white shadow">
        <div className="  w-full  m-2 p-4 space-y-4">
          <Input
            className=" p-1 "
            type="text"
            placeholder="Write exam prompt..."
            onChange={(event) => setPrompt(event.target.value)}
          />

          <Input
            className=" p-1 "
            type="number"
            placeholder="Add expected study hours..."
            onChange={(event) => setExpectedStudyHour(Number(event.target.value))}
          />

          <div className="p-2 flex justify-end">
            <Button
              className=" col-span-2 bg-blue-500 hover:bg-blue-600"
              onClick={onGenerateExam}
              disabled={isLoading}
            >
              {isLoading ? "Loading..." : "Generate"}
            </Button>
          </div>
        </div>
      </div>
    </>
  );
};

export default ExamPrompt;