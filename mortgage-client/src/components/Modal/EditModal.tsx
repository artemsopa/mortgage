import { render } from "@testing-library/react";
import React, { Dispatch, SetStateAction, useState } from "react";
import { bankInfo } from "../../models/bankInfo";
import BankService from "../../services/bankService";
import "./Modal.css";

type ModalProps = {
  active: boolean;
  setActive: Dispatch<SetStateAction<boolean>>;
  getAllBanksInfo: () => Promise<void>;
  bank: bankInfo;
};

const EditModal: React.FC<ModalProps> = (props) => {
  const [title, setTitle] = useState<string>(props!.bank.title);
  const [rate, setRate] = useState<number>(props!.bank.rate);
  const [maxLoan, setMaxLoan] = useState<number>(props!.bank.maxLoan);
  const [minPayment, setMinPayment] = useState<number>(props!.bank.minPayment);
  const [term, setTerm] = useState<number>(props!.bank.loanTerm);

  const getFixed = (num: number): number => {
    return Math.round(num * Math.pow(10, 2)) / Math.pow(10, 2);
  };

  const submit = async (e: any) => {
    e.preventDefault();

    let bank = new bankInfo(
      props.bank.id,
      title,
      getFixed(rate),
      maxLoan,
      minPayment,
      term,
      props.bank.userId
    );
    await BankService.updateBank(bank);
    await props.getAllBanksInfo();
    props.setActive(false);
  };

  return (
    <div
      className={
        props.active
          ? "md-bgt modal modal-signin position-absolute d-block py-5"
          : ""
      }
      tabIndex={-1}
      role="dialog"
      id="modalSignin"
    >
      {props.active && (
        <div className="modal-dialog" role="document">
          <div className="modal-content rounded-5 shadow">
            <div className="modal-header p-5 pb-4 border-bottom-0">
              <h2 className="fw-bold mb-0">Update Bank Info</h2>
              <button
                onClick={() => props.setActive(false)}
                type="button"
                className="btn-close"
                data-bs-dismiss="modal"
                aria-label="Close"
              ></button>
            </div>
            <div
              className="modal-body p-5 pt-0"
              onClick={(e) => e.stopPropagation()}
            >
              <form>
                <div className="form-floating mb-1">
                  <input
                    type="text"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    className="form-control rounded-4"
                    id="floatingTitle"
                    placeholder="Bank title"
                  />
                  <label htmlFor="floatingTitle">Bank title</label>
                </div>
                <div className="form-floating mb-1">
                  <input
                    type="number"
                    value={rate * 100}
                    onChange={(e) => setRate(Number(e.target.value) / 100)}
                    min={1}
                    className="form-control rounded-4"
                    id="floatingRate"
                    placeholder="Interest rate"
                  />
                  <label htmlFor="floatingRate">Interest rate (%)</label>
                </div>
                <div className="form-floating mb-1">
                  <input
                    type="number"
                    value={maxLoan}
                    onChange={(e) => setMaxLoan(Number(e.target.value))}
                    min={1}
                    className="form-control rounded-4"
                    id="floatingMaximum"
                    placeholder="Maximum loan"
                  />
                  <label htmlFor="floatingMaximum">Maximum loan ($)</label>
                </div>
                <div className="form-floating mb-1">
                  <input
                    type="number"
                    value={minPayment}
                    onChange={(e) => setMinPayment(Number(e.target.value))}
                    min={1}
                    className="form-control rounded-4"
                    id="floatingMinimum"
                    placeholder="Minimum down payment"
                  />
                  <label htmlFor="floatingMinimum">
                    Minimum down payment ($)
                  </label>
                </div>
                <div className="form-floating mb-1">
                  <input
                    type="number"
                    value={term}
                    onChange={(e) => setTerm(Number(e.target.value))}
                    min={1}
                    className="form-control rounded-4"
                    id="floatingLoan"
                    placeholder="Loan term"
                  />
                  <label htmlFor="floatingLoan">Loan term (months)</label>
                </div>
                <button
                  onClick={(e) => submit(e)}
                  className="w-100 mb-2 btn btn-lg rounded-4 btn-outline-warning"
                  type="submit"
                >
                  Edit Info
                </button>
              </form>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default EditModal;
