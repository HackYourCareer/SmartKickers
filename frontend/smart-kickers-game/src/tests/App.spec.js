import { render } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import WS from 'jest-websocket-mock';
import { getElementWhichContain } from './helpers';
import App from '../App';

let ws;
describe('<App />', () => {
  beforeEach(() => {
    ws = new WS('ws://localhost:3006/csc');
  });

  afterEach(() => {
    WS.clean();
  });

  it('should render correctly', async () => {
    render(<App />);

    expect(getElementWhichContain('Blue:')).toBeDefined();
    expect(getElementWhichContain('White:')).toBeDefined();
  });

  it('should update score on score message', async () => {
    render(<App />);
    await ws.connected;

    ws.send(JSON.stringify({ blueScore: 10, whiteScore: 14 }));

    expect(getElementWhichContain('Blue:')).toHaveTextContent('10');
    expect(getElementWhichContain('White:')).toHaveTextContent('14');
  });
});
