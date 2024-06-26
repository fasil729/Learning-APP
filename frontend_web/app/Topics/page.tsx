'use client';
import SideBar from "@/components/shared/sideBar/sideBar";
import TopHeader from "@/components/shared/top_header/top_header";
import TopicCard from "@/components/ui/topicCard";
import { topics } from "./topics";
import { TopicCreateForm } from "./components/topic-create-form";
import { useSelector } from "react-redux";
import { useRouter } from "next/navigation";

export default function Topics() {
    const authStatus = useSelector((state: any) => state.auth.data.isAuthenticated);
    const router = useRouter();

    if (!authStatus)
        router.push("/login")
    return <div>
            {/* <SideBar></SideBar> */}
            <TopHeader></TopHeader>
            <main className="ml-[18%] p-7">
                <h1 className="text-xl text-titleColor px-10">Topics</h1>
                <TopicCreateForm/>
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
            </main>
         </div>
}