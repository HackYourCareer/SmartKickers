import { useEffect, useState } from 'react'
import './App.css'
import {Button} from "./components/button"

import * as WebSocket from "websocket"

function App() {

  let [blueScore, setBlueScore] = useState(0);
  let [whiteScore, setWhiteScore] = useState(0);

  useEffect(() => {
    const socket = new WebSocket.w3cwebsocket('ws://localhost:3006/csc');

    socket.onopen = function () {
      console.log("connected");
      //send to server
      socket.send("Hello from client")
      socket.onmessage = (msg) => {
        console.log(msg);
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