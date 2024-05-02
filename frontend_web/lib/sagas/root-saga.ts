import { all, fork } from "redux-saga/effects";
import { watchCreateSubjectSaga } from "./createTopicSaga";
import { watchSignInSaga } from "./signInSaga";
import { watchSignUpSaga } from "./signUpSaga";
import { watchExperimentForChapterSaga } from "./getExperimentsSaga";
import { watchAddNoteSaga } from "./addNoteSaga";


const rootSaga = function* () {
    yield all([
        fork(watchCreateSubjectSaga),
        fork(watchSignInSaga),
        fork(watchSignUpSaga),
        fork(watchExperimentForChapterSaga),
        fork(watchAddNoteSaga)
    ])
}

export default rootSaga;