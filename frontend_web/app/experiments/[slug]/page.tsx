import { CustomMDX } from '@/components/mdx-remote';
import { getExperimentDetial } from '@/lib/utils';


interface Props {
    params: {slug: string}
}



export default async function ExperimentDetail({params: { slug }} : Props) {
    var data = await getExperimentDetial(slug);

    return <div className='px-[20px] py-[20px]'>
        <CustomMDX  source={data.content}/>
    </div>
}
