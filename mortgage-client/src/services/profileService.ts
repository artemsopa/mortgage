import $api from "../http";
import { AxiosResponse } from "axios";
import { IPassword } from "../models/passwordInput";
import { IUser } from "../models/user";

export default class ProfileService {
    static async getProfile(): Promise<AxiosResponse<IUser>> {
        return await $api.get<IUser>("profile");
    }

    static async changePassword(password: IPassword): Promise<AxiosResponse<Response>> {
        return await $api.put<Response>("profile/password", password);
    }

    static async deleteProfile(password: IPassword): Promise<AxiosResponse<Response>> {
        return await $api.delete<Response>("profile");
    }
}
