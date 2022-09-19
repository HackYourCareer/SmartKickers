import React from 'react';
import './StatisticItem.css';
import { TeamID } from '../../../../constants/score.js';

function FastestShot({ statistics }) {
  return (
    <>
      <div className="table-item">{statistics?.teamID[TeamID.Team_blue]?.fastestShot}</div>
      <div className="table-item">fastest shot of the game</div>
      <div className="table-item">{statistics?.teamID[TeamID.Team_white]?.fastestShot}</div>
    </>
  );
}

export default FastestShot;
