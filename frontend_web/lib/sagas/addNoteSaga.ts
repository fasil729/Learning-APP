import { PayloadAction } from "@reduxjs/toolkit";
import axios, { AxiosResponse } from "axios";
import { put, takeLatest } from "redux-saga/effects";
import { addNoteAction, addNoteFailureAction, addNoteSuccessAction } from "../features/topics/noteSlice";

function *AddNoteSaga(action: PayloadAction<any>) {
    let response: AxiosResponse;
    try {
        response = yield axios.post(`https://learning-app-idt8.onrender.com/subjects/create`,
            action.payload,
            {
            headers: {
                Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhc2lsQGdtYWlsLmNvbSIsImV4cCI6MTcxNDQ4NjM2OSwicm9sZSI6InN0dWRlbnQiLCJzdWIiOjIsInVzZXJuYW1lIjoiZmFzaWwifQ.FIR3HRHzFvorWuv0eMIOmVD2XaQB5TDss3swfXVtFCE`
            }
        });
        yield put(addNoteSuccessAction(response.data))
    } catch(error) {
        yield put(addNoteFailureAction(error));
    }
}

export function *watchAddNoteSaga() {
    yield takeLatest(addNoteAction.type, AddNoteSaga);
}