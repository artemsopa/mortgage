import axios from "axios";
import { userInputSignIn } from "../models/loginInput";
import { userInputSignUp } from "../models/registerInput";

export default class AuthService {
    static async signIn(loginInput: userInputSignIn) {
        return await axios.post("auth/sign-in", loginInput, {
            withCredentials: true,
        });
    }

    static async signUp(registerInput: userInputSignUp) {
        return await axios.post("auth/sign-up", registerInput);
    }

    static async refreshSession() {
        return await await axios.post("auth/refresh", {}, {
            withCredentials: true
        })
    }

    static async logout() {
        return await axios.post("auth/logout", {}, {
            withCredentials: true
        })
    }
}
