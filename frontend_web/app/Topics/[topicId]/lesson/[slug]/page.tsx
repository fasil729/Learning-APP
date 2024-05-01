// app/page.js
// 'use client';
// import { CustomMDX } from '../../../../../components/mdx-remote';
// import fs from 'fs';
// import path from 'path';
import { CustomMDX } from '@/components/mdx-remote';
import { Button } from '@/components/ui/button';
import { getTopicDetail } from '@/lib/utils';


interface Props {
    params: {slug: string}
}


// console.log("this page slug data");
export default async function LessonDetail({params: { slug }} : Props) {
    var data = await getTopicDetail(slug);

    return <div className='p-[20px]'><CustomMDX  source={data}/>
            <div className='flex gap-[3%] py-[3%]'>
                <Button>Generate Quiz</Button>
                <Button>Exam Preparation</Button>
            </div>
        </div>  
}
