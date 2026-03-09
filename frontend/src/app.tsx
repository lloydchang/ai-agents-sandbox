import { createApp } from '@backstage/app-defaults';
import { FlatRoutes } from '@backstage/core-app-api';
import { Route, BrowserRouter } from 'react-router-dom';
import { TemporalIntegrationPage } from './plugins/temporal-integration/index';

const app = createApp({
  apis: [],
  plugins: [],
});

const App = app.createRoot(
  <BrowserRouter>
    <FlatRoutes>
      <Route path="/" element={<div><h1>Welcome to Backstage + Temporal Sandbox</h1><p>Navigate to <a href="/temporal">/temporal</a> for workflow management</p><p>Catalog integration configured - see catalog-info.yaml</p></div>} />
      <Route path="/catalog" element={<div><h1>Software Catalog</h1><p>Catalog entities are configured in catalog-info.yaml</p><p>Components: backstage-temporal, backstage-temporal-frontend</p><p>System: backstage-temporal-system</p><p>APIs: temporal-workflow-api, backstage-temporal-api</p></div>} />
      <Route path="/temporal" element={<TemporalIntegrationPage />} />
    </FlatRoutes>
  </BrowserRouter>
);

export default App;
