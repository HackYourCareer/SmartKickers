import React from 'react';
import './GameStatistics.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Button } from '../../Button/Button.js';
import { getStatistics } from '../../../apis/getStatistics.js';
import { useEffect, useState } from 'react';
import { TeamID } from '../../../constants/score.js';
import NumberOfShots from './NumberOfShots.js';

function GameStatistics({ finalScores, onNewGameRequested }) {
  const [statistics, setStatistics] = useState(null);

  const handleGetStatistics = async () => {
    const result = await getStatistics();
    if (result?.error) alert(result.error);
    setStatistics(result);
  };

  function returnFastestShot(teamID) {
    if (!statistics?.fastestShot) return;
    const { speed, team } = statistics.fastestShot;
    return team === teamID ? speed.toFixed(2) + ' km/h' : '😵';
  }

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

  useEffect(() => {
    handleGetStatistics();
  }, []);

  return (
    <>
      <h2>
        <em>Statistics</em>
      </h2>
      <div className="table-with-stats">
        <div className="table-item">
          <FontAwesomeIcon className="blue-team-icon" icon="fa-person" />
          Blue
        </div>
        <div className="table-item"></div>
        <div className="table-item">
          <FontAwesomeIcon className="white-team-icon" icon="fa-person" />
          White
        </div>
        <div className="table-item">{finalScores.blue}</div>
        <div className="table-item">score</div>
        <div className="table-item">{finalScores.white}</div>
        <div className="table-item">{returnFastestShot(TeamID.Team_blue)}</div>
        <div className="table-item">fastest shot of the game</div>
        <div className="table-item">{returnFastestShot(TeamID.Team_white)}</div>
        <div className="table-item">{getManualAddedGoals(TeamID.Team_blue)}</div>
        <div className="table-item">Manually added goals</div>
        <div className="table-item">{getManualAddedGoals(TeamID.Team_white)}</div>
        <div className="table-item">{getManualSubstractedGoals(TeamID.Team_blue)}</div>
        <div className="table-item">Manually substracted goals</div>
        <div className="table-item">{getManualSubstractedGoals(TeamID.Team_white)}</div>
        <NumberOfShots statistics={statistics} />
      </div>

      <Button
        className="btn--primary new-game-btn"
        onClick={() => {
          onNewGameRequested();
        }}
      >
        New game
      </Button>
    </>
  );
}

export default GameStatistics;
