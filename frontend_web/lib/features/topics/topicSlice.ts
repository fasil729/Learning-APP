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
        data: null,
        errors: null,
    }
};

export const topicSlice = createSlice({
    name: 'subjects',
    initialState: topicInitalState,
    reducers: {
        createSubjectAction: (state: TopicState, action: PayloadAction<FormDataModel | null>) => {
            state.topic.data = null;
            state.topic.errors = "";
            state.topic.isLoading = true;
          },
      
          createSubjectSuccessAction: (state: TopicState, action: PayloadAction<Topic | null>) => {
            state.topic.isLoading = false;
            state.topic.data = action.payload;
            console.log(action.payload);
          },
          createSubjectFailureAction: (state: TopicState, action: PayloadAction) => {
            state.topic.isLoading = false;
            state.topic.errors = "error encountered"
          },
    }
})

export const { createSubjectAction, createSubjectSuccessAction, createSubjectFailureAction } = topicSlice.actions;

export default topicSlice.reducer;