import React from 'react';
import { Button } from './Button.js';
import { useNavigate } from 'react-router-dom';

function NewGameButton() {
  const navigate = useNavigate();
  return (
    <Button
      className="btn--primary new-game-btn"
      onClick={() => {
        navigate('/');
      }}
    >
      New game
    </Button>
  );
}

export default NewGameButton;
