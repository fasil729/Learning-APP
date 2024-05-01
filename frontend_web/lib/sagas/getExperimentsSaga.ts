import { put, takeLatest } from "redux-saga/effects";
import axios, { AxiosResponse } from "axios";
import { 
    getExperimentsForChapterAction, 
    getExperimentsForChapterFailureAction, 
    getExperimentsForChapterSuccessAction 
} from "../features/experiments/experimentSlice";


function *getExperimentForChapterSaga(action: any) {
    let response: AxiosResponse;
    const token = localStorage.getItem("accessToken")
    console.log("getting experiment")

    try {
        response = yield axios.get(`https://learning-app-idt8.onrender.com/experiments/chapter/${action.payload}`,
     {headers: {
            Authorization: `Bearer ${token}`
        }});
        console.log(response);
        yield put(getExperimentsForChapterSuccessAction(response.data));
    }catch (error) {
        yield put(getExperimentsForChapterFailureAction(error));
    }
}

export function *watchExperimentForChapterSaga() {
  yield takeLatest(getExperimentsForChapterAction.type, getExperimentForChapterSaga);
}