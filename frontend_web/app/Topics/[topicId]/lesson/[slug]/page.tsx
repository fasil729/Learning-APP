// // app/page.js
// // 'use client';
// // import { CustomMDX } from '../../../../../components/mdx-remote';
// // import fs from 'fs';
// // import path from 'path';
// import { getTopicDetail } from '@/lib/utils';
// import { MdxComponent } from './mdx';
// // import matter from 'gray-matter'


// // const getContent = () => {
// //   const filePath = path.join(process.cwd(), './lib/test.md')
// //   const fileContent = fs.readFileSync(filePath, 'utf-8')
// // //   const { content } = matter(fileContent)
// //   return fileContent
// // }


// interface Props {
//     params: {slug: string}
// }


// console.log("this page slug data");
// export default async function LessonDetail({params: { slug }} : Props) {
//     console.log("slug data" ,slug);
//     var data = await getTopicDetail(slug);
//     // var data ="**Testing**"
//     console.log(data);

//   return (
//     <MdxComponent data={data}/>
//   )
// }

export default function Page() {
    console.log('workning now')
    return <div>working</div>
}
