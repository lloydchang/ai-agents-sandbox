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
      <Route path="/" element={<div><h1>Welcome to Backstage + Temporal Sandbox</h1><p>Navigate to <a href="/temporal">/temporal</a> for workflow management</p></div>} />
      <Route path="/temporal" element={<TemporalIntegrationPage />} />
    </FlatRoutes>
  </BrowserRouter>
);

export default App;
