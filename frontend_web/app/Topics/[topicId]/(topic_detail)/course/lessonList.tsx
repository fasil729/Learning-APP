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
import { getExperimentsForChapterAction } from "@/lib/features/experiments/experimentSlice";
import { useRouter } from "next/navigation";
import ExperimentButton from "@/components/ui/experimentButton";


interface Props {
  isLoading: boolean,
  data: any,
  errors: any
}
export function LessonList(props: Props) {
  const router = useRouter();

  if (props.data.isLoading)
    return <div>...Loading</div>
  function onExperimentRequest(chapterName: string) {
    dispatch(getExperimentsForChapterAction(chapterName));
    router.push(`/experiments?chapterName=${chapterName}`)
  }

   return (<Accordion type="single" collapsible className="w-full">
            {props.data.Chapters && props.data.Chapters.map((chapter: any, index: number) => {
              return <AccordionItem key={index} value={`${index}`} className="data-[state=open]:border-none">
                <AccordionTrigger className= 'text-[21px]'>{chapter.name}</AccordionTrigger>
                <AccordionContent>
                  {chapter.lessons.map((lesson: any, ind: number) => {
                    return <LessonLink key={ind} name={`${lesson}`} url={`/topics/1/lesson/${lesson}`} status={lesson} />
                  })}
                  <div className="pl-[2%]">
                    <ExperimentButton message="Experiments" clickEvent={onExperimentRequest} chapterName={chapter.name} />
                  </div>
                </AccordionContent>
              </AccordionItem>
            })}
          </Accordion>
          )
}
function dispatch(arg0: any) {
  throw new Error("Function not implemented.");
}

