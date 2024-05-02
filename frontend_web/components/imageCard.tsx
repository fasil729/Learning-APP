import Image from 'next/image'
import React from 'react'
interface ImageCardProps{
  url:any;
}
const ImageCard:React.FC<ImageCardProps> = ({url}) => {
  return (<div>
    <Image height={200} width={200} src={url} alt=''className='rounded-[5px]'/>
  </div>
  )
}

export default ImageCard
