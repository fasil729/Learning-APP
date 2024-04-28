import { FormDataModel } from "@/types/formData";
import { Topic } from "@/types/topic";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const baseQuery = fetchBaseQuery({
    baseUrl: 'https://learning-app-idt8.onrender.com/subjects',
})

export const topicsApi = createApi({
    baseQuery,
    endpoints: (builder) => ({
        getAllTopics : builder.query<Topic, void>({
            query: () => ``
        }),
        createTopic: builder.mutation<Topic, FormDataModel>({
            query(body) {
                return {
                    url: `create`,
                    method: 'POST',
                    body
                }
            }
        })
    })
})

export const {
    useGetAllTopicsQuery,
    useCreateTopicMutation
} =  topicsApi;