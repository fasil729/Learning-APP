import { put, takeLatest } from 'redux-saga/effects';
import {
    createExperimentSuccessAction,
    createExperimentFailureAction,
    createExperimentAction,
} from "../features/experiments/experimentSlice";
import axios, { Axios, AxiosResponse } from 'axios';
import { Topic } from '@/types/topic';
import { PayloadAction } from '@reduxjs/toolkit';

const token=localStorage.getItem("accessToken")
console.log(token);

function* createExperimentSaga(action: PayloadAction<any>) {
    let response: AxiosResponse;
    try {
        response = yield axios.post(`https://learning-app-idt8.onrender.com/subjects/create`,
            action.payload,
            {
            headers: {
                Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDQ4NjM2OSwicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.FIR3HRHzFvorWuv0eMIOmVD2XaQB5TDss3swfXVtFCE`
            }
            }
        );
        yield put(createExperimentSuccessAction(response.data));
    } catch (error) {
        yield put(createExperimentFailureAction());    
    }
}

export function* watchCreateExperimentSaga() {
    yield takeLatest(createExperimentAction.type, createExperimentSaga);
}