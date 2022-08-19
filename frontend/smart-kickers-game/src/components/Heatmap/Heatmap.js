import React from 'react';
import HeatMap from 'react-heatmap-grid';
import './Heatmap.css';

function Heatmap({ heatMapTable }) {
  const heatmapDim = Object.keys(heatMapTable.data).length;
  const array = new Array(heatmapDim).fill(0).map(() => '');
  let numbersCopy = JSON.parse(JSON.stringify(heatMapTable.data));

  for (let i = 0; i < heatmapDim; i++)
    for (let j = 0; j <= i; j++) {
      numbersCopy[i][j] = numbersCopy[i][j] + numbersCopy[j][i] - numbersCopy[j][i];
      numbersCopy[j][i] = numbersCopy[i][j];
    }

  return (
    <div className="heatmap-parent">
      <div className="heatmap-container">
        <HeatMap
          xLabels={array}
          yLabels={array}
          data={numbersCopy}
          cellStyle={(background, value, min, max) => ({
            background: `rgb(0, 151, 230, ${1 - (max - value) / (max - min)})`,
            fontSize: '0px',
            color: '#444',
            width: '4px',
            height: '3px',
          })}
          cellRender={(value) => value && <div>{value}</div>}
        />
      </div>
    </div>
  );
}
export default Heatmap;
