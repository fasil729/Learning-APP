import { configureStore } from "@reduxjs/toolkit";
import createSagaMiddlware from '@redux-saga/core';
import rootReducers from "./reducer";
import rootSaga from "./sagas/root-saga";
import { authApi } from "./features/auth/authApi";
import {quizApi} from "./features/quiz/quizApi"
import {quizSlice} from "./features/quiz/quizSlice";
import {
  persistStore,
  persistReducer,
} from 'redux-persist' ;
import storage from 'redux-persist/lib/storage'



const persistConfig = {
  key: 'root',
  version: 1,
  storage,
  blacklist: ['isSuccess']
}
;


const sagaMiddleware  = createSagaMiddlware();

const persistedReducer = persistReducer(persistConfig, rootReducers)

export const store = configureStore({
  reducer: persistedReducer,
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(sagaMiddleware,quizApi.middleware),
  devTools:false,
})

sagaMiddleware.run(rootSaga);


export type RootState = ReturnType<typeof store.getState>;
export  type AppDispatch = typeof store.dispatch;