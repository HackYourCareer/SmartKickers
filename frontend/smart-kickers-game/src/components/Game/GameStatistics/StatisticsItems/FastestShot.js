import React from 'react';
import './StatisticItem.css';
import { TeamID } from '../../../../constants/score.js';

function FastestShot({ statistics }) {
  function returnFastestShot(teamID) {
    return statistics?.teamID[teamID]?.fastestShot.toFixed(2) + ' km/h';
  }
  return (
    <>
      <div className="table-item">{returnFastestShot(TeamID.Team_blue)}</div>
      <div className="table-item">Fastest shot of the game</div>
      <div className="table-item">{returnFastestShot(TeamID.Team_white)}</div>
    </>
  );
}

export default FastestShot;
