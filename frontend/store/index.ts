import { configureStore } from "@reduxjs/toolkit";
import { useDispatch } from "react-redux";
import userSlice from "./user-slice";

const store = configureStore({
    reducer: {
        user: userSlice,
    }
})

type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch;
export const useUserDispatch = () => useDispatch<AppDispatch>();
export const selectUser = (state:RootState) => state.user.user
export const selectToken = (state:RootState) => state.user.token
export default store;
