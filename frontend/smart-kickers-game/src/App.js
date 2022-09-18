import React from 'react';
import './App.css';
import GameStatistics from './components/Game/GameStatistics/GameStatistics.js';
import CurrentGameplay from './components/Game/CurrentGameplay/CurrentGameplay';
import GameHistory from './components/Game/GameHistory/GameHistory';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Heatmap from './components/Heatmap/Heatmap';
import StatsHeader from './components/Game/Header/StatsHeader';
import NewGameButton from './components/Button/NewGameButton.js';

export default function Router() {
  return (
    // TODO move to separate component
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

function App() {
  return (
    <>
      <h1>Smart Kickers</h1>
      <CurrentGameplay />
    </>
  );
}
