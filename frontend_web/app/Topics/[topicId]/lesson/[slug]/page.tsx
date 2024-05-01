// app/page.js
// 'use client';
// import { CustomMDX } from '../../../../../components/mdx-remote';
// import fs from 'fs';
// import path from 'path';
import { CustomMDX } from '@/components/mdx-remote';
import { getTopicDetail } from '@/lib/utils';


interface Props {
    params: {slug: string}
}


// console.log("this page slug data");
export default async function LessonDetail({params: { slug }} : Props) {
    var data = await getTopicDetail(slug);

    return <CustomMDX  source={data}/>
}
