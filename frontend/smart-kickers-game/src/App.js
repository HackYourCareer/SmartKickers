import { useEffect, useState } from "react";
import "./App.css";
import { Button } from "./components/button";

function App() {
  let [blueScore, setBlueScore] = useState(0);
  let [whiteScore, setWhiteScore] = useState(0);

  const socket = new WebSocket("ws://localhost:3006/csc");

  useEffect(() => {
    socket.onopen = function () {
      socket.onmessage = (msg) => {
        msg = JSON.parse(msg.data);
        setBlueScore(msg.blueScore);
        setWhiteScore(msg.whiteScore);
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
        <Button onClick={() => socket.send(JSON.stringify(true))}>
          Reset game
        </Button>
      </center>
    </>
  );
}

export default App;
