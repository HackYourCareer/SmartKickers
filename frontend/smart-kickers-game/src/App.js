import { useEffect, useState, useRef } from 'react';
import './App.css';
import { resetGame } from './apis/resetGame';
import GameStatistics from './components/Game/GameStatistics/GameStatistics.js';
import config from './config';
import CurrentGameplay from './components/Game/CurrentGameplay/CurrentGameplay';
import { Goal, TeamID } from './constants/score.js';
import { useStopwatch } from 'react-timer-hook';
import GameHistory from './components/Game/GameHistory/GameHistory';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Heatmap from './components/Heatmap/Heatmap';
import StatsHeader from './components/Game/Header/StatsHeader';
import { useGameDataContext } from './contexts/GameDataContext';
import StatsContextProvider from './contexts/StatsContext.js';

export default function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route
          path="/stats/heatmap"
          element={
            <StatsContextProvider>
              [<StatsHeader />, <Heatmap />]
            </StatsContextProvider>
          }
        />
        <Route
          path="/stats/gameHistory"
          element={
            <StatsContextProvider>
              [<StatsHeader />, <GameHistory />]{' '}
            </StatsContextProvider>
          }
        />
        <Route
          path="/stats"
          element={
            <StatsContextProvider>
              [<StatsHeader />, <GameStatistics />]
            </StatsContextProvider>
          }
        />
        <Route path="*" element={<App />} />
      </Routes>
    </BrowserRouter>
  );
}

function App() {
  const { seconds, minutes, isRunning, start, pause, reset } = useStopwatch({ autoStart: false });
  const [isVisible, setIsVisible] = useState(false);

  const { blueScore, setBlueScore, whiteScore, setWhiteScore, finalScores, setFinalScores, goalsArray, setGoalsArray } = useGameDataContext();
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

  const ScorePrevious = (value) => {
    const ref = useRef();
    useEffect(() => {
      ref.current = value;
    });
    return ref.current;
  };

  const prevBlueScore = ScorePrevious(blueScore);
  const prevWhiteScore = ScorePrevious(whiteScore);

  useEffect(() => {
    if (!isRunning) return;
    if (prevBlueScore > blueScore) {
      const indexOfLastBlue = goalsArray.indexOf(goalsArray.findLast((e) => e.teamID === TeamID.Team_blue));
      goalsArray.splice(indexOfLastBlue, 1);
    } else {
      goalsArray.push(new Goal(TeamID.Team_blue, 'time: ' + minutes.toString().padStart(2, '0') + ':' + seconds.toString().padStart(2, '0')));
    }
  }, [blueScore]);
  useEffect(() => {
    if (!isRunning) return;
    if (prevWhiteScore > whiteScore) {
      const indexOfLastWhite = goalsArray.indexOf(goalsArray.findLast((e) => e.teamID === TeamID.Team_white));
      goalsArray.splice(indexOfLastWhite, 1);
    } else {
      goalsArray.push(new Goal(TeamID.Team_white, 'time: ' + minutes.toString().padStart(2, '0') + ':' + seconds.toString().padStart(2, '0')));
    }
  }, [whiteScore]);

  const handleStartGame = () => {
    resetGoalsArray();
    handleResetGame();
    setIsVisible(true);
    start();
  };

  const resetGoalsArray = () => {
    setGoalsArray([]);
  };

  const handleResetGame = () => {
    resetGame().then((data) => {
      if (data.error) alert(data.error);
    });
    reset();
    start();
    resetGoalsArray();
  };
  const handleEndGame = () => {
    setFinalScores({ blueScore: blueScore, whiteScore: whiteScore });
    pause();
  };

  return (
    <>
      <h1>Smart Kickers</h1>
      <CurrentGameplay
        blueScore={blueScore}
        whiteScore={whiteScore}
        handleStartGame={handleStartGame}
        handleResetGame={handleResetGame}
        handleEndGame={handleEndGame}
        isVisible={isVisible}
      />
    </>
  );
}
