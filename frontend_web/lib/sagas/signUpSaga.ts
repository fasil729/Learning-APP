import { put, takeLatest } from "redux-saga/effects";
import { signUpAction, signUpFailureAction, signUpSuccessAction } from "../features/auth/authSlice";
import axios, { AxiosResponse } from "axios";

function *signUpSaga(action: any) {
    let response: AxiosResponse;
    try {
        response = yield axios.post(`https://learning-app-idt8.onrender.com/users/login`,
        action.payload);
        yield put(signUpSuccessAction(response.data));

    }catch (error) {
        yield put(signUpFailureAction());
    }
}

export function *watchSignUpSaga() {
  yield takeLatest(signUpAction.type, signUpSaga);
}