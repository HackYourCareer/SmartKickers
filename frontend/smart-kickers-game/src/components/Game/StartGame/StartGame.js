import React from 'react';
import useCurrentGameplay from '../../../hooks/useCurrentGameplay';
import { Button } from '../../Button/Button';

function StartGame() {
  const { handleStartGame } = useCurrentGameplay();
  return (
    <center>
      <Button id="start-game" onClick={() => handleStartGame()}>
        Start game
      </Button>
    </center>
  );
}

export default StartGame;
