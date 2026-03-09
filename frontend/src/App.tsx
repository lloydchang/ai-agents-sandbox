import React from 'react';
import { Route } from 'react-router-dom';
import { FlatRoutes } from '@backstage/core-app-api';
import { ApiDocsPage } from '@backstage/plugin-api-docs';
import { CatalogEntityPage, CatalogIndexPage } from '@backstage/plugin-catalog';
import { CatalogImportPage } from '@backstage/plugin-catalog-import';
import { ScaffolderPage } from '@backstage/plugin-scaffolder';
import { SearchPage } from '@backstage/plugin-search';
import { TechRadarPage } from '@backstage/plugin-tech-radar';
import { TechDocsIndexPage, TechDocsReaderPage } from '@backstage/plugin-techdocs';
import { UserSettingsPage } from '@backstage/plugin-user-settings';
import { TemporalIntegrationPage } from './plugins/temporal-integration';

import Root from './components/Root';

const app = (
  <Root>
    <FlatRoutes>
      <Route path="/" element={<CatalogIndexPage />} />
      <Route path="/catalog/:namespace/:kind/:name" element={<CatalogEntityPage />} />
      <Route path="/catalog-import" element={<CatalogImportPage />} />
      <Route path="/docs" element={<TechDocsIndexPage />} />
      <Route path="/docs/:namespace/:kind/:name/*" element={<TechDocsReaderPage />} />
      <Route path="/create" element={<ScaffolderPage />} />
      <Route path="/api-docs" element={<ApiDocsPage />} />
      <Route path="/tech-radar" element={<TechRadarPage />} />
      <Route path="/search" element={<SearchPage />} />
      <Route path="/settings" element={<UserSettingsPage />} />
      <Route path="/temporal" element={<TemporalIntegrationPage />} />
    </FlatRoutes>
  </Root>
);

export default app;
