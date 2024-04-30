import { put, takeLatest } from "redux-saga/effects";
import { signInAction, signInFailureAction, signInSuccessAction } from "../features/auth/authSlice";
import axios, { AxiosResponse } from "axios";
import { cookies } from "next/headers";

function *signInSaga(action: any) {
    let response: AxiosResponse;

    try {
        response = yield axios.post(`https://learning-app-idt8.onrender.com/users/login`,
        action.payload);
        cookies().set("accessToken", response.data?.accessToken, { secure: true });
        cookies().set("username", response.data?.username, { secure: true })
        yield put(signInSuccessAction(response.data));

    }catch (error) {
        yield put(signInFailureAction());
    }
}

export function *watchSignInSaga() {
  yield takeLatest(signInAction.type, signInSaga);
}