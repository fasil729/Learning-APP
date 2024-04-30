import Link from 'next/link'
import React from 'react'
import { IoCheckmarkSharp } from 'react-icons/io5';
import { MdOutlinePlayCircle } from "react-icons/md";


interface LessonLinkProps{
  url: string;
  name: string;

  status: string;
  
}
const LessonLink:React.FC<LessonLinkProps> = ({url,status,name}) => {
  return (
     <Link href={url}  className="hover:text-[#4C6FFF] flex justify-between transition duration-300 ">
     <div className='flex gap-2'>
        <MdOutlinePlayCircle className='text-[#4C6FFF]' size={20}/>
        <p>{name}</p>
      </div>

      {status==="completed"?<IoCheckmarkSharp className='text-[#30A876]' size={20}/>:''}
     </Link>
  )
}

export default LessonLink