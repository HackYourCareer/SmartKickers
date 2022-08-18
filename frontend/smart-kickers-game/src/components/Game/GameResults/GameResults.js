import React from 'react';
import { Button } from '../../Button/Button';
import { pointsManipulation } from '../../../apis/goalManipulation.js';
import './GameResults.css';
import { TeamID, ScoreChange } from '../../../constants/score.js';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function GameResults({ blueScore, whiteScore }) {
  return (
    <div>
      <div className="game-result-container">
        <div className="game-result-item">
          <Button
            onClick={() => {
              pointsManipulation(TeamID.Team_blue, ScoreChange.Add_goal);
            }}
          >
            +
          </Button>
          <div className="icon-result">
            <FontAwesomeIcon className="blueTeamIcon" icon="fa-person" />
            Blue: <span className="score"> {blueScore}</span>
          </div>
          <Button
            onClick={() => {
              pointsManipulation(TeamID.Team_blue, ScoreChange.Sub_goal);
            }}
            disabled={blueScore === 0}
          >
            -
          </Button>
        </div>

        <div className="game-result-item">
          <Button
            onClick={() => {
              pointsManipulation(TeamID.Team_white, ScoreChange.Add_goal);
            }}
          >
            +
          </Button>
          <div className="icon-result">
            <FontAwesomeIcon className="whiteTeamIcon" icon="fa-person" />
            White: <span className="score">{whiteScore}</span>
          </div>
          <Button
            onClick={() => {
              pointsManipulation(TeamID.Team_white, ScoreChange.Sub_goal);
            }}
            disabled={whiteScore === 0}
          >
            -
          </Button>
        </div>
      </div>
    </div>
  );
}

export default GameResults;
