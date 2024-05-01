"use client"
import { store } from '@/lib/store'
import React from 'react'
import { Provider } from 'react-redux';
import { PersistGate } from 'redux-persist/integration/react';
import {
  persistStore,

} from 'redux-persist'


let persistor = persistStore(store)
const AppProvider = ({children}:{children:React.ReactNode}) => {
  return (<Provider store={store}><PersistGate loading={null} persistor={persistor}></PersistGate>{children}</Provider>
    
  )
}

export default AppProvider;