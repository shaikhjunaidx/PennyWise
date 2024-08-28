
import React from "react";
import Navbar from "../components/Navbar";
import art from "../assets/financeBackground.jpg";
import "./homepage.css"

const homepage=()=>{
    return(
        <>
        <Navbar/>
        <div className="logincontainer">
            <h1 className="motto">Spend Your Pennies Wisely </h1>
            <button className="loginButton">
                login
            </button>
            <button className="signUpButton">
                Sign-up
            </button>
            <img src={art} className="homeart"/>
        </div>
        </>
    );
}

export default homepage;