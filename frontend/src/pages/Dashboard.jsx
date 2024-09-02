import React, { useState,useEffect } from "react";
import NavbarLoggedIn from "../components/NavbarLoggedIn";
import './Dashboard.css';
import './AddTransactionForm.css';
import AddTransactionForm from "../components/AddTransactionForm";
import BudgetSummary from "../components/BudgetSummary";
import './BudgetSummary.css';
import AddBudgetForm from "../components/AddCategory";
import './AddCategory.css';

const Dashboard = () => {
  const [showAll, setShowAll] = useState(false);
  const [transactions, setTransactions] = useState([]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showAddTransForm, setShowAddTransForm] = useState(false);
  const [showAddBudgetForm, setShowAddBudgetForm] = useState(false);

    const [budget, setBudget] = useState({
        total: 250, 
        spent: 200 
      });

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          throw new Error("No token found");
        }

        const response = await fetch("http://localhost:8080/api/transactions", {
          method: "GET",
          headers: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });

        if (!response.ok) {
          throw new Error("Failed to fetch transactions");
        }

        const data = await response.json();
        setTransactions(data);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    fetchTransactions();
  }, []);

  const handleShowMore = () => {
    setShowAll(true);
  };

  const handleAddTransactionClick = () => {
    setShowAddTransForm(true); 
  };

   const closehandleAddTransactionClick = (newTransaction) => {
    setShowAddTransForm(false); 

    if (newTransaction) {
      fetchTransactions();
    }
  };

  const handleAddBudgetClick = () => {
    setShowAddBudgetForm(true);
  };

  const closeHandleAddBudgetClick = () => {
    setShowAddBudgetForm(false);
  };

  const displayedTransactions = showAll
  ? (transactions || []).slice(0, 6)
  : (transactions || []).slice(0, 6);

  return (
    <>
      <NavbarLoggedIn />
      <div className="DashboardCont">
      {showAddTransForm && <AddTransactionForm onAddTransaction={closehandleAddTransactionClick} />}
      {showAddBudgetForm && (<AddBudgetForm onAddBudget={closeHandleAddBudgetClick}  />)}
      
        <section id="overall" className="overalls">
            
            <BudgetSummary budget={budget} heading="Months Budget" color="hsl(355, 57%, 57%)" />
            <h1 className="sectionHeader">Categories</h1>
        <div className="categoriesCont">
            <BudgetSummary budget={budget} heading="Health" color="hsl(355, 57%, 57%)" />
            <BudgetSummary budget={budget} heading="Finance" color="hsl(355, 57%, 57%)" />
            <BudgetSummary budget={budget} heading="Random BS" color="hsl(355, 57%, 57%)" />
            <BudgetSummary budget={budget} heading="Months Budget" color="hsl(355, 57%, 57%)" />
        </div>
        <button className="AddCategoryButton" onClick={handleAddBudgetClick}>Add New Budget</button>
        </section>

        <section id="Transactions" className="Transactions">
          <h1 className="sectionHeader">Transactions</h1>
          <div className={`transactionsTableContainer ${showAll ? 'scrollable' : ''}`}>
            <table>
              <thead>
                <tr>
                  <th>Transaction ID</th>
                  <th>Category</th>
                  <th>Amount</th>
                  <th>Description</th>
                  <th>Transaction Date</th>
                </tr>
              </thead>
              <tbody>
              {(displayedTransactions || []).map((transaction)  => (
                  <tr key={transaction.id}>
                    <td>{transaction.id}</td>
                    <td>{transaction.category_name}</td>
                    <td>{transaction.amount.toFixed(2)}</td>
                    <td>{transaction.description || "N/A"}</td>
                    <td>{new Date(transaction.transaction_date).toLocaleString()}</td>
                    <td><button className="deleteTransactionButton">🗑️</button></td>
                  </tr>
                ))}
              </tbody>
            </table>
            <div className="buttonsContainer ">
                {!showAll && (<button className="showMoreButton" onClick={handleShowMore}>Show More</button>
                    )}
                <button className="AddTransaction" onClick={handleAddTransactionClick}>Add Transaction</button>
          </div>
          </div>
         
        </section>
      </div>
    </>
  );
};

export default Dashboard;
