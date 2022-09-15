import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { initLibs } from './appConfig';
import GameDataContextProvider from './contexts/GameDataContext.js';

initLibs();
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <GameDataContextProvider>
      <App />
    </GameDataContextProvider>
  </React.StrictMode>
);
