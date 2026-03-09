import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

const app = (
  <BrowserRouter>
    <Routes>
      <Route path="/" element={
        <div>
          <h1>Welcome to Backstage + Temporal Sandbox</h1>
          <p>Basic routing working - Backstage integration next step</p>
        </div>
      } />
      <Route path="/temporal" element={
        <div>
          <h1>Temporal Integration</h1>
          <p>Temporal workflow management page coming soon...</p>
        </div>
      } />
    </Routes>
  </BrowserRouter>
);

export default app;
