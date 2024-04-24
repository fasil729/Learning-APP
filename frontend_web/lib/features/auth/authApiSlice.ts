// import { fetchBaseQuery } from "@reduxjs/toolkit/query";
// import { createApi } from "@reduxjs/toolkit/query/react";
// // import { apiSlice } from "../../api";

// const baseEndPoint = process.env.NEXT_PUBLIC_NODE_ENV === "production" ? process.env.NEXT_PUBLIC_BACKEND_PROD : process.env.NEXT_PUBLIC_BACKEND_DEV;

// export const authApiSlice = createApi({
//   baseQuery: fetchBaseQuery({ baseUrl: baseEndPoint }),
//   endpoints: (builder) => ({
//     login: builder.mutation({
//       query: (credentials: { username: string; password: string }) => ({
//         url: '/user/login',
//         method: 'POST',
//         body: credentials,
//       }),
//     }),
//   }),
// });

// export const userCreateApiSlice = createApi({
//   baseQuery: fetchBaseQuery({ baseUrl: baseEndPoint }),
//   endpoints: (builder) => ({
//     createAccount: builder.mutation({
//       query: (credentials: { username: string; email: string; password: string }) => ({
//         url: `/user/create`,
//         method: 'POST',
//         body: credentials,
//       }),
//     }),
//   }),
// });

// export const { useCreateAccountMutation } = userCreateApiSlice;
// export const { useLoginMutation } = authApiSlice;


