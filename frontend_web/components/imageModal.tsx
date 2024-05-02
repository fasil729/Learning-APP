"use client"
import { ReactNode, useState } from "react";
import { Button } from "./ui/button";

interface ModalProps{
  label?:ReactNode;
 
  children:ReactNode;
  title?:string;
  setOpen:() => void;
  isOpen:boolean;
}
const CModal:React.FC<ModalProps> = ({label,children,title,setOpen,isOpen}) => {
  
  
  return ( <>
  <div onClick={setOpen} >{label}</div>
  <div className={`fixed  h-screen w-screen top-0 bg-gray-700 dark:bg-opacity-80 bg-opacity-60 botton-0 left-0 right-0  flex justify-center items-center ${isOpen? "block opacity-1":"hidden opacity-0"} `}>
   <div className={`flex flex-col  min-w-[500px] bg-white shadow-lg rounded-[10px]   max-w-[800px]
    ${isOpen? "translate-y-0" : "translate-y-full"
              } transition-transform duration-500`}>
      <div className="flex justify-between p-2">
     <h1 className="text-lg font-semibold leading-6 p-2">Image Generators</h1>
      <Button onClick={setOpen} variant="destructive">X</Button>
      
      </div> 
      <div className="">{children}</div>
    </div>
  
  </div> </>);
}
 
export default CModal;
