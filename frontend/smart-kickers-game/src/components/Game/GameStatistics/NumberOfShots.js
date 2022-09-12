import React from 'react';
import './GameStatistics.css';

function NumberOfShots({ statistics }) {
  return (
    <>
      <div className="table-item">{statistics ? statistics.blueShotsCount : '0'}</div>
      <div className="table-item">number of all shots in the game</div>
      <div className="table-item">{statistics ? statistics.blueShotsCount : '0'}</div>
    </>
  );
}

export default NumberOfShots;
