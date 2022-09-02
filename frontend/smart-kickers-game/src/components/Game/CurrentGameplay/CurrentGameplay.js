import React from 'react';
import { Button } from '../../Button/Button';
import GameResults from '../GameResults/GameResults.js';

function CurrentGameplay({ blueScore, whiteScore, handleStartGame, handleResetGame, handleEndGame, isVisible }) {
  return (
    <div>
      <GameResults blueScore={blueScore} whiteScore={whiteScore} isVisible={isVisible} />
      <center className="game-ending-buttons">
        <Button id="start-game" hidden={isVisible} onClick={() => handleStartGame()}>
          Start game
        </Button>
        <br />
        <Button hidden={!isVisible} onClick={() => handleResetGame()}>
          Reset game
        </Button>
        <br />
        <Button
          hidden={!isVisible}
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
