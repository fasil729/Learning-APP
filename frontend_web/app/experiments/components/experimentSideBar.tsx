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
    
    return <aside className="fixed overflow-y-scroll top-0 bg-gradient-to-b from-[#ddebf5] to-[#dedae0] left-0 w-[26%] h-screen px-4 py-8 border-white rtl:border-r-0 rtl:border-l dark:bg-gray-900 dark:border-gray-700 mt-[100px]">
              <Accordion type="single" collapsible className="w-full  mt-[100px]">
              {data && data.map((experiment: any, index: number) => {
                return <AccordionItem key={index} value={`${index}`} className="text-start">
                  <AccordionTrigger className=' text-start text-[21px]'>{experiment.experiment_name}</AccordionTrigger>
                </AccordionItem>
              })}
            </Accordion>
          </aside>
}