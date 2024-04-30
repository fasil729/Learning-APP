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
import { useSelector } from 'react-redux';
const CourseDetailComponent = () => {
  const {isSuccess, isLoading, errors, data} = useSelector((state: any) => state.subjects.topic);
  console.log(data);

  console.log(data);
  return (<div className="bg-[#e5e5e5] min-h-screen w-full flex xl:p-20 gap-2 lg:justify-center">

    <div className="w-full  flex   justify-center">

      <div className="space-y-6  min-h-screen w-full">

        {/* <div className="space-y-2">
          <h1 className='font-semibold text-lg'>Discover</h1>
          <div className="flex gap">
            <p className='text-gray-600 flex'><span>Course</span> <span className='p-2'><FaChevronRight size={12} /></span></p>
            <h2 className='text-gray-800 font-semibold'>Artificial Intelligence & Machine Learning</h2>
          </div>
        </div> */}

        <div className="bg-white flex  flex-col justify-center items-center ">
          {/* <div className="flex justify-center p-4 w-full md:w-11/12 xl:w-10/12" >
            <VideoComponent />
          </div> */}

          <h1 className='font-semibold text-lg'>{data.Subject?.Name ?? "Test topic"}</h1>
          {/* <p className='text-gray-600 text-sm'>By  <span className='text-gray-800 text-base'>Simon Shaw</span>, Illustrator and 3D designer</p> */}

          <div className="py-10 w-full">
            <Content isLoading={isLoading} errors={errors} data={data}/>
            {/* <Tabs defaultValue="information" className="w-full">
              <TabsList className="grid w-full grid-cols-2 md:grid-cols-4 gap-2 px-6">
                <TabsTrigger value="information" className="data-[state=active]:bg-[#FF7800] data-[state=active]:text-white rounded-full">Information</TabsTrigger>
                <TabsTrigger value="content" className="data-[state=active]:bg-[#FF7800] data-[state=active]:text-white rounded-full">Content</TabsTrigger>

                <TabsTrigger disabled value="community" className='text-gray-500 font-normal ursor-not-allowed flex gap-2'>
                  <CiLock />
                  <p>Community</p>
                </TabsTrigger>

                <TabsTrigger disabled value="students" className='text-gray-500 font-normal  cursor-not-allowed flex gap-2'>
                  <CiLock />
                  <p>Students</p>
                </TabsTrigger>
              </TabsList>


              <TabsContent value="information" className='mt-10  md:mt-0'>
                <Information />
              </TabsContent>

              <TabsContent value="content" className='mt-10 md:mt-0'>
                <Content />
              </TabsContent>

              <TabsContent value="community" className='mt-10 md:mt-0'>
                <Community />
              </TabsContent>

              <TabsContent value="students" className='mt-10 md:mt-0'>
                <h1>students</h1>
              </TabsContent>

            </Tabs> */}
          </div>


        </div>
      </div>

    </div>

    {/* <div className="hidden lg:block t lg:w-4/12 relative xl:p-20">
      <div className="fixed flex justify-center">
        <CourseDescribe />
      </div>
    </div> */}

  </div>
  )
}

export default CourseDetailComponent