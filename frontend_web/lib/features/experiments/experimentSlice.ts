import { FormDataModel } from "@/types/formData";
import { Topic, TopicState } from "@/types/topic";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const experimentInitalState = {
   isLoading: false,
   data : null,
   errors: null
};

export const experimentSlice = createSlice({
    name: 'experiments',
    initialState: experimentInitalState,
    reducers: {
        createExperimentAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = true;
            state.data.experimentName = null;
          },
      
          createExperimentSuccessAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = false;
            state.data.experimentname = action.payload.experiment_name
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

export const { createExperimentAction, createExperimentSuccessAction, createExperimentFailureAction, getExperimentsForChapterAction, getExperimentsForChapterSuccessAction, getExperimentsForChapterFailureAction } = experimentSlice.actions;

export default experimentSlice.reducer;