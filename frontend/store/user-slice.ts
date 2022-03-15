import { createSlice, PayloadAction } from "@reduxjs/toolkit"
import { NewUser, User } from "../src/generated/generated"

interface UserSliceState {
    user: User | null
    token: string | null
}

const initialState:UserSliceState = {
    user: null,
    token: null
}

const userSlice = createSlice({
    name: 'user',
    initialState, 
    reducers: {
        login: (state, action:PayloadAction<UserSliceState>) => {
            state.user = action.payload.user
            state.token = action.payload.token
        },
        logout: state => {
            state.user = null
            state.token = null
        }
    }
})

export const {login, logout} = userSlice.actions;


export default userSlice.reducer