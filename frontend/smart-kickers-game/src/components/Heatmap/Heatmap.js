import React from 'react';
import HeatMap from 'react-heatmap-grid';
import './Heatmap.css';

function Heatmap({ heatMapTable }) {
  const xLabels = new Array(Object.keys(heatMapTable.data).length).fill(0).map((_, i) => ``);

  // Display only even labels
  const xLabelsVisibility = new Array(24).fill(0).map((_, i) => (i % 2 === 0 ? true : false));

  let i = 0;
  const yLabels = new Array(Object.keys(heatMapTable.data).length).fill(0).map((_, i) => ``);
  const data = new Array(yLabels.length).fill(0).map(() => new Array(xLabels.length).fill(0).map(() => Math.floor(i++)));
  console.log(heatMapTable);

  let numbersCopy = JSON.parse(JSON.stringify(heatMapTable.data));

  let n = 100;
  for (let i = 0; i < n; i++)
    for (let j = 0; j <= i; j++) numbersCopy[i][j] = numbersCopy[i][j] + numbersCopy[j][i] - (numbersCopy[j][i] = numbersCopy[i][j]);

  return (
    <div className="heatmap-parent">
      <div className="heatmap-container">
        <HeatMap
          xLabels={xLabels}
          yLabels={yLabels}
          xLabelsLocation={'bottom'}
          xLabelsVisibility={xLabelsVisibility}
          xLabelWidth={60}
          data={numbersCopy}
          height={100}
          cellStyle={(background, value, min, max, data, x, y) => ({
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
