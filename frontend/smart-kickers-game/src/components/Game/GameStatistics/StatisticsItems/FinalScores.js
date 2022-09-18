import React from 'react';
import './StatisticItem.css';
import { useGameDataContext } from '../../../../contexts/GameDataContext';

function FinalScores() {
  const finalScores = useGameDataContext().finalScores;
  return (
    <>
      <div className="table-item">{finalScores.blueScore}</div>
      <div className="table-item">score</div>
      <div className="table-item">{finalScores.whiteScore}</div>
    </>
  );
}

export default FinalScores;
