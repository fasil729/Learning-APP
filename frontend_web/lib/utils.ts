import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export async function getTopicDetail(lessonName: string) {
  const res = await fetch(`https://learning-app-idt8.onrender.com/lessons/${lessonName}/content`, {headers: {
    Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDY1NzAyNywicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.sn3TEh8GxP8jD5DmLUmNuQRraspepshIQrcBNPEUiBc`
  }});
  console.log(res);
  return res.text();
}

export async function getExamPrepDetail(payload: { prompt: string, readTime: number, topics: string[] }) {
  const res = await fetch('https://learning-app-idt8.onrender.com/examprep/generate', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDY1NzAyNywicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.sn3TEh8GxP8jD5DmLUmNuQRraspepshIQrcBNPEUiBc'
    },
    body: JSON.stringify(payload)
  });
  
  if (!res.ok) {
    throw new Error(`Failed to fetch exam prep detail: ${res.statusText}`);
  }
  
  return res.json();
}


export async function getExperimentDetial(experimentName: string) {
  const res = await fetch(`https://learning-app-idt8.onrender.com/experiments/${experimentName}/content`, {headers: {
    Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDY1NzAyNywicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.sn3TEh8GxP8jD5DmLUmNuQRraspepshIQrcBNPEUiBc`
    }});
  console.log(res);
  return res.json();
}