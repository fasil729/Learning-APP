import authSlice from './features/auth/authSlice';
import quizSlice from './features/quiz/quizSlice';
import lessonSlice from './features/topics/lessonSlice';
import subjectesReducer from './features/topics/topicSlice';
import {quizApi} from "./features/quiz/quizApi"
;
const rootReducers = {
    subjects: subjectesReducer,
    [quizApi.reducerPath]:quizApi.reducer,

    quiz: quizSlice,
    auth: authSlice,
    lesson: lessonSlice
};

export default rootReducers;