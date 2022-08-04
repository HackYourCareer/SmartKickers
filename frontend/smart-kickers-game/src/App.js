import { useEffect, useState } from 'react';
import './App.css';
import * as GameAPI from './apis/Game';
import { Button } from './components/Button';
import config from './config';

function App() {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);

  useEffect(() => {
    const socket = new WebSocket(`${config.wsBaseUrl}/csc`);

    socket.onopen = function () {
      //send to server
      socket.send('Hello from client');
      socket.onmessage = (msg) => {
        msg = JSON.parse(msg.data);
        setBlueScore(msg.blueScore);
        setWhiteScore(msg.whiteScore);
      };
    };
  }, []);

  return (
    <>
      <h1>Smart Kickers</h1>
      <div className="game-result-container">
        <p className="game-result-item">Blue: {blueScore}</p>
        <p className="game-result-item">White: {whiteScore}</p>
      </div>
      <center>
        <Button onClick={() => GameAPI.resetGame()}>Reset game</Button>
      </center>
    </>
  );
}

export default App;
