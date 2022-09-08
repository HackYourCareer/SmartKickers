import React from 'react';
import './StatisticItem.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function TeamIcons() {
  return (
    <>
      <div className="table-item">
        <FontAwesomeIcon className="blue-team-icon" icon="fa-person" />
        Blue
      </div>
      <div className="table-item"></div>
      <div className="table-item">
        <FontAwesomeIcon className="white-team-icon" icon="fa-person" />
        White
      </div>
    </>
  );
}

export default TeamIcons;
