
"use client"
import { useState } from 'react';

import { Button } from '@/components/ui/button';

import { InputForm } from './form';

export function ButtonDemo() {
  const [isClicked, setIsClicked] = useState(false);

  const handleClick = () => {
    setIsClicked(true);
  };

  return <Button onClick={handleClick}>Button</Button>;
}

// Main page
export default function MainPage() {
  const [isClicked, setIsClicked] = useState(false);

  const handleClick = () => {
    setIsClicked(! isClicked);
  };

  return (
    <div className=" justify-center items-center bg-gray-100 h-screen break-words">
      <div className="  ml-12 ">
      {isClicked ? (<div className=" flex  items-center justify-between">
        <div className="break-normal w-100 mr-12 w-2/3">
        <p className=" break-words text-justify  text-gray-900">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Tempora, aliquid. Natus deserunt iste magni ab cupiditate cumque officiis esse sint aliquam totam nesciunt, blanditiis sapiente hic optio inventore quasi consequatur.
          Lorem ipsum dolor sit, amet consectetur adipisicing elit. Laborum accusantium quis autem quia, nisi suscipit harum commodi incidunt aliquid assumenda tempora recusandae, sit aut veritatis architecto distinctio a quidem cupiditate.
          </p>
      </div>
               <div className='w-1/3	'> <InputForm /></div>

      </div>):          <div className="break-normal w-100 mr-12 ">
        <p className=" break-words text-justify  text-gray-900">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Tempora, aliquid. Natus deserunt iste magni ab cupiditate cumque officiis esse sint aliquam totam nesciunt, blanditiis sapiente hic optio inventore quasi consequatur.
          Lorem ipsum dolor sit, amet consectetur adipisicing elit. Laborum accusantium quis autem quia, nisi suscipit harum commodi incidunt aliquid assumenda tempora recusandae, sit aut veritatis architecto distinctio a quidem cupiditate.
          </p>
      </div>}
        <div className="break-words">
        <Button onClick={handleClick}>Button</Button>

        </div>
      </div>
    </div>
  );
}