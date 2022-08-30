import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import './GameHistory.css';
import { TeamID } from '../../../constants/score.js';

function GameHistory({ goalsArray }) {
  return (
    <div>
      <h2>
        <em>Game History</em>
      </h2>
      <div className="table-with-stats history">
        <div className="table-item">
          <FontAwesomeIcon className="blue-team-icon" icon="fa-person" />
          Blue
        </div>
        <div className="table-item">
          <FontAwesomeIcon className="white-team-icon" icon="fa-person" />
          White
        </div>
        {goalsArray.map((item) =>
          item.teamID === TeamID.Team_blue ? (
            <>
              <div key={item} className="table-item">
                {item.timestamp}
              </div>
              <div key={item + item.timestamp} className="table-item"></div>
            </>
          ) : (
            <>
              <div key={item + item.timestamp} className="table-item"></div>{' '}
              <div key={item} className="table-item">
                {item.timestamp}
              </div>
            </>
          )
        )}
      </div>
    </div>
  );
}

export default GameHistory;
