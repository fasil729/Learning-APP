import SideBar from "@/components/shared/sideBar/sideBar";
import TopHeader from "@/components/shared/top_header/top_header";
import Card from "@/components/ui/card";
import { topics } from "./topics";

export default function Topics() {
    return <div>
            <SideBar></SideBar>
            <TopHeader></TopHeader>
            <main className="ml-[18%] p-7">
                <h1 className="text-lg text-titleColor  text-[34px]">Topics</h1>
                <div className="">
                    <h1 className="text-[18px]">Recommended</h1>
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
                <div className="text-[18px]">
                    <h1>Created Topics</h1>
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