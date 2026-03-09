import React from 'react';
import { Route } from 'react-router-dom';
import { FlatRoutes } from '@backstage/core-app-api';

import Root from './components/Root';

const app = (
  <Root>
    <FlatRoutes>
      <Route path="/" element={<div><h1>Welcome to Backstage + Temporal Sandbox</h1></div>} />
    </FlatRoutes>
  </Root>
);

export default app;
