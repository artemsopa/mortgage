import axios from "axios";
import { useDispatch } from "react-redux";
import { useActions } from "../hooks/useActions";
import { useTypedSelector } from "../hooks/useTypedSelectors";
import { LoginActionTypes } from "../store/actions/loginAction";

axios.defaults.baseURL = "http://localhost:8000/api/v1/"


// axios.interceptors.response.use(resp => resp, async error => {

//     const dispatch = useDispatch();
//     if (error.response.status === 401) {
//         dispatch({type: LoginActionTypes.LOGIN_ERROR, payload: "something went wrong"})

//         const { status } = await axios.post("auth/refresh", {}, {
//             withCredentials: true
//         })

//         if (status === 200) {
//             dispatch({type: LoginActionTypes.LOGIN, payload: "session refreshed"})
//             return axios(error.config)
//         } else if (status === 500) {
//             dispatch({type: LoginActionTypes.LOGIN_ERROR, payload: "cookies expired"})
//             return axios(error.config)
//         }
//     }
//     return error;
// });
