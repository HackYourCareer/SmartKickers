import axios from 'axios';
import config from '../config';

export async function getHeatmapData() {
  return axios.get(`${config.apiBaseUrl}/stats`).catch((err) => {
    alert(err);
  });
}
