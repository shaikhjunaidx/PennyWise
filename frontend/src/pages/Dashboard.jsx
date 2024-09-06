import React, { useState,useEffect } from "react";
import { useNavigate } from 'react-router-dom';
import NavbarLoggedIn from "../components/NavbarLoggedIn";
import './Dashboard.css';
import './AddTransactionForm.css';
import AddTransactionForm from "../components/AddTransactionForm";
import BudgetSummary from "../components/BudgetSummary";
import './BudgetSummary.css';
import AddBudgetForm from "../components/AddCategory";
import './AddCategory.css';
import { fetchCategories } from "../utils/fetchCategories,jsx";
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  LineElement,
  PointElement,
  CategoryScale,
  LinearScale,
  Title,
  Tooltip,
  Legend
} from 'chart.js';
import './TransacChart.css';


ChartJS.register(LineElement, PointElement, CategoryScale, LinearScale, Title, Tooltip, Legend);


const Dashboard = () => {
  const [showAll, setShowAll] = useState(false);
  const [transactions, setTransactions] = useState([]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showAddTransForm, setShowAddTransForm] = useState(false);
  const [showAddBudgetForm, setShowAddBudgetForm] = useState(false);
  const [categoryBudgets, setCategoryBudgets] = useState([]);
  const [overallBudget, setOverallBudget] = useState(null);
  const [categories, setCategories] = useState([]);
  const [categoryMap, setCategoryMap] = useState({});
  const [weeklyData, setWeeklyData] = useState([]);
  const navigate = useNavigate();
  const handleCategoryClick = (categoryId) => {
    console.log("Category clicked")
    navigate(`/category-summary/${categoryId}`);
  };

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

  useEffect(() => {
    const fetchCategoryBudgets = async () => {
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          throw new Error("No token found");
        }

        const response = await fetch("http://localhost:8080/api/budgets", {
          method: "GET",
          headers: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });

        if (!response.ok) {
          throw new Error("Failed to fetch category budgets");
        }

        const data = await response.json();
        setCategoryBudgets(data);
      } catch (error) {
        setError(error.message);
      }
    };

    fetchCategoryBudgets();
  }, []);

  useEffect(() => {
    const fetchOverallBudget = async () => {
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          throw new Error("No token found");
        }

        const response = await fetch("http://localhost:8080/api/budgets/overall", {
          method: "GET",
          headers: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
          }
        });

        if (!response.ok) {
          throw new Error("Failed to fetch overall budget");
        }

        const data = await response.json();
        setOverallBudget(data);
      } catch (error) {
        setError(error.message);
      }
    };

    fetchOverallBudget();
  }, []);

  useEffect(() => {
    const token = localStorage.getItem('token');

    fetch('http://localhost:8080/api/transactions/weekly', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`, 
      },
    })
      .then((response) => response.json())
      .then((data) => setWeeklyData(data))
      .catch((error) => console.error('Error fetching weekly data:', error));
  }, []);

  const chartData = {
    labels: weeklyData.map((entry) => `Week ${entry.week}`),
    datasets: [
      {
        label: 'Total Spent',
        data: weeklyData.map((entry) => entry.total_spent),
        borderColor: 'rgba(75, 192, 192, 1)',
        backgroundColor: 'rgba(75, 192, 192, 0.2)',
        fill: true,
        tension: 0.1,
      },
    ],
  };


  useEffect(() => {
    const getCategories = async () => {
      try {
        const data = await fetchCategories();
        setCategories(data);
        const categoryLookup = {};
        data.forEach(category => {
          categoryLookup[category.id] = category.name; // Assume category object has 'id' and 'name'
        });
        setCategoryMap(categoryLookup); // Store the lookup object in state
      } catch (error) {
        console.error("Error fetching categories:", error.message);
      }
    };

    getCategories();
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
             <h1 className="sectionHeader">Monthly Budget Summary</h1>
          {overallBudget && (<BudgetSummary
              budget={{
                total: overallBudget.amount_limit,
                spent: overallBudget.spent_amount,
                remaining: overallBudget.remaining_amount
              }}
              heading="Overall Monthly Budget"
              color="hsl(355, 57%, 57%)"
            />
          )}
            <h2 className="subSectionHeader">Categories</h2>
        <div className="categoriesCont">
        {categoryBudgets.map((budget) => (
            <BudgetSummary
              key={budget.id}
              budget={{
                total: budget.amount_limit,
                spent: budget.spent_amount,
                remaining: budget.remaining_amount
              }}
              heading={categoryMap[budget.category_id] || "Unknown Category"}
              color="hsl(355, 57%, 57%)"
              onClick={() => handleCategoryClick(budget.category_id)}
              style={{ cursor: 'pointer' }}
            />
          ))}
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
          <h2>Weekly Spending Overview</h2>
          {weeklyData.length > 0 && (
              <div className="weekly-spending-chart">
                
                <Line data={chartData} />
              </div>
            )}
        </section>
      </div>
    </>
  );
};

export default Dashboard;