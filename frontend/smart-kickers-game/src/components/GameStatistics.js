import React from 'react';
import './GameStatistics.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
function GameStatistics({ blueScore, whiteScore }) {
  return (
    <>
      <h2>
        <em>Statistics</em>
      </h2>
      <table>
        <thead>
          <tr>
            <th>
              <FontAwesomeIcon className="bluePerson" icon="fa-person" />
              Blue
            </th>
            <th></th>
            <th>
              <FontAwesomeIcon className="whitePerson" icon="fa-person" />
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
