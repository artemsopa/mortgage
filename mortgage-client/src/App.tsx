import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.css';
import BlankLayout from './layouts/BlankLayout';
import Layout from './layouts/Layout';

const App: React.FC = () => {
  return (
    <BrowserRouter>
        <Routes>
          <Route path='/' element={<Layout />} />
          <Route path='/auth' element={<BlankLayout />} />
        </Routes>
      </BrowserRouter>
  );
}

export default App;
