import { createApp } from '@backstage/app-defaults';
import { FlatRoutes } from '@backstage/core-app-api';
import { Route, BrowserRouter } from 'react-router-dom';
import { TemporalPage } from './plugins/temporal-integration/components/TemporalPage';

const app = createApp({
  apis: [],
  components: {
    Router: props => <BrowserRouter>{props.children}</BrowserRouter>,
  },
});

const App = app.createRoot(
  <FlatRoutes>
    <Route path="/" element={<div><h1>Welcome to Backstage + Temporal Sandbox</h1><p>Navigate to <a href="/temporal">/temporal</a> for workflow management</p></div>} />
    <Route path="/temporal" element={<TemporalPage />} />
  </FlatRoutes>
);

export default App;
