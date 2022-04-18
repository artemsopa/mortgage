import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useActions } from "../../hooks/useActions";
import { useTypedSelector } from "../../hooks/useTypedSelectors";
import { userInputSignIn } from "../../models/loginInput";
import "./SignIn.css";

const SignIn: React.FC = () => {
  const [login, setLogin] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const navigate = useNavigate();

  const { isAuthed, payload } = useTypedSelector((state) => state.login);
  const { signInCreator } = useActions();

  const handleClick = async (e: any) => {
    e.preventDefault();

    signInCreator(new userInputSignIn(login, password));

    navigate("/");
  };

  return (
    <div className="main">
      {isAuthed && <h1>You are already logged in. Logout</h1>}
      {!isAuthed && (
        <div className="form-signin">
          <form>
            <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

            <div className="form-floating">
              <input
                value={login}
                onChange={(e) => setLogin(e.target.value)}
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
                onChange={(e) => setPassword(e.target.value)}
                type="password"
                className="form-control"
                id="floatingPassword"
                placeholder="Password"
              />
              <label htmlFor="floatingPassword">Password</label>
            </div>

            <button
              onClick={handleClick}
              className="b-w-p w-100 btn btn-lg btn-primary"
              type="submit"
            >
              Sign in
            </button>
          </form>
        </div>
      )}
    </div>
  );
};

export default SignIn;
