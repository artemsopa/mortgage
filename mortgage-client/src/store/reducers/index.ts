import { combineReducers } from "redux";
import { loginReducer } from "./authReducer";

export const rootReducer = combineReducers({
    login: loginReducer
})

export type RootState = ReturnType<typeof rootReducer>