import axios from 'axios';
import config from '../config';

export async function getHeatmapData() {

return await axios.get(`${config.apiBaseUrl}/stats`).catch((err) => {
    alert(err);
  });
}