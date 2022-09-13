export const Colors = {
  purple: {
    red: 237,
    green: 2,
    blue: 245,
    opacity: 0.5,
  },
  green: {
    red: 2,
    green: 255,
    blue: 20,
    opacity: 0.75,
  },
  red: {
    red: 255,
    green: 0,
    blue: 0,
    opacity: 1,
  },
  yellow: {
    red: 242,
    green: 239,
    blue: 52,
    opacity: 1,
  },
  blue: {
    red: 57,
    green: 191,
    blue: 187,
    opacity: 0.3,
  },
  none: {
    red: 0,
    green: 0,
    blue: 0,
    opacity: 0,
  },
};

export function chooseColor(value) {
  switch (true) {
    case value <= 2:
      return returnCellColor(Colors.none);
    case value <= 5:
      return returnCellColor(Colors.blue);
    case value <= 10:
      return returnCellColor(Colors.purple);
    case value <= 15:
      return returnCellColor(Colors.green);
    case value <= 25:
      return returnCellColor(Colors.yellow);
    case value > 25:
      return returnCellColor(Colors.red);
    default:
      return returnCellColor(Colors.none);
  }
}

function returnCellColor({ red, green, blue, opacity }) {
  return `rgba(${red}, ${green},${blue},  ${opacity} )`;
}
