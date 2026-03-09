import React from 'react';
import { HotKeys } from 'react-hotkeys';
import {
  createApp,
  AlertDisplay,
  OAuthRequestDialog,
  Progress,
  AppProvider,
  ConfigProvider,
  ErrorBoundary,
} from '@backstage/core-app-api';
import { ApiExplorer } from '@backstage/core-plugin-api';
import { ThemeProvider } from '@material-ui/core/styles';
import LightIcon from '@material-ui/icons/Brightness7';
import DarkIcon from '@material-ui/icons/Brightness4';

const app = createApp({
  apis: [
    // Add API providers here
  ],
  themes: [
    {
      id: 'light',
      title: 'Light Theme',
      variant: 'light',
      Provider: ThemeProvider,
      Icon: LightIcon,
    },
    {
      id: 'dark',
      title: 'Dark Theme',
      variant: 'dark',
      Provider: ThemeProvider,
      Icon: DarkIcon,
    },
  ],
});

interface RootProps {
  children: React.ReactNode;
}

const Root: React.FC<RootProps> = ({ children }) => (
  <ErrorBoundary>
    <ConfigProvider>
      <AppProvider>
        <HotKeys>
          <AlertDisplay />
          <OAuthRequestDialog />
          <Progress />
          {children}
          <ApiExplorer />
        </HotKeys>
      </AppProvider>
    </ConfigProvider>
  </ErrorBoundary>
);

export default app.createRoot(Root);
