import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';
import config from '../config';
import { resetGame } from './Game';

describe('resetGame API caller', () => {
  let mock;

  beforeAll(() => {
    mock = new MockAdapter(axios);
  });

  afterEach(() => {
    mock.reset();
  });

  it('should send reset game request', async () => {
    mock.onPost(`${config.apiBaseUrl}/reset`).reply(200);

    const result = await resetGame();

    expect(result.status).toBe(200);
    expect(result.error).not.toBeDefined();
  });

  it('should return error with correct status code if received error from backend', async () => {
    mock.onPost(`${config.apiBaseUrl}/reset`).reply(500);

    const result = await resetGame();

    expect(result.error).toBeDefined();
    expect(result.status).toBe(500);
  });
});
