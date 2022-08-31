import { getHeatmapData } from '../apis/heatmap.js';
import useAxios from 'axios-hooks';
import config from '../config';

const useHeatmap = () => {
  const [{ data, loading, error }] = useAxios({
    method: 'get',
    url: `${config.apiBaseUrl}/stats`,
  });
  return [{ data, loading, error }];
};

export default useHeatmap;
