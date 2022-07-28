import axios from 'axios';

const serverUrl = 'http://localhost:3006';

export async function addPoint(teamId) {
  try {
    await axios.post(`${serverUrl}/addPoint/${teamId}`);
  } catch (error) {}
}

export async function subPoint(teamId) {
  try {
    await axios.post(`${serverUrl}/subPoint/${teamId}`);
  } catch (error) {}
}
