import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const inintalState = {
    isLoading : false,
    data: null,
    errors: null,
}
export const noteSlice = createSlice({
    name: 'note',
    initialState: inintalState,
    reducers: {
        addNoteAction: (state: any) => {
            state.isLoading = true;
            state.data = null;
            state.errors = null;
        },
        addNoteSuccessAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = false;
            state.data = action.payload;
            state.errors = null;
        },
        addNoteFailureAction: (state: any, action: PayloadAction<any>) => {
            state.isLoading = false;
            state.data = null;
            state.errors = action.payload;
        },

    }
});

export const { addNoteAction, addNoteSuccessAction, addNoteFailureAction } = noteSlice.actions;
export default noteSlice.reducer;