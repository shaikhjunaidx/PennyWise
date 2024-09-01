import React, { useState,useEffect } from "react";
import Navbar from "../components/Navbar";
import './Dashboard.css';
import './AddTransactionForm.css';
import AddTransactionForm from "../components/AddTransactionForm";

const transactions = [
  {
    id: 1,
    user_id: 101,
    category_id: 5,
    amount: 150.75,
    description: "Health",
    transaction_date: "2024-08-29 14:30:00",
  },
  {
    id: 2,
    user_id: 101,
    category_id: 3,
    amount: 20.0,
    description: "Food",
    transaction_date: "2024-08-30 09:15:00",
  }
];

const Dashboard = () => {
  const [showAll, setShowAll] = useState(false);
  const [transactions, setTransactions] = useState([]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showAddTransForm, setShowAddTransForm] = useState(false);

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          throw new Error("No token found");
        }

        const response = await fetch("/api/transactions", {
          method: "GET",
          headers: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json",
            "Category":0
          }
        });

        if (!response.ok) {
          throw new Error("Failed to fetch transactions");
        }

        const data = await response.json();
        setTransactions(data.transactions);
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

  const closehandleAddTransactionClick = () => {
    setShowAddTransForm(false); 
  };



  const displayedTransactions = showAll ? transactions : transactions.slice(0, 6);

  return (
    <>
      <Navbar />
      <div className="DashboardCont">
      {showAddTransForm && <AddTransactionForm closeForm={closehandleAddTransactionClick} />}
      
        <section id="overall">
          <progress className="overallProgress" value={80} max={100}></progress>
        </section>

        <section id="Transactions" className="Transactions">
          <h1 className="sectionHeader">Transactions</h1>
          <div className={`transactionsTableContainer ${showAll ? 'scrollable' : ''}`}>
            <table>
              <thead>
                <tr>
                  <th>Transaction ID</th>
                  <th>User ID</th>
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
                    <td>{transaction.user_id}</td>
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
