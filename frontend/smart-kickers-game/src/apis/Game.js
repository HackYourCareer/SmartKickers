import axios from 'axios';
import config from '../config';

export async function resetGame() {
  try {
    let result = await axios.put(`${config.apiBaseUrl}/reset`);

    return {
      status: result.status,
    };
  } catch (e) {
    return {
      error: e,
      status: e.response.status,
    };
  }
}
