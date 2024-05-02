"use client"
import React, { useState } from 'react'
import ImageCard from './imageCard';
import { Button } from './ui/button';
import Image from 'next/image';
import { ScrollArea } from './ui/scroll-area';
import CModal from './imageModal';

const ImagePopup = () => {
  const [isOpen,setOpen]=useState(false)
  const [gen,setGen]=useState(false)
  return (<div className="flex min-h-screen justify-center items-center">
    <CModal 
    label={<div className="h-[400] flex justify-center items-center">
      <Image
        height={400}
        width={300}
        src='https://img.freepik.com/free-photo/scientific-microscope-laboratory-desk-with-researching-instruments_482257-13974.jpg?size=626&ext=jpg&ga=GA1.1.2093102158.1709920269&semt=sph'
        alt='image'
        className='rounded-[10px]' />
    </div>} 
     setOpen={()=>setOpen((prev)=>!prev)} 
     isOpen={isOpen}>


  <div className=' bg-white grid grid-cols-2 gap-10 p-4'>
  <div className="h-full flex justify-center items-center">
    <Image 
    height={400}
     width={300} 
     src={`https://picsum.photos/200/200?random=${3}`}
      alt='image'
      className='rounded-[10px]'
      />
  </div>
 
  <div className="flex flex-col pt-6 gap-10">
    <Button 
    onClick={()=>setGen((prev)=>!prev)}
    className='
    hover:bg-blue-500
    hover:text-white
    z-20
    transition
    duration-300
    bg-slate-100
     rounded-full
     border-blue-500 border-2 
     text-blue-500 font-semibold'>Generate Image</Button>

   <ScrollArea className={`h-[500px] ${gen?'opacity-100 translate-y-0':'opacity-0 -translate-y-[100px]'} transition duration-1000`}>
   <div className="grid grid-cols-2 p-4 gap-10">
     <ImageCard url={`https://img.freepik.com/free-photo/scientific-microscope-laboratory-desk-with-researching-instruments_482257-13974.jpg?size=626&ext=jpg&ga=GA1.1.2093102158.1709920269&semt=sph`}/>
     <ImageCard url={`https://img.freepik.com/free-photo/world-science-day-arrangement-with-microscope_23-2149112622.jpg?size=626&ext=jpg&ga=GA1.1.2093102158.1709920269&semt=ais`}/>
     <ImageCard url={`https://picsum.photos/200/200?random=${3}`}/>
     <ImageCard url={`https://picsum.photos/200/200?random=${6}`}/>
     <ImageCard url={`https://picsum.photos/200/200?random=${2}`}/>
     <ImageCard url={`https://picsum.photos/200/200?random=${7}`}/>
    </div>
   </ScrollArea>
  </div>
</div>
    </CModal>
   


  </div>
  )
}

export default ImagePopup;
