import { put, takeLatest } from "redux-saga/effects";
import { signInAction, signInFailureAction, signInSuccessAction } from "../features/auth/authSlice";
import axios, { AxiosResponse } from "axios";

function* signInSaga(action: any) {
    let response: AxiosResponse;
    console.log("signing in")

    try {
        response = yield axios.post(`https://learning-app-idt8.onrender.com/users/login`,
        action.payload);
        localStorage.setItem("accessToken", response.data?.accessToken);
        localStorage.setItem("username", response.data?.username);
        yield put(signInSuccessAction(response.data));

    }catch (error) {
        yield put(signInFailureAction());
    }
}

export function* watchSignInSaga() {
  yield takeLatest(signInAction.type, signInSaga);
}