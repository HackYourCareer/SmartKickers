import { useEffect, useState } from 'react';
import './App.css';
import { resetGame } from './apis/resetGame';
import GameStatistics from './components/Game/GameStatistics/GameStatistics.js';
import Heatmap from './components/Heatmap/Heatmap';
import config from './config';
import CurrentGameplay from './components/Game/CurrentGameplay/CurrentGameplay';
import { getHeatmapData } from './apis/heatmap';

function App() {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);
  const [clicked, setClicked] = useState(0);
  const [heatmap, setHeatmap] = useState([]);
  const [isStatisticsDisplayed, setIsStatisticsDisplayed] = useState(false);
  const [finalScores, setFinalScores] = useState({ blue: 0, white: 0 });
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

  const handleResetGame = () => {
    resetGame().then((data) => {
      if (data.error) alert(data.error);
    });
  };
  const handleEndGame = () => {
    setFinalScores({ blue: blueScore, white: whiteScore });
    setIsStatisticsDisplayed(!isStatisticsDisplayed);
    handleResetGame();
  };
  let heatMapTable = [];
  async function getHeatmap() {
    heatMapTable = await getHeatmapData();
    setClicked(true);
    console.log(heatMapTable);
  }

  return (
    <>
      <h1>Smart Kickers</h1>
      {isStatisticsDisplayed ? (
        <GameStatistics finalScores={finalScores} setIsStatisticsDisplayed={setIsStatisticsDisplayed} />
      ) : (
        <CurrentGameplay blueScore={blueScore} whiteScore={whiteScore} handleResetGame={handleResetGame} handleEndGame={handleEndGame} />
      )}
    </>
  );
}

export default App;
