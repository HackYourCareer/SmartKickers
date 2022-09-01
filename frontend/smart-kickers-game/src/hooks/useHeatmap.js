import useAxios from 'axios-hooks';
import config from '../config';
import { useMemo } from 'react';

const useHeatmap = () => {
  const [{ data, loading, error }] = useAxios({
    method: 'get',
    url: `${config.apiBaseUrl}/stats`,
  });

  const heatmap = useMemo(() => (data?.Heatmap ? mirrorHeatmap() : []), [data]);

  function mirrorHeatmap() {
    const heatmapDim = data.Heatmap.length;
    const array = new Array(heatmapDim).fill(0).map(() => '');
    let numbersCopy = JSON.parse(JSON.stringify(data.Heatmap));

    // transpone
    const transpose = (matrix) => {
      for (let row = 0; row < matrix.length; row++) {
        for (let column = 0; column < row; column++) {
          let temp = matrix[row][column];
          matrix[row][column] = matrix[column][row];
          matrix[column][row] = temp;
        }
      }
      return matrix;
    };

    // transpose(numbersCopy);
    // reverse
    // const reverse = (matrix) => matrix.map((row) => row.reverse());
    // rotate
    // const rotate = (matrix) => transpose(reverse(matrix));

    // rotate(numbersCopy);

    return { array, numbersCopy };
  }

  return [{ loading, error, heatmap }];
};

export default useHeatmap;
