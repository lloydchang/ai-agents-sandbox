import React, { useState } from 'react';
import { Button, Table, TableBody, TableCell, TableHead, TableRow, Typography } from '@material-ui/core';

const TemporalPage = () => {
  const [workflows, setWorkflows] = useState<{id: string, status: string}[]>([]);

  const startWorkflow = async () => {
    try {
      const response = await fetch('http://localhost:8081/workflow/start', { method: 'POST' });
      const id = await response.text();
      setWorkflows([...workflows, { id, status: 'Running' }]);
    } catch (error) {
      console.error('Error starting workflow:', error);
    }
  };

  const checkStatus = async (id: string) => {
    try {
      const response = await fetch(`http://localhost:8081/workflow/status?id=${id}`);
      const status = await response.text();
      setWorkflows(workflows.map(w => w.id === id ? { ...w, status } : w));
    } catch (error) {
      console.error('Error checking status:', error);
    }
  };

  return (
    <div>
      <Typography variant="h4">Temporal Integration</Typography>
      <Button variant="contained" color="primary" onClick={startWorkflow}>Start Workflow</Button>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell>Status</TableCell>
            <TableCell>Actions</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {workflows.map(w => (
            <TableRow key={w.id}>
              <TableCell>{w.id}</TableCell>
              <TableCell>{w.status}</TableCell>
              <TableCell>
                <Button onClick={() => checkStatus(w.id)}>Check Status</Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};

export { TemporalPage };
