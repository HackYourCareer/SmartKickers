import React, { createContext, useState } from 'react';

const GameDataContext = createContext({});

export function GameDataContextProvider() {
  const [blueScore, setBlueScore] = useState(0);
  const [whiteScore, setWhiteScore] = useState(0);
  const [finalScores, setFinalScores] = useState({ blueScore: 0, whiteScore: 0 });
  const [goalsArray, setGoalsArray] = useState([]);

  return <GameDataContext.Provider></GameDataContext.Provider>;
}
