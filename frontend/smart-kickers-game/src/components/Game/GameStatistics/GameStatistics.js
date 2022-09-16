import React, { useEffect, useState } from 'react';
import './GameStatistics.css';
import { getStatistics } from '../../../apis/getStatistics.js';
import { TeamID } from '../../../constants/score.js';

import NumberOfShots from './NumberOfShots.js';
import FinalScores from './StatisticsItems/FinalScores.js';
import FastestShot from './StatisticsItems/FastestShot.js';
import ManualChangedGoals from './StatisticsItems/ManualChangedGoals.js';
import TeamIcons from './StatisticsItems/TeamIcons.js';
import { useGameDataContext } from '../../../contexts/GameDataContext';

function GameStatistics() {
  const [statistics, setStatistics] = useState(null);
  const finalScores = useGameDataContext().finalScores;

  const handleGetStatistics = async () => {
    const result = await getStatistics();
    result?.error ? alert(result.error) : setStatistics(result);
  };

  useEffect(() => {
    handleGetStatistics();
  }, []);

  return (
    <>
      <h2>
        <em>Statistics</em>
      </h2>
      {statistics && statistics?.fastestShot?.speed !== 0 ? (
        <div className="table-with-stats">
          <TeamIcons />
          <FinalScores blue={finalScores.blueScore} white={finalScores.whiteScore} />
          <FastestShot fastestShot={statistics.fastestShot} />
          <ManualChangedGoals blue={TeamID.Team_blue} white={TeamID.Team_white} statistics={statistics} />
          <NumberOfShots statistics={statistics} />{' '}
        </div>
      ) : (
        <div className="no-statistics">Something went wrong, statistics went on the vacation and we don't have it</div>
      )}
    </>
  );
}

export default GameStatistics;
