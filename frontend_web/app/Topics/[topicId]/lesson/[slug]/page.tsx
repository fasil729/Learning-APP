// app/page.js
// 'use client';
// import { CustomMDX } from '../../../../../components/mdx-remote';
// import fs from 'fs';
// import path from 'path';
import { getTopicDetail } from '@/lib/utils';
import { MdxComponent } from './mdx';
// import matter from 'gray-matter'


// const getContent = () => {
//   const filePath = path.join(process.cwd(), './lib/test.md')
//   const fileContent = fs.readFileSync(filePath, 'utf-8')
// //   const { content } = matter(fileContent)
//   return fileContent
// }


interface Props {
    params: {slug: string}
}



export default async function Home({params: { slug }} : Props) {
    console.log("slug data" ,slug);
    var data = await getTopicDetail(slug);

  return (
    <MdxComponent data={data}/>
  )
}
