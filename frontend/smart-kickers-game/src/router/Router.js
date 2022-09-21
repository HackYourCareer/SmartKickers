import React from 'react';
import GameStatistics from '../components/Game/GameStatistics/GameStatistics.js';
import GameHistory from '../components/Game/GameHistory/GameHistory';
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';
import Heatmap from '../components/Heatmap/Heatmap';
import App from '../App.js';
import StatsItem from './StatsItem.js';
import { useGameDataContext } from '../contexts/GameDataContext.js';

export default function Router() {
  const { isGameEnded } = useGameDataContext();
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        {isGameEnded && (
          <>
            <Route
              path="/stats/heatmap"
              element={
                <StatsItem>
                  <Heatmap />
                </StatsItem>
              }
            />
            <Route
              path="/stats/game-history"
              element={
                <StatsItem>
                  <GameHistory />
                </StatsItem>
              }
            />
            <Route
              path="/stats"
              element={
                <StatsItem>
                  <GameStatistics />
                </StatsItem>
              }
            />
          </>
        )}
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </BrowserRouter>
  );
}
