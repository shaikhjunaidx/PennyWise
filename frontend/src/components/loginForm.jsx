import React, { useState } from "react";

const LoginForm = ({ closeForm }) => {


    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log("Username:", username);
        console.log("Password:", password);
        closeForm();
    };
    console.log("Here")
    return (
        <div className="loginFormBackground" onClick={closeForm}>
            <div className="loginFormContainer" onClick={(e) => e.stopPropagation()}>
                <h2>Login</h2>
                <form onSubmit={handleSubmit}>
                    <div className="inputGroup">
                        <label htmlFor="username">Username:</label>
                        <input
                            type="text"
                            id="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                        />
                    </div>
                    <div className="inputGroup">
                        <label htmlFor="password">Password:</label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </div>
                    <button type="submit" className="loginFormSubmitButton">Login</button>
                </form>
            </div>
        </div>
    );
};

export default LoginForm;
