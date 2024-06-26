"use client"
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '@/components/ui/accordion'
import { useTopicList } from '@/lib/features/quiz/quizSlice'
import { lessonTestData } from '@/lib/lesson.test.data'
import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'

const QuizContent = () => {
  const [checkedLessons, setCheckedLessons] = useState<string[]>([]);
  const [accordIndex, setAccordIndex] = useState<string[]>([]);
  const { isSuccess, isLoading, errors, data } = useSelector((state: any) => state.subjects.topic);

  useEffect(() => {
    if (data?.Chapters?.length) {
      const indexes = data?.Chapters?.map((chapter: any, index: any) => index.toString());
      console.log("index:", indexes)
      setAccordIndex(indexes);
    }
  }, [data?.Chapters]);

  const dispatch = useDispatch();

  useEffect(() => {
    // eslint-disable-next-line react-hooks/rules-of-hooks
    dispatch(useTopicList({
      topics: checkedLessons
    }))
  }, [checkedLessons])




  const handleChapterChange = (event: React.ChangeEvent<HTMLInputElement>, lessons: any[]) => {
    const { checked } = event.target;

    if (checked) {
      // If checkbox is checked, add the lessons to the checkedLessons array
      setCheckedLessons(prevLessons => [...prevLessons, ...lessons.map(lesson => lesson)]);
    } else {
      // If checkbox is unchecked, remove the lessons from the checkedLessons array
      setCheckedLessons(prevLessons => prevLessons.filter(name => !lessons.map(lesson => lesson).includes(name)));
    }
  };



  const handleCheckboxChange = (event: any, lesson: string) => {
    if (event.target.checked) {
      // If checkbox is checked, add the lesson to the checkedLessons array
      setCheckedLessons([...checkedLessons, lesson]);
    } else {
      // If checkbox is unchecked, remove the lesson from the checkedLessons array
      setCheckedLessons(checkedLessons.filter(name => name !== lesson));
    }
  };

  console.log("topics", checkedLessons)

  return ( <Accordion
    type="multiple"
    defaultValue={["0", "1", "2"]}
    className="w-full bg-white shadow p-10"
  >
    {data?.Chapters?.map((chapter: any, index: number) => (
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
            {chapter.name}
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
  )
}

export default QuizContent