import { FormDataModel } from "@/types/formData";
import { Topic, TopicState } from "@/types/topic";
import { User, UserResponse } from "@/types/user";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const topicInitalState = {
      isLoading: false,
      data: {
          token: null,
          isAuthenticated: false,
          user: null
      },
      errors: null,
};

export const authSlice = createSlice({
    name: 'auth',
    initialState: topicInitalState,
    reducers: {
        signInAction: (state: any, action: PayloadAction<User | null>) => {
            state.data.user = action.payload;
            state.errors = "";
            state.isLoading = true;
          },
      
          signInSuccessAction: (state: any, action: PayloadAction<UserResponse | null>) => {
            console.log("here singin in user")
            state.data = action.payload;
            state.data.isAuthenticated = true;            
            state.isLoading = false;
            console.log(action.payload);
          },
          signInFailureAction: (state: any) => {
            state.isLoading = false;
            state.errors = "error encountered"
          },
          signUpAction: (state: any, action: any) => {
            state.data = null;
            state.errors = "";
            state.isLoading = true;
          },
      
          signUpSuccessAction: (state: any, action: PayloadAction<Topic | null>) => {
            state.isLoading = false;
            state.data = action.payload;
            console.log(action.payload);
          },
          signUpFailureAction: (state: any, action: PayloadAction) => {
            state.isLoading = false;
            state.errors = "error encountered"
          },
          logoutAction: (state: any) => {
            state.data.isAuthenticated = false;
            state.data.accessToken = null;
            state.data.user = null;
          }
    }
})

export const {
    signInAction,
    signInSuccessAction,
    signInFailureAction,
    signUpAction,
    signUpSuccessAction,
    signUpFailureAction,
    logoutAction
} = authSlice.actions;

export default authSlice.reducer;