import React, { createContext, useCallback, useContext, useState } from 'react';
import useHeatmap from '../hooks/useHeatmap';

const StatsContext = createContext({});
//Todo Remove this component
export default function StatsContextProvider({ children }) {
  const { loading, error, heatmap } = useHeatmap();
  return (
    <StatsContext.Provider
      value={{
        heatmapData: { loading, error, heatmap },
      }}
    >
      {children}
    </StatsContext.Provider>
  );
}

export function useStatsContext() {
  return useContext(StatsContext);
}
