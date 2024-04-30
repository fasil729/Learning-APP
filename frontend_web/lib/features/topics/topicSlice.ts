import { FormDataModel } from "@/types/formData";
import { Topic, TopicState } from "@/types/topic";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const topicInitalState = {
    topics: {
        isLoading: false,
        data: null,
        errors: null,
    },
    topic: {
        isLoading: false,
        data: {
          chapters: null,
          subject: null,
        },
        errors: null,
        isSuccess: false,
    }
};

export const topicSlice = createSlice({
    name: 'subjects',
    initialState: topicInitalState,
    reducers: {
        createSubjectAction: (state: any, action: PayloadAction<FormDataModel | null>) => {
            state.topic.data = action.payload;
            state.topic.errors = "";
            state.topic.isLoading = true;
          },
      
          createSubjectSuccessAction: (state: any, action: PayloadAction<Topic | null>) => {
            state.topic.isLoading = false;
            state.topic.data = action.payload;
            state.topic.isSuccess = true;
            console.log(action.payload);
          },
          createSubjectFailureAction: (state: any) => {
            state.topic.isLoading = false;
            state.topic.errors = "error encountered"
          },
    }
})

export const { createSubjectAction, createSubjectSuccessAction, createSubjectFailureAction } = topicSlice.actions;

export default topicSlice.reducer;