import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export async function getTopicDetail(lessonName: string) {
  const res = await fetch(`https://learning-app-idt8.onrender.com/lessons/54/content`, {headers: {
    Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDUxNzgwNywicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.2SC0NyNCF3H-0A7xYkbjblV8FPmAOeQKRIAQm-rgNoM`
  
  }});
  console.log("getting resopnse", res);
  return res.body;
}