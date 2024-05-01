import { configureStore } from "@reduxjs/toolkit";
import createSagaMiddlware from '@redux-saga/core';
// import { topicsApi } from "./features/topics/topicsApi";
import rootReducers from "./reducer";
import rootSaga from "./sagas/root-saga";
import { topicsApi } from "./features/topics/topicsApi";
import { authApi } from "./features/auth/authApi";
import {quizApi} from "./features/quiz/quizApi"
import {quizSlice} from "./features/quiz/quizSlice"



const sagaMiddleware  = createSagaMiddlware();


export const store = configureStore({
  reducer: rootReducers,
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(sagaMiddleware,quizApi.middleware),
  devTools:false,
})

sagaMiddleware.run(rootSaga);


export type RootState = ReturnType<typeof store.getState>;
export  type AppDispatch = typeof store.dispatch;