import React from 'react';
import HeatMap from 'react-heatmap-grid';
import './Heatmap.css';
import { chooseColor } from './Colors';
import useHeatmap from '../../hooks/useHeatmap';

function Heatmap() {
  const { loading, error, heatmap } = useHeatmap();
  if (loading) return <div className="heatmap-status">Loading...</div>;
  if (error) return <div className="heatmap-status">Error</div>;

  return (
    <div className="heatmap-parent">
      <div className="heatmap-container">
        <HeatMap
          xLabels={heatmap.array}
          yLabels={heatmap.array}
          data={heatmap.numbersCopy}
          cellStyle={(background, value, min, max) => ({
            background: `${chooseColor(value)} `,
            fontSize: '0px',
            color: '#444',
            width: '3.28px',
            height: '3.28px',
            margin: '0px',
          })}
          cellRender={(value) => value && <div>{value}</div>}
        />
      </div>
    </div>
  );
}
export default Heatmap;
