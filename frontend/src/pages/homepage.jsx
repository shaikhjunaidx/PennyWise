
import React, { useState } from "react";
import Navbar from "../components/Navbar";
import art from "../assets/financeBackground.jpg";
import "./homepage.css"
import "./loginForm.css"
import LoginForm from "../components/LoginForm";
import SignUpForm from "../components/signUpForm";

const homepage=()=>{
    const [isLoginForm,setIsLoginForm]=useState(false);
    const [isSignupForm,setIsSignupForm]=useState(false);

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
    
    const openSignupForm=()=>
            { 
                setIsSignupForm(true);
            };
        
    const closeSignupForm=()=>
            {   
                setIsSignupForm(false);
            };
        

    return(
        <>
        <Navbar/>
        {isLoginForm && <LoginForm closeForm={closeLoginForm} />}
        {isSignupForm && <SignUpForm closeForm={closeSignupForm} />}
        <div className="logincontainer">
            <h1 className="motto">Spend Your Pennies Wisely </h1>
            <button className="loginButton" onClick={openLoginForm}>
                Login
            </button>
            <button className="signUpButton" onClick={openSignupForm}>
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