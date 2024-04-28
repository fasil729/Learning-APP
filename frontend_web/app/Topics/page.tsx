'use client';
import SideBar from "@/components/shared/sideBar/sideBar";
import TopHeader from "@/components/shared/top_header/top_header";
import TopicCard from "@/components/ui/topicCard";
import { topics } from "./topics";
import { useGetAllTopicsQuery } from "@/lib/features/topics/topicsApi";
import { TopicCreateForm } from "./components/topic-create-form";

export default function Topics() {
    const {data, error, isLoading} = useGetAllTopicsQuery();
    return <div>
            <SideBar></SideBar>
            <TopHeader></TopHeader>
            <main className="ml-[18%] p-7">
                <h1 className="text-xl text-titleColor">Topics</h1>
                <TopicCreateForm/>
                {isLoading && <div>Loading...</div>}
                {error && <div>{JSON.stringify(error)}</div>}
                {data &&
                    <>
                    <div className="">
                        <h4 className="text-sm font-medium py-5">Recommended</h4>
                        <div className="flex flex-wrap gap-10">
                            {topics.map((topic) => (
                                <TopicCard 
                                key ={topic.id}
                                src={topic.src}
                                alt={topic.alt}
                                title={topic.title}
                                lessons={topic.lessons}
                                price={topic.price}
                                ></TopicCard>
                            ))}
                        </div>
                    </div>
                    <div className="py-5">
                        <h4 className="text-sm font-meidum py-5">Created Topics</h4>
                        <div className="flex flex-wrap gap-10">
                            {topics.map((topic) => (
                                <TopicCard 
                                key ={topic.id}
                                id={topic.id}
                                src={topic.src}
                                alt={topic.alt}
                                title={topic.title}
                                lessons={topic.lessons}
                                price={topic.price}
                                ></TopicCard>
                            ))}
                        </div>
                    </div>
                </>
            }
            </main>
         </div>
}