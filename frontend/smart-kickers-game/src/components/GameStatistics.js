import React from 'react';
import './GameStatistics.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
function GameStatistics({ blueScore, whiteScore }) {
  return (
    <>
      <table>
        <thead>
          <tr>
            <th>
              <FontAwesomeIcon icon="fa-person" />
              Blue
            </th>
            <th></th>
            <th>
              <FontAwesomeIcon icon="fa-person" />
              White
            </th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>{blueScore}</td>
            <td>
              <strong>score</strong>
            </td>
            <td>{whiteScore}</td>
          </tr>
        </tbody>
      </table>
    </>
  );
}

export default GameStatistics;
