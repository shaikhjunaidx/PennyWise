import React, { useState, useEffect } from "react";
import { jwtDecode } from 'jwt-decode';
import { fetchCategories } from "../utils/fetchCategories,jsx";

const AddTransactionForm = ({ onAddTransaction }) => {
  const [formData, setFormData] = useState({
    category_id: 0,
    description: '',
    amount: '',
    transaction_date: ''
  });

  const [categories, setCategories] = useState([]);
  const [userName, setUserName] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      const decodedToken = jwtDecode(token);
      setUserName(decodedToken.sub); 
    }
  }, []);

  useEffect(() => {
    const getCategories = async () => {
      try {
        const data = await fetchCategories();
        setCategories(data);
      } catch (error) {
        console.error("Error fetching categories:", error.message);
      }
    };

    getCategories();
  }, []);


  const handleChange = (e) => {
    const { name, value } = e.target;
    const updatedValue =
      name === "category_id" ? parseInt(value, 10) : name === "amount" ? parseFloat(value) : value;

    setFormData({
      ...formData,
      [name]: updatedValue,
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
        amount: parseFloat(formData.amount),
        transaction_date: new Date(formData.transaction_date).toISOString(),
        username: userName,
    };
    
    console.log(dataToSubmit)

    try {
      const response = await fetch("http://localhost:8080/api/transactions", {
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
        category_id: '',
        description: '',
        amount: '',
        transaction_date: ''
      });
    } catch (error) {
      console.error(error.message);
    }
  };

  const handleBackgroundClick = (e) => {
    if (e.target === e.currentTarget) {
      onAddTransaction();
    }
  };

  return (
    <div className="addTransFormBackground" onClick={handleBackgroundClick}>
    <div className="addTransactionForm">
      <h2 className="formTitle">Add Transaction</h2>
      <form onSubmit={handleSubmit} className="AddTransFormDiv">
        <div className="formGroup">
        <label htmlFor="category_id">Category:</label>
            <select
              id="category_id"
              name="category_id"
              value={formData.category_id}
              onChange={handleChange}
              required
            >
              <option value="">Select a category</option>
              {categories.map((category) => (
                <option key={category.id} value={category.id}>
                  {category.name}
                </option>
              ))}
            </select>
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
