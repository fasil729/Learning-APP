import { configureStore } from "@reduxjs/toolkit";
import createSagaMiddlware from '@redux-saga/core';
// import { topicsApi } from "./features/topics/topicsApi";
import rootReducers from "./reducer";
import rootSaga from "./sagas/root-saga";



const sagaMiddleware  = createSagaMiddlware();

export const store = configureStore({
  reducer: rootReducers,
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(sagaMiddleware)
    // [topicsApi.reducerPath]: topicsApi.reducer,
  // middleware: (getsDefaultMiddleware) => getsDefaultMiddleware().concat(topicsApi.middleware),
})
sagaMiddleware.run(rootSaga);


export type RootState = ReturnType<typeof store.getState>;
export  type AppDispatch = typeof store.dispatch;