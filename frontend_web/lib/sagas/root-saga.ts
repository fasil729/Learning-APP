import { all, fork } from "redux-saga/effects";
import { watchCreateSubjectSaga } from "./createTopicSaga";


const rootSaga = function* () {
    yield all([
        fork(watchCreateSubjectSaga),
    ])
}

export default rootSaga;