import { useEffect, useState } from 'react'
import './App.css'
import {Button} from "./components/button"

function App() {

  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:3006/csc');
    //const socket = new WebSocket("ws://localhost:3000")
    socket.onopen = function () {
      //send to server
      socket.send("Hello from client")
      socket.onmessage = (msg) => {
        msg = JSON.parse(msg.data)
        setBlueScore(msg.blueScore)
        setWhiteScore(msg.whiteScore)
      };
    };
  });

  return (
    <>
      <h1>Smart Kickers</h1>
      <div className="game-result-container">
        <p className="game-result-item">Blue: {blueScore}</p>
        <p className="game-result-item">White: {whiteScore}</p>
      </div>
      <center>
        <Button>Reset game</Button>
      </center>
    </>
  );
}

export default App