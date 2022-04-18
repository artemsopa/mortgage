import axios from "axios";
import { bankInfo, bankInput } from "../models/bankInfo";
import { calculateInput } from "../models/calcInput";

export default class BankService {
    static async getAll() {
        return await axios.get("banks");
    }

    static async getMyBanks() {
        return await axios.get("banks/profile", {
            withCredentials: true,
        });
    }

    static async deleteBank(id: string) {
        return await axios.delete(`banks/profile/${id}`, {
            withCredentials: true
        });
    }

    static async createBank(bank: bankInput) {
        return await axios.post("banks/profile", bank, {
            withCredentials: true
        })
    }

    static async getMortgage(calc: calculateInput) {
        return await axios.post("banks/mortgage", calc)
    }

    static async updateBank(bank: bankInfo) {
        return await axios.put("banks/profile", bank, {
            withCredentials: true
        })
    }
}
