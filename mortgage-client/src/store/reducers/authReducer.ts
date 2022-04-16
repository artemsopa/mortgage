import { Action, AuthActionTypes, AuthState } from "../actions/authAction";

const initalState: AuthState = {
    isAuthed: false,
}

export const authReducer = (state = initalState, action: Action): AuthState => {
    switch (action.type) {
        case AuthActionTypes.LOGIN: return { isAuthed: true };
        case AuthActionTypes.LOGIN_SUCCESS: return { isAuthed: true, payload: action.payload };
        case AuthActionTypes.LOGIN_ERROR: return { isAuthed: false, payload: action.payload };
        default: return state;
    }
}
