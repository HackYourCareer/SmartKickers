import React, { createContext, useCallback, useContext, useState } from 'react';
import useHeatmap from '../hooks/useHeatmap';

const StatsContext = createContext({});

export default function StatsContextProvider({ children }) {
  const { loading, error, heatmap } = useHeatmap();

  return (
    <StatsContext.Provider
      value={{
        heatmap,
        loading,
        error,
      }}
    >
      {children}
    </StatsContext.Provider>
  );
}

export function useStatsContext() {
  return useContext(StatsContext);
}
