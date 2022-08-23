import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';
import config from '../config';
import { pointsManipulation } from './goalManipulation.js';
import { TeamID, ScoreChange } from '../constants/score.js';

describe('pointsManipulation API caller', () => {
  let mock;

  beforeAll(() => {
    mock = new MockAdapter(axios);
  });

  afterEach(() => {
    mock.reset();
  });

  it.each([
    [TeamID.Team_blue, ScoreChange.Add_goal, 200],
    [TeamID.Team_blue, ScoreChange.Sub_goal, 200],
    [TeamID.Team_white, ScoreChange.Add_goal, 200],
    [TeamID.Team_white, ScoreChange.Sub_goal, 200],
  ])(`should return proper result when passed arguments are: %i, %i`, async (teamID, action, expectedResult) => {
    mock.onPost(`${config.apiBaseUrl}/goal?action=${action}&team=${teamID}`).reply(expectedResult);
    mock.onPost(`${config.apiBaseUrl}/goal?action=${ScoreChange.Add_goal}&team=${TeamID.Team_white}`).reply(200);
    mock.onPost(`${config.apiBaseUrl}/goal?action=${ScoreChange.Sub_goal}&team=${TeamID.Team_blue}`).reply(200);
    mock.onPost(`${config.apiBaseUrl}/goal?action=${ScoreChange.Sub_goal}&team=${TeamID.Team_white}`).reply(200);

    const result = await pointsManipulation(teamID, action);
    expect(result.status).toBe(expectedResult);
    expect(result.error).not.toBeDefined();
  });

  it.each([
    [TeamID.Team_blue, ScoreChange.Add_goal],
    [TeamID.Team_blue, ScoreChange.Sub_goal],
    [TeamID.Team_white, ScoreChange.Add_goal],
    [TeamID.Team_white, ScoreChange.Sub_goal],
  ])(`should return error with correct status code if received error from backend when passed arguments are: %i, %i`, async (teamID, action) => {
    mock.onPost(`${config.apiBaseUrl}/goal?action=${ScoreChange.Add_goal}&team=${TeamID.Team_blue}`).reply(500);
    mock.onPost(`${config.apiBaseUrl}/goal?action=${ScoreChange.Add_goal}&team=${TeamID.Team_white}`).reply(500);
    mock.onPost(`${config.apiBaseUrl}/goal?action=${ScoreChange.Sub_goal}&team=${TeamID.Team_blue}`).reply(500);
    mock.onPost(`${config.apiBaseUrl}/goal?action=${ScoreChange.Sub_goal}&team=${TeamID.Team_white}`).reply(500);

    const result = await pointsManipulation(teamID, action);
    expect(result.status).toBe(500);
    expect(result.error).toBeDefined();
  });
});
