import React from 'react';
import { Button } from '../../Button/Button';
import GameResults from '../GameResults/GameResults.js';
import { useNavigate } from 'react-router-dom';

function CurrentGameplay({ blueScore, whiteScore, handleStartGame, handleResetGame, handleEndGame, isVisible }) {
  const navigate = useNavigate();
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
              navigate('/stats');
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
