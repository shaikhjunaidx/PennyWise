import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const LoginForm = ({ closeForm }) => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate=useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch("http://localhost:8080/api/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ username, password }),
            });

            if (response.ok) {
                const data = await response.json();
                console.log("Login successful:", data);
                localStorage.setItem("token", data.token);

                closeForm();
                navigate("/dashboard");
            } else {
                console.log("Login failed");
                alert("Invalid username or password");
            }
        } catch (error) {
            console.error("Error during login:", error);
            alert("An error occurred during login. Please try again.");
        }
    };
    
    console.log("Here")
    return (
        <div className="loginFormBackground" onClick={closeForm}>
            <div className="loginFormContainer" onClick={(e) => e.stopPropagation()}>
                <h2 className="formTitle">Login</h2>
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
