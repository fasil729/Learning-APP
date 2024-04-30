import { type ClassValue, clsx } from "clsx"
import { cookies } from "next/headers";
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

const token = cookies().get("accessToken");
export async function getTopicDetail(lessonName: string) {
  const res = await fetch(`https://learning-app-idt8.onrender.com/lessons/54/content`, { headers: {
    Authorization: `Bearer ${token}`
  }});
  console.log(res)
  return res.body;
}