import { configureStore } from "@reduxjs/toolkit";
import { topicsApi } from "./features/topics/topicsApi";
import { authApi } from "./features/auth/authApi";
import {quizApi} from "./features/quiz/quizApi"
import {quizSlice} from "./features/quiz/quizSlice"


export const store = configureStore({
  reducer:{
    quiz: quizSlice.reducer,
    [topicsApi.reducerPath]: topicsApi.reducer,
    [authApi.reducerPath]: authApi.reducer,
    [quizApi.reducerPath]: quizApi.reducer,
    
  


  },
  devTools:false,
  middleware: (getsDefaultMiddleware) => getsDefaultMiddleware().concat(topicsApi.middleware,authApi.middleware,quizApi.middleware)

})

export type RooState = ReturnType<typeof store.getState>;
export  type AppDispatch = typeof store.dispatch;