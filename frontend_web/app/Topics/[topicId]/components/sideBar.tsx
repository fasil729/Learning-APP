'use client';
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion"
import LessonLink from "../course/lessonLink"
import { useDispatch, useSelector } from "react-redux";
import { getExperimentsForChapterAction } from "@/lib/features/experiments/experimentSlice";
import { useRouter } from "next/navigation";
import ExperimentButton from "@/components/ui/experimentButton";
import { Button } from "@/components/ui/button";


export const TopicSideBar = () => {
  const { isSuccess, isLoading, errors, data } = useSelector((state: any) => state.subjects.topic);
  console.log(data)
  const dispatch = useDispatch();
  const router = useRouter();


  function onExperimentRequest(chapterName: string) {
    dispatch(getExperimentsForChapterAction(chapterName));
    router.push(`/experiments?chapterName=${chapterName}`);
    console.log("chapterName", chapterName)
  }

  function onGenerateQuiz() {
    if (data && data.Chapters) {
      router.push("/start-quiz");
    }
  }
  
  function onExamGenerate() {
   
  }

  return <aside className="fixed overflow-y-scroll top-0 bg-gradient-to-b from-[#ddebf5] to-[#dedae0] left-0 w-[26%] h-screen px-4 border-white rtl:border-r-0 rtl:border-l dark:bg-gray-900 dark:border-gray-700 mt-[100px]">
    <Accordion type="single" collapsible className="w-full  mt-[100px]">
      {data.Chapters && data.Chapters.map((chapter: any, index: number) => {
        return <AccordionItem key={index} value={`${index}`} className="text-start">
          <AccordionTrigger className=' text-start text-[21px]'>{chapter.name}</AccordionTrigger>
          <AccordionContent className="pl-8">
            {chapter.lessons.map((lesson: any, ind: number) => {
              return <LessonLink key={ind} name={`${lesson}`} url={`/topics/1/lesson/${lesson}`} status={lesson} />
            })}
            <div className="pl-[2%]">
              <ExperimentButton message="Experiments" onClick={() => onExperimentRequest(chapter.name)} />
            </div>
          </AccordionContent>
        </AccordionItem>
      }
      )}
      <div className="flex flex-col w-1/2 gap-3 m-auto">
        <Button onClick={() => onGenerateQuiz()}>Generate Quiz</Button>
        <Button onClick={() => onExamGenerate()}>Exam Preparation</Button>
      </div>
    </Accordion>
  </aside>
}