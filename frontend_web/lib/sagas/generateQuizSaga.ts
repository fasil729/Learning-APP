import axios, { AxiosResponse } from 'axios';
import {
  put,
  takeLatest,
} from 'redux-saga/effects';

import { PayloadAction } from '@reduxjs/toolkit';

import {
  createSubjectAction,
  createSubjectFailureAction,
  createSubjectSuccessAction,
} from '../features/topics/topicSlice';

const token=localStorage.getItem("accessToken")
console.log(token);

function* generateQuizSaga(action: PayloadAction<any>) {
    let response: AxiosResponse;
    const token = localStorage.getItem("accessToken")
    try {
        response = yield axios.post(`https://learning-app-idt8.onrender.com/quiz/generate`,
            action.payload,
            {
            headers: {
                Authorization: `Bearer ${token}`
            }
            }
        );
        yield put(createSubjectSuccessAction(response.data));
    } catch (error) {
        yield put(createSubjectFailureAction());    
    }
}

export function* watchCreateSubjectSaga() {
    yield takeLatest(createSubjectAction.type, generateQuizSaga);
}