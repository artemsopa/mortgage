export class calculateInput {
    loan!: number;
    payment!: number;
    bankId!: string
    constructor(loan: number, payment: number, bankId: string) {
        this.loan = loan;
        this.payment = payment;
        this.bankId = bankId;
    }
}