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
    mock.onPost(`http://localhost:3000/reset`).reply(200);

    const result = await resetGame();

    expect(result.status).toBe(200);
  });

  it('should ignore backend errors (result not updated)', async () => {
    mock.onPost(`http://localhost:3000/reset`).reply(500);

    const result = await resetGame();

    expect(result.status).toBe(500);
  });
});
