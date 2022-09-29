import useAxios from 'axios-hooks';
import config from '../config';
import { useMemo } from 'react';

const useHeatmap = () => {
  const [{ data, loading, error }] = useAxios(
    {
      method: 'get',
      url: `${config.apiBaseUrl}/heatmap`,
    },
    { useCache: false }
  );
  const heatmap = useMemo(() => data && { array: new Array(data.heatmap.length).fill(''), numbersCopy: data.heatmap }, [data]);

  return { loading, error, heatmap };
};

export default useHeatmap;
