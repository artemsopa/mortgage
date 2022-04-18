import React, { useEffect } from "react";
import { useState } from "react";
import { bankInfo, bankInput } from "../../models/bankInfo";
import BankService from "../../services/bankService";
import AddModal from "../Modal/Add.Modal";
import EditModal from "../Modal/EditModal";
import "./MyBanks.css";

const MyBanks: React.FC = () => {
    const [modalAddActivate, setModalAddActive] = useState(false);

    const [modalActivate, setModalActive] = useState(false);
    const emptyBank = new bankInput("", 0.01, 1, 1, 1)
    const [currentBank, setCurrentBank] = useState<bankInfo>();

    const openModal = (bank: bankInfo) => {
        setCurrentBank(bank)
        setModalActive(true);
    }

    const editBanks = async () => {
        // const banks = await getAllBanks();
        // if (banks) setBanks(banks);
    }

    const removeBank = async (e: any, id: string, index: number) => {
        e.preventDefault();

        await BankService.deleteBank(id);
        getAllBanksInfo();
    }

    const [banks, setBanks] = useState<Array<bankInfo>>(Array<bankInfo>());

    const getAllBanks = async () => {
        return await (await BankService.getMyBanks()).data
    };

    const getAllBanksInfo = async () => {
        const banks = await getAllBanks();
        if (banks) setBanks(banks);
    };

    

    useEffect(() => {
        getAllBanksInfo();
    }, []);

    return (
        <div className="album py-5">
            <div className="add-b">
                <div onClick={() => setModalAddActive(true)} className="add-b-d btn btn-success btn-gm">Add New Bank</div>
            </div>
            {currentBank &&
            <EditModal
                  getAllBanksInfo={getAllBanksInfo} 
                  bank={currentBank}
                  active={modalActivate}
                  setActive={setModalActive}
                />}
            <AddModal
            getAllBanksInfo={getAllBanksInfo} 
            bank={emptyBank} 
            active={modalAddActivate} 
            setActive={setModalAddActive} />
            <div className="container">
                <div className="row row-cols-1 row-cols-sm-2 row-cols-md-3 g-3">
                    {banks.map((bank, index) => {
                        return (
                            <div key={index}>
                                <div className="col">
                                    <div className="card shadow-sm">
                                        <div className="card-body">
                                            <div className="mb-5px">
                                            <p 
                                            onClick={(e) => removeBank(e, bank.id, index)}
                                            className="db text-center card-text">Delete</p>
                                            </div>
                                            <p className="text-center card-text">{bank.title}</p>
                                            <p className="card-text">Interest rate: {bank.rate * 100}%</p>
                                            <p className="card-text">Maximum loan: {bank.maxLoan}$</p>
                                            <p className="card-text">Minimum down payment: {bank.minPayment}$</p>
                                            <p className="card-text">Loan term: {bank.loanTerm} months</p>
                                            <div>
                                                <div onClick={() => openModal(bank)} className="btn btn-warning btn-gm">Edit Info</div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>)
                    })}
                </div>
            </div>
        </div>
    );
};

export default MyBanks;
