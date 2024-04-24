import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const baseQuery = fetchBaseQuery({
    baseUrl: 'http://localhost:4000/',
})

export const topicsApi = createApi({
    baseQuery,
    endpoints: (builder) => ({
        getAllTopics : builder.query<Topic, void>({
            query: () => `topics`
        })
    })
})

export const {
    useGetAllTopicsQuery
} =  topicsApi;