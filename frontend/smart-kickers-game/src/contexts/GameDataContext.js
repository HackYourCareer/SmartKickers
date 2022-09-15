import React, { createContext, useCallback, useContext, useState } from 'react';

const GameDataContext = createContext({});

export default function GameDataContextProvider({ children }) {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);
  const [goalsArray, setGoalsArray] = useState([]);
  const [finalScores, setFinalScores] = useState({ blueScore: 0, whiteScore: 0 });

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
