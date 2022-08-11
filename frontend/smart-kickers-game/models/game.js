class Game {
  constructor(blueScore, whiteScore, blueAdded, whiteAdded) {
    this.blueScore = blueScore;
    this.whiteScore = whiteScore;
    this.whiteAdded = whiteAdded;
    this.blueAdded = blueAdded;
  }
  get blueScore() {
    return this.blueScore;
  }
  set blueScore(blueScore) {
    this.blueScore = blueScore;
  }
  get whiteScore() {
    return this.whiteScore;
  }
  set whiteScore(whiteScore) {
    this.whiteScore = whiteScore;
  }
  get blueAdded() {
    return this.blueAdded;
  }
  set blueAdded(blueAdded) {
    this.blueAdded = blueAdded;
  }
  get whiteAdded() {
    return this.whiteAdded;
  }
  set whiteAdded(whiteAdded) {
    this.whiteAdded = whiteAdded;
  }
}
