"use client"
import React from 'react';
import { LessonList } from './lessonList';

const ChooseChapters = () => {
    return (
        <div className='p-4 md:p-10 xl:p-20 space-y-10 bg-primaryShade flex flex-col justify-center items-center'>
            <h2 className="text-l font-bold mb-4">You can generate your Exam as you need by filling the following questions:</h2>
            <div className="bg-white rounded-md shadow-md p-6 w-80 sm:w-96">
                <div className="mb-3">
                    <label htmlFor="readingTime" className="block text-sm font-medium text-titleColor">Reading Time</label>
                    <input
                        type="text"
                        id="readingTime"
                        name="readingTime"
                        className="mt-1 block w-full px-4 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring focus:ring-mainColor focus:border-mainColor"
                        placeholder="Enter your potential reading time"
                    />
                </div>
                <div>
                    <label htmlFor="suggestion" className="block text-sm font-medium text-titleColor">Suggestion</label>
                    <textarea
                        id="suggestion"
                        name="suggestion"
                        rows={5}
                        className="mt-1 block w-full px-4 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring focus:ring-mainColor focus:border-mainColor"
                        placeholder="Add your suggestion to us!"
                    />
                </div>
            </div>
            
            <LessonList/>
            <button className="mt-7 px-6 py-3 bg-mainColor text-white font-medium rounded-md shadow-md hover:bg-secondayColor focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-mainColor">Generate Exam</button>
        </div>
    )
}

export default ChooseChapters;
