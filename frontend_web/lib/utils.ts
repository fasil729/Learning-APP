import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export async function getTopicDetail(lessonName: string) {
  const res = await fetch(`https://learning-app-idt8.onrender.com/lessons/${lessonName}/content`, {headers: {
    Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDcyODE1Nywicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.OQQ1gcLWtFqrUwcFhZb_XMCRhmL8XKA1jg7briLpa8g`
  }});
  // console.log(res); 
  return res.text();
}

export async function getExamPrepDetail(payload: { prompt: string, readTime: number, topics: string[] }) {
  console.log("payload", payload);
  const res = await fetch('https://learning-app-idt8.onrender.com/examprep/generate', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDcyODE1Nywicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.OQQ1gcLWtFqrUwcFhZb_XMCRhmL8XKA1jg7briLpa8g'
    },
    body: JSON.stringify(payload)
  });

  if (!res.ok) {
    throw new Error(`Failed to fetch exam prep detail: ${res.statusText}`);
  }

  const text = await res.text();
  console.log("response body text", text);

  return text;
}
