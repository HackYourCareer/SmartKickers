export const Colors = {
  red: {
    red: 202,
    green: 3,
    blue: 0,
    opacity: 0.75,
  },
  yellow: {
    red: 225,
    green: 101,
    blue: 25,
    opacity: 1,
  },
  green: {
    red: 202,
    green: 206,
    blue: 23,
    opacity: 0.75,
  },
  purple: {
    red: 56,
    green: 140,
    blue: 4,
    opacity: 0.5,
  },
  blue: {
    red: 4,
    green: 115,
    blue: 49,
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
