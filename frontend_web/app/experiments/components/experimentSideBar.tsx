'use client';
import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
  } from "@/components/ui/accordion"
import LessonLink from "../../topics/[topicId]/course/lessonLink";
import { useDispatch, useSelector } from "react-redux";
import { getExperimentsForChapterAction } from "@/lib/features/experiments/experimentSlice";

 

export const ExperimentSideBar  = () => {
    const {isLoading, data, errors}= useSelector((state: any) => state.experiment);
    const dispatch = useDispatch();
    
    return <aside className="fixed top-0 bg-gradient-to-b from-[#418CD1] to-[#9C41D1] left-0 w-[18%] h-screen px-4 py-8 border-white rtl:border-r-0 rtl:border-l dark:bg-gray-900 dark:border-gray-700">
            <a href="#" className="mx-auto">
                <img className="w-auto h-6 sm:h-7 text-white" src="https://merakiui.com/images/full-logo.svg" alt=""/>
            </a>
            <Accordion type="single" collapsible className="w-full ml-auto mr-auto">
            {data && data.map((experiment: any, index: number) => {
              return <AccordionItem key={index} value={`${index}`} className="data-[state=open]:border-none">
                <AccordionTrigger>{experiment.experiment_name}</AccordionTrigger>
                {/* <AccordionContent className="pl-8">
                  {chapter.lessons.map((lesson: any, ind: number) => {
                    return <LessonLink key={ind} name={`${lesson}`} url={`/topics/1/lesson/${lesson}`} status={lesson} />
                  })}
                </AccordionContent> */}
              </AccordionItem>
            })}
          </Accordion>
    </aside>
}