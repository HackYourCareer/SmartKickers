import React, { createContext, useContext, useState } from 'react';
const CurrentGameplayContext = createContext({});

export default function CurrentGameplayContextProvider({ children }) {
  const [goalsArray, setGoalsArray] = useState([]);
  const [finalScores, setFinalScores] = useState({ blueScore: 0, whiteScore: 0 });
  return (
    <CurrentGameplayContext.Provider
      value={{
        finalScores,
        setFinalScores,
        goalsArray,
        setGoalsArray,
      }}
    >
      {children}
    </CurrentGameplayContext.Provider>
  );
}

export function useGameDataContext() {
  return useContext(CurrentGameplayContext);
}
