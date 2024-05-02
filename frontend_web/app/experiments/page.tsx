'use client';
import { useRouter } from "next/navigation";
import { useSelector } from "react-redux";


interface Props {
    searchParams: { chapterName: string }
}


export default function Page({ searchParams: { chapterName } }: Props) {
    const { isLoading, data, errors } = useSelector((state: any) => state.experiment);
    const authStatus = useSelector((state: any) => state.auth.data.isAuthenticated);
    const router = useRouter();

    if (!authStatus)
        router.push("/login")

    if (isLoading)
        return <div>Loading...</div>

    return (

        <div className="w-full px-10  p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 dark:bg-gray-800 dark:border-gray-700">
            <h5 className="mb-3 text-sm font-semibold pb-10 text-gray-900 md:text-lg dark:text-white">
            {chapterName} Experiments
            </h5>
            <ul className="my-4 space-y-3">
            {data && data.map((experiment: any, index: number) => {
                                return (<li key={index}>
                                    <a href={`/experiments/${experiment.experiment_name}`} className="flex items-center p-3 text-base font-semibold text-gray-900 rounded-lg bg-gray-50 hover:bg-gray-100 group hover:shadow dark:bg-gray-600 dark:hover:bg-gray-500 dark:text-white">
                                        <span className="flex-1 ms-3 whitespace-nowrap">{experiment.experiment_name}</span>
                                    </a>
                                </li>)
            })}
                
            </ul>
        </div>)
}   