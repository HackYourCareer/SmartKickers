import React from 'react';
import HeatMap from 'react-heatmap-grid';
import './Heatmap.css';
import { Colors } from './Colors';
import { useRef, useEffect } from 'react';

function Heatmap({ heatmap }) {
  const ref = useRef(null);
  const heatmapDim = heatmap.length;
  const array = new Array(heatmapDim).fill(0).map(() => '');
  let numbersCopy = JSON.parse(JSON.stringify(heatmap));

  useEffect(() => {
    const cells = document.querySelectorAll('.heatmap-container > div > div > div > div');
    cells.forEach((element) => {
      element.classList.add('heatmap-cell');
    });
    const leftCells = document.querySelectorAll('.heatmap-container > div > div > div > div:first-child');
    leftCells.forEach((element) => {
      element.classList.add('heatmap-left-cells');
    });
  }, []);

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
      <div className="heatmap-container" ref={ref}>
        <HeatMap
          xLabels={array}
          yLabels={array}
          data={numbersCopy}
          cellStyle={(background, value, min, max) => ({
            background: `${chooseColor(value)} `,
            fontSize: '0px',
            color: '#444',
            width: '3.28px',
            height: '2px',
          })}
          cellRender={(value) => value && <div>{value}</div>}
        />
      </div>
    </div>
  );
}
export default Heatmap;
