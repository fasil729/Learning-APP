import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const baseQuery = fetchBaseQuery({
    baseUrl: 'http://localhost:4000/',
    headers: {Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDMzMzQzNywicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.gEZC9qbBV8Qf8m_vOxAXOtGeUV8dGm327H14en5EVag`}
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