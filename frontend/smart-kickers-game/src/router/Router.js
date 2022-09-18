import React from 'react';
import GameStatistics from '../components/Game/GameStatistics/GameStatistics.js';
import GameHistory from '../components/Game/GameHistory/GameHistory';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Heatmap from '../components/Heatmap/Heatmap';
import StatsHeader from '../components/Game/Header/StatsHeader';
import NewGameButton from '../components/Button/NewGameButton.js';
import App from '../App.js';

export default function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />} />
        <Route
          path="/stats/heatmap"
          element={
            //Todo Wrap this
            <>
              <StatsHeader /> <Heatmap /> <NewGameButton />
            </>
          }
        />
        <Route
          path="/stats/gameHistory"
          element={
            //Todo Wrap this
            <>
              <StatsHeader /> <GameHistory /> <NewGameButton />
            </>
          }
        />
        <Route
          path="/stats"
          element={
            //Todo Wrap this
            <>
              <StatsHeader /> <GameStatistics /> <NewGameButton />
            </>
          }
        />
        <Route path="*" element={<App />} />
      </Routes>
    </BrowserRouter>
  );
}
