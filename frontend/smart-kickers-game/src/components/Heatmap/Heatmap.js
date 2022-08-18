import React from 'react'
import HeatMap from "react-heatmap-grid";



function Heatmap ({heatMapTable}) {
  const xLabels = new Array(Math.sqrt(heatMapTable.length)).fill(0).map((_, i) => ``);

  // Display only even labels
  const xLabelsVisibility = new Array(24)
    .fill(0)
    .map((_, i) => (i % 2 === 0 ? true : false));
  
    let i = 0
  const yLabels = new Array(Math.sqrt(heatMapTable.length)).fill(0).map((_, i) => ``);
  const data = new Array(yLabels.length)
    .fill(0)
    .map(() =>
      new Array(xLabels.length).fill(0).map(() => Math.floor(i++))
    );
    console.log(heatMapTable)

  return (
    <div style={{ fontSize: "13px" }}>
      <HeatMap
        xLabels={xLabels}
        yLabels={yLabels}
        xLabelsLocation={"bottom"}
        xLabelsVisibility={xLabelsVisibility}
        xLabelWidth={60}
        data={heatMapTable}
        squares
        height={45}
        onClick={(x, y) => alert(`Clicked ${x}, ${y}`)}
        cellStyle={(background, value, min, max, data, x, y) => ({
          background: `rgb(0, 151, 230, ${1 - (max - value) / (max - min)})`,
          fontSize: "10px",
          color: "#444",
          width: "12px",
          height: "122px"
        })}
        cellRender={value => value && <div>{value}</div>}
      />
    </div>
  );
}
export default Heatmap