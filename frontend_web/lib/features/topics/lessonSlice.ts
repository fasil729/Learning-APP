import { FormDataModel } from "@/types/formData";
import { Topic, TopicState } from "@/types/topic";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const lessonInitalState = {
    isLoading: false,
    data: null,
    errors : null,
    isSuccess: false
};

export const lessonSlice = createSlice({
    name: 'lesson',
    initialState: lessonInitalState,
    reducers: {
        getTopicDetailAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = false;
            state.data = null;
            state.isSuccess = true;
          },
      
        getTopicDetailSuccessAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = false;
            state.data = action.payload;
          },
        getTopicDetailFaliureAction: (state: any) => {
            state.isLoading = false;
            state.errors = "error encountered"
          },
    }
})

export const { getTopicDetailAction, getTopicDetailSuccessAction, getTopicDetailFaliureAction } = lessonSlice.actions;

export default lessonSlice.reducer;