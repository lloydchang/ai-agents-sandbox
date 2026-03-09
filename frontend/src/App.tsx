import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Button, Paper, Typography, Container, Grid } from '@material-ui/core';
import { TemporalPage } from './components/TemporalPage';

const app = (
  <BrowserRouter>
    <Container maxWidth="lg">
      <Routes>
        <Route path="/" element={
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <Paper style={{ padding: 24 }}>
                <Typography variant="h4" component="h1" gutterBottom>
                  Welcome to Backstage + Temporal Sandbox
                </Typography>
                <Typography variant="body1" gutterBottom>
                  Basic routing working - Backstage integration next step
                </Typography>
                <Button 
                  variant="contained" 
                  color="primary"
                  onClick={() => window.location.href = '/temporal'}
                >
                  Go to Temporal Integration
                </Button>
              </Paper>
            </Grid>
          </Grid>
        } />
        <Route path="/temporal" element={
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <Paper style={{ padding: 24 }}>
                <Typography variant="h4" component="h1" gutterBottom>
                  Temporal Integration
                </Typography>
                <Typography variant="body1" gutterBottom>
                  Manage your Temporal workflows below
                </Typography>
                <Button 
                  variant="outlined"
                  onClick={() => window.location.href = '/'}
                  style={{ marginBottom: 24 }}
                >
                  Back to Home
                </Button>
                <TemporalPage />
              </Paper>
            </Grid>
          </Grid>
        } />
      </Routes>
    </Container>
  </BrowserRouter>
);

export default app;
