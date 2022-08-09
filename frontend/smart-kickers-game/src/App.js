import { useEffect, useState } from 'react';
import './App.css';
import { resetGame } from './apis/Game';
import { Button } from './components/Button';
import GameResults from './components/GameResults.js';

import config from './config';

function App() {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);

  const socket = new WebSocket("ws://localhost:3006/csc");

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

  return (
    <>
      <h1>Smart Kickers</h1>
      <GameResults blueScore={blueScore} whiteScore={whiteScore} />
      <center>
        <Button onClick={() => handleResetGame()}>Reset game</Button>
      </center>
    </>
  );
}

export default App;
