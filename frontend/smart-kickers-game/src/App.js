import { useEffect, useState } from 'react';
import './App.css';
import { Button } from './components/Button';
import GameResults from './components/GameResults.js';

function App() {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:3000/score');

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

  return (
    <>
      <h1>Smart Kickers</h1>
      <GameResults blueScore={blueScore} whiteScore={whiteScore} />
      <center>
        <Button>Reset game</Button>
      </center>
    </>
  );
}

export default App;
