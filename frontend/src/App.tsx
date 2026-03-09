import React from 'react';
import { FlatRoutes } from '@backstage/core-app-api';
import { Route } from 'react-router-dom';
import { TemporalPage } from './plugins/temporal-integration/components/TemporalPage';

const app = (
  <FlatRoutes>
    <Route path="/temporal" element={<TemporalPage />} />
  </FlatRoutes>
);

export default app;
