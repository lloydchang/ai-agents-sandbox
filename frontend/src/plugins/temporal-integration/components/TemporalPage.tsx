import React, { useState } from 'react';
import { Button, Paper, Typography, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from '@material-ui/core';

interface WorkflowStatus {
  id: string;
  status: string;
  createdAt: string;
}

export const TemporalPage = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string>();
  const [workflows, setWorkflows] = useState<WorkflowStatus[]>([]);
  const backendUrl = 'http://localhost:8081';

  const startWorkflow = async () => {
    setLoading(true);
    setError(undefined);
    
    try {
      const response = await fetch(`${backendUrl}/workflow/start`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      
      if (!response.ok) {
        throw new Error(`Failed to start workflow: ${response.statusText}`);
      }
      
      const workflowId = await response.text();
      const newWorkflow: WorkflowStatus = {
        id: workflowId,
        status: 'RUNNING',
        createdAt: new Date().toISOString(),
      };
      
      setWorkflows(prev => [newWorkflow, ...prev]);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
    } finally {
      setLoading(false);
    }
  };

  const checkWorkflowStatus = async (workflowId: string) => {
    try {
      const response = await fetch(`${backendUrl}/workflow/status?id=${workflowId}`);
      
      if (!response.ok) {
        throw new Error(`Failed to get workflow status: ${response.statusText}`);
      }
      
      const status = await response.text();
      
      setWorkflows(prev => 
        prev.map(w => 
          w.id === workflowId ? { ...w, status } : w
        )
      );
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
    }
  };

  if (error) {
    return (
      <Paper style={{ padding: 24 }}>
        <Typography variant="h6" color="error">
          Error: {error}
        </Typography>
        <Button onClick={() => setError(undefined)} variant="contained">
          Clear Error
        </Button>
      </Paper>
    );
  }

  return (
    <div>
      <Paper style={{ padding: 24, marginBottom: 24 }}>
        <Typography variant="h6" gutterBottom>
          Workflow Management
        </Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={startWorkflow}
          disabled={loading}
          style={{ marginRight: 16 }}
        >
          {loading ? 'Starting...' : 'Start HelloBackstage Workflow'}
        </Button>
        <Button
          variant="outlined"
          onClick={() => workflows.forEach(w => checkWorkflowStatus(w.id))}
        >
          Refresh All Status
        </Button>
      </Paper>
      
      <Paper style={{ padding: 24 }}>
        <Typography variant="h6" gutterBottom>
          Workflow Status
        </Typography>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Workflow ID</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Created At</TableCell>
                <TableCell>Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {workflows.map((workflow) => (
                <TableRow key={workflow.id}>
                  <TableCell>{workflow.id}</TableCell>
                  <TableCell>{workflow.status}</TableCell>
                  <TableCell>{new Date(workflow.createdAt).toLocaleString()}</TableCell>
                  <TableCell>
                    <Button
                      size="small"
                      onClick={() => checkWorkflowStatus(workflow.id)}
                    >
                      Check Status
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
              {workflows.length === 0 && (
                <TableRow>
                  <TableCell colSpan={4} align="center">
                    No workflows started yet. Click "Start HelloBackstage Workflow" to begin.
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </TableContainer>
      </Paper>
    </div>
  );
};
