import useAxios from 'axios-hooks';
import config from '../config';
import { useMemo } from 'react';

function mirrorHeatmap(data) {
  const heatmapDim = data.length;
  const array = new Array(heatmapDim).fill('');

  const transpose = (matrix) => {
    const numbersCopy = structuredClone(matrix);
    for (let row = 0; row < numbersCopy.length; row++) {
      for (let column = 0; column < row; column++) {
        let temp = numbersCopy[row][column];
        numbersCopy[row][column] = numbersCopy[column][row];
        numbersCopy[column][row] = temp;
      }
    }
    return numbersCopy;
  };

  return { array, numbersCopy: transpose(data) };
}

const useHeatmap = () => {
  const [{ data, loading, error }] = useAxios(
    {
      method: 'get',
      url: `${config.apiBaseUrl}/heatmap`,
    },
    { useCache: false }
  );
  const heatmap = useMemo(() => data && mirrorHeatmap(data), [data]);

  return { loading, error, heatmap };
};

export default useHeatmap;
