import React, { Dispatch, SetStateAction, useState } from "react";
import { bankInfo } from "../../models/bankInfo";
import { calculateInput } from "../../models/calcInput";
import BankService from "../../services/bankService";
import "./Modal.css";

type ModalProps = {
  active: boolean;
  setActive: Dispatch<SetStateAction<boolean>>;
  bank: bankInfo;
};

const ValueModal: React.FC<ModalProps> = (props) => {
  const [loan, setLoan] = useState<number>(props.bank.minPayment + 5000);
  const [payment, setPayment] = useState<number>(props.bank.minPayment);

  const [message, setMessage] = useState<string>("");

  const close = (e: any) => {
    e.preventDefault();
    setLoan(props.bank.minPayment + 5000);
    setPayment(props.bank.minPayment);
    setMessage("");
    props.setActive(false);
  };

  const submit = async (e: any) => {
    e.preventDefault();

    let calc = new calculateInput(loan, payment, props.bank.id);
    const response = await BankService.getMortgage(calc);
    setMessage(response.data.message);
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
                onClick={(e) => close(e)}
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
                  <h3 className="h-mb text-center">{props.bank!.title}</h3>
                </div>
                <div className="form-floating mb-1">
                  <input
                    value={loan}
                    onChange={(e) => setLoan(Number(e.target.value))}
                    type="number"
                    min={1}
                    className="form-control rounded-4"
                    id="floatingInitialloan"
                    placeholder="Initial loan"
                  />
                  <label htmlFor="floatingInitialloan">
                    Initial loan: max - {props.bank!.maxLoan}$
                  </label>
                </div>
                <div className="form-floating mb-1">
                  <input
                    value={payment}
                    onChange={(e) => setPayment(Number(e.target.value))}
                    type="number"
                    min={1}
                    className="form-control rounded-4"
                    id="floatingDownpayment"
                    placeholder="Down payment"
                  />
                  <label htmlFor="floatingDownpayment">
                    Down payment: min - {props.bank!.minPayment}$
                  </label>
                </div>

                <button
                  onClick={(e) => submit(e)}
                  className="w-100 mb-2 btn btn-lg rounded-4 btn-outline-primary"
                  type="submit"
                >
                  Calculate
                </button>
                <div className="mt-10px">
                  <h3 className="text-center">{message}</h3>
                </div>
              </form>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default ValueModal;
