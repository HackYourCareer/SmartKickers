import React from 'react';
import { Button } from '../components/Button';
import { addPoint, subPoint } from '../apis/goalManipulation.js';
import './GameResults.css';

function GameResults({ blueScore, whiteScore }) {
  return (
    <div>
      <div className="game-result-container">
        <div className="game-result-item">
          <Button
            onClick={() => {
              addPoint(2);
            }}
          >
            +
          </Button>
          Blue: {blueScore}
          <Button
            onClick={() => {
              subPoint(2);
            }}
          >
            -
          </Button>
        </div>

        <div className="game-result-item">
          <Button
            onClick={() => {
              addPoint(1);
            }}
          >
            +
          </Button>
          White: {whiteScore}
          <Button
            onClick={() => {
              subPoint(1);
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
