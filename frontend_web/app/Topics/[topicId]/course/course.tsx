'use client';
import React from 'react'
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "@/components/ui/tabs"
import { FaChevronRight } from "react-icons/fa6";
import Information from './information';
import Content from './content';
import Community from './community';
import VideoComponent from '../../../../components/video';
import { CiLock } from "react-icons/ci";
import CourseDescribe from './courseDescribe';
import { useDispatch, useSelector } from 'react-redux';


const CourseDetailComponent = () => {
  const {isSuccess, isLoading, errors, data} = useSelector((state: any) => state.subjects.topic);
  
  return (<div className="bg-[#e5e5e5] min-h-screen w-full flex xl:p-20 gap-2 lg:justify-center">

    <div className="w-full  flex   justify-center">

      <div className="space-y-6  min-h-screen w-full">

        <div className="bg-white flex  flex-col justify-center items-center ">

          <h1 className='font-semibold text-xl'>{data.Subject?.Name.toUpperCase() ?? "Test topic"}</h1>

          <div className="py-10 w-full">
            <Content isLoading={isLoading} errors={errors} data={data}/>
            
          </div>


        </div>
      </div>

    </div>

  </div>
  )
}

export default CourseDetailComponent