import { CustomMDX } from '@/components/mdx-remote';
import { getExamPrepDetail } from '@/lib/utils';

interface SearchParams {
  prompt: string;
  readTime: string;
  topics: string;
}

export default async function ExamPrepDetail(props: { searchParams?: SearchParams }) {
  const { searchParams } = props;

  if (!searchParams) {
    return <div>Loading...</div>; // Or handle the case when searchParams is undefined
  }

  const readTime = Number(searchParams.readTime) || 0;
  const topics = searchParams.topics.split(',').map(topic => topic.trim());

  const data = await getExamPrepDetail({ prompt: searchParams.prompt, readTime, topics });

  return <div className='px-[20px] py-[20px]'>
            <CustomMDX  source={data}/>
          </div>
}