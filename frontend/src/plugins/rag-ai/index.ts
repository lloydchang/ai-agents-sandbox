import {
  createPlugin,
  createRoutableExtension,
} from '@backstage/core-plugin-api';

import { rootRouteRef } from './routes';

export const ragAiPlugin = createPlugin({
  id: 'rag-ai',
  routes: {
    root: rootRouteRef,
  },
});

export const RagAiPage = ragAiPlugin.provide(
  createRoutableExtension({
    name: 'RagAiPage',
    component: () =>
      import('./components/RagAiPage').then(m => m.RagAiPage),
    mountPoint: rootRouteRef,
  }),
);
