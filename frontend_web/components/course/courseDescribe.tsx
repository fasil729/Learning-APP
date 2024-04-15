import React from 'react'
import { IconType } from 'react-icons'
import { PiUploadSimple } from "react-icons/pi";
import { PiStudentLight } from "react-icons/pi";
import { MdOutlinePlayLesson } from "react-icons/md";
import { LiaBookSolid } from "react-icons/lia";
import { GrDocumentDownload } from "react-icons/gr";
import { IoIosPhonePortrait } from "react-icons/io";
import { HiOutlineSpeakerWave } from "react-icons/hi2";
import { FiBarChart } from "react-icons/fi";
import { Button } from '../ui/button';
interface ContentListProps{
  icon:IconType;
  name:string;
}
const ContentList:React.FC<ContentListProps>=({icon:Icon,name})=>{
  return(<div className='text-xs text-gray-500 flex gap-2 w-full p-2'>
    <Icon size={16}/>
    <p className=''>{name}</p>

  </div>)

}
const CourseDescribe = () => {
  return (
    <div className='flex flex-col justify-center items-center gap-10 p-4'>
      
     <div className="bg-white shadow-lg  w-[300px] flex  flex-col justify-center items-center">
    <div className="border-y border-dashed p-4">
    <h3 className='text-xl font-semibold text-gray-900'>What you’ll learn</h3>
     <p className='text-gray-500 text-sm'>Access to IBM Cloud modelling what is I was looking for and Aaron got me the solution.</p>
    </div>

    <div className="py-6">
      <h6 className='text-lg font-medium'>This course includes:</h6>
      <ContentList icon={PiUploadSimple} name='100% Positive reviews (45)'/>
      <ContentList icon={PiStudentLight} name='2167 students'/>
      <ContentList icon={MdOutlinePlayLesson} name='59 Lessons (9h 11m)'/>
      <ContentList icon={LiaBookSolid} name='6 courses'/>
      <ContentList icon={GrDocumentDownload} name='70 downloads (70 files)'/>
      <ContentList icon={IoIosPhonePortrait} name='Available from the app'/>
      <ContentList icon={HiOutlineSpeakerWave} name='Audio: Spanish'/>
      <div className='text-xs text-gray-500 flex gap-2 w-full p-2'>
    <FiBarChart size={16}/>
    <div className=''><span>Level:</span> <span className='p-2 text-[#FF7800] bg-[#FFEFE0] rounded-[10px]'>Beginner</span></div>

  </div>

    </div>
     </div>

     <Button  className='bg-[#4C6FFF] rounded-[5px] text-white hover:text-black'>Publish review</Button>
     </div>
  )
}

export default CourseDescribe