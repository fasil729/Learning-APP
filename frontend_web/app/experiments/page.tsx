'use client';
import { useSelector } from "react-redux";
import Content from "../topics/[topicId]/course/content";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import LessonLink from "../topics/[topicId]/course/lessonLink";
import Link from "next/link";

export default function Page() {
    const {isLoading, data, errors}= useSelector((state: any) => state.experiment);
    const state = useSelector((state: any) => state);

    if (isLoading)
        return <div>Loading...</div>

    return (<div className="bg-[#e5e5e5] min-h-screen w-full flex xl:p-20 gap-2 lg:justify-center">
                <div className="w-full  flex   justify-center">
                    <div className="space-y-6  min-h-screen w-full">
                        <div className="bg-white flex  flex-col justify-center items-center ">
                            <h1 className='font-semibold text-lg'>{data.experiment_name ?? "Test Experiments"}</h1>
                            <div className="py-10 w-full flex flex-col">
                            {data && data.map((experiment: any, index: number) => {
                                return <Link key={index} href={`/experiments/${experiment.experiment_name}`} className="w-full"><button>{experiment.experiment_name}</button></Link>
                            })}
                            {/* <Accordion type="single" collapsible className="w-full">
                                {data && data.map((experiment: any, index: number) => {
                                return <AccordionItem key={index} value={`${index}`} className="data-[state=open]:border-none">
                                    <AccordionTrigger>{experiment.experiment_name}</AccordionTrigger>
                                    <AccordionContent>
                                    {chapter.lessons.map((lesson: any, ind: number) => {
                                        return <LessonLink key={ind} name={`${lesson}`} url={`/topics/1/lesson/${lesson}`} status={lesson} />
                                    })}
                                    </AccordionContent>
                                </AccordionItem>
                                })}
                            </Accordion> */}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
    )
}   