import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
const token=localStorage.getItem("accessToken")

export const baseQuery = fetchBaseQuery({
    baseUrl: 'https://learning-app-idt8.onrender.com/'
});

export const quizApi= createApi({
    reducerPath: 'quizApi',
    baseQuery: baseQuery,
    endpoints:(builder)=>({
        
quizGenerate:builder.mutation<any,any>({
    query: (quizData) => ({
        url: 'quiz/generate',
        method: 'POST',
        body: quizData,
        headers: {
            Authorization: `Bearer ${token}`,
        },
    }),
})
        })

    })


    export const {useQuizGenerateMutation}=quizApi;