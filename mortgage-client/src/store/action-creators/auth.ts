import { Dispatch } from "redux"
import { ILogin } from "../../models/loginInput"
import { AuthActionTypes } from "../actions/authAction"
import AuthService from "../../services/authService"
import { Response } from "../../models/response"

export const Login = (loginInput: ILogin) => {
    return async (dispatch: Dispatch) => {
        try {
            dispatch({type: AuthActionTypes.LOGIN})
            const reponse = await AuthService.signIn(loginInput)
            dispatch({type: AuthActionTypes.LOGIN_SUCCESS, payload: reponse.data.message})
        } catch(e) {
            dispatch({type: AuthActionTypes.LOGIN_ERROR, payload: "something went wrong"})
        }
    }
}
