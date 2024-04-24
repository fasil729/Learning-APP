import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const baseQuery = fetchBaseQuery({
    baseUrl: 'http://localhost:4000/'
});

export const authApi = createApi({
    reducerPath: 'authApi',
    baseQuery: baseQuery,
    endpoints: (builder) => ({
        signUp: builder.mutation({
            query: (userData) => ({
                url: 'api/signup',
                method: 'POST',
                body: userData,
            }),
        }),
        login: builder.mutation({
            query: (loginData) => ({
                url: 'api/login',
                method: 'POST',
                body: loginData,
            }),
        }),
        logout: builder.mutation({
            query: () => ({
                url: 'api/logout',
                method: 'POST',
            }),
        }),
        getAuthData: builder.query({
            query: (token) => ({
                url: 'api/auth-details',
                method: 'GET',
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            }),
        }),
    }),
});

export const {
    useSignUpMutation,
    useLoginMutation,
    useLogoutMutation,
    useGetAuthDataQuery
} = authApi;
