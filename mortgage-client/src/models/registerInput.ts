export class userInputSignUp {
    nick!: string; 
    email!: string; 
    password!: string; 
    confirm!: string;

    constructor(nick: string, email: string, password: string, confirm: string) {
        this.nick = nick;
        this.email = email;
        this.password = password;
        this.confirm = confirm;
    }
}
