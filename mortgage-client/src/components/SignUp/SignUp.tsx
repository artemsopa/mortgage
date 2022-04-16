import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { useTypedSelector } from '../../hooks/typedSelectors';
import './SignUp.css'

const SignUp: React.FC = () => {
    const [login, setLogin] = useState<string>('')
    const [password, setPassword] = useState<string>('')

    const { isAuthed, payload } = useTypedSelector(state => state.auth)

    const dispatch = useDispatch();

    return (
        <div className="form-signin">
            <form>
                <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

                <div className="form-floating">
                    <input
                        value={login}
                        onChange={e => setLogin(e.target.value)}
                        type="email"
                        className="form-control"
                        id="floatingInput"
                        placeholder="name@example.com"
                    />
                    <label htmlFor="floatingInput">Email address</label>
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

                <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
            </form>
        </div>
    )
}

export default SignUp;