// lessonList.tsx
import React, { useState } from "react";
import { lessonTestData } from "@/lib/lesson.test.data";
import LessonLink from "./lessonLink";

export function LessonList() {
  const [checkedChapters, setCheckedChapters] = useState<boolean[]>(Array(lessonTestData.length).fill(false));
  const [checkedUnits, setCheckedUnits] = useState<Record<string, boolean>>({});

  const handleChapterCheckboxChange = (index: number) => {
    const newCheckedChapters = [...checkedChapters];
    newCheckedChapters[index] = !newCheckedChapters[index];
    setCheckedChapters(newCheckedChapters);

    const chapterLessons = lessonTestData[index].lessons;
    chapterLessons.forEach((lesson, ind) => {
      const unitName = `U${ind + 1}: ${lesson.lesson_name}`;
      setCheckedUnits((prevCheckedUnits) => ({
        ...prevCheckedUnits,
        [unitName]: newCheckedChapters[index],
      }));
    });
  };

  const handleUnitCheckboxChange = (unitName: string) => {
    setCheckedUnits((prevCheckedUnits) => ({
      ...prevCheckedUnits,
      [unitName]: !prevCheckedUnits[unitName],
    }));
  };

  return (
    <div className="w-full">
      {lessonTestData.map((chapter, index) => {
        const isChapterChecked = checkedChapters[index];

        return (
          <div key={index} className="mb-4">
            <div className='flex gap-2 items-center'> {/* Added checkbox */}
              <input
                type="checkbox"
                checked={isChapterChecked}
                onChange={() => handleChapterCheckboxChange(index)}
                className='text-[#4C6FFF]'
              />
              <h2 className="text-base font-medium mb-2"> {/* Decreased font size */}
                Chapter {index + 1} - {chapter.name}
              </h2>
            </div>
            <div className="ml-4">
              {chapter.lessons.map((lesson, ind) => {
                const unitName = `U${ind + 1}: ${lesson.lesson_name}`;
                const isChecked = checkedUnits[unitName];

                return (
                  <LessonLink
                    key={ind}
                    name={unitName}
                    onChange={() => handleUnitCheckboxChange(unitName)}
                    checked={isChecked}
                  />
                );
              })}
            </div>
          </div>
        );
      })}
    </div>
  );
}
