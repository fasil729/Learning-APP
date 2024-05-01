'use client';
import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
  } from "@/components/ui/accordion"
import LessonLink from "../course/lessonLink"
import { useDispatch, useSelector } from "react-redux";

// const data = {Chapters: [
//     {name: "chapter 1", lessons: ["lesson1", "lesson2", "lesson3"]},
//     {name: "chapter 1", lessons: ["lesson1", "lesson2", "lesson3"]},
//     {name: "chapter 1", lessons: ["lesson1", "lesson2", "lesson3"]},
//     {name: "chapter 1", lessons: ["lesson1", "lesson2", "lesson3"]},
//     {name: "chapter 1", lessons: ["lesson1", "lesson2", "lesson3"]}
// ]}



export const TopicSideBar  = () => {
    const {isSuccess, isLoading, errors, data} = useSelector((state: any) => state.subjects.topic);
    
    return <aside className="fixed overflow-y-scroll top-0 bg-gradient-to-b from-[#ddebf5] to-[#dedae0] left-0 w-[26%] h-screen px-4 py-8 border-white rtl:border-r-0 rtl:border-l dark:bg-gray-900 dark:border-gray-700 mt-[100px]">
            {/* <a href="#" className="mx-auto">
                <img className="w-auto h-6 sm:h-7 text-white" src="https://merakiui.com/images/full-logo.svg" alt=""/>
            </a> */}
            <Accordion type="single" collapsible className="w-full  mt-[100px]">
            {data.Chapters && data.Chapters.map((chapter: any, index: number) => {
              return <AccordionItem key={index} value={`${index}`} className="text-start">
                <AccordionTrigger className=' text-start text-[21px]'>{chapter.name}</AccordionTrigger>
                <AccordionContent className="pl-8">
                  {chapter.lessons.map((lesson: any, ind: number) => {
                    return <LessonLink  key={ind} name={`${lesson}`} url={`/topics/1/lesson/${lesson}`} status={lesson} />
                  })}
                  <button className="pl-8">Experiments</button>
                </AccordionContent>
              </AccordionItem>
            })}
          </Accordion>
    </aside>
}