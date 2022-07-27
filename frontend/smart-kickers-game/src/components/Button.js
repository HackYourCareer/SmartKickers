import React from 'react';
import './Button.css';

export const Button = ({ children, type, onClick, className = 'btn--primary btn--medium' }) => {
  return (
    <button className={`btn ${className}`} onClick={onClick} type={type}>
      {children}
    </button>
  );
};
