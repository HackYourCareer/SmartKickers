import { useEffect, useState } from "react";
import "./App.css";
import { Button } from "./components/button";

function App() {
  let [blueScore, setBlueScore] = useState(0);
  let [whiteScore, setWhiteScore] = useState(0);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:3006/csc");
    //const socket = new WebSocket("ws://localhost:3000")
    socket.onopen = function () {
      console.log("connected");
      //send to server
      socket.send("Hello from client");
      socket.onmessage = (msg) => {
        console.log(msg);
        msg = JSON.parse(msg.data);
        console.log(msg);
        setBlueScore(msg.blueScore);
        setWhiteScore(msg.whiteScore);

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
      <center>
        <Button>Reset game</Button>
      </center>
    </>
  );
}

export default App;
