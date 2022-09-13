import React from 'react';
import './StatisticItem.css';

function FinalScores({ blue, white }) {
  return (
    <>
      <div className="table-item">{blue}</div>
      <div className="table-item">score</div>
      <div className="table-item">{white}</div>
    </>
  );
}

export default FinalScores;
