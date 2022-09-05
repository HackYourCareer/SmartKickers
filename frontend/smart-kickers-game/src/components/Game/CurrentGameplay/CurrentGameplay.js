import React from 'react';
import { Button } from '../../Button/Button';
import GameResults from '../GameResults/GameResults.js';

function CurrentGameplay({ blueScore, whiteScore, handleStartGame, handleResetGame, handleEndGame, isVisible }) {
  return (
    <div>
      <GameResults blueScore={blueScore} whiteScore={whiteScore} isVisible={isVisible} />
      <center className="game-ending-buttons">
        {!isVisible && (
          <Button id="start-game" onClick={() => handleStartGame()}>
            Start game
          </Button>
        )}
        <br />
        {isVisible && <Button onClick={() => handleResetGame()}>Reset game</Button>}
        <br />
        {isVisible && (
          <Button
            onClick={() => {
              handleEndGame();
            }}
          >
            End game
          </Button>
        )}
      </center>
    </div>
  );
}

export default CurrentGameplay;
