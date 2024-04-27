'use client';
import SideBar from "@/components/shared/sideBar/sideBar";
import TopHeader from "@/components/shared/top_header/top_header";
import Card from "@/components/ui/card";
import { topics } from "./topics";
import { useGetAllTopicsQuery } from "@/lib/features/topics/topicsApi";

export default function Topics() {
    const {data, error, isLoading} = useGetAllTopicsQuery();
    if(isLoading) return <div>Loading...</div>
    
    if (error) return <div>{JSON.stringify(error)}</div>

    return <div>
            <SideBar></SideBar>
            <TopHeader></TopHeader>
            <main className="ml-[18%] p-7">
                <h1 className="text-xl text-titleColor">Topics</h1>
                <div className="">
                    <h4 className="text-sm font-medium py-5">Recommended</h4>
                    <div className="flex flex-wrap gap-10">
                        {topics.map((topic) => (
                            <Card 
                            key ={topic.id}
                            src={topic.src}
                            alt={topic.alt}
                            title={topic.title}
                            lessons={topic.lessons}
                            price={topic.price}
                            ></Card>
                        ))}
                    </div>
                </div>
                <div className="py-5">
                    <h4 className="text-sm font-meidum py-5">Created Topics</h4>
                    <div className="flex flex-wrap gap-10">
                        {topics.map((topic) => (
                            <Card 
                            key ={topic.id}
                            src={topic.src}
                            alt={topic.alt}
                            title={topic.title}
                            lessons={topic.lessons}
                            price={topic.price}
                            ></Card>
                        ))}
                    </div>
                </div>
            </main>
         </div>
}