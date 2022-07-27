import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';
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
    const gameId = 'anyGameId';
    mock.onPost(`http://localhost:3006/reset/${gameId}`).reply(200);

    const result = await resetGame(gameId);

    expect(result.status).toBe(200);
  });
});
