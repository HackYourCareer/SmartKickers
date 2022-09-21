import React, { useEffect, useState } from 'react';
import './GameStatistics.css';
import { getStatistics } from '../../../apis/getStatistics.js';
import NumberOfShots from './NumberOfShots.js';
import FinalScores from './StatisticsItems/FinalScores.js';
import FastestShot from './StatisticsItems/FastestShot.js';
import ManualChangedGoals from './StatisticsItems/ManualChangedGoals.js';
import TeamIcons from './StatisticsItems/TeamIcons.js';

function GameStatistics() {
  const [statistics, setStatistics] = useState(null);

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
          <FinalScores />
          <FastestShot fastestShot={statistics.fastestShot} />
          <ManualChangedGoals statistics={statistics} />
          <NumberOfShots statistics={statistics} />{' '}
        </div>
      ) : (
        <div className="no-statistics">Something went wrong, statistics went on the vacation and we don't have it</div>
      )}
    </>
  );
}

export default GameStatistics;
