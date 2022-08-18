import React from 'react';
import { Button } from '../../Button/Button';
import GameResults from '../GameResults/GameResults.js';

function CurrentGameplay({ blueScore, whiteScore, handleResetGame, handleEndGame }) {
  return (
    <div>
      <GameResults blueScore={blueScore} whiteScore={whiteScore} />
      <center className="game-ending-buttons">
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
