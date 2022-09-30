import React from 'react';
import './StatisticItem.css';
import { useGameDataContext } from '../../../../contexts/GameDataContext';

function FinalScores() {
  const { finalScores } = useGameDataContext();
  return (
    <>
      <div className="table-item">{finalScores.blueScore}</div>
      <div className="table-item">Score</div>
      <div className="table-item">{finalScores.whiteScore}</div>
    </>
  );
}

export default FinalScores;
