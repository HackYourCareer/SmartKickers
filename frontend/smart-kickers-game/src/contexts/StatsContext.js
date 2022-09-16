import React, { createContext, useCallback, useContext, useState } from 'react';
import useHeatmap from '../hooks/useHeatmap';

const StatsContext = createContext({});

export default function StatsContextProvider({ children }) {
  const { loading, error, heatmap } = useHeatmap();
  //const heaatmaaap = { loading, error, heatmap };
  return (
    <StatsContext.Provider
      value={{
        heaatmaaap: { loading, error, heatmap },
      }}
    >
      {children}
    </StatsContext.Provider>
  );
}

export function useStatsContext() {
  return useContext(StatsContext);
}
