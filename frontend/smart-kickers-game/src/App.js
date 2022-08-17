import { useEffect, useState } from 'react';
import './App.css';
import { Button } from './components/Button/Button';
import GameResults from './components/Game/GameResults.js';
import { resetGame } from './apis/resetGame';
import GameStatistics from './components/Game/GameStatistics.js';
import { initLibs } from './appConfig';
import config from './config';

function App() {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);
  const [toggleGameScore, setToggleGameScore] = useState(false);
  const [finalScores, setFinalScores] = useState({ blue: 0, white: 0 });
  useEffect(() => {
    initLibs();
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
    setToggleGameScore(!toggleGameScore);
  }

  return (
    <>
      <h1>Smart Kickers</h1>
      {toggleGameScore === false ? (
        <>
          {<GameResults blueScore={blueScore} whiteScore={whiteScore} />}{' '}
          {
            <>
              <center className="game-ending-buttons">
                <Button onClick={() => handleResetGame()}>Reset game</Button>
                <br />
                <Button
                  onClick={() => {
                    setFinalScores({ blue: blueScore, white: whiteScore });
                    handleEndGame();
                  }}
                >
                  End game
                </Button>
              </center>
            </>
          }{' '}
        </>
      ) : (
        <GameStatistics finalScores={finalScores} handleEndGame={handleEndGame} handleResetGame={handleResetGame} />
      )}
    </>
  );
}

export default App;
