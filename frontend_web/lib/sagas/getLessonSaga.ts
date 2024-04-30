import axios, { AxiosResponse } from "axios";
import { getTopicDetailSuccessAction,  getTopicDetailFaliureAction, getTopicDetailAction } from "../features/topics/lessonSlice";
import { put, takeLatest } from "redux-saga/effects";


const token=localStorage.getItem("accessToken")

function *getTopicDetailSaga(action: any) {
    let response : AxiosResponse;
    try {
        response = yield axios.post('https://learning-app-idt8.onrender.com/subjects/create', "", 
        {
            headers: {
                Authorization: `Bearer ${token}`
            }
        }
    )
        put(getTopicDetailSuccessAction(response.data));
    } catch (error)
    {
        put(getTopicDetailFaliureAction())
    }
}

export function *watchGetTopicDetailSaga() {
    yield takeLatest(getTopicDetailAction.type, getTopicDetailSaga)
}