import { useEffect, useState } from 'react'
import './App.css'
import {Button} from "./components/button"

function App() {

  let [blueScore, setBlueScore] = useState(0);
  let [whiteScore, setWhiteScore] = useState(0);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:3000")

    socket.onopen = function () {
      console.log("connected");
      socket.onmessage = (msg) => {
        msg = JSON.parse(msg.data)
        console.log(msg)
        if (msg.type === "blueGoal") {
          setBlueScore(blueScore + 1)
        }
        if (msg.type === "whiteGoal") {
          setWhiteScore(whiteScore + 1)
        }
     
        console.log(msg);
      };
    };
  });

  return (
    <>
      <h1>Smart Kickers</h1>
      <div className="game-result-container" data-testid="blue-team-score">
        Blue: {blueScore}
          {"   "}
        White: {whiteScore}
      </div>
      <center><Button >Reset game</Button></center>
    </>
  )
}

export default App