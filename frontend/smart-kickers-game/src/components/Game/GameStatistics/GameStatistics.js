import React from 'react';
import './GameStatistics.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Button } from '../../Button/Button.js';

function GameStatistics({ finalScores, setIsStatisticsDisplayed }) {
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
      </div>
      <Button
        className="btn--primary new-game-btn"
        onClick={() => {
          setIsStatisticsDisplayed(false);
        }}
      >
        New game
      </Button>
    </>
  );
}

export default GameStatistics;
