export const TeamID = {
  Team_white: 1,
  Team_blue: 2,
};

export const ScoreChange = {
  Add_goal: 'add',
  Sub_goal: 'sub',
};

export class Goal {
  constructor(teamID, timestamp) {
    this.teamID = teamID;
    this.timestamp = timestamp;
  }
}
