export interface LoginAction {
    type: string;
    payload?: string;
}

export enum LoginActionTypes {
    LOGIN = 'LOGIN',
    LOGIN_SUCCESS = 'LOGIN_SUCCESS',
    LOGIN_ERROR = 'LOGIN_ERROR',
    LOGOUT = 'LOGOUT'
}

interface LoginInitAction {
    type: LoginActionTypes.LOGIN
}

interface LoginSuccessAction {
    type: LoginActionTypes.LOGIN_SUCCESS
}

interface LoginErrorAction {
    type: LoginActionTypes.LOGIN_ERROR
}

interface LogoutInitAction {
    type: LoginActionTypes.LOGOUT
}

export type AuthAction = LoginInitAction | LoginSuccessAction | LoginErrorAction | LogoutInitAction
