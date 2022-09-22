import React from 'react';
import { Button } from '../../Button/Button';
import { useNavigate } from 'react-router-dom';
import './StatsHeader.css';
import { useLocation } from 'react-router-dom';

function StatsHeader() {
  const navigate = useNavigate();
  const location = useLocation();
  return (
    <>
      <h1>Smart Kickers</h1>
      <div className="nav-buttons">
        <Button onClick={() => navigate('/stats')} disabled={location.pathname === '/stats'}>
          Statistics
        </Button>
        <Button onClick={() => navigate('/stats/heatmap')} disabled={location.pathname === '/stats/heatmap'}>
          Heatmap
        </Button>
        <Button onClick={() => navigate('/stats/game-history')} disabled={location.pathname === '/stats/game-history'}>
          Game history
        </Button>
      </div>
    </>
  );
}

export default StatsHeader;
