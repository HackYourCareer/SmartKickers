import React from 'react';
import './GameStatistics.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Button } from '../Button/Button.js';

function GameStatistics({ finalScores, setIsStatisticsDisplayed }) {
  return (
    <>
      <h2>
        <em>Statistics</em>
      </h2>
      <div class="table-with-stats">
        <div class="table-item">
          <FontAwesomeIcon className="blueTeamIcon" icon="fa-person" />
          Blue
        </div>
        <div class="table-item"></div>
        <div class="table-item">
          <FontAwesomeIcon className="whiteTeamIcon" icon="fa-person" />
          White
        </div>
        <div class="table-item">{finalScores.blue}</div>
        <div class="table-item">score</div>
        <div class="table-item">{finalScores.white}</div>
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
