import React, { createContext, useContext, useState, useEffect } from 'react';
import config from '../config';
import { useStopwatch } from 'react-timer-hook';
import { resetGame } from '../apis/resetGame';
const GameDataContext = createContext({});

export default function GameDataContextProvider({ children }) {
  const { seconds, minutes, start, reset } = useStopwatch({ autoStart: false });
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);
  const [goalsArray, setGoalsArray] = useState([]);
  const [finalScores, setFinalScores] = useState({ blueScore: 0, whiteScore: 0 });
  const [isGameEnded, setIsGameEnded] = useState(false);
  const [isGameStarted, setIsGameStarted] = useState(false);

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

  return (
    <GameDataContext.Provider
      value={{
        blueScore,
        setBlueScore,
        whiteScore,
        setWhiteScore,
        finalScores,
        setFinalScores,
        goalsArray,
        setGoalsArray,
        isGameEnded,
        setIsGameEnded,
        isGameStarted,
        setIsGameStarted,
        handleResetGame,
        seconds,
        minutes,
        start,
        reset,
        resetGoalsArray,
      }}
    >
      {children}
    </GameDataContext.Provider>
  );
}

export function useGameDataContext() {
  return useContext(GameDataContext);
}
