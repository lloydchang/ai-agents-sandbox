import React from 'react';
import { Grid, Paper, Typography, Button, Box } from '@mui/material';
import {
  ViewColumn,
  Timeline,
  AccountTree,
  Speed,
} from '@mui/icons-material';

export const SplitScreenLandingPage: React.FC = () => {
  const features = [
    {
      icon: <ViewColumn fontSize="large" />,
      title: 'Split Screen View',
      description: 'View Backstage and Temporal side-by-side for seamless workflow management',
      path: '/split-screen'
    },
    {
      icon: <Timeline fontSize="large" />,
      title: 'Temporal Workflows',
      description: 'Manage and monitor your Temporal workflows with advanced AI agent orchestration',
      path: '/temporal'
    },
    {
      icon: <AccountTree fontSize="large" />,
      title: 'Backstage Catalog',
      description: 'Explore your software catalog and manage components, APIs, and infrastructure',
      path: '/catalog'
    },
    {
      icon: <Speed fontSize="large" />,
      title: 'AI Agent Integration',
      description: 'Leverage AI agents for automated compliance, security, and cost optimization',
      path: '/split-screen'
    }
  ];

  return (
    <Box sx={{ flexGrow: 1, p: 3 }}>
      <Typography variant="h3" gutterBottom align="center">
        Backstage + Temporal AI Agent Sandbox
      </Typography>
      
      <Typography variant="h6" color="textSecondary" paragraph align="center">
        Experiment with AI-powered workflows in a safe, local environment
      </Typography>

      <Grid container spacing={4} sx={{ mt: 2 }}>
        {features.map((feature, index) => (
          <Grid item xs={12} sm={6} md={3} key={index}>
            <Paper
              sx={{
                p: 3,
                height: '100%',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                textAlign: 'center',
                cursor: 'pointer',
                transition: 'transform 0.2s, box-shadow 0.2s',
                '&:hover': {
                  transform: 'translateY(-4px)',
                  boxShadow: 4,
                },
              }}
              onClick={() => window.location.href = feature.path}
            >
              <Box sx={{ color: 'primary.main', mb: 2 }}>
                {feature.icon}
              </Box>
              
              <Typography variant="h6" gutterBottom>
                {feature.title}
              </Typography>
              
              <Typography variant="body2" color="textSecondary" paragraph>
                {feature.description}
              </Typography>
              
              <Button
                variant="outlined"
                size="small"
                sx={{ mt: 'auto' }}
              >
                Launch
              </Button>
            </Paper>
          </Grid>
        ))}
      </Grid>

      <Box sx={{ mt: 6, p: 3, bgcolor: 'background.paper', borderRadius: 2 }}>
        <Typography variant="h5" gutterBottom>
          Quick Start Guide
        </Typography>
        
        <Typography variant="body1" paragraph>
          1. <strong>Split Screen</strong>: Use the split-screen view to work with both Backstage and Temporal simultaneously
        </Typography>
        
        <Typography variant="body1" paragraph>
          2. <strong>Mobile Friendly</strong>: On mobile devices, toggle between Backstage and Temporal views
        </Typography>
        
        <Typography variant="body1" paragraph>
          3. <strong>AI Workflows</strong>: Start AI agent workflows for compliance, security, and cost optimization scenarios
        </Typography>
        
        <Typography variant="body1">
          4. <strong>Real-time Monitoring</strong>: Monitor workflow execution and agent performance in real-time
        </Typography>
      </Box>
    </Box>
  );
};
