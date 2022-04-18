export class bankInfo {
    id!: string;
    title!: string;
    rate!: number;
    maxLoan!: number;
    minPayment!: number;
    loanTerm!: number;
    userId!: string;

    constructor(id: string, title: string, rate: number, maxLoan: number, minPayment: number, loanTerm: number, userId: string) {
        this.id = id;
        this.title = title;
        this.rate = rate;
        this.maxLoan = maxLoan;
        this.minPayment = minPayment;
        this.loanTerm = loanTerm;
        this.userId = userId;
    }
}

export class bankInput {
    title!: string;
    rate!: number;
    maxLoan!: number;
    minPayment!: number;
    loanTerm!: number;

    constructor(title: string, rate: number, maxLoan: number, minPayment: number, loanTerm: number) {
        this.title = title;
        this.rate = rate;
        this.maxLoan = maxLoan;
        this.minPayment = minPayment;
        this.loanTerm = loanTerm;
    }
}
