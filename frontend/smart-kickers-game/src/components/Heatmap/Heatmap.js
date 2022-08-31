import React from 'react';
import HeatMap from 'react-heatmap-grid';
import './Heatmap.css';
import { Colors } from './Colors';
import useHeatmap from '../../hooks/useHeatmap';

function Heatmap() {
  const [{ loading, error, heatmap }] = useHeatmap();
  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error</div>;

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
          xLabels={heatmap.array}
          yLabels={heatmap.array}
          data={heatmap.numbersCopy}
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
