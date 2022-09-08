import React from 'react';
import './GameStatistics.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Button } from '../../Button/Button.js';
import { getStatistics } from '../../../apis/getStatistics.js';
import { useEffect, useState } from 'react';
import { TeamID } from '../../../constants/score.js';
import StatisticsItem from './StatisticsItems/StatisticItem';

function GameStatistics({ finalScores, onNewGameRequested }) {
  const [statistics, setStatistics] = useState(null);

  const handleGetStatistics = async () => {
    const result = await getStatistics();
    if (result?.error) alert(result.error);
    setStatistics(result);
  };

  function returnFastestShot(teamID) {
    if (!statistics?.fastestShot) return '😞';
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
        <StatisticsItem>
          {' '}
          <FontAwesomeIcon className="blue-team-icon" icon="fa-person" />
          Blue
        </StatisticsItem>
        <StatisticsItem />
        <StatisticsItem>
          {' '}
          <FontAwesomeIcon className="white-team-icon" icon="fa-person" />
          White
        </StatisticsItem>
        <StatisticsItem>{finalScores.blue}</StatisticsItem>
        <StatisticsItem>score</StatisticsItem>
        <StatisticsItem>{finalScores.white}</StatisticsItem>
        <StatisticsItem>{returnFastestShot(TeamID.Team_blue)}</StatisticsItem>
        <StatisticsItem>fastest shot of the game</StatisticsItem>
        <StatisticsItem>{returnFastestShot(TeamID.white)}</StatisticsItem>
        <StatisticsItem>{getManualAddedGoals(TeamID.Team_blue)}</StatisticsItem>
        <StatisticsItem>Manually added goals</StatisticsItem>
        <StatisticsItem>{getManualAddedGoals(TeamID.Team_white)}</StatisticsItem>
        <StatisticsItem>{getManualSubstractedGoals(TeamID.Team_blue)}</StatisticsItem>
        <StatisticsItem>Manually substracted goals</StatisticsItem>
        <StatisticsItem>{getManualSubstractedGoals(TeamID.Team_white)}</StatisticsItem>
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
