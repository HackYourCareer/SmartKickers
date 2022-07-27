import axios from 'axios';

export function resetGame(gameId) {
  return axios.post(`http://localhost:3006/reset/${gameId}`);
}
