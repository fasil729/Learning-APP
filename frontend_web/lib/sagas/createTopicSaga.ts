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
console.log(token);

function* createSubjectSaga(action: PayloadAction<any>) {
    let response: AxiosResponse;
    try {
        response = yield axios.post(`https://learning-app-idt8.onrender.com/subjects/create`,
            action.payload,
            {
            headers: {
                Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDQ3OTI3OCwicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.odLQP8H_yBXuAWMilh_LU_XUi4YO6BIoYfihXjFeCNs`
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