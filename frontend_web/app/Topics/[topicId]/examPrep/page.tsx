import { CustomMDX } from '@/components/mdx-remote';
import { useSelector } from 'react-redux';



// console.log("this page slug data");
export default async function LessonDetail() {
    const { isSuccess, errors, data } = useSelector((state: any) => state.examPrep);

    return <CustomMDX  source={data}/>
}