package types

import "time"

// Enhanced types for V2 implementation with better error handling and monitoring

// Enhanced workflow request with validation
type EnhancedWorkflowRequest struct {
	TargetResource     string                 `json:"targetResource" validate:"required"`
	WorkflowType       string                 `json:"workflowType" validate:"required,oneof=enhanced-compliance security-scan cost-optimization"`
	Parameters         map[string]interface{} `json:"parameters"`
	Priority           string                 `json:"priority" validate:"oneof=low normal high critical"`
	DueDate            *time.Time             `json:"dueDate,omitempty"`
	RequesterID        string                 `json:"requesterId" validate:"required"`
	Tags               []string               `json:"tags"`
	Metadata           map[string]interface{} `json:"metadata"`
	NotificationConfig *NotificationConfig    `json:"notificationConfig,omitempty"`
}

type NotificationConfig struct {
	Email     []string `json:"email"`
	Webhook   string   `json:"webhook"`
	Slack     string   `json:"slack"`
	OnSuccess bool     `json:"onSuccess"`
	OnFailure bool     `json:"onFailure"`
}

// Enhanced workflow response with detailed status
type EnhancedWorkflowResponse struct {
	WorkflowID      string                    `json:"workflowId"`
	RunID           string                    `json:"runId"`
	Status          string                    `json:"status"`
	StartTime       time.Time                 `json:"startTime"`
	EndTime         *time.Time                `json:"endTime,omitempty"`
	Duration        *time.Duration            `json:"duration,omitempty"`
	Request         EnhancedWorkflowRequest   `json:"request"`
	Result          *WorkflowResult           `json:"result,omitempty"`
	Error           *WorkflowError            `json:"error,omitempty"`
	Metrics         *WorkflowMetricsV2        `json:"metrics,omitempty"`
	ExecutionPath   []ExecutionStep           `json:"executionPath"`
	Notifications   []NotificationEvent       `json:"notifications"`
}

type WorkflowResult struct {
	Type           string                 `json:"type"`
	Data           map[string]interface{} `json:"data"`
	Summary        string                 `json:"summary"`
	Confidence     float64                `json:"confidence"`
	RequiresAction bool                   `json:"requiresAction"`
	ActionItems    []string               `json:"actionItems"`
}

type WorkflowError struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Details   map[string]interface{} `json:"details"`
	Retryable bool                   `json:"retryable"`
	Timestamp time.Time              `json:"timestamp"`
}

type ExecutionStep struct {
	StepID      string                 `json:"stepId"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Status      string                 `json:"status"`
	StartTime   time.Time              `json:"startTime"`
	EndTime     *time.Time             `json:"endTime,omitempty"`
	Duration    *time.Duration         `json:"duration,omitempty"`
	Input       map[string]interface{} `json:"input"`
	Output      map[string]interface{} `json:"output,omitempty"`
	Error       *WorkflowError         `json:"error,omitempty"`
	Retries     int                    `json:"retries"`
}

type NotificationEvent struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Channel   string    `json:"channel"`
	Recipient string    `json:"recipient"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
}

// Enhanced batch workflow request
type BatchWorkflowRequest struct {
	BatchID       string                    `json:"batchId"`
	Requests      []EnhancedWorkflowRequest  `json:"requests"`
	BatchConfig   *BatchConfig              `json:"batchConfig,omitempty"`
	NotificationConfig *NotificationConfig  `json:"notificationConfig,omitempty"`
}

type BatchConfig struct {
	MaxConcurrent   int           `json:"maxConcurrent"`
	FailurePolicy   string        `json:"failurePolicy"` // continue, stop, retry
	RetryPolicy     *RetryPolicy  `json:"retryPolicy,omitempty"`
	Timeout         time.Duration `json:"timeout"`
	ProgressUpdates bool          `json:"progressUpdates"`
}

type RetryPolicy struct {
	MaxAttempts     int           `json:"maxAttempts"`
	InitialInterval  time.Duration `json:"initialInterval"`
	BackoffCoefficient float64     `json:"backoffCoefficient"`
	MaximumInterval  time.Duration `json:"maximumInterval"`
}

type BatchWorkflowResponse struct {
	BatchID       string                    `json:"batchId"`
	Status        string                    `json:"status"`
	StartTime     time.Time                 `json:"startTime"`
	EndTime       *time.Time                `json:"endTime,omitempty"`
	Duration      *time.Duration            `json:"duration,omitempty"`
	TotalCount    int                       `json:"totalCount"`
	CompletedCount int                      `json:"completedCount"`
	FailedCount   int                       `json:"failedCount"`
	Responses     []EnhancedWorkflowResponse `json:"responses"`
	Progress      float64                   `json:"progress"`
	Notifications []NotificationEvent       `json:"notifications"`
}

// Enhanced metrics for V2
type WorkflowMetricsV2 struct {
	WorkflowID       string                `json:"workflowId"`
	WorkflowType     string                `json:"workflowType"`
	StartTime        time.Time             `json:"startTime"`
	EndTime          *time.Time           `json:"endTime,omitempty"`
	Duration         *time.Duration       `json:"duration,omitempty"`
	Status           string                `json:"status"`
	
	// Performance metrics
	TotalDuration    time.Duration         `json:"totalDuration"`
	ActivityDuration time.Duration         `json:"activityDuration"`
	WaitDuration     time.Duration         `json:"waitDuration"`
	RetryDuration    time.Duration         `json:"retryDuration"`
	
	// Resource metrics
	ActivitiesRun    int                   `json:"activitiesRun"`
	ActivitiesFailed int                   `json:"activitiesFailed"`
	ActivitiesRetried int                  `json:"activitiesRetried"`
	
	// Quality metrics
	SuccessRate      float64               `json:"successRate"`
	AverageConfidence float64              `json:"averageConfidence"`
	ErrorRate        float64               `json:"errorRate"`
	
	// Cost metrics
	EstimatedCost    float64               `json:"estimatedCost"`
	ActualCost       float64               `json:"actualCost"`
	
	// Custom metrics
	CustomMetrics    map[string]interface{} `json:"customMetrics"`
}

// Enhanced agent result with more details
type EnhancedAgentResult struct {
	AgentID           string                 `json:"agentId"`
	AgentType         string                 `json:"agentType"`
	AgentVersion      string                 `json:"agentVersion"`
	Status            string                 `json:"status"`
	Score             float64                `json:"score"`
	Confidence        float64                `json:"confidence"`
	Findings          []EnhancedFinding      `json:"findings"`
	Recommendations   []EnhancedRecommendation `json:"recommendations"`
	RiskAssessment    *RiskAssessment        `json:"riskAssessment,omitempty"`
	ExecutionDetails  ExecutionDetails       `json:"executionDetails"`
	CostAnalysis      *CostAnalysis          `json:"costAnalysis,omitempty"`
	ComplianceStatus  ComplianceStatus       `json:"complianceStatus"`
	Metadata          map[string]interface{} `json:"metadata"`
	ExecutedAt        time.Time              `json:"executedAt"`
	ExecutionDuration time.Duration          `json:"executionDuration"`
	Tags              []string               `json:"tags"`
}

type EnhancedFinding struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Severity    string                 `json:"severity"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Location    string                 `json:"location"`
	Impact      string                 `json:"impact"`
	Remediation string                 `json:"remediation"`
	References  []string               `json:"references"`
	Metadata    map[string]interface{} `json:"metadata"`
	Confidence  float64                `json:"confidence"`
}

type EnhancedRecommendation struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Priority    string                 `json:"priority"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Justification string               `json:"justification"`
	Steps       []string               `json:"steps"`
	Effort      string                 `json:"effort"`
	Impact      string                 `json:"impact"`
	Timeline    time.Duration          `json:"timeline"`
	Dependencies []string               `json:"dependencies"`
	CostSavings float64                `json:"costSavings,omitempty"`
	ROI         float64                `json:"roi,omitempty"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type ExecutionDetails struct {
	StartTime       time.Time              `json:"startTime"`
	EndTime         *time.Time             `json:"endTime,omitempty"`
	Duration        *time.Duration         `json:"duration,omitempty"`
	Inputs          map[string]interface{} `json:"inputs"`
	Outputs         map[string]interface{} `json:"outputs,omitempty"`
	Steps           []ExecutionStep        `json:"steps"`
	ResourcesUsed   ResourceUsage          `json:"resourcesUsed"`
	Environment     string                 `json:"environment"`
	Version         string                 `json:"version"`
}

type ResourceUsage struct {
	CPUTime     time.Duration `json:"cpuTime"`
	MemoryUsed  int64         `json:"memoryUsed"`
	NetworkIO   int64         `json:"networkIo"`
	StorageUsed int64         `json:"storageUsed"`
	APICalls    int           `json:"apiCalls"`
}

type CostAnalysis struct {
	EstimatedCost    float64            `json:"estimatedCost"`
	ActualCost       float64            `json:"actualCost"`
	CostBreakdown    map[string]float64 `json:"costBreakdown"`
	SavingsOpportunity float64          `json:"savingsOpportunity"`
	Currency         string             `json:"currency"`
	BillingPeriod    string             `json:"billingPeriod"`
}

type ComplianceStatus struct {
	OverallStatus    string              `json:"overallStatus"`
	ComplianceScore  float64             `json:"complianceScore"`
	StandardsChecked []StandardCheck     `json:"standardsChecked"`
	Violations       []Violation         `json:"violations"`
	PoliciesPassed   int                 `json:"policiesPassed"`
	PoliciesFailed   int                 `json:"policiesFailed"`
	LastChecked      time.Time           `json:"lastChecked"`
	NextCheckDue     time.Time           `json:"nextCheckDue"`
}

type StandardCheck struct {
	StandardName string    `json:"standardName"`
	Version      string    `json:"version"`
	Status       string    `json:"status"`
	Score        float64   `json:"score"`
	CheckedAt    time.Time `json:"checkedAt"`
	ExpiryDate   time.Time `json:"expiryDate"`
}

type Violation struct {
	ID          string                 `json:"id"`
	Standard    string                 `json:"standard"`
	Control     string                 `json:"control"`
	Severity    string                 `json:"severity"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Impact      string                 `json:"impact"`
	Remediation string                 `json:"remediation"`
	DiscoveredAt time.Time             `json:"discoveredAt"`
	Status      string                 `json:"status"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// Enhanced health check response
type HealthCheckResponse struct {
	Status      string                 `json:"status"`
	Timestamp   time.Time              `json:"timestamp"`
	Version     string                 `json:"version"`
	Uptime      time.Duration          `json:"uptime"`
	Components  ComponentHealth        `json:"components"`
	Metrics     SystemMetrics          `json:"metrics"`
	Environment map[string]interface{} `json:"environment"`
}

type ComponentHealth struct {
	Temporal    ComponentStatus `json:"temporal"`
	Database    ComponentStatus `json:"database"`
	HTTPServer  ComponentStatus `json:"httpServer"`
	Metrics     ComponentStatus `json:"metrics"`
	Security    ComponentStatus `json:"security"`
}

type ComponentStatus struct {
	Status      string                 `json:"status"`
	LastChecked time.Time              `json:"lastChecked"`
	Message     string                 `json:"message,omitempty"`
	Details     map[string]interface{} `json:"details,omitempty"`
}

type SystemMetrics struct {
	Goroutines   int                 `json:"goroutines"`
	Memory       MemoryStats         `json:"memory"`
	CPU          CPUStats            `json:"cpu"`
	HTTP         HTTPStats           `json:"http"`
	Workflows    WorkflowStats       `json:"workflows"`
	Custom       map[string]interface{} `json:"custom"`
}

type MemoryStats struct {
	Alloc        uint64 `json:"alloc"`
	TotalAlloc   uint64 `json:"totalAlloc"`
	Sys          uint64 `json:"sys"`
	Lookups      uint64 `json:"lookups"`
	Mallocs      uint64 `json:"mallocs"`
	Frees        uint64 `json:"frees"`
	HeapAlloc    uint64 `json:"heapAlloc"`
	HeapSys      uint64 `json:"heapSys"`
	HeapIdle     uint64 `json:"heapIdle"`
	HeapInuse    uint64 `json:"heapInuse"`
	HeapReleased uint64 `json:"heapReleased"`
	HeapObjects  uint64 `json:"heapObjects"`
	StackInuse   uint64 `json:"stackInuse"`
	StackSys     uint64 `json:"stackSys"`
	MSpanInuse   uint64 `json:"mspanInuse"`
	MSpanSys     uint64 `json:"mspanSys"`
	MCacheInuse  uint64 `json:"mcacheInuse"`
	MCacheSys    uint64 `json:"mcacheSys"`
	BuckHashSys  uint64 `json:"buckHashSys"`
	GCSys        uint64 `json:"gcSys"`
	OtherSys     uint64 `json:"otherSys"`
	GCCPUFraction float64 `json:"gcCpuFraction"`
}

type CPUStats struct {
	UserTime    time.Duration `json:"userTime"`
	SystemTime  time.Duration `json:"systemTime"`
	IdleTime    time.Duration `json:"idleTime"`
	PercentUsed float64       `json:"percentUsed"`
}

type HTTPStats struct {
	RequestsTotal    int64         `json:"requestsTotal"`
	RequestsActive   int64         `json:"requestsActive"`
	ResponseTime     time.Duration `json:"averageResponseTime"`
	RequestsPerSecond float64      `json:"requestsPerSecond"`
	ErrorRate        float64       `json:"errorRate"`
	StatusCodes      map[int]int64 `json:"statusCodes"`
}

type WorkflowStats struct {
	ActiveWorkflows int64   `json:"activeWorkflows"`
	CompletedToday  int64   `json:"completedToday"`
	FailedToday     int64   `json:"failedToday"`
	AverageDuration time.Duration `json:"averageDuration"`
	SuccessRate     float64 `json:"successRate"`
	Throughput      float64 `json:"throughput"`
}
