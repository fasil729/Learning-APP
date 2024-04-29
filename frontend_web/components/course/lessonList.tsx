/* eslint-disable react-hooks/rules-of-hooks */
"use client"
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion"
import { lessonTestData } from "@/lib/lesson.test.data"
import LessonLink from "./lessonLink"
import { useEffect, useState } from "react";
import { Checkbox } from "../ui/checkbox";
import { useDispatch } from "react-redux";
import { useTopicList } from "@/lib/features/quiz/quizSlice";

export function LessonList() {
const dispatch=useDispatch()
  const [selectedData, setSelectedData] = useState<any[]>([]);


  useEffect(() => {
    dispatch(useTopicList({ topics: selectedData }));
  }, [dispatch, selectedData]);

  const handleCheckboxChange = (event:any) => {
    const isChecked = event.target.checked;
    const value = event.target.value;

    if (isChecked) {
      setSelectedData((prev) => [...prev, value]);
    } else {
      setSelectedData((prev) => prev.filter((item) => item !== value));
    }
  };

  // function handleSubmit(event:any) {
  //   event.preventDefault();
  //   const formData = { Topics: selectedTopics };
  //   console.log(formData); // Do something with the form data
  // }

  console.log("topics",selectedData)
  return (
    <Accordion type="single" collapsible className="w-full">
      {lessonTestData.map((chapter,index)=>{
       return  <AccordionItem key={index} value={`${index}`} className="data-[state=open]:border-none">
        <AccordionTrigger>Course {index+1} - {chapter.name}</AccordionTrigger>
        <AccordionContent>
        {chapter.lessons.map((lesson,ind)=>{
          return  <div key={ind} className='w-full flex p-2 py-3 border-b  gap-2 '>
          <div className="">
          <input type="checkbox"  id={lesson?.id}  value={lesson?.lesson_name} onChange={handleCheckboxChange}/>
          </div> <LessonLink 
         
          key={ind}
           name={`U${ind+1}: ${lesson.lesson_name}`} 
           url={`/${lesson.id}`} 
           status={lesson.status}/>
           </div>
        })}
        </AccordionContent>
      </AccordionItem>
      })}
    </Accordion>
  )
}
