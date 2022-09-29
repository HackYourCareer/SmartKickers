import useAxios from 'axios-hooks';
import config from '../config';
import { useMemo } from 'react';

function mirrorHeatmap(data) {
  const heatmapDim = data.length;
  const array = new Array(heatmapDim).fill('');

  return { array, numbersCopy: data };
}

const useHeatmap = () => {
  const [{ data, loading, error }] = useAxios(
    {
      method: 'get',
      url: `${config.apiBaseUrl}/heatmap`,
    },
    { useCache: false }
  );
  const heatmap = useMemo(() => data && mirrorHeatmap(data.heatmap), [data]);

  return { loading, error, heatmap };
};

export default useHeatmap;
