import axios from "axios";
import React, { useState } from "react";
import { useEffect } from "react";
import { bankInfo } from "../../models/bankInfo";
import BankService from "../../services/bankService";
import ValueModal from "../Modal/ValueModal";
import "./AllBanks.css";

const AllBanks: React.FC = () => {
  const [modalActivate, setModalActive] = useState(false);
  const [currentBank, setCurrentBank] = useState<bankInfo>();

  const openModal = (bank: bankInfo) => {
    setCurrentBank(bank);
    setModalActive(true);
  };

  const [banks, setBanks] = useState<Array<bankInfo>>(Array<bankInfo>());

  const getAllBanks = async () => {
    return await (await BankService.getAll()).data;
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
      {currentBank && (
        <ValueModal
          bank={currentBank}
          active={modalActivate}
          setActive={setModalActive}
        />
      )}
      <div className="apy container">
        <div className="row row-cols-1 row-cols-sm-2 row-cols-md-5 g-3">
          {banks.map((bank, index) => {
            return (
              <div key={index}>
                <div className="col">
                  <div className="card shadow-sm">
                    <div className="card-body">
                      <p className="text-center card-text">{bank.title}</p>
                      <div>
                        <div
                          onClick={() => openModal(bank)}
                          className="btn btn-primary btn-gm"
                        >
                          Get Mortgage
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default AllBanks;
