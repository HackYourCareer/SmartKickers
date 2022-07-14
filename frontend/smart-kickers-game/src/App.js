import { useEffect } from 'react'
import './App.css'
import {Button} from "./components/button"

import * as WebSocket from "websocket"

function App() {

  var blueScore = 0
  var whiteScore = 0

  useEffect(() => {
    const socket = new WebSocket.w3cwebsocket('ws://localhost:3000');

    socket.onopen = function () {
      console.log("connected");
      socket.onmessage = (msg) => {
        msg = JSON.parse(msg.toString())
        if (msg.type === "blueGoal") {
          blueScore++
        }
        if (msg.type === "whiteGoal") {
          whiteScore++
        }
     
        console.log(msg);
      };
    };
  });

  return (
    <>
      <h1>Smart Kickers</h1>
      <div className="game-result-container">
        Blue: {blueScore}
          {"   "}
        White: {whiteScore}
      </div>
      <center><Button >Reset game</Button></center>
    </>
  )
}

export default App