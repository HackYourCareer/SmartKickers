import React from 'react';
import './Button.css';

export const Button = ({ children, onClick, className = 'btn--primary btn--medium', ...props }) => {
  return (
    <button className={`btn ${className}`} onClick={onClick} {...props}>
      {children}
    </button>
  );
};
