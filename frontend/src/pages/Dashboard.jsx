import React, { useState,useEffect } from "react";
import NavbarLoggedIn from "../components/NavbarLoggedIn";
import './Dashboard.css';
import './AddTransactionForm.css';
import AddTransactionForm from "../components/AddTransactionForm";
import BudgetSummary from "../components/BudgetSummary";
import './BudgetSummary.css';

const Dashboard = () => {
  const [showAll, setShowAll] = useState(false);
  const [transactions, setTransactions] = useState([]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showAddTransForm, setShowAddTransForm] = useState(false);

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


  const displayedTransactions = showAll ? transactions : transactions.slice(0, 6);
  const progressPercentage = (budget.spent / budget.total) * 100;

  return (
    <>
      <NavbarLoggedIn />
      <div className="DashboardCont">
      {showAddTransForm && <AddTransactionForm onAddTransaction={closehandleAddTransactionClick} />}
      
        <section id="overall" className="overalls">
            <BudgetSummary budget={budget} heading="Months Budget" color="hsl(355, 57%, 57%)" />
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
                {displayedTransactions.map((transaction) => (
                  <tr key={transaction.id}>
                    <td>{transaction.id}</td>
                    <td>{transaction.category_id}</td>
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
