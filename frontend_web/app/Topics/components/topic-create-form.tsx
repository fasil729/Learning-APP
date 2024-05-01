'use client';
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { useState } from "react";
import { IoAdd } from "react-icons/io5"
import { FormModal } from "./FormModal";

export const  TopicCreateForm = () => {
    const [modalIsOpen, setModalIsOpen] = useState(false);
    const closeModal = () => setModalIsOpen(false);


    return <div className= "h-[6.2rem] flex gap-10 items-center px-10">
                <Input type="email" className="w-1/5" placeholder="Search Topics"></Input>
                <div className="flex gap-9">
                    <Button className="bg-mainColor flex gap-2" onClick={() =>{console.log("pushed"), setModalIsOpen(true)}}><IoAdd className="text-white" size="1.5rem"/> Create</Button>
                    <FormModal modalIsOpen={modalIsOpen}  closeModal={closeModal}></FormModal>
                </div>
            </div>
}