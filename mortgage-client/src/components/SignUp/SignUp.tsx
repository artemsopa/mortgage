import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useTypedSelector } from '../../hooks/useTypedSelectors';
import { userInputSignUp } from '../../models/registerInput';
import AuthService from '../../services/authService';
import './SignUp.css'

const SignUp: React.FC = () => {
    const [nick, setNick] = useState<string>('')
    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const [confirm, setConfirm] = useState<string>('')

    const { isAuthed, payload } = useTypedSelector((state) => state.login);

    const navigate = useNavigate();

    const handleClick = (e: any) => {
        e.preventDefault();

        try {
            AuthService.signUp(new userInputSignUp(nick, email, password, confirm)).then(
                () => navigate(`/sign-in`)
            )
        } catch (error) {
            console.log(error)
        }
    }

    return (
        <div className='main'>
            {isAuthed && <h1>You are already logged in. Logout</h1>}
            {!isAuthed && <div className="form-signin">
                <form>
                    <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

                    <div className="form-floating">
                        <input
                            value={nick}
                            onChange={e => setNick(e.target.value)}
                            type="text"
                            className="form-control"
                            id="floatingNick"
                            placeholder="nick"
                        />
                        <label htmlFor="floatingInput">Nick</label>
                    </div>
                    <div className="form-floating">
                        <input
                            value={email}
                            onChange={e => setEmail(e.target.value)}
                            type="email"
                            className="form-control"
                            id="floatingEmail"
                            placeholder="name@example.com"
                        />
                        <label htmlFor="floatingInput">Email</label>
                    </div>
                    <div className="form-floating">
                        <input
                            value={password}
                            onChange={e => setPassword(e.target.value)}
                            type="password"
                            className="form-control"
                            id="floatingPassword"
                            placeholder="Password"
                        />
                        <label htmlFor="floatingPassword">Password</label>
                    </div>
                    <div className="form-floating">
                        <input
                            value={confirm}
                            onChange={e => setConfirm(e.target.value)}
                            type="password"
                            className="form-control"
                            id="floatingConfirm"
                            placeholder="Confirm"
                        />
                        <label htmlFor="floatingPassword">Password Repeat</label>
                    </div>

                    <button onClick={ handleClick } className="w-100 btn btn-lg btn-primary" type="submit">Sign Up</button>
                </form>
            </div>}
        </div>
    )
}

export default SignUp;