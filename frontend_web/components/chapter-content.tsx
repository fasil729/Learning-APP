import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useRouter } from "next/navigation";



const ExamPrepContent = () => {
  const router = useRouter();
  const [checkedLessons, setCheckedLessons] = useState<string[]>([]);
  const [accordIndex, setAccordIndex] = useState<string[]>([]);
  const [prompt, setPrompt] = useState("");
  const [expectedReadTime, setExpectedReadTime] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const dispatch = useDispatch();

  const { isSuccess, errors, data } = useSelector(
    (state: any) => state.subjects.topic
  );
  const chaptersData = data.Chapters;

  useEffect(() => {
    if (chaptersData?.length) {
      const indexes = chaptersData.map((_, index: any) => index.toString());
      setAccordIndex(indexes);
    }
  }, [chaptersData]);

  const handleChapterChange = (
    event: React.ChangeEvent<HTMLInputElement>,
    lessons: any[]
  ) => {
    const { checked } = event.target;

    if (checked) {
      setCheckedLessons((prevLessons) => [
        ...prevLessons,
        ...lessons.map((lesson) => lesson),
      ]);
    } else {
      setCheckedLessons((prevLessons) =>
        prevLessons.filter(
          (name) => !lessons.map((lesson) => lesson).includes(name)
        )
      );
    }
  };

  const handleCheckboxChange = (event: any, lesson: string) => {
    if (event.target.checked) {
      setCheckedLessons([...checkedLessons, lesson]);
    } else {
      setCheckedLessons(checkedLessons.filter((name) => name !== lesson));
    }
  };

  const generateExamPrep = async () => {
    try {
      setIsLoading(true);
      const queryParams = new URLSearchParams({
        prompt: prompt,
        readTime: expectedReadTime,
        topics: JSON.stringify(checkedLessons),
      }).toString();
      router.push(`examPrep?${queryParams}`);
    } catch (error) {
      console.error("Failed to generate exam prep:", error);
    } finally {
      setIsLoading(false);
    }
  };


  return (
    <>
      <div className=" relative bg-white shadow">
        <div className="w-full m-2 p-4 space-y-4">
          <Input
            className="p-1"
            type="text"
            placeholder="Write exam prompt..."
            value={prompt}
            onChange={(event) => setPrompt(event.target.value)}
          />

          <Input
            className="p-1"
            type="number"
            placeholder="Add expected read time..."
            value={expectedReadTime}
            onChange={(event) => setExpectedReadTime(event.target.value)}
          />

          <div className="p-2 flex justify-end">
            <Button
              className="col-span-2 bg-blue-500 hover:bg-blue-600"
              onClick={generateExamPrep}
              disabled={isLoading}
            >
              {isLoading ? "Loading..." : "Generate Exam Prep"}
            </Button>
          </div>
        </div>
      </div>

      <Accordion
        type="multiple"
        defaultValue={["0", "1", "2"]}
        className="w-full bg-white shadow p-10"
      >
        {chaptersData.map((chapter: any, index: number) => (
          <AccordionItem
            key={index}
            value={`${index}`}
            className="m-2 p-2 data-[state=open]:border-none"
          >
            <div className="flex gap-2">
              <input
                type="checkbox"
                name={`lesson_${index}`}
                checked={chapter.lessons.every((lesson: any) =>
                  checkedLessons.includes(lesson)
                )}
                onChange={(event) =>
                  handleChapterChange(event, chapter.lessons)
                }
                className="h-4 mt-5 w-4"
              />
              <AccordionTrigger
                className={`${
                  chapter.lessons.every((lesson: any) =>
                    checkedLessons.includes(lesson)
                  )
                    ? "text-blue-500"
                    : ""
                }`}
              >
                Chapter {index + 1} - {chapter.name}
              </AccordionTrigger>
            </div>
            <AccordionContent className="px-4 space-y-3">
              {chapter.lessons.map((lesson: any, ind: number) => (
                <div
                  key={ind}
                  className="hover:text-[#4C6FFF] flex gap-2 transition duration-300"
                >
                  <input
                    type="checkbox"
                    name={`lesson_${ind}`}
                    checked={checkedLessons.includes(lesson)}
                    onChange={(event) => handleCheckboxChange(event, lesson)}
                    className="h-4 w-4"
                  />
                  <div className="flex gap-6">
                    <p
                      className={`pl-2 ${
                        checkedLessons.includes(lesson) ? "text-blue-500" : ""
                      }`}
                    >
                      {lesson}
                    </p>
                  </div>
                </div>
              ))}
            </AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>
    </>
  );
};

export default ExamPrepContent;
