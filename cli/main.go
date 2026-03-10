package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/lloydchang/backstage-temporal/backend/mcp"
)

var (
	serverURL   string
	apiKey      string
	verbose     bool
	interactive bool
)

type CLIClient struct {
	serverURL string
	apiKey    string
	httpClient *http.Client
}

type WorkflowStatus struct {
	WorkflowID string                 `json:"workflowId"`
	RunID      string                 `json:"runId"`
	Status     string                 `json:"status"`
	StartTime  *time.Time             `json:"startTime,omitempty"`
	CloseTime  *time.Time             `json:"closeTime,omitempty"`
	Details    map[string]interface{} `json:"details,omitempty"`
}

type WorkflowResult struct {
	WorkflowID string      `json:"workflowId"`
	Status     string      `json:"status"`
	Result     interface{} `json:"result,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func NewCLIClient(serverURL, apiKey string) *CLIClient {
	return &CLIClient{
		serverURL: strings.TrimSuffix(serverURL, "/"),
		apiKey:    apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *CLIClient) makeRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	url := c.serverURL + endpoint
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	req.Header.Set("Content-Type", "application/json")

	if verbose {
		log.Printf("Making %s request to %s", method, url)
	}

	return c.httpClient.Do(req)
}

func (c *CLIClient) startWorkflow(workflowType string, targetResource string, params map[string]interface{}) (*WorkflowResult, error) {
	var endpoint string
	var requestBody map[string]interface{}

	switch workflowType {
	case "compliance":
		endpoint = "/workflow/start-enhanced-compliance"
		requestBody = map[string]interface{}{
			"targetResource": targetResource,
			"parameters":     params,
			"priority":       "normal",
		}
	case "batch":
		endpoint = "/workflow/start-batch"
		requestBody = []map[string]interface{}{
			{
				"targetResource": targetResource,
				"parameters":     params,
			},
		}
	default:
		return nil, fmt.Errorf("unknown workflow type: %s", workflowType)
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest("POST", endpoint, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result WorkflowResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *CLIClient) getWorkflowStatus(workflowID string) (*WorkflowStatus, error) {
	endpoint := "/workflow/status?id=" + workflowID

	resp, err := c.makeRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var status WorkflowStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}

func (c *CLIClient) listWorkflows() ([]WorkflowStatus, error) {
	// This would require a new endpoint - for now return empty
	return []WorkflowStatus{}, nil
}

func (c *CLIClient) signalWorkflow(workflowID, signalName, signalValue string) error {
	endpoint := fmt.Sprintf("/workflow/signal/%s", workflowID)

	requestBody := map[string]interface{}{
		"signal": signalName,
		"value":  signalValue,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	resp, err := c.makeRequest("POST", endpoint, strings.NewReader(string(jsonBody)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

func (c *CLIClient) getHealth() (map[string]interface{}, error) {
	resp, err := c.makeRequest("GET", "/health", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var health map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&health); err != nil {
		return nil, err
	}

	return health, nil
}

func (c *CLIClient) getMetrics() (map[string]interface{}, error) {
	resp, err := c.makeRequest("GET", "/metrics", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var metrics map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&metrics); err != nil {
		return nil, err
	}

	return metrics, nil
}

var rootCmd = &cobra.Command{
	Use:   "temporal-agents",
	Short: "Temporal AI Agents CLI",
	Long:  `Command-line interface for managing Temporal AI agent workflows and MCP interactions.`,
}

var startCmd = &cobra.Command{
	Use:   "start [workflow-type] [target-resource]",
	Short: "Start a new workflow",
	Long:  `Start a new AI agent workflow. Supported types: compliance, security, cost-analysis`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client := NewCLIClient(serverURL, apiKey)

		workflowType := args[0]
		targetResource := args[1]

		fmt.Printf("Starting %s workflow for %s...\n", workflowType, targetResource)

		result, err := client.startWorkflow(workflowType, targetResource, map[string]interface{}{})
		if err != nil {
			log.Fatalf("Failed to start workflow: %v", err)
		}

		fmt.Printf("Workflow started successfully!\n")
		fmt.Printf("Workflow ID: %s\n", result.WorkflowID)
		fmt.Printf("Run ID: %s\n", result.RunID)
		fmt.Printf("Status: %s\n", result.Status)
	},
}

var statusCmd = &cobra.Command{
	Use:   "status [workflow-id]",
	Short: "Get workflow status",
	Long:  `Get the current status of a running workflow.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := NewCLIClient(serverURL, apiKey)

		workflowID := args[0]

		status, err := client.getWorkflowStatus(workflowID)
		if err != nil {
			log.Fatalf("Failed to get workflow status: %v", err)
		}

		fmt.Printf("Workflow Status:\n")
		fmt.Printf("ID: %s\n", status.WorkflowID)
		fmt.Printf("Run ID: %s\n", status.RunID)
		fmt.Printf("Status: %s\n", status.Status)

		if status.StartTime != nil {
			fmt.Printf("Started: %s\n", status.StartTime.Format(time.RFC3339))
		}
		if status.CloseTime != nil {
			fmt.Printf("Completed: %s\n", status.CloseTime.Format(time.RFC3339))
		}

		if status.Details != nil && verbose {
			fmt.Printf("Details: %+v\n", status.Details)
		}
	},
}

var signalCmd = &cobra.Command{
	Use:   "signal [workflow-id] [signal-name] [signal-value]",
	Short: "Send a signal to a workflow",
	Long:  `Send a signal to a running workflow to trigger state changes.`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		client := NewCLIClient(serverURL, apiKey)

		workflowID := args[0]
		signalName := args[1]
		signalValue := args[2]

		fmt.Printf("Sending signal '%s' with value '%s' to workflow %s...\n",
			signalName, signalValue, workflowID)

		err := client.signalWorkflow(workflowID, signalName, signalValue)
		if err != nil {
			log.Fatalf("Failed to send signal: %v", err)
		}

		fmt.Println("Signal sent successfully!")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List workflows",
	Long:  `List all workflows and their current status.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := NewCLIClient(serverURL, apiKey)

		workflows, err := client.listWorkflows()
		if err != nil {
			log.Fatalf("Failed to list workflows: %v", err)
		}

		if len(workflows) == 0 {
			fmt.Println("No workflows found.")
			return
		}

		fmt.Printf("%-40s %-15s %-20s\n", "Workflow ID", "Status", "Started")
		fmt.Println(strings.Repeat("-", 75))

		for _, wf := range workflows {
			startTime := "N/A"
			if wf.StartTime != nil {
				startTime = wf.StartTime.Format("2006-01-02 15:04:05")
			}
			fmt.Printf("%-40s %-15s %-20s\n", truncateString(wf.WorkflowID, 40), wf.Status, startTime)
		}
	},
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check server health",
	Long:  `Check the health status of the Temporal AI Agents server.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := NewCLIClient(serverURL, apiKey)

		health, err := client.getHealth()
		if err != nil {
			log.Fatalf("Failed to get health status: %v", err)
		}

		fmt.Println("Server Health:")
		if status, ok := health["status"].(string); ok {
			fmt.Printf("Status: %s\n", status)
		}
		if version, ok := health["version"].(string); ok {
			fmt.Printf("Version: %s\n", version)
		}
		if uptime, ok := health["uptime"].(float64); ok {
			fmt.Printf("Uptime: %.2f seconds\n", uptime)
		}
		if goroutines, ok := health["goroutines"].(float64); ok {
			fmt.Printf("Goroutines: %.0f\n", goroutines)
		}

		if verbose {
			fmt.Printf("Full health info: %+v\n", health)
		}
	},
}

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Get server metrics",
	Long:  `Get performance metrics from the Temporal AI Agents server.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := NewCLIClient(serverURL, apiKey)

		metrics, err := client.getMetrics()
		if err != nil {
			log.Fatalf("Failed to get metrics: %v", err)
		}

		fmt.Println("Server Metrics:")
		for key, value := range metrics {
			fmt.Printf("%s: %v\n", key, value)
		}
	},
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Start interactive mode",
	Long:  `Start an interactive shell for managing workflows.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := NewCLIClient(serverURL, apiKey)
		runInteractiveMode(client)
	},
}

func runInteractiveMode(client *CLIClient) {
	fmt.Println("Temporal AI Agents Interactive Mode")
	fmt.Println("Type 'help' for available commands, 'quit' to exit")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		args := strings.Fields(line)
		command := args[0]

		switch command {
		case "help":
			printInteractiveHelp()
		case "start":
			handleInteractiveStart(client, args[1:])
		case "status":
			handleInteractiveStatus(client, args[1:])
		case "signal":
			handleInteractiveSignal(client, args[1:])
		case "list":
			handleInteractiveList(client)
		case "health":
			handleInteractiveHealth(client)
		case "metrics":
			handleInteractiveMetrics(client)
		case "quit", "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Printf("Unknown command: %s. Type 'help' for available commands.\n", command)
		}
		fmt.Println()
	}
}

func printInteractiveHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  start <type> <resource>     - Start a workflow")
	fmt.Println("  status <workflow-id>        - Get workflow status")
	fmt.Println("  signal <id> <name> <value>  - Send signal to workflow")
	fmt.Println("  list                        - List workflows")
	fmt.Println("  health                      - Check server health")
	fmt.Println("  metrics                     - Get server metrics")
	fmt.Println("  help                        - Show this help")
	fmt.Println("  quit                        - Exit interactive mode")
}

func handleInteractiveStart(client *CLIClient, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: start <type> <resource>")
		return
	}

	workflowType := args[0]
	targetResource := args[1]

	fmt.Printf("Starting %s workflow for %s...\n", workflowType, targetResource)

	result, err := client.startWorkflow(workflowType, targetResource, map[string]interface{}{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Workflow started: %s\n", result.WorkflowID)
}

func handleInteractiveStatus(client *CLIClient, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: status <workflow-id>")
		return
	}

	workflowID := args[0]

	status, err := client.getWorkflowStatus(workflowID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Status: %s\n", status.Status)
	if status.StartTime != nil {
		fmt.Printf("Started: %s\n", status.StartTime.Format(time.RFC3339))
	}
}

func handleInteractiveSignal(client *CLIClient, args []string) {
	if len(args) < 3 {
		fmt.Println("Usage: signal <workflow-id> <signal-name> <signal-value>")
		return
	}

	workflowID := args[0]
	signalName := args[1]
	signalValue := args[2]

	err := client.signalWorkflow(workflowID, signalName, signalValue)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Signal sent successfully")
}

func handleInteractiveList(client *CLIClient) {
	workflows, err := client.listWorkflows()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(workflows) == 0 {
		fmt.Println("No workflows found")
		return
	}

	for _, wf := range workflows {
		fmt.Printf("%s - %s\n", wf.WorkflowID, wf.Status)
	}
}

func handleInteractiveHealth(client *CLIClient) {
	health, err := client.getHealth()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if status, ok := health["status"].(string); ok {
		fmt.Printf("Status: %s\n", status)
	}
}

func handleInteractiveMetrics(client *CLIClient) {
	metrics, err := client.getMetrics()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for key, value := range metrics {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&serverURL, "server", "s", "http://localhost:8081", "Server URL")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", "", "API key for authentication")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Start in interactive mode")

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(signalCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(healthCmd)
	rootCmd.AddCommand(metricsCmd)
	rootCmd.AddCommand(interactiveCmd)
}

func main() {
	if interactive {
		client := NewCLIClient(serverURL, apiKey)
		runInteractiveMode(client)
	} else {
		if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
