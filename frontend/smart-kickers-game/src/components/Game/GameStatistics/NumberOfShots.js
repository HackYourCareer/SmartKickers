import React from 'react';
import './GameStatistics.css';
import { TeamID } from '../../../constants/score.js';

function NumberOfShots({ statistics }) {
  return (
    <>
      <div className="table-item">{statistics ? statistics?.teamID[TeamID.Team_blue]?.shotsCount : '0'}</div>
      <div className="table-item">Number of all shots in the game</div>
      <div className="table-item">{statistics ? statistics?.teamID[TeamID.Team_white]?.shotsCount : '0'}</div>
    </>
  );
}

export default NumberOfShots;
