import axios from 'axios';
import config from '../config';

export async function pointsManipulation(teamID, action) {
  try {
    await axios.post(`${config.apiBaseUrl}/goal?action=${action}&team=${teamID}`);
  } catch(error) {
    alert(error);
  }
}