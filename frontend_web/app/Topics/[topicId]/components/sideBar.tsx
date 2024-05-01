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
    
    return <aside className="fixed top-0 bg-gradient-to-b from-[#418CD1] to-[#9C41D1] left-0 w-[18%] h-screen px-4 py-8 border-white rtl:border-r-0 rtl:border-l dark:bg-gray-900 dark:border-gray-700">
            <a href="#" className="mx-auto">
                <img className="w-auto h-6 sm:h-7 text-white" src="https://merakiui.com/images/full-logo.svg" alt=""/>
            </a>
            <Accordion type="single" collapsible className="w-full ml-auto mr-auto">
            {data.Chapters && data.Chapters.map((chapter: any, index: number) => {
              return <AccordionItem key={index} value={`${index}`} className="data-[state=open]:border-none">
                <AccordionTrigger>{chapter.name}</AccordionTrigger>
                <AccordionContent className="pl-8">
                  {chapter.lessons.map((lesson: any, ind: number) => {
                    return <LessonLink key={ind} name={`${lesson}`} url={`/topics/1/lesson/${lesson}`} status={lesson} />
                  })}
                  <button className="pl-8">Experiments</button>
                </AccordionContent>
              </AccordionItem>
            })}
          </Accordion>
    </aside>
}