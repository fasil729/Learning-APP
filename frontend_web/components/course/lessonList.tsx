import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion"
import { lessonTestData } from "@/lib/lesson.test.data"
import LessonLink from "./lessonLink"

export function LessonList() {
  return (
    <Accordion type="single" collapsible className="w-full">
      {lessonTestData.map((chapter,index)=>{
       return  <AccordionItem key={index} value={`${index}`} className="data-[state=open]:border-none">
        <AccordionTrigger>Course {index+1} - {chapter.name}</AccordionTrigger>
        <AccordionContent>
        {chapter.lessons.map((lesson,ind)=>{
          return  <LessonLink key={ind} name={`U${ind+1}: ${lesson.lesson_name}`} url={`/${lesson.id}`} status={lesson.status}/>
        })}
        </AccordionContent>
      </AccordionItem>
      })}
    </Accordion>
  )
}
