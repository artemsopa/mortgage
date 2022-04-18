import { Dispatch } from "redux"
import { userInputSignIn } from "../../models/loginInput"
import { LoginActionTypes } from "../actions/loginAction"
import AuthService from "../../services/authService"

export const signInCreator = (loginInput: userInputSignIn) => {
    return async (dispatch: Dispatch) => {
        try {
            //dispatch({type: LoginActionTypes.LOGIN})
            const response = await AuthService.signIn(loginInput)
            dispatch({type: LoginActionTypes.LOGIN_SUCCESS, payload: response.data.message})
        } catch(e) {
            dispatch({type: LoginActionTypes.LOGIN_ERROR, payload: "something went wrong"})
        }
    }
}

export const logoutCreator = () => {
    return async (dispatch: Dispatch) => {
        try {
            const response = await AuthService.logout()
            dispatch({type: LoginActionTypes.LOGOUT})
        } catch(e) {
            console.log(e)
        }
    }
}

export const refreshCreator = () => {
    return async (dispatch: Dispatch) => {
        try {
            const response = await AuthService.refreshSession()
            dispatch({type: LoginActionTypes.LOGIN})
        } catch(e) {
            console.log(e)
        }
    }
}
