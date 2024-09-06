import React from 'react';


const BudgetSummary = ({ budget, heading, color,onClick }) => {
  const progressPercentage = (budget.spent / budget.total) * 100;

  return (
    <div className="budget" 
    style={{
      "--accent": "white", 
      "--border-color": color, 
      "--progress-color": color,
    }}
    onClick={onClick}
    > 
      <div className="progress-text-box">
        <h3>{heading}</h3>
        <p>${budget.total.toFixed(2)} Budgeted</p>
      </div>
      <progress className="overallProgress" value={budget.spent} max={budget.total}>
        {progressPercentage.toFixed(0)}%
      </progress>
      <div className="progress-text-box">
        <small>${budget.spent.toFixed(2)} spent</small>
        <small>${(budget.total - budget.spent).toFixed(2)} remaining</small>
      </div>
    </div>
  );
};

export default BudgetSummary;
