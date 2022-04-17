import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.css';
import AllBanks from './components/AllBanks/AllBanks';
import Nav from './components/Nav/Nav';
import SignIn from './components/SignIn/SignIn';
import SignUp from './components/SignUp/SignUp';
import { useTypedSelector } from './hooks/useTypedSelectors';

const App: React.FC = () => {
  const { isAuthed, payload } = useTypedSelector(state => state.login)

  return (
    <div className='main'>
      <BrowserRouter>
        <Nav />
        <Routes>
          {isAuthed && <Route path='/' element={<AllBanks />} /> }
          {!isAuthed && <Route>
            <Route path='/sign-in' element={<SignIn />} />
            <Route path='/sign-up' element={<SignUp />} />
          </Route> }
          <Route path="*" element={<h2>Not found</h2>} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
