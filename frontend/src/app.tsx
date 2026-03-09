import { createApp } from '@backstage/app-defaults';
import { FlatRoutes } from '@backstage/core-app-api';
import { Route, BrowserRouter } from 'react-router-dom';
import { TemporalIntegrationPage } from './plugins/temporal-integration/index';

// Import catalog plugin
import { CatalogEntityPage, CatalogIndexPage } from '@backstage/plugin-catalog';

const app = createApp({
  apis: [],
  plugins: [
    // Add catalog plugin
    import('@backstage/plugin-catalog').then(m => m.plugin),
  ],
});

const App = app.createRoot(
  <BrowserRouter>
    <FlatRoutes>
      <Route path="/" element={<div><h1>Welcome to Backstage + Temporal Sandbox</h1><p>Navigate to <a href="/temporal">/temporal</a> for workflow management</p><p>Navigate to <a href="/catalog">/catalog</a> for software catalog</p></div>} />
      <Route path="/catalog" element={<CatalogIndexPage />} />
      <Route path="/catalog/:namespace/:kind/:name" element={<CatalogEntityPage />} />
      <Route path="/temporal" element={<TemporalIntegrationPage />} />
    </FlatRoutes>
  </BrowserRouter>
);

export default App;
