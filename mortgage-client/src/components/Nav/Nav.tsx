import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useActions } from '../../hooks/useActions';
import { useTypedSelector } from '../../hooks/useTypedSelectors';
import './Nav.css'

const Nav: React.FC = () => {
    const navigate = useNavigate();

    const { isAuthed, payload } = useTypedSelector(state => state.login)
    const { logoutCreator } = useActions();

    const handleClick = async (e: any) => {
        e.preventDefault();

        logoutCreator()
        
        navigate("/")
    }

    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
                <div className="collapse navbar-collapse" id="navbarCollapse">
                    <ul className="navbar-nav me-auto mb-2 mb-md-0">
                        <li className="nav-item">
                            <Link to="/" className="nav-link">All Banks</Link>
                        </li>
                        {isAuthed && <li className="nav-item">
                            <Link to="/my-banks" className="nav-link">My Banks</Link>
                        </li>}
                    </ul>
                </div>
                <div>
                    { !isAuthed &&
                    <ul className="navbar-nav me-auto mb-2 mb-md-0">
                        <li className="nav-item">
                            <Link to="/sign-in" className="nav-link">Sign In</Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/sign-up" className="nav-link">Sign Up</Link>
                        </li>
                    </ul> }
                    { isAuthed &&
                    <ul className="navbar-nav me-auto mb-2 mb-md-0">
                        <li className="nav-item div-b">
                            <div onClick={ handleClick } className="nav-link">Logout</div>
                        </li>
                    </ul> }
                </div>
            </div>
        </nav>
    )
}

export default Nav;