import axios from 'axios';

export async function resetGame() {
  try {
    let result = await axios.post(`http://localhost:3000/reset`);

    return {
      status: result.status,
      data: result.data,
    };
  } catch (e) {
    return {
      error: e.response.data,
      status: e.response.status,
    };
  }
}
