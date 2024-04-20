import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { IoAdd, IoSearch } from "react-icons/io5";

export default function TopHeader() {
        return   <div className="ml-[18%] bg-white h-[6.2rem] flex justify-between items-center px-14 ">
                    <Input type="email" className="w-1/5" placeholder="Search Anything"></Input>
                    <div className="flex gap-9">
                        <Button className="bg-mainColor flex gap-2"><IoAdd className="text-white" size="1.5rem"/> Create</Button>
                        <div className="flex gap-4">
                            <div>
                                <img src="https://cdn-icons-png.flaticon.com/512/25/25231.png" alt="user" className="w-10 h-10 rounded-full"/>
                            </div>
                            <div>
                                <h3>John Doe</h3>
                                <p>Level 3</p>
                            </div>
                        </div>
                    </div>
        </div>
    }