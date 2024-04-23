// lessonLink.tsx
import React from 'react'
import { IoCheckmarkSharp } from "react-icons/io5";

interface LessonLinkProps{
  name: string;
  checked: boolean;
  onChange: () => void;
}

const LessonLink: React.FC<LessonLinkProps> = ({ name, checked, onChange }) => {
  return (
    <div className='flex gap-2 items-center'> {/* Changed from Link to div */}
      <input type="checkbox" checked={checked} onChange={onChange} className='text-[#4C6FFF]' /> {/* Display checkbox */}
      <p>{name}</p>
      {checked && <IoCheckmarkSharp className='text-[#30A876]' size={20}/>} {/* Show checkmark if checked */}
    </div>
  )
}

export default LessonLink
