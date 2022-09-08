import React from 'react';
import './GameStatistics.css';
import { Button } from '../../Button/Button.js';
import { getStatistics } from '../../../apis/getStatistics.js';
import { useEffect, useState } from 'react';
import { TeamID } from '../../../constants/score.js';
import FinalScores from './StatisticsItems/FinalScores.js';
import FastestShot from './StatisticsItems/FastestShot.js';
import ManualChangedGoals from './StatisticsItems/ManualChangedGoals.js';
import TeamIcons from './StatisticsItems/TeamIcons.js';

function GameStatistics({ finalScores, onNewGameRequested }) {
  const [statistics, setStatistics] = useState(null);

  const handleGetStatistics = async () => {
    const result = await getStatistics();
    if (result?.error) alert(result.error);
    setStatistics(result);
  };

  useEffect(() => {
    handleGetStatistics();
  }, []);

  return (
    <>
      <h2>
        <em>Statistics</em>
      </h2>
      <div className="table-with-stats">
        <TeamIcons />
        <FinalScores blue={finalScores.blue} white={finalScores.white} />
        <FastestShot blue={TeamID.Team_blue} white={TeamID.Team_white} statistics={statistics} />
        <ManualChangedGoals blue={TeamID.Team_blue} white={TeamID.Team_white} statistics={statistics} />
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
