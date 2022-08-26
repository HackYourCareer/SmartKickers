import { useEffect, useState } from 'react';
import './App.css';
import { resetGame } from './apis/resetGame';
import GameStatistics from './components/Game/GameStatistics/GameStatistics.js';
import config from './config';
import CurrentGameplay from './components/Game/CurrentGameplay/CurrentGameplay';
import { Goal, TeamID } from './constants/score.js';
import { useStopwatch } from 'react-timer-hook';

function App() {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);
  const [isStatisticsDisplayed, setIsStatisticsDisplayed] = useState(false);
  const [finalScores, setFinalScores] = useState({ blue: 0, white: 0 });
  let [goalsArray, setGoalsArray] = useState([]);
  const { seconds, minutes, isRunning, start, reset } = useStopwatch({ autoStart: false });

  useEffect(() => {
    const socket = new WebSocket(`${config.wsBaseUrl}/score`);

    socket.onopen = () => {
      // Send to server
      socket.send('Hello from client');
      socket.onmessage = (msg) => {
        msg = JSON.parse(msg.data);
        setBlueScore(msg.blueScore);
        setWhiteScore(msg.whiteScore);
      };
    };
  }, []);
  useEffect(() => {
    goalsArray.push(new Goal(TeamID.Team_blue, 'time: ' + minutes + ':' + seconds));
  }, [blueScore]);
  useEffect(() => {
    goalsArray.push(new Goal(TeamID.Team_white, 'time: ' + minutes + ':' + seconds));
  }, [whiteScore]);

  const handleStartGame = () => {
    resetGoalsArray();
    handleResetGame();
    start();
    alert('Game started');
  };

  const resetGoalsArray = () => {
    setGoalsArray([]);
  };

  const handleResetGame = () => {
    resetGame().then((data) => {
      if (data.error) alert(data.error);
    });
  };
  const handleEndGame = () => {
    setFinalScores({ blue: blueScore, white: whiteScore });
    setIsStatisticsDisplayed(!isStatisticsDisplayed);
    reset();
  };
  return (
    <>
      <h1>Smart Kickers</h1>
      {isStatisticsDisplayed ? (
        <GameStatistics
          finalScores={finalScores}
          setIsStatisticsDisplayed={setIsStatisticsDisplayed}
          handleResetGame={handleResetGame}
          stopwatchStart={start}
          resetGoalsArray={resetGoalsArray}
          goalsArray={goalsArray}
        />
      ) : (
        <CurrentGameplay
          blueScore={blueScore}
          whiteScore={whiteScore}
          handleStartGame={handleStartGame}
          handleResetGame={handleResetGame}
          handleEndGame={handleEndGame}
        />
      )}
    </>
  );
}

export default App;
