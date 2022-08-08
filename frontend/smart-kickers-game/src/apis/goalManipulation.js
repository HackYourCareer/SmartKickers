import axios from 'axios';
import config from '../config';

export function pointsManipulation(teamID, action) {
  axios.put(`${config.apiBaseUrl}/goal?action=${action}&team=${teamID}`).catch((err) => {
    alert(err);
  });
}
