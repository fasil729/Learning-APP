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


interface Props {
  isLoading: boolean,
  data: any,
  errors: any
}
export function LessonList(props: Props) {

  if (props.data.isLoading)
    return <div>...Loading</div>

   return (<Accordion type="single" collapsible className="w-full">
            {props.data.Chapters && props.data.Chapters.map((chapter: any, index: number) => {
              return <AccordionItem key={index} value={`${index}`} className="data-[state=open]:border-none">
                <AccordionTrigger className= 'text-[21px]'>{chapter.name}</AccordionTrigger>
                <AccordionContent>
                  {chapter.lessons.map((lesson: any, ind: number) => {
                    return <LessonLink key={ind} name={`${lesson}`} url={`/topics/1/lesson/${lesson}`} status={lesson} />
                  })}
                </AccordionContent>
              </AccordionItem>
            })}
            {/* {isSuccess && data.chapters.map((chapter, index) => <p key={index}>{chapter.name}</p>)} */}
            {/* {topic && topic.data && topic.data.chapters.map((chapter: Chapter, index: number) => <p key={index}>{chapter.name}</p>)} */}
          </Accordion>
          )
}
