import React from 'react';
import './StatisticItem.css';

function FastestShot({ blue, white, statistics }) {
  function returnFastestShot(teamID) {
    if (!statistics?.fastestShot) return '😞';
    const { speed, team } = statistics.fastestShot;
    return team === teamID ? speed.toFixed(2) + ' km/h' : '😵';
  }
  return (
    <>
      <div className="table-item">{returnFastestShot(blue)}</div>
      <div className="table-item">fastest shot of the game</div>
      <div className="table-item">{returnFastestShot(white)}</div>
    </>
  );
}

export default FastestShot;
