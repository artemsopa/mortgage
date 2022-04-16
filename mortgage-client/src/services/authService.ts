import $api from "../http";
import { AxiosResponse } from "axios";
import { ILogin } from "../models/loginInput";
import { IRegister } from "../models/registerInput";
import { Response } from "../models/response";

export default class AuthService {
    static async signIn(loginInput: ILogin): Promise<AxiosResponse<Response>> {
        return await $api.post<Response>("auth/sign-in", loginInput);
    }

    static async signUp(registerInput: IRegister): Promise<AxiosResponse<Response>> {
        return await $api.post<Response>("auth/sign-up", registerInput);
    }

    static async logout(): Promise<AxiosResponse<Response>> {
        return await $api.post<Response>("auth/logout")
    }
}
