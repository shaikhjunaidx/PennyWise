import React from 'react';
import logo from './finance.png';
import './NavbarLoggedIn.css';

function NavbarLoggedIn() {
  const handleLogout = () => {
    localStorage.removeItem('token');
    window.location.href = 'http://localhost:5173/';
  };

  return (
    <nav className="navbar">
      <div className="navcontainer">
        <a className="navbar-brand" href="#">
          <img src={logo} alt="Logo" className="navlogo" />
          PennyWise
        </a>
        <button className="logout-button" onClick={handleLogout}>
          Logout
        </button>
      </div>
    </nav>
  );
}

export default NavbarLoggedIn;
