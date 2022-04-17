import axios from "axios";
import { IPassword } from "../models/passwordInput";

export default class ProfileService {
    static async getProfile() {
        return await axios.get("profile");
    }

    static async changePassword(password: IPassword) {
        return await axios.put("profile/password", password);
    }

    static async deleteProfile(password: IPassword) {
        return await axios.delete("profile");
    }
}
