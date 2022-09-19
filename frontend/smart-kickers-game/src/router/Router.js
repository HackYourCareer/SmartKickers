import React from 'react';
import GameStatistics from '../components/Game/GameStatistics/GameStatistics.js';
import GameHistory from '../components/Game/GameHistory/GameHistory';
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';
import Heatmap from '../components/Heatmap/Heatmap';
import App from '../App.js';
import RouterItem from './RouterItem.js';

export default function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route
          path="/stats/heatmap"
          element={
            <RouterItem>
              <Heatmap />
            </RouterItem>
          }
        />
        <Route
          path="/stats/gameHistory"
          element={
            <RouterItem>
              <GameHistory />
            </RouterItem>
          }
        />
        <Route
          path="/stats"
          element={
            <RouterItem>
              <GameStatistics />
            </RouterItem>
          }
        />
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </BrowserRouter>
  );
}
