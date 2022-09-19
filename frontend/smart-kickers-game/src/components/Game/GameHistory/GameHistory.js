import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import './GameHistory.css';
import { TeamID } from '../../../constants/score.js';
import { useGameDataContext } from '../../../contexts/GameDataContext';

function GameHistory() {
  const goalsArray = useGameDataContext().goalsArray;
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
        {goalsArray.map((item) => (
          <GoalsArrayItem item={item} />
        ))}
      </div>
    </div>
  );
}

function GoalsArrayItem({ item }) {
  return item.teamID === TeamID.Team_blue ? (
    <React.Fragment key={item.teamID + item.timestamp}>
      <div className="table-item">{item.timestamp}</div>
      <div className="table-item"></div>
    </React.Fragment>
  ) : (
    <React.Fragment key={item.teamID + item.timestamp}>
      <div className="table-item"></div>
      <div className="table-item">{item.timestamp}</div>
    </React.Fragment>
  );
}

export default GameHistory;
