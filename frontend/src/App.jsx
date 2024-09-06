// src/App.jsx
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from './pages/HomePage';
import Dashboard from './pages/Dashboard';
import CategorySummary from './pages/CategorySummary';


const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/Dashboard" element={<Dashboard />} />
        <Route path="/category-summary/:categoryId" element={<CategorySummary />} />
      </Routes>
    </Router>
  );
};

export default App;
