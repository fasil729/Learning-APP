import React from 'react'
interface SuccessMessageProps{
    message: string;
}
const SuccessMessage:React.FC<SuccessMessageProps> = ({message}) => {
  return (<div className="flex justify-center rounded-md bg-green-100 p-2">
  <p className='text-green-400  text-sm leading-6'>{message}</p>

 </div>
  )
}

export default SuccessMessage