import { put, takeLatest } from 'redux-saga/effects';
import {
    createSubjectSuccessAction,
    createSubjectFailureAction,
    createSubjectAction,
} from "../features/topics/topicSlice";
import axios, { Axios, AxiosResponse } from 'axios';
import { Topic } from '@/types/topic';
import { PayloadAction } from '@reduxjs/toolkit';

const token=localStorage.getItem("accessToken")

function* createSubjectSaga(action: PayloadAction<any>) {
    let response: AxiosResponse;
    try {
        response = yield axios.post(`https://learning-app-idt8.onrender.com/subjects/create`,
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
    yield takeLatest(createSubjectAction.type, createSubjectSaga);
}