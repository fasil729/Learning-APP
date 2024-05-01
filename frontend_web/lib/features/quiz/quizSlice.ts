import { createSlice } from "@reduxjs/toolkit";

export const quizSlice=createSlice({
   name:"quiz",
    initialState:{
        number_of_quizzes: 5,
        prompt: "",
        topics: [],
        quizData:[]
      },

    reducers:{
        useGenerateQuiz(state,action){
            state.number_of_quizzes=action.payload.number_of_quizzes
            state.prompt=action.payload.prompt
          

        },
        useTopicList(state, action) {
            state.topics = action.payload.topics;
          },
          setQuizData(state, action) {
            state.quizData = action.payload.quizData;
          }


    }


});


export const {useGenerateQuiz,useTopicList,setQuizData}=quizSlice.actions;

export default quizSlice.reducer;