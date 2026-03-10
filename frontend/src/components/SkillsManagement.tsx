import React, { useState, useEffect } from 'react';
import {
  Box,
  Typography,
  List,
  ListItem,
  ListItemText,
  Button,
  TextField,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Chip,
  Card,
  CardContent,
  CardActions,
  Grid,
  Alert,
  CircularProgress,
} from '@mui/material';
import { PlayArrow, Info, Refresh } from '@mui/icons-material';

interface Skill {
  name: string;
  description: string;
  scope: string;
  argumentHint?: string;
  userInvocable: boolean;
  allowedTools?: string[];
  context?: string;
  model?: string;
}

interface SkillExecution {
  skillName: string;
  executionId: string;
  forkRequired: boolean;
  agentType: string;
  content: string;
  arguments: string[];
}

const SkillsManagement: React.FC = () => {
  const [skills, setSkills] = useState<Skill[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [selectedSkill, setSelectedSkill] = useState<Skill | null>(null);
  const [executionDialog, setExecutionDialog] = useState(false);
  const [executionArgs, setExecutionArgs] = useState('');
  const [executing, setExecuting] = useState(false);
  const [lastExecution, setLastExecution] = useState<SkillExecution | null>(null);

  const API_BASE = 'http://localhost:8081';

  const fetchSkills = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch(`${API_BASE}/api/skills`);
      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }
      const data = await response.json();
      setSkills(data.skills || []);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch skills');
    } finally {
      setLoading(false);
    }
  };

  const executeSkill = async (skillName: string, args: string[]) => {
    setExecuting(true);
    setError(null);
    try {
      const response = await fetch(`${API_BASE}/api/skills/${skillName}/execute`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ arguments: args }),
      });

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}: ${response.statusText}`);
      }

      const execution = await response.json();
      setLastExecution(execution);

      return execution;
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to execute skill');
      throw err;
    } finally {
      setExecuting(false);
    }
  };

  const handleExecuteSkill = async () => {
    if (!selectedSkill) return;

    const args = executionArgs.trim() ? executionArgs.split(' ') : [];
    try {
      await executeSkill(selectedSkill.name, args);
      setExecutionDialog(false);
      setExecutionArgs('');
    } catch (err) {
      // Error already set in executeSkill
    }
  };

  const openExecutionDialog = (skill: Skill) => {
    setSelectedSkill(skill);
    setExecutionDialog(true);
  };

  const closeExecutionDialog = () => {
    setExecutionDialog(false);
    setSelectedSkill(null);
    setExecutionArgs('');
  };

  useEffect(() => {
    fetchSkills();
  }, []);

  return (
    <Box sx={{ p: 3 }}>
      <Typography variant="h4" gutterBottom>
        AI Agent Skills
      </Typography>

      <Box sx={{ mb: 3, display: 'flex', gap: 2 }}>
        <Button
          variant="contained"
          startIcon={<Refresh />}
          onClick={fetchSkills}
          disabled={loading}
        >
          {loading ? <CircularProgress size={20} /> : 'Refresh Skills'}
        </Button>
      </Box>

      {error && (
        <Alert severity="error" sx={{ mb: 3 }}>
          {error}
        </Alert>
      )}

      {lastExecution && (
        <Alert severity="success" sx={{ mb: 3 }}>
          Skill "{lastExecution.skillName}" executed successfully!
          {lastExecution.forkRequired && (
            <> Running in forked context with agent: {lastExecution.agentType}</>
          )}
        </Alert>
      )}

      <Grid container spacing={3}>
        {skills.map((skill) => (
          <Grid item xs={12} md={6} lg={4} key={skill.name}>
            <Card sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
              <CardContent sx={{ flexGrow: 1 }}>
                <Typography variant="h6" gutterBottom>
                  /{skill.name}
                </Typography>
                <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                  {skill.description}
                </Typography>

                <Box sx={{ mb: 1 }}>
                  <Chip
                    label={skill.scope}
                    size="small"
                    color={skill.scope === 'repo' ? 'primary' : 'secondary'}
                  />
                  {skill.context && (
                    <Chip
                      label={`Context: ${skill.context}`}
                      size="small"
                      variant="outlined"
                      sx={{ ml: 1 }}
                    />
                  )}
                  {skill.model && (
                    <Chip
                      label={`Model: ${skill.model}`}
                      size="small"
                      variant="outlined"
                      sx={{ ml: 1 }}
                    />
                  )}
                </Box>

                {skill.argumentHint && (
                  <Typography variant="caption" color="text.secondary">
                    Args: {skill.argumentHint}
                  </Typography>
                )}

                {skill.allowedTools && skill.allowedTools.length > 0 && (
                  <Box sx={{ mt: 1 }}>
                    <Typography variant="caption" color="text.secondary">
                      Tools: {skill.allowedTools.join(', ')}
                    </Typography>
                  </Box>
                )}
              </CardContent>

              <CardActions>
                <Button
                  size="small"
                  startIcon={<Info />}
                  onClick={() => setSelectedSkill(skill)}
                  disabled={!skill.userInvocable}
                >
                  Info
                </Button>
                <Button
                  size="small"
                  variant="contained"
                  startIcon={<PlayArrow />}
                  onClick={() => openExecutionDialog(skill)}
                  disabled={!skill.userInvocable || executing}
                >
                  Execute
                </Button>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>

      {skills.length === 0 && !loading && (
        <Box sx={{ textAlign: 'center', py: 4 }}>
          <Typography variant="h6" color="text.secondary">
            No skills available
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Skills will appear here once the backend discovers them from the .agents/skills or .claude/skills directories.
          </Typography>
        </Box>
      )}

      {/* Skill Info Dialog */}
      <Dialog open={!!selectedSkill && !executionDialog} onClose={() => setSelectedSkill(null)} maxWidth="md" fullWidth>
        <DialogTitle>
          Skill: {selectedSkill?.name}
        </DialogTitle>
        <DialogContent>
          {selectedSkill && (
            <Box>
              <Typography variant="body1" sx={{ mb: 2 }}>
                {selectedSkill.description}
              </Typography>

              <Grid container spacing={2}>
                <Grid item xs={6}>
                  <Typography variant="subtitle2">Scope:</Typography>
                  <Chip label={selectedSkill.scope} size="small" />
                </Grid>
                <Grid item xs={6}>
                  <Typography variant="subtitle2">User Invocable:</Typography>
                  <Chip
                    label={selectedSkill.userInvocable ? 'Yes' : 'No'}
                    size="small"
                    color={selectedSkill.userInvocable ? 'success' : 'default'}
                  />
                </Grid>
                {selectedSkill.argumentHint && (
                  <Grid item xs={12}>
                    <Typography variant="subtitle2">Arguments:</Typography>
                    <Typography variant="body2">{selectedSkill.argumentHint}</Typography>
                  </Grid>
                )}
                {selectedSkill.allowedTools && selectedSkill.allowedTools.length > 0 && (
                  <Grid item xs={12}>
                    <Typography variant="subtitle2">Allowed Tools:</Typography>
                    <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 1 }}>
                      {selectedSkill.allowedTools.map((tool) => (
                        <Chip key={tool} label={tool} size="small" variant="outlined" />
                      ))}
                    </Box>
                  </Grid>
                )}
                {selectedSkill.context && (
                  <Grid item xs={6}>
                    <Typography variant="subtitle2">Context:</Typography>
                    <Typography variant="body2">{selectedSkill.context}</Typography>
                  </Grid>
                )}
                {selectedSkill.model && (
                  <Grid item xs={6}>
                    <Typography variant="subtitle2">Model:</Typography>
                    <Typography variant="body2">{selectedSkill.model}</Typography>
                  </Grid>
                )}
              </Grid>
            </Box>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setSelectedSkill(null)}>Close</Button>
          {selectedSkill?.userInvocable && (
            <Button variant="contained" onClick={() => openExecutionDialog(selectedSkill)}>
              Execute Skill
            </Button>
          )}
        </DialogActions>
      </Dialog>

      {/* Skill Execution Dialog */}
      <Dialog open={executionDialog} onClose={closeExecutionDialog} maxWidth="sm" fullWidth>
        <DialogTitle>
          Execute Skill: {selectedSkill?.name}
        </DialogTitle>
        <DialogContent>
          <Typography variant="body2" sx={{ mb: 2 }}>
            {selectedSkill?.description}
          </Typography>

          {selectedSkill?.argumentHint && (
            <Typography variant="caption" color="text.secondary" sx={{ mb: 1, display: 'block' }}>
              Arguments: {selectedSkill.argumentHint}
            </Typography>
          )}

          <TextField
            fullWidth
            label="Arguments (space-separated)"
            value={executionArgs}
            onChange={(e) => setExecutionArgs(e.target.value)}
            placeholder={selectedSkill?.argumentHint || "Enter arguments..."}
            sx={{ mt: 1 }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={closeExecutionDialog}>Cancel</Button>
          <Button
            variant="contained"
            onClick={handleExecuteSkill}
            disabled={executing}
            startIcon={executing ? <CircularProgress size={16} /> : <PlayArrow />}
          >
            {executing ? 'Executing...' : 'Execute'}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default SkillsManagement;
