import React from 'react';
import HeatMap from 'react-heatmap-grid';
import './Heatmap.css';
import { Colors } from './Colors';

function Heatmap({ heatMapTable }) {
  const heatmapDim = Object.keys(heatMapTable.data).length;
  const array = new Array(heatmapDim).fill(0).map(() => '');
  let numbersCopy = JSON.parse(JSON.stringify(heatMapTable.data));

  for (let i = 0; i < heatmapDim; i++)
    for (let j = 0; j <= i; j++) {
      numbersCopy[i][j] = numbersCopy[i][j] + numbersCopy[j][i] - numbersCopy[j][i];
      numbersCopy[j][i] = numbersCopy[i][j];
    }

  function chooseColor(value) {
    let chosenColor = Colors.none;
    if (value > 2 && value <= 5) {
      chosenColor = Colors.blue;
    } else if (value > 5 && value <= 10) {
      chosenColor = Colors.purple;
    } else if (value > 10 && value <= 15) {
      chosenColor = Colors.green;
    } else if (value > 15 && value <= 25) {
      chosenColor = Colors.yellow;
    } else if (value > 25) {
      chosenColor = Colors.red;
    }
    return `rgb(${chosenColor.red}, ${chosenColor.green},${chosenColor.blue},  ${chosenColor.opacity} )`;
  }

  return (
    <div className="heatmap-parent">
      <div className="heatmap-container">
        <HeatMap
          xLabels={array}
          yLabels={array}
          data={numbersCopy}
          cellStyle={(background, value, min, max) => ({
            background: `${chooseColor(value)} `,
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
