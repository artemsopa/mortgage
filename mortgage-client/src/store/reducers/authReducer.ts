import { LoginAction, LoginActionTypes } from "../actions/loginAction";

export interface AuthState {
    isAuthed: boolean;
    payload?: string;
}

const initalState: AuthState = {
    isAuthed: false,
}

export const loginReducer = (state = initalState, action: LoginAction): AuthState => {
    switch (action.type) {
        case LoginActionTypes.LOGIN: return { isAuthed: true };
        case LoginActionTypes.LOGIN_SUCCESS: return { isAuthed: true, payload: action.payload };
        case LoginActionTypes.LOGIN_ERROR: return { isAuthed: false, payload: action.payload };
        case LoginActionTypes.LOGOUT: return { isAuthed: false };
        default: return state;
    }
}
