import React from 'react';
import './StatisticItem.css';
import { TeamID } from '../../../../constants/score.js';

function ManualChangedGoals({ statistics }) {
  function getManualChangedGoals(teamID) {
    if (!statistics?.manualGoals) return;
    return statistics.manualGoals[teamID];
  }

  function getManualSubstractedGoals(teamID) {
    return getManualChangedGoals(teamID)?.sub || 0;
  }

  function getManualAddedGoals(teamID) {
    return getManualChangedGoals(teamID)?.add || 0;
  }
  return (
    <>
      <div className="table-item">{getManualAddedGoals(TeamID.Team_blue)}</div>
      <div className="table-item">Manually added goals</div>
      <div className="table-item">{getManualAddedGoals(TeamID.Team_white)}</div>
      <div className="table-item">{getManualSubstractedGoals(TeamID.Team_blue)}</div>
      <div className="table-item">Manually substracted goals</div>
      <div className="table-item">{getManualSubstractedGoals(TeamID.Team_white)}</div>
    </>
  );
}

export default ManualChangedGoals;
