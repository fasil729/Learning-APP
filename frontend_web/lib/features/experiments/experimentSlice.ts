import { FormDataModel } from "@/types/formData";
import { Topic, TopicState } from "@/types/topic";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const experimentInitalState = {
   isLoading: false,
   data : {
    contentLink: null,
    experimentName: null,
    topicName: null,
   },
   errors: null
};

export const experimentSlice = createSlice({
    name: 'subjects',
    initialState: experimentInitalState,
    reducers: {
        createExperimentAction: (state: any, action: PayloadAction<FormDataModel | null>) => {
            state.topic.data = action.payload;
            state.topic.errors = "";
            state.topic.isLoading = true;
          },
      
          createExperimentSuccessAction: (state: any, action: PayloadAction<Topic | null>) => {
            state.topic.isLoading = false;
            state.topic.data = action.payload;
            state.topic.isSuccess = true;
            console.log(action.payload);
          },
          createExperimentFailureAction: (state: any) => {
            state.topic.isLoading = false;
            state.topic.errors = "error encountered"
          },
          getExperimentsForChapterAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = true;
            state.data = null;

          },
          getExperimentsForChapterSuccessAction: (state: any, action: any) => {
            state.isLoading = false;
            state.data = action.payload;

          },
          getExperimentsForChapterFailureAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = false;
            state.errors = action.payload;

          },
          getExperimentsForSubjectAction: (state: any, action: any) => {
            state.isLoading = true;
            state.data = null;
          },

          getExperimentsForSubjectSuccessAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = false;
            state.data = action.payload

          },
          getExperimntsForSubjectFailureAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = false;
            state.errors = action.payload
          }

    }
})

export const { createExperimentAction, createExperimentSuccessAction, createExperimentFailureAction } = experimentSlice.actions;

export default experimentSlice.reducer;