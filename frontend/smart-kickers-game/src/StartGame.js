import React from 'react';
import './StartGame.css';
import { Button } from './components/Button/Button';
import { useNavigate } from 'react-router-dom';
import { useGameDataContext } from './contexts/GameDataContext';

export default function StartGame() {
  const navigate = useNavigate();
  const { start, handleResetGame, resetGoalsArray, setIsGameStarted } = useGameDataContext();
  const handleStartGame = () => {
    setIsGameStarted(true);
    handleResetGame();
    resetGoalsArray();
    start();
    navigate('/game');
  };
  return (
    <>
      <h1>Smart Kickers</h1>
      <center>
        <Button id="start-game" onClick={() => handleStartGame()}>
          Start game
        </Button>
      </center>
    </>
  );
}
