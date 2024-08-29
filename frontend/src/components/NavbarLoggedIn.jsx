// src/components/Navbar.jsx

import React from 'react';
import logo from './finance.png';
import './Navbar.css'


function Navbar() {
  return (
    <nav className="navbar">
      <div className="navcontainer">
        <a className="navbar-brand" href="#">
          <img src={logo} alt="Logo" className="navlogo" />
          PennyWise
        </a>
      </div>
    </nav>
  );
}

export default Navbar;