import { configureStore } from "@reduxjs/toolkit";
import { topicsApi } from "./features/topics/topicsApi";
// import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";
// import { apiSlice } from "./api";
// import authReducer from "./features/auth/authSlice";


export const store = configureStore({
  reducer:{
    [topicsApi.reducerPath]: topicsApi.reducer,
    // [apiSlice.reducerPath]: apiSlice.reducer,
    // auth: authReducer,


  },
  middleware: (getsDefaultMiddleware) => 
    getsDefaultMiddleware().concat(topicsApi.middleware),

})

export type RooState = ReturnType<typeof store.getState>;
export  type AppDispatch = typeof store.dispatch;