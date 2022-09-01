import axios from 'axios';
import config from '../config';

export async function getStatistics() {
  try {
    const result = await axios.get(`${config.apiBaseUrl}/stats`);

    return {
      data: result.data,
    };
  } catch (e) {
    return {
      error: e,
      status: e.response.status,
    };
  }
}
