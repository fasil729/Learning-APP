import authSlice from './features/auth/authSlice';
import quizSlice from './features/quiz/quizSlice';
import subjectesReducer from './features/topics/topicSlice';
;
const rootReducers = {
    subjects: subjectesReducer,
    quiz: quizSlice,
    auth: authSlice
};

export default rootReducers;