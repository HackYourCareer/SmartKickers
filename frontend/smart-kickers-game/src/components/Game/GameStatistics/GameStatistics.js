import React from 'react';
import './GameStatistics.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Button } from '../../Button/Button.js';
import { getStatistics } from '../../../apis/getStatistics.js';
import { useEffect, useState } from 'react';
import { TeamID } from '../../../constants/score.js';

function GameStatistics({ finalScores, onNewGameRequested }) {
  const [data, setData] = useState({});
  useEffect(() => {
    getStatistics().then((result) => {
      if (result?.error) alert(result.error);
      setData(result.data);
    });
  });
  function checkIfFastestShot(teamID) {
    return data?.FastestShot.Team === teamID ? data?.FastestShot.Speed.toFixed(2) + ' km/h' : '';
  }
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
        {data?.FastestShot && <div className="table-item">{checkIfFastestShot(TeamID.Team_blue)}</div>}
        <div className="table-item">fastest shot of the game</div>
        {data?.FastestShot && <div className="table-item">{checkIfFastestShot(TeamID.Team_white)}</div>}
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
