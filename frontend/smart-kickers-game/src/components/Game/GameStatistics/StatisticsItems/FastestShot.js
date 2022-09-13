import React from 'react';
import './StatisticItem.css';
import { TeamID } from '../../../../constants/score.js';

function FastestShot({ fastestShot }) {
  function returnFastestShot(teamID) {
    const { speed, team } = fastestShot;
    return team === teamID ? speed.toFixed(2) + ' km/h' : '😵';
  }
  return (
    <>
      <div className="table-item">{returnFastestShot(TeamID.Team_blue)}</div>
      <div className="table-item">fastest shot of the game</div>
      <div className="table-item">{returnFastestShot(TeamID.Team_white)}</div>
    </>
  );
}

export default FastestShot;
