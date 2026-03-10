import { createApp } from '@backstage/app-defaults';
import { FlatRoutes } from '@backstage/core-app-api';
import { Route, BrowserRouter } from 'react-router-dom';
import { StylesProvider, createGenerateClassName } from '@material-ui/core/styles';
import './utils/makeStyles'; // Import the monkey patch
import { TemporalIntegrationPage } from './plugins/temporal-integration/index';
import { SplitScreenPage } from './components/SplitScreenPage';
import { SplitScreenLandingPage } from './components/SplitScreenLandingPage';
import SkillsManagement from './components/SkillsManagement';

// Import RAG AI component
import { RagAiPage } from './plugins/rag-ai';

// Import catalog plugin
import { CatalogEntityPage, CatalogIndexPage } from '@backstage/plugin-catalog';

const app = createApp({
  apis: [],
  plugins: [],
});

const App = app.createRoot(
  <StylesProvider 
    generateClassName={createGenerateClassName()}
    injectFirst={true}
  >
    <BrowserRouter>
      <FlatRoutes>
        <Route path="/" element={<SplitScreenLandingPage />} />
        <Route path="/catalog" element={<CatalogIndexPage />} />
        <Route path="/catalog/:namespace/:kind/:name" element={<CatalogEntityPage />} />
        <Route path="/temporal" element={<TemporalIntegrationPage />} />
        <Route path="/split-screen" element={<SplitScreenPage />} />
        <Route path="/skills" element={<SkillsManagement />} />
        <Route path="/rag-ai" element={<RagAiPage />} />
      </FlatRoutes>
    </BrowserRouter>
  </StylesProvider>
);

export default App;
