import './index.css';

import React from 'react';
import ReactDOM from 'react-dom/client';
import app from './App';

const rootElement = document.getElementById('root') as HTMLElement;
const root = ReactDOM.createRoot(rootElement);

root.render(
  <React.StrictMode>
    {app}
  </React.StrictMode>
);
