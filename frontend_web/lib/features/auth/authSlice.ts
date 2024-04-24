// import { createSlice, PayloadAction } from "@reduxjs/toolkit";
// import TokenService from "@/utils/Token.service";
// import jwtDecode from "jwt-decode";

// interface AuthState {
//   user: any; // Define the type of user object according to your application's user model
//   token: string | null;
// }

// const initialState: AuthState = {
//   user: TokenService.getUser(),
//   token: TokenService.getToken()
// };

// const authSlice = createSlice({
//   name: 'auth',
//   initialState,
//   reducers: {
//     setCredentials: (state, action: PayloadAction<{ response: any }>) => {
//       const { response } = action.payload;
//       if (response) {
//         const decodedUser = jwtDecode(response?.accessToken);
//         state.user = decodedUser;
//         state.token = response?.accessToken;
//         TokenService.updateLocalAccessToken(response);
//       }
//     },
//     logOut: (state) => {
//       state.user = null;
//       state.token = null;
//       TokenService.removeUser();
//     }
//   }
// });

// export const { setCredentials, logOut } = authSlice.actions;

// export default authSlice.reducer;

// export const selectCurrentUser = (state: { auth: AuthState }) => {
//   return state.auth.user;
// };

// export const selectCurrentToken = (state: { auth: AuthState }) => {
//   return state.auth.token;
// };
