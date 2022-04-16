export interface AuthState {
    isAuthed: boolean;
    payload?: string;
}

export interface Action {
    type: string;
    payload?: string;
}

export enum AuthActionTypes {
    LOGIN = 'LOGIN',
    LOGIN_SUCCESS = 'LOGIN_SUCCESS',
    LOGIN_ERROR = 'AUTH_ERROR'
}

interface AuthInitAction {
    type: AuthActionTypes.LOGIN
}

interface AuthSuccessAction {
    type: AuthActionTypes.LOGIN_SUCCESS
}

interface AuthErrorAction {
    type: AuthActionTypes.LOGIN_ERROR
}

export type AuthAction = AuthInitAction | AuthSuccessAction | AuthErrorAction
