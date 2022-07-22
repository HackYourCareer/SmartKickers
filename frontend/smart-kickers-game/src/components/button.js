import React from 'react';
import './button.css';

export const Button = ({ children, type, onClick, className = 'btn--primary btn--medium' }) => {
  return (
    <button className={`btn ${className}`} onClick={onClick} type={type}>
      {children}
    </button>
  );
};
