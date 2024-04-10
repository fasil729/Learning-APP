import { configureStore } from "@reduxjs/toolkit";

export const store=configureStore({
  reducer:{

  }
})


export type RooState=ReturnType<typeof store.getState>;
export  type AppDispatch=typeof store.dispatch;