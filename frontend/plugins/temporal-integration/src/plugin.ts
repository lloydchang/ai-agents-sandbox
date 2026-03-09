import { createPlugin, createRoutableExtension } from '@backstage/core-plugin-api';

export const temporalIntegrationPlugin = createPlugin({
  id: 'temporal-integration',
});

export const TemporalIntegrationPage = temporalIntegrationPlugin.provide(
  createRoutableExtension({
    name: 'TemporalIntegrationPage',
    component: () => import('./components/TemporalPage').then(m => m.TemporalPage),
  }),
);
