import { combineReducers } from '@reduxjs/toolkit';
import authSlice from './features/auth/authSlice';
import quizSlice from './features/quiz/quizSlice';
import lessonSlice from './features/topics/lessonSlice';
import subjectesReducer from './features/topics/topicSlice';
import experimentSlice from './features/experiments/experimentSlice';

const rootReducers = combineReducers({
    subjects: subjectesReducer,
    quiz: quizSlice,
    auth: authSlice,
    lesson: lessonSlice,
    experiment: experimentSlice
});

export default rootReducers;