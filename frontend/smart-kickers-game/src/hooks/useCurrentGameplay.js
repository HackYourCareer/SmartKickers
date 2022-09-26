import { useEffect, useRef, useState } from 'react';
import { Goal, TeamID } from '../constants/score.js';
import { useGameDataContext } from '../contexts/GameDataContext';

export default function useCurrentGameplay() {
  const { blueScore, whiteScore, setFinalScores, goalsArray, minutes, seconds, reset, start, handleResetGame, handleStartGame, setIsGameEnded } =
    useGameDataContext();

  const ScorePrevious = (value) => {
    const ref = useRef();
    useEffect(() => {
      ref.current = value;
    });
    return ref.current;
  };
  const prevBlueScore = ScorePrevious(blueScore);
  const prevWhiteScore = ScorePrevious(whiteScore);

  useEffect(() => {
    if (prevBlueScore !== undefined) {
      if (prevBlueScore > blueScore) {
        const indexOfLastBlue = goalsArray.indexOf(goalsArray.findLast((e) => e.teamID === TeamID.Team_blue));
        goalsArray.splice(indexOfLastBlue, 1);
      } else {
        goalsArray.push(new Goal(TeamID.Team_blue, 'time: ' + minutes.toString().padStart(2, '0') + ':' + seconds.toString().padStart(2, '0')));
      }
    }
  }, [blueScore]);

  useEffect(() => {
    if (prevWhiteScore !== undefined) {
      if (prevWhiteScore > whiteScore) {
        const indexOfLastWhite = goalsArray.indexOf(goalsArray.findLast((e) => e.teamID === TeamID.Team_white));
        goalsArray.splice(indexOfLastWhite, 1);
      } else {
        goalsArray.push(new Goal(TeamID.Team_white, 'time: ' + minutes.toString().padStart(2, '0') + ':' + seconds.toString().padStart(2, '0')));
      }
    }
  }, [whiteScore]);

  const handleEndGame = () => {
    setFinalScores({ blueScore: blueScore, whiteScore: whiteScore });
    setIsGameEnded(true);
  };

  return { handleStartGame, handleResetGame, handleEndGame, start, reset };
}
