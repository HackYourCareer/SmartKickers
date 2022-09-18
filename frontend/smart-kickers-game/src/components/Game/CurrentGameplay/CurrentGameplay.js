import React from 'react';
import { Button } from '../../Button/Button';
import GameResults from '../GameResults/GameResults.js';
import { useNavigate } from 'react-router-dom';
import useCurrentGameplay from '../../../hooks/useCurrentGameplay';

function CurrentGameplay() {
  const { blueScore, whiteScore, isVisible, handleStartGame, handleResetGame, handleEndGame } = useCurrentGameplay();
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
