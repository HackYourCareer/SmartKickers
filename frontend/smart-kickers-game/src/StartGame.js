import React from 'react';
import './StartGame.css';
import useCurrentGameplay from './hooks/useCurrentGameplay';
import { Button } from './components/Button/Button';

export default function StartGame() {
  const { handleStartGame } = useCurrentGameplay();
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
