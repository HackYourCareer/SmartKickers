import React from 'react';
import './GameStatistics.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Button } from '../components/Button.js';

function GameStatistics({ finalScores, handleEndGame }) {
  function handleNewGame() {
    handleEndGame();
    console.log('dupa');
  }
  return (
    <>
      <h2>
        <em>Statistics</em>
      </h2>
      <table>
        <thead>
          <tr>
            <th>
              <FontAwesomeIcon className="blueTeamIcon" icon="fa-person" />
              Blue
            </th>
            <th></th>
            <th>
              <FontAwesomeIcon className="whiteTeamIcon" icon="fa-person" />
              White
            </th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>{finalScores.blue}</td>
            <td>
              <strong>score</strong>
            </td>
            <td>{finalScores.white}</td>
          </tr>
        </tbody>
      </table>
      <Button className="btn--primary new-game-btn" onClick={() => handleNewGame()}>
        New game
      </Button>
    </>
  );
}

export default GameStatistics;
