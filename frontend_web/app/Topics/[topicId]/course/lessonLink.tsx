import Link from 'next/link'
import React from 'react'
import { MdOutlinePlayCircle } from "react-icons/md";
import { IoCheckmarkSharp } from "react-icons/io5";
interface LessonLinkProps{
  url: string;
  name: string;
  status: string;
}
const LessonLink:React.FC<LessonLinkProps> = ({url,status,name}) => {
  return (
    <Link href={url} className='w-full flex p-2 py-3 border-b hover:text-[#4C6FFF] justify-between transition duration-300'>
      <div className='flex gap-2'>
        <MdOutlinePlayCircle className='text-[#4C6FFF]' size={20}/>
        <p>{name}</p>
      </div>

      {status==="completed"?<IoCheckmarkSharp className='text-[#30A876]' size={20}/>:''}
    </Link>
  )
}

export default LessonLink