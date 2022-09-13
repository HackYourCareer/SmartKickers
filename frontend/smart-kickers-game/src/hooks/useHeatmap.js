import useAxios from 'axios-hooks';
import config from '../config';
import { useMemo } from 'react';

function mirrorHeatmap(data) {
  const heatmapDim = data.length;
  const array = new Array(heatmapDim).fill('');
  const numbersCopy = JSON.parse(JSON.stringify(data));

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

  transpose(numbersCopy);

  return { array, numbersCopy };
}

const useHeatmap = () => {
  const [{ data, loading, error }] = useAxios(
    {
      method: 'get',
      url: `${config.apiBaseUrl}/heatmap`,
    },
    { useCache: false }
  );
  console.log(data);

  const heatmap = useMemo(() => data && mirrorHeatmap(data), [data]);

  return { loading, error, heatmap };
};

export default useHeatmap;
