export const Colors = {
  green: {
    red: 150,
    green: 230,
    blue: 50,
    opacity: 50,
  },
  yellow: {
    red: 240,
    green: 230,
    blue: 20,
    opacity: 50,
  },
  orangeLight: {
    red: 240,
    green: 180,
    blue: 20,
    opacity: 50,
  },
  orangeDark: {
    red: 240,
    green: 130,
    blue: 20,
    opacity: 50,
  },
  redLight: {
    red: 240,
    green: 80,
    blue: 20,
    opacity: 50,
  },
  redDark: {
    red: 240,
    green: 30,
    blue: 20,
    opacity: 50,
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
      return returnCellColor(Colors.green);
    case value <= 10:
      return returnCellColor(Colors.yellow);
    case value <= 15:
      return returnCellColor(Colors.orangeLight);
    case value <= 25:
      return returnCellColor(Colors.orangeDark);
    case value <= 30:
      return returnCellColor(Colors.redLight);
    case value > 30:
      return returnCellColor(Colors.redDark);
    default:
      return returnCellColor(Colors.none);
  }
}

function returnCellColor({ red, green, blue, opacity }) {
  return `rgba(${red}, ${green},${blue},  ${opacity} )`;
}
