import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";
import WS from "jest-websocket-mock";
import App from "./App";

let ws;
describe("<App />", () => {
  beforeEach(() => {
    ws = new WS("ws://localhost:3000");
  });

  afterEach(() => {
    WS.clean();
  });

  it("should render correctly", () => {
    const { container } = render(<App />);

    expect(screen.getByTestId("blue-team-score")).toHaveTextContent("0");
  });

  it("should update score on blue goal message", async () => {
    const { container } = render(<App />);
    await ws.connected;

    ws.send(JSON.stringify({ type: "blueGoal" }));

    expect(screen.getByTestId("blue-team-score")).toHaveTextContent("Blue: 1");
  });

  it("should update score on white goal message", async () => {
    const { container } = render(<App />);
    await ws.connected;

    ws.send(JSON.stringify({ type: "whiteGoal" }));

    expect(screen.getByTestId("blue-team-score")).toHaveTextContent("White: 1");
  });
});
