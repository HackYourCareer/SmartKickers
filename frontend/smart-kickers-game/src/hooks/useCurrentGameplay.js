import { useEffect, useRef, useState } from 'react';
import { Goal, TeamID } from '../constants/score.js';
import { useStopwatch } from 'react-timer-hook';
import { useGameDataContext } from '../contexts/GameDataContext';
import { resetGame } from '../apis/resetGame';
import { useNavigate } from 'react-router-dom';

export default function useCurrentGameplay() {
  const navigate = useNavigate();
  const [isVisible, setIsVisible] = useState(false);

  const { blueScore, whiteScore, setFinalScores, goalsArray, setGoalsArray, setIsGameEnded, setIsGameStarted } = useGameDataContext();

  const { seconds, minutes, isRunning, start, pause, reset } = useStopwatch({ autoStart: false });
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
    if (!isRunning) return;
    if (prevBlueScore > blueScore) {
      const indexOfLastBlue = goalsArray.indexOf(goalsArray.findLast((e) => e.teamID === TeamID.Team_blue));
      goalsArray.splice(indexOfLastBlue, 1);
    } else {
      goalsArray.push(new Goal(TeamID.Team_blue, 'time: ' + minutes.toString().padStart(2, '0') + ':' + seconds.toString().padStart(2, '0')));
    }
  }, [blueScore]);
  useEffect(() => {
    if (!isRunning) return;
    if (prevWhiteScore > whiteScore) {
      const indexOfLastWhite = goalsArray.indexOf(goalsArray.findLast((e) => e.teamID === TeamID.Team_white));
      goalsArray.splice(indexOfLastWhite, 1);
    } else {
      goalsArray.push(new Goal(TeamID.Team_white, 'time: ' + minutes.toString().padStart(2, '0') + ':' + seconds.toString().padStart(2, '0')));
    }
  }, [whiteScore]);

  const handleStartGame = () => {
    navigate('/game');
    resetGoalsArray();
    handleResetGame();
    setIsVisible(true);
    start();
    setIsGameStarted(true);
  };

  const resetGoalsArray = () => {
    setGoalsArray([]);
  };

  const handleResetGame = () => {
    resetGame().then((data) => {
      if (data.error) alert(data.error);
    });
    reset();
    start();
    resetGoalsArray();
  };
  const handleEndGame = () => {
    setFinalScores({ blueScore: blueScore, whiteScore: whiteScore });
    pause();
    setIsGameEnded(true);
  };

  return { isVisible, handleStartGame, handleResetGame, handleEndGame };
}
