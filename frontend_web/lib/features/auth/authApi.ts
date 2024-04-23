import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const baseQuery = fetchBaseQuery({
    baseUrl: 'http://localhost:4000/'
})

export const authApi = createApi({
    baseQuery,
    endpoints: (builder) => ({
        signUp: builder.mutation({
            query: () => ''
        }),

        login: builder.mutation({
            query: () => ''
        }),

        logout: builder.mutation({
            query: () => ''
        })
    })
})


export const {
    useSignUpMutation,
    useLoginMutation,
    useLogoutMutation
} = authApi;