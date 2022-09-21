import React from 'react';
import { Button } from '../../Button/Button';
import { useNavigate } from 'react-router-dom';
import './StatsHeader.css';

function StatsHeader() {
  const navigate = useNavigate();
  return (
    <>
      <h1>Smart Kickers</h1>
      <div className="nav-buttons">
        <Button onClick={() => navigate('/stats')}>Statistics</Button>
        <Button onClick={() => navigate('/stats/heatmap')}>Heatmap</Button>
        <Button onClick={() => navigate('/stats/game-history')}>Game history</Button>
      </div>
    </>
  );
}

export default StatsHeader;
