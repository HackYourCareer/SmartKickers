import { useEffect, useState } from 'react';
import './App.css';
import { resetGame } from './apis/Game';
import { Button } from './components/Button';
import GameResults from './components/GameResults.js';
import GameStatistics from './components/GameStatistics.js';
import { library } from '@fortawesome/fontawesome-svg-core';
import { faPerson } from '@fortawesome/free-solid-svg-icons';

import config from './config';

library.add(faPerson);

function App() {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);

  useEffect(() => {
    const socket = new WebSocket(`${config.wsBaseUrl}/score`);

    socket.onopen = function () {
      // Send to server
      socket.send('Hello from client');
      socket.onmessage = (msg) => {
        msg = JSON.parse(msg.data);
        setBlueScore(msg.blueScore);
        setWhiteScore(msg.whiteScore);
      };
    };
  }, []);

  function handleResetGame() {
    resetGame().then((data) => {
      if (data.error) alert(data.error);
    });
  }

  function handleEndGame() {
    console.log('dupsko');
  }

  return (
    <>
      <h1>Smart Kickers</h1>
      <GameResults blueScore={blueScore} whiteScore={whiteScore} />
      <center className="game-ending-buttons">
        <Button onClick={() => handleResetGame()}>Reset game</Button>
        <br />
        <Button onClick={() => handleEndGame()}>End game</Button>
      </center>
      <GameStatistics blueScore={0} whiteScore={0} />
    </>
  );
}

export default App;
