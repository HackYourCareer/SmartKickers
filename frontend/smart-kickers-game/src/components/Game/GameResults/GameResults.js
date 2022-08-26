import React from 'react';
import { Button } from '../../Button/Button';
import { updateScores } from '../../../apis/updateScores.js';
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
              updateScores(TeamID.Team_blue, ScoreChange.Add_goal).then((data) => {
                if (data.error) alert(data.error);
              });
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
              updateScores(TeamID.Team_blue, ScoreChange.Sub_goal).then((data) => {
                if (data.error) alert(data.error);
              });
            }}
            disabled={blueScore === 0}
          >
            -
          </Button>
        </div>

        <div className="game-result-item">
          <Button
            onClick={() => {
              updateScores(TeamID.Team_white, ScoreChange.Add_goal).then((data) => {
                if (data.error) alert(data.error);
              });
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
              updateScores(TeamID.Team_white, ScoreChange.Sub_goal).then((data) => {
                if (data.error) alert(data.error);
              });
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
