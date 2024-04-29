import { FormDataModel } from "@/types/formData";
import { Topic } from "@/types/topic";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const baseQuery = fetchBaseQuery({
    baseUrl: 'https://learning-app-idt8.onrender.com/subjects',
    headers: {Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDMzNzI0NCwicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.QC4ZFNIE6fp0-eULVJXauai5S88IQ9uJBCyMIV_4jWk`}
})

export const topicsApi = createApi({
    baseQuery,
    endpoints: (builder) => ({
        getAllTopics : builder.query<Topic, void>({
            query: () => ``
        }),
        createTopic: builder.mutation<Topic, FormDataModel>({
            query : (body) => ({
                    url: `create`,
                    method: 'POST',
                    body
            })
        })
    })
})

export const {
    useGetAllTopicsQuery,
    useCreateTopicMutation
} =  topicsApi;