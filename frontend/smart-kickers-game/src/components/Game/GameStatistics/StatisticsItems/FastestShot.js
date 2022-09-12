import React from 'react';
import './StatisticItem.css';

function FastestShot({ blue, white, fastestShot }) {
  function returnFastestShot(teamID) {
    const { speed, team } = fastestShot;
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
