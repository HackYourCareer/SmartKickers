import React from 'react';
import { Button } from '../../Button/Button';
import GameResults from '../GameResults/GameResults.js';
import { useNavigate } from 'react-router-dom';
import useCurrentGameplay from '../../../hooks/useCurrentGameplay';

function CurrentGameplay() {
  const { handleResetGame, handleEndGame } = useCurrentGameplay();
  const navigate = useNavigate();

  return (
    <div>
      <GameResults />
      <center className="game-ending-buttons">
        <br />
        <Button onClick={() => handleResetGame()}>Reset game</Button>
        <br />

        <Button
          onClick={() => {
            handleEndGame();
            navigate('/stats');
          }}
        >
          End game
        </Button>
      </center>
    </div>
  );
}

export default CurrentGameplay;
