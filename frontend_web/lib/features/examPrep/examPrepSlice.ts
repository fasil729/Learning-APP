import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface ExamPrepState {
  isLoading: boolean;
  data: string;
  errors: string | null;
}

const initialState: ExamPrepState = {
  isLoading: false,
  data: "",
  errors: null,
};

const examPrepSlice = createSlice({
  name: 'examPrep',
  initialState,
  reducers: {
    generateExamPrepRequest: (state) => {
      state.isLoading = true;
      state.errors = null;
    },
    setExamPrepData: (state, action: PayloadAction<string>) => {
      state.isLoading = false;
      state.data = action.payload;
      state.errors = null;
    },
    generateExamPrepFailure: (state, action: PayloadAction<string>) => {
      state.isLoading = false;
      state.errors = action.payload;
    },
  },
});

export const { generateExamPrepRequest, setExamPrepData, generateExamPrepFailure } = examPrepSlice.actions;

export default examPrepSlice.reducer;
