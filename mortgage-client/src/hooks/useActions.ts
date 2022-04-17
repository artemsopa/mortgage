import { useDispatch } from "react-redux"
import { bindActionCreators } from "redux"
import * as AuthActionCreators from '../store/action-creators/auth'

export const useActions = () => {
    const dispatch = useDispatch()
    return bindActionCreators(AuthActionCreators, dispatch)
}