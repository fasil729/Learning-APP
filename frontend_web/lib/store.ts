import { configureStore } from "@reduxjs/toolkit";
import { topicsApi } from "./features/topics/topicsApi";

export const store = configureStore({
  reducer:{
    [topicsApi.reducerPath]: topicsApi.reducer,
  },
  middleware: (getsDefaultMiddleware) => getsDefaultMiddleware().concat(topicsApi.middleware),
})

export type RooState = ReturnType<typeof store.getState>;
export  type AppDispatch = typeof store.dispatch;