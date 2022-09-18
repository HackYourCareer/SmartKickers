import React from 'react';
import StatsHeader from '../components/Game/Header/StatsHeader';
import NewGameButton from '../components/Button/NewGameButton.js';

export default function RouterItem({ children }) {
  return (
    <>
      <StatsHeader /> {children} <NewGameButton />
    </>
  );
}
