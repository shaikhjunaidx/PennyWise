
import React, { useState } from "react";
import Navbar from "../components/Navbar";
import art from "../assets/financeBackground.jpg";
import "./homepage.css"
import LoginForm from "../components/LoginForm";

const homepage=()=>{
    const [isLoginForm,setIsLoginForm]=useState(false);

    const openLoginForm=()=>
        { 
            console.log("OpenloginForm Clicked")
            setIsLoginForm(true);
        };
    
    const closeLoginForm=()=>
        {   
            console.log("CloseForm Clicked")
            setIsLoginForm(false);
        };


    return(
        <>
        <Navbar/>
        {isLoginForm && <LoginForm closeForm={closeLoginForm} />}
        <div className="logincontainer">
            <h1 className="motto">Spend Your Pennies Wisely </h1>
            <button className="loginButton" onClick={openLoginForm}>
                login
            </button>
            <button className="signUpButton">
                Sign-up
            </button>
            <img src={art} className="homeart"/>
            <div className="align"></div>
            <footer className="footer">
                <div className="waves">
                <div className="wave" id="wave1"></div>
                <div className="wave" id="wave2"></div>
                <div className="wave" id="wave3"></div>
                <div className="wave" id="wave4"></div>
                </div>
            </footer>
        </div>
        </>
    );
}

export default homepage;