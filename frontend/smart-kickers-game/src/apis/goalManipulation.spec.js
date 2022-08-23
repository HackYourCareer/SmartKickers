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

    const result = await pointsManipulation(teamID, action);
    expect(result.status).toBe(expectedResult);
    expect(result.error).not.toBeDefined();
  });

  it.each([
    [TeamID.Team_blue, ScoreChange.Add_goal, 500],
    [TeamID.Team_blue, ScoreChange.Sub_goal, 500],
    [TeamID.Team_white, ScoreChange.Add_goal, 500],
    [TeamID.Team_white, ScoreChange.Sub_goal, 500],
  ])(
    `should return error with correct status code if received error from backend when passed arguments are: %i, %i`,
    async (teamID, action, expectedResult) => {
      mock.onPost(`${config.apiBaseUrl}/goal?action=${action}&team=${teamID}`).reply(expectedResult);

      const result = await pointsManipulation(teamID, action);
      expect(result.status).toBe(500);
      expect(result.error).toBeDefined();
    }
  );
});
