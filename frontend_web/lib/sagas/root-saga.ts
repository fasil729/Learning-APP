import { all, fork } from "redux-saga/effects";
import { watchCreateSubjectSaga } from "./createTopicSaga";
import { watchSignInSaga } from "./signInSaga";
import { watchSignUpSaga } from "./signUpSaga";


const rootSaga = function* () {
    yield all([
        fork(watchCreateSubjectSaga),
        fork(watchSignInSaga),
        fork(watchSignUpSaga)
    ])
}

export default rootSaga;