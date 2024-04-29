import { put, takeLatest } from 'redux-saga/effects';
import {
    createSubjectSuccessAction,
    createSubjectFailureAction,
    createSubjectAction,
} from "../features/topics/topicSlice";
import axios, { Axios, AxiosResponse } from 'axios';
import { Topic } from '@/types/topic';
import { PayloadAction } from '@reduxjs/toolkit';


function* createSubjectSaga(action: PayloadAction<Topic>) {
    console.log("here creating subject");
    console.log(action.payload);
    let response: AxiosResponse;
    try {
        response = yield axios.post(`https://learning-app-idt8.onrender.com/subjects/create`,
            action.payload,
            {
            headers: {
                Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDM5MTY0NSwicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.QRRebYPcQkN4_Tzciryc3XepYWlstmhCCM9zOTijFa0`
            }
            }
        );
        console.log(response)
        yield put(createSubjectSuccessAction(response.data));
    } catch (error) {
        yield put(createSubjectFailureAction());   
        console.log(error) 
    }
}

export function* watchCreateSubjectSaga() {
    yield takeLatest(createSubjectAction.type, createSubjectSaga);
}