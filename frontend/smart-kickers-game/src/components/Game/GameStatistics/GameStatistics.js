import React from 'react';
import './GameStatistics.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Button } from '../../Button/Button.js';
import { getStatistics } from '../../../apis/getStatistics.js';
import { useEffect, useState } from 'react';
import { TeamID } from '../../../constants/score.js';

function GameStatistics({ finalScores, onNewGameRequested }) {
  const [statistics, setStatistics] = useState(null);

  const handleGetStatistics = async () => {
    const result = await getStatistics();
    if (result?.error) alert(result.error);
    setStatistics(result);
  };

  function returnFastestShot(teamID) {
    if (!statistics) return;
    const { Speed, Team } = statistics.FastestShot;
    return Team === teamID ? Speed.toFixed(2) + ' km/h' : ':(';
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
