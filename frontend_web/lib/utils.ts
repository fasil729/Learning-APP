import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export async function getTopicDetail(lessonName: string) {
  const res = await fetch(`https://learning-app-idt8.onrender.com/lessons/${lessonName}/content`, {headers: {
    Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDYxOTY1MSwicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.7cknUvx3qKGYMStDZn7Bh9lQIHESQEnLKBrv3adgaPA`
  }});
  console.log(res);
  return res.text();
}