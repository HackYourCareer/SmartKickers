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
    console.log('dipa');
    const heatmapDim = data.Heatmap.length;
    const array = new Array(heatmapDim).fill(0).map(() => '');
    let numbersCopy = JSON.parse(JSON.stringify(data.Heatmap));

    for (let i = 0; i < heatmapDim; i++)
      for (let j = 0; j <= i; j++) {
        numbersCopy[i][j] = numbersCopy[i][j] + numbersCopy[j][i] - numbersCopy[j][i];
        numbersCopy[j][i] = numbersCopy[i][j];
      }

    return { array, numbersCopy };
  }

  return [{ loading, error, heatmap }];
};

export default useHeatmap;
