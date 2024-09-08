import React, { useState,useEffect } from "react";
import { useParams } from 'react-router-dom';
import NavbarLoggedIn from '../components/NavbarLoggedIn';
import './CategorySummary.css';
import { Bar } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';
import './BudgetHistoryChart.css';
ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

const CategorySummary = () => {
  const { categoryId } = useParams();
  const [transactions, setTransactions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [historyData, setHistoryData] = useState([]);


    useEffect(() => {
      const fetchTransactions = async () => {
        try {
          const token = localStorage.getItem("token");
          if (!token) {
            throw new Error("No token found");
          }
  
          const response = await fetch(`http://localhost:8080/api/transactions/category/${categoryId}`, {
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
      const token = localStorage.getItem('token');
  
      fetch(`http://localhost:8080/api/budgets/category/${categoryId}/history`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      })
        .then((response) => response.json())
        .then((data) => setHistoryData(data.history)) // Extracting history array from the response
        .catch((error) => console.error('Error fetching history data:', error));
    }, [categoryId]);
  
    const chartData = {
      labels: historyData.map((entry) => `${entry.month} ${entry.year}`), // Combine month and year for labels
      datasets: [
        {
          label: 'Spent Amount',
          data: historyData.map((entry) => entry.spent_amount), // Using the spent_amount field
          backgroundColor: 'rgba(54, 162, 235, 0.6)',
          borderColor: 'rgba(54, 162, 235, 1)',
          borderWidth: 2,
          borderRadius: 4,
          hoverBackgroundColor: 'rgba(75, 192, 192, 0.8)',
        },
      ],
    };
  
    const options = {
      responsive: true,
      plugins: {
        legend: {
          display: true,
          position: 'top',
          labels: {
            color: '#333',
            font: {
              size: 14,
              family: 'Arial, sans-serif',
              style: 'italic',
            },
          },
        },
        title: {
          display: true,
          text: 'Monthly Spending History',
          font: {
            size: 20,
            family: 'Arial, sans-serif',
            weight: 'bold',
          },
          color: '#333',
        },
        tooltip: {
          enabled: true,
          backgroundColor: 'rgba(0,0,0,0.8)',
          bodyColor: '#fff',
          bodyFont: {
            size: 12,
            family: 'Arial, sans-serif',
          },
          callbacks: {
            label: function (tooltipItem) {
              return `Spent: $${tooltipItem.raw}`;
            },
          },
        },
      },
      scales: {
        x: {
          grid: {
            display: false,
          },
          ticks: {
            color: '#333',
            font: {
              size: 12,
            },
          },
        },
        y: {
          grid: {
            borderDash: [5, 5],
            color: '#ccc',
          },
          beginAtZero: true,
          ticks: {
            color: '#333',
            font: {
              size: 12,
            },
          },
        },
      },
    };

  



  return (
    <div>
        <NavbarLoggedIn></NavbarLoggedIn>
        <div className='CategoryDetailsSummaryCont'>
        
        <section id="Transactions" className="Transactions">
          <h1 className="sectionHeader">Transactions</h1>
          <div className={"categoryTransactionTC"}>
            <table>
              <thead>
                <tr>
                  <th>Transaction ID</th>
                  <th>Amount</th>
                  <th>Description</th>
                  <th>Transaction Date</th>
                </tr>
              </thead>
              <tbody>
              {(transactions || []).map((transaction)  => (
                  <tr key={transaction.id}>
                    <td>{transaction.id}</td>
                    <td>{transaction.amount.toFixed(2)}</td>
                    <td>{transaction.description || "N/A"}</td>
                    <td>{new Date(transaction.transaction_date).toLocaleString()}</td>
                    <td><button className="deleteTransactionButton">🗑️</button></td>
                  </tr>
                ))}
              </tbody>
            </table>
           </div>
           </section>
           {historyData.length > 0 && (
        <div className="history-chart">
          <Bar data={chartData} options={options} />
        </div>
      )}
        </div>
    </div>
  );
};

export default CategorySummary;