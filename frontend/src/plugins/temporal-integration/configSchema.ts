import { ConfigSchema } from '@backstage/config';

export const configSchema: ConfigSchema = {
  configSchema: {
    temporal: {
      type: 'object',
      properties: {
        backendUrl: {
          type: 'string',
          description: 'URL of the Temporal backend service',
        },
      },
    },
  },
};
