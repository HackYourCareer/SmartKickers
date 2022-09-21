import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import Router from './router/Router';
import { initLibs } from './appConfig';
import GameDataContextProvider from './contexts/GameDataContext.js';

initLibs();
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <GameDataContextProvider>
      <Router />
    </GameDataContextProvider>
  </React.StrictMode>
);
