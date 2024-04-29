import React from 'react'
interface ErrorMessageProps{
    message: string;
}
const ErrorMessage:React.FC<ErrorMessageProps> = ({message}) => {
  return (<div className="flex justify-center rounded-md bg-red-100 p-2">
  <p className='text-red-400  text-sm leading-6'>{message}</p>

 </div>
  )
}

export default ErrorMessage