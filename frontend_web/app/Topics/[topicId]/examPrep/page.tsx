import { CustomMDX } from '@/components/mdx-remote';
import { getExamPrepDetail } from '@/lib/utils';

interface SearchParams {
  prompt: string;
  readTime: string;
  topics: string;
}

export default function ExamPrepDetail(props: { searchParams?: SearchParams }) {
  const { searchParams } = props;

  if (!searchParams) {
    return <div>Loading...</div>; // Or handle the case when searchParams is undefined
  }

  const readTime = Number(searchParams.readTime) || 0;
  const topics = searchParams.topics.split(',').map(topic => topic.trim());

  const data = getExamPrepDetail({ prompt: searchParams.prompt, readTime, topics });

  // debug 

  return  <CustomMDX source={data} />
}