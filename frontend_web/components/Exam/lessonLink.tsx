import React from 'react'
import { IoCheckmarkSharp } from "react-icons/io5";

interface LessonLinkProps{
  name: string;
  checked: boolean;
  onChange: () => void;
}

const LessonLink: React.FC<LessonLinkProps> = ({ name, checked, onChange }) => {
  return (
    <div className='flex gap-2  text-titleColor items-center text-center'> {/* Changed from Link to div, Centered horizontally */}
      <input type="checkbox" checked={checked} onChange={onChange} className='text-mainColor' /> {/* Display checkbox */}
      <p>{name}</p>
      {checked && <IoCheckmarkSharp className='text-secondayColor' size={20}/>} {/* Show checkmark if checked */}
    </div>
  )
}

export default LessonLink
