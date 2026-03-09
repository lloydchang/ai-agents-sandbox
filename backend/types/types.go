package types

import "time"

// Shared types for workflows and activities

type InfrastructureResult struct {
	ResourceID        string                 `json:"resourceId"`
	ResourceType      string                 `json:"resourceType"`
	Properties        map[string]interface{} `json:"properties"`
	Emulated          bool                   `json:"emulated"`
	DiscoveryTime     time.Time              `json:"discoveryTime"`
	ValidationStatus  string                 `json:"validationStatus"`
}

type AgentResult struct {
	AgentID           string                 `json:"agentId"`
	AgentType         string                 `json:"agentType"`
	Status            string                 `json:"status"`
	Score             float64                `json:"score"`
	Confidence        float64                `json:"confidence"`
	Findings          []string               `json:"findings"`
	Recommendations   []string              `json:"recommendations"`
	Metadata          map[string]interface{} `json:"metadata"`
	ExecutedAt        time.Time              `json:"executedAt"`
	ExecutionDuration time.Duration          `json:"executionDuration"`
	Tags              []string               `json:"tags"`
}

type AggregatedResult struct {
	OverallScore        float64             `json:"overallScore"`
	ConfidenceScore     float64             `json:"confidenceScore"`
	AgentResults        []AgentResult       `json:"agentResults"`
	RequiresHumanReview bool               `json:"requiresHumanReview"`
	RiskLevel           string              `json:"riskLevel"`
	Summary             string              `json:"summary"`
	HumanReviewResult   *HumanReviewResult  `json:"humanReviewResult,omitempty"`
	AggregationMethod   string              `json:"aggregationMethod"`
	ProcessedAt         time.Time           `json:"processedAt"`
}

type HumanReviewResult struct {
	ReviewerID      string                 `json:"reviewerId"`
	Approved        bool                   `json:"approved"`
	Decision        string                 `json:"decision"`
	Comments        string                 `json:"comments"`
	ReviewedAt      time.Time              `json:"reviewedAt"`
	ReviewDuration  time.Duration          `json:"reviewDuration"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	Confidence      float64                `json:"confidence"`
	EscalationLevel int                    `json:"escalationLevel"`
}

type ComplianceReport struct {
	ID              string            `json:"id"`
	TargetResource  string            `json:"targetResource"`
	OverallStatus   string            `json:"overallStatus"`
	Score           float64           `json:"score"`
	Confidence      float64           `json:"confidence"`
	AgentResults    []AgentResult     `json:"agentResults"`
	RiskAssessment  RiskAssessment    `json:"riskAssessment"`
	Recommendations []string          `json:"recommendations"`
	GeneratedAt     time.Time         `json:"generatedAt"`
	ReportVersion   string            `json:"reportVersion"`
	ComplianceFramework string         `json:"complianceFramework"`
}

type RiskAssessment struct {
	Level          string   `json:"level"`
	CriticalItems  []string `json:"criticalItems"`
	WarningItems   []string `json:"warningItems"`
	InfoItems      []string `json:"infoItems"`
}

type ComplianceRequest struct {
	TargetResource     string            `json:"targetResource"`
	ComplianceType     string            `json:"complianceType"`
	Parameters         map[string]string `json:"parameters"`
	RequesterID        string            `json:"requesterId"`
	Priority           string            `json:"priority"`
	DueDate            time.Time         `json:"dueDate"`
	AutoApproval       bool              `json:"autoApproval"`
	RequiredScore      float64           `json:"requiredScore"`
	EscalationPolicy   string            `json:"escalationPolicy"`
}

type ComplianceResult struct {
	Report         ComplianceReport       `json:"report"`
	Approved       bool                   `json:"approved"`
	CompletedAt    time.Time              `json:"completedAt"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	ProcessingTime time.Duration          `json:"processingTime"`
	AutoApproved   bool                   `json:"autoApproved"`
}

type HumanTask struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Priority    string                `json:"priority"`
	AssignedTo  string                `json:"assignedTo"`
	DueAt       time.Time             `json:"dueAt"`
	Status      HumanTaskStatus       `json:"status"`
	Data        map[string]interface{} `json:"data"`
}

type HumanTaskStatus struct {
	State      string    `json:"state"`
	UpdatedAt  time.Time `json:"updatedAt"`
	UpdatedBy  string    `json:"updatedBy"`
	Notes      string    `json:"notes"`
}

type HumanTaskResult struct {
	TaskID      string                 `json:"taskId"`
	Approved    bool                   `json:"approved"`
	Decision    string                 `json:"decision"`
	CompletedAt time.Time              `json:"completedAt"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type CollaborationRequest struct {
	TaskID           string   `json:"taskId"`
	PrimaryAgent     string   `json:"primaryAgent"`
	ValidationAgents []string `json:"validationAgents"`
	Data             map[string]interface{} `json:"data"`
	ConsensusType    string   `json:"consensusType"`
}

type CollaborationResult struct {
	TaskID           string                 `json:"taskId"`
	ConsensusResult  ConsensusResult        `json:"consensusResult"`
	Recommendation   string                 `json:"recommendation"`
	Confidence       float64                `json:"confidence"`
	AgentResults     []AgentResult          `json:"agentResults"`
	Metadata         map[string]interface{} `json:"metadata"`
	CompletedAt      time.Time              `json:"completedAt"`
}

type ConsensusResult struct {
	ConsensusLevel    string    `json:"consensusLevel"`
	AgreementScore    float64   `json:"agreementScore"`
	ConflictingItems  []string  `json:"conflictingItems"`
	ResolvedItems     []string  `json:"resolvedItems"`
	RequiresEscalation bool     `json:"requiresEscalation"`
	ResolvedAt        time.Time `json:"resolvedAt"`
	ResolutionMethod  string    `json:"resolutionMethod"`
}

// Enhanced monitoring and metrics types
type WorkflowMetrics struct {
	WorkflowID       string                    `json:"workflowId"`
	WorkflowType     string                    `json:"workflowType"`
	StartTime        time.Time                 `json:"startTime"`
	EndTime          time.Time                 `json:"endTime"`
	Duration         time.Duration             `json:"duration"`
	Status           string                    `json:"status"`
	Stages           map[string]*StageMetrics  `json:"stages"`
	AgentResults     []AgentResult             `json:"agentResults"`
	ErrorCount       int                       `json:"errorCount"`
	RetryCount       int                       `json:"retryCount"`
	ResourceUsage    ResourceUsage             `json:"resourceUsage"`
	Cost             float64                   `json:"cost"`
}

type StageMetrics struct {
	Name        string        `json:"name"`
	StartTime   time.Time     `json:"startTime"`
	EndTime     time.Time     `json:"endTime"`
	Status      string        `json:"status"`
	Duration    time.Duration `json:"duration"`
	ErrorCount  int           `json:"errorCount"`
	RetryCount  int           `json:"retryCount"`
	Inputs      interface{}   `json:"inputs"`
	Outputs     interface{}   `json:"outputs"`
}

type ResourceUsage struct {
	CPUUsage      float64 `json:"cpuUsage"`
	MemoryUsage   float64 `json:"memoryUsage"`
	NetworkIO     int64   `json:"networkIO"`
	StorageIO     int64   `json:"storageIO"`
	AgentCount    int     `json:"agentCount"`
	ParallelTasks int     `json:"parallelTasks"`
}

type Alert struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Severity    string                 `json:"severity"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	WorkflowID  string                 `json:"workflowId"`
	Timestamp   time.Time              `json:"timestamp"`
	Status      string                 `json:"status"`
	Metadata    map[string]interface{} `json:"metadata"`
	ResolvedAt  *time.Time             `json:"resolvedAt,omitempty"`
}

// Enhanced configuration types
type WorkflowConfig struct {
	Name                string        `json:"name"`
	Version             string        `json:"version"`
	Timeout             time.Duration `json:"timeout"`
	RetryPolicy         RetryPolicy   `json:"retryPolicy"`
	MaxConcurrent       int           `json:"maxConcurrent"`
	EnableMonitoring    bool          `json:"enableMonitoring"`
	EnableAutoApproval  bool          `json:"enableAutoApproval"`
	RequiredScore       float64       `json:"requiredScore"`
	EscalationPolicy    string        `json:"escalationPolicy"`
}

type RetryPolicy struct {
	InitialInterval    time.Duration `json:"initialInterval"`
	BackoffCoefficient float64       `json:"backoffCoefficient"`
	MaximumInterval    time.Duration `json:"maximumInterval"`
	MaximumAttempts    int32         `json:"maximumAttempts"`
	NonRetryableErrors []string      `json:"nonRetryableErrors"`
}

// Enhanced notification types
type Notification struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Recipient   string                 `json:"recipient"`
	Subject     string                 `json:"subject"`
	Message     string                 `json:"message"`
	WorkflowID  string                 `json:"workflowId"`
	Timestamp   time.Time              `json:"timestamp"`
	Status      string                 `json:"status"`
	Channels    []string               `json:"channels"`
	Metadata    map[string]interface{} `json:"metadata"`
	ReadAt      *time.Time             `json:"readAt,omitempty"`
}
