import axios from 'axios';
import config from '../config';

export async function getStatistics() {
  try {
    const result = await axios.get(`${config.apiBaseUrl}/stats`);

    return result.data;
  } catch (e) {
    return {
      error: e,
    };
  }
}
