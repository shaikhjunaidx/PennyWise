import React, { useState, useRef, useEffect } from "react";

const AddBudgetForm = ({ onAddBudget}) => {
  const [formData, setFormData] = useState({
    category_name: "",
    description: "",
    amount: "",
    budget_year: "",
    budget_month: "",
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const formRef = useRef(null);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    const token = localStorage.getItem("token");
    console.log(formData);

    try {
      const categoryResponse = await fetch("http://localhost:8080/api/categories", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: formData.category_name,
          description: formData.description,
        }),
        
      });

      if (!categoryResponse.ok) {
        throw new Error("Failed to create category");
      }

      const categoryData = await categoryResponse.json();
      const category_id = categoryData.id;

      const budgetResponse = await fetch("http://localhost:8080/api/budgets", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          amount_limit: parseFloat(formData.amount),
          budget_year: parseInt(formData.budget_year),
          budget_month: parseInt(formData.budget_month),
          category_id: category_id,
        }),
      });

      if (!budgetResponse.ok) {
        throw new Error("Failed to create budget");
      }

      const budgetData = await budgetResponse.json();

      if (onAddBudget) {
        onAddBudget(budgetData);
      }

      setFormData({
        category_name: "",
        description: "",
        amount: "",
        budget_year: "",
        budget_month: "",
      });

      if (onClose) {
        onClose(); 
      }
    } catch (error) {
      setError(error.message);
    } finally {
      setLoading(false);
    }
  };

  const handleBackgroundClick = (e) => {
    if (e.target === e.currentTarget) {
        onAddBudget();
    }
  };

  return (
    <div className="addBudgetFormContainer" onClick={handleBackgroundClick}>
        <div className="addBudgetForm">
      <h2 className="budgetFormTitle">Add Budget</h2>
      <form onSubmit={handleSubmit}>
        <div className="budgetFormGroup">
          <label htmlFor="category_name">Category Name:</label>
          <input
            type="text"
            id="category_name"
            name="category_name"
            value={formData.category_name}
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
          <label htmlFor="budget_year">Budget Year:</label>
          <input
            type="number"
            id="budget_year"
            name="budget_year"
            value={formData.budget_year}
            onChange={handleChange}
            required
          />
        </div>
        <div className="formGroup">
          <label htmlFor="budget_month">Budget Month:</label>
          <input
            type="number"
            id="budget_month"
            name="budget_month"
            value={formData.budget_month}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="submitButton" disabled={loading}>
          {loading ? "Submitting..." : "Add Budget"}
        </button>
        {error && <p className="error">{error}</p>}
      </form>
      </div>
    </div>
  );
};

export default AddBudgetForm;
