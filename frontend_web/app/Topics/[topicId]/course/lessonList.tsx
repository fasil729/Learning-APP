import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion"
import { lessonTestData } from "@/lib/lesson.test.data"
import LessonLink from "./lessonLink"
import { useSelector } from "react-redux"
import { Chapter, TopicState } from "@/types/topic";

export function LessonList() {
  
  const topic = useSelector((state: TopicState) => state.topic);
  console.log(topic);
  if (topic.isLoading)
    return <div>...Loading</div>

   return (<Accordion type="single" collapsible className="w-full">
            {lessonTestData.map((chapter, index) => {
              return <AccordionItem key={index} value={`${index}`} className="data-[state=open]:border-none">
                <AccordionTrigger>Course {index + 1} - {chapter.name}</AccordionTrigger>
                <AccordionContent>
                  {chapter.lessons.map((lesson, ind) => {
                    return <LessonLink key={ind} name={`U${ind + 1}: ${lesson.lesson_name}`} url={`/${lesson.id}`} status={lesson.status} />
                  })}
                </AccordionContent>
              </AccordionItem>
            })}
            {topic.data && topic.data.chapters.map((chapter: Chapter, index: number) => <p key={index}>{chapter.name}</p>)}
          </Accordion>
          )
}
