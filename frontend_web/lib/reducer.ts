import quizSlice from './features/quiz/quizSlice';
import subjectesReducer from './features/topics/topicSlice';
;
const rootReducers = {
    subjects: subjectesReducer,
    quiz: quizSlice
};

export default rootReducers;