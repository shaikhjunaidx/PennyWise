import React, { useState, useEffect } from "react";
import { jwtDecode } from 'jwt-decode';

const AddTransactionForm = ({ onAddTransaction }) => {
    console.log('here')
  const [formData, setFormData] = useState({
    category_id: '',
    description: '',
    amount: '',
    transaction_date: ''
  });

  const [userName, setUserName] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      const decodedToken = jwtDecode(token);
      setUserName(decodedToken.sub); 
    }
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!userName) {
      console.error("Username is not available");
      return;
    }

    const dataToSubmit = {
      ...formData,
      username: userName
    };

    try {
      const response = await fetch("/api/transactions", {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json"
        },
        body: JSON.stringify(dataToSubmit)
      });

      if (!response.ok) {
        throw new Error("Failed to add transaction");
      }

      const result = await response.json();
      if (onAddTransaction) {
        onAddTransaction(result.transaction);
      }
      setFormData({
        category: '',
        description: '',
        amount: '',
        transaction_date: ''
      });
    } catch (error) {
      console.error(error.message);
    }
  };

  return (
    <div className="addTransFormBackground" onClick={onAddTransaction}>
    <div className="addTransactionForm">
      <h2 className="formTitle">Add Transaction</h2>
      <form onSubmit={handleSubmit} className="AddTransFormDiv">
        <div className="formGroup">
          <label htmlFor="category">Category:</label>
          <input
            type="text"
            id="category"
            name="category"
            value={formData.category}
            onChange={handleChange}
            required
          />
        </div>
        <div className="formGroup">
          <label htmlFor="description">Description:</label>
          <input
            type="text"
            id="description"
            name="description"
            value={formData.description}
            onChange={handleChange}
            required
          />
        </div>
        <div className="formGroup">
          <label htmlFor="amount">Amount:</label>
          <input
            type="number"
            id="amount"
            name="amount"
            value={formData.amount}
            onChange={handleChange}
            step="0.01"
            required
          />
        </div>
        <div className="formGroup">
          <label htmlFor="transaction_date">Transaction Date:</label>
          <input
            type="datetime-local"
            id="transaction_date"
            name="transaction_date"
            value={formData.transaction_date}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="submitButton">Add Transaction</button>
      </form>
    </div>
    </div>
  );
};

export default AddTransactionForm;
