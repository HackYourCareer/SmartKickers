import React from 'react';
import { Button } from '../components/Button';
import { pointsManipulation } from '../apis/goalManipulation.js';
import './GameResults.css';
import { TeamID, ScoreChange} from '../constants/score.js';

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
          Blue: {blueScore}
          <Button
            onClick={() => {
              pointsManipulation(TeamID.Team_blue, ScoreChange.Sub_goal);
            }}
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
          White: {whiteScore}
          <Button
            onClick={() => {
              pointsManipulation(TeamID.Team_white, ScoreChange.Sub_goal);
            }}
          >
            -
          </Button>
        </div>
      </div>
    </div>
  );
}

export default GameResults;
