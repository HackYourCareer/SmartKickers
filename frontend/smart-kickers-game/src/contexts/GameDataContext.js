import React, { createContext, useContext, useState, useEffect } from 'react';
import config from '../config';
const GameDataContext = createContext({});

export default function GameDataContextProvider({ children }) {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);
  const [goalsArray, setGoalsArray] = useState([]);
  const [finalScores, setFinalScores] = useState({ blueScore: 0, whiteScore: 0 });

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
      }}
    >
      {children}
    </GameDataContext.Provider>
  );
}

export function useGameDataContext() {
  return useContext(GameDataContext);
}
