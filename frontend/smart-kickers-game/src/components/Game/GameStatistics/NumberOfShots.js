import React from 'react';
import './GameStatistics.css';

function NumberOfShots({ statistics }) {
  console.log(statistics);
  return (
    <>
      <div className="table-item">{statistics?.blueShotsCount}</div>
      <div className="table-item">number of all shots in the game</div>
      <div className="table-item">{statistics?.whiteShotsCount}</div>
    </>
  );
}

export default NumberOfShots;
