import React from 'react';
import { Button } from '../../Button/Button';
import GameResults from '../GameResults/GameResults.js';

function CurrentGameplay({ blueScore, whiteScore, handleStartGame, handleResetGame, handleEndGame }) {
  return (
    <div>
      <GameResults blueScore={blueScore} whiteScore={whiteScore} />
      <center className="game-ending-buttons">
        <Button onClick={() => handleStartGame()}>Start game</Button>
        <br />
        <Button onClick={() => handleResetGame()}>Reset game</Button>
        <br />
        <Button
          onClick={() => {
            handleEndGame();
          }}
        >
          End game
        </Button>
      </center>
    </div>
  );
}

export default CurrentGameplay;
