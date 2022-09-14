import React from 'react';
import { Button } from '../../Button/Button';

function StatsHeader() {
  return (
    <>
      <h1>Smart Kickers</h1>
      <Button>Statistics</Button>
      <Button>Heatmap</Button>
      <Button>Game history</Button>
    </>
  );
}

export default StatsHeader;
