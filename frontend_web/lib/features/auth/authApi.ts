import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const baseQuery = fetchBaseQuery({
    baseUrl: 'https://learning-app-idt8.onrender.com/'
});

export const authApi = createApi({
    reducerPath: 'authApi',
    baseQuery: baseQuery,
    endpoints: (builder) => ({
        signUp: builder.mutation<any,any>({
            query: (userData) => ({
                url: 'users/register',
                method: 'POST',
                body: userData,
            }),
        }),
        login: builder.mutation<any, any>({
            query: (loginData) => ({
                url: 'users/login',
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
