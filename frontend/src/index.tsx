import React from 'react';
import ReactDOM from 'react-dom/client';

const rootElement = document.getElementById('root') as HTMLElement;
const root = ReactDOM.createRoot(rootElement);

root.render(
  <React.StrictMode>
    <div>
      <h1>Backstage + Temporal Sandbox</h1>
      <p>Loading...</p>
    </div>
  </React.StrictMode>,
);
