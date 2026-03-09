package types

import "time"

// Shared types for workflows and activities

type InfrastructureResult struct {
	ResourceID   string                 `json:"resourceId"`
	ResourceType string                 `json:"resourceType"`
	Properties   map[string]interface{} `json:"properties"`
	Emulated     bool                   `json:"emulated"`
}

type AgentResult struct {
	AgentID        string                 `json:"agentId"`
	AgentType      string                 `json:"agentType"`
	Status         string                 `json:"status"`
	Score          float64                `json:"score"`
	Findings       []string               `json:"findings"`
	Recommendations []string              `json:"recommendations"`
	Metadata       map[string]interface{} `json:"metadata"`
	ExecutedAt     time.Time              `json:"executedAt"`
}

type AggregatedResult struct {
	OverallScore       float64             `json:"overallScore"`
	AgentResults       []AgentResult       `json:"agentResults"`
	RequiresHumanReview bool               `json:"requiresHumanReview"`
	RiskLevel          string              `json:"riskLevel"`
	Summary            string              `json:"summary"`
	HumanReviewResult  *HumanReviewResult  `json:"humanReviewResult,omitempty"`
}

type HumanReviewResult struct {
	ReviewerID string    `json:"reviewerId"`
	Approved   bool      `json:"approved"`
	Decision   string    `json:"decision"`
	Comments   string    `json:"comments"`
	ReviewedAt time.Time `json:"reviewedAt"`
}

type ComplianceReport struct {
	ID              string            `json:"id"`
	TargetResource  string            `json:"targetResource"`
	OverallStatus   string            `json:"overallStatus"`
	Score           float64           `json:"score"`
	AgentResults    []AgentResult     `json:"agentResults"`
	RiskAssessment  RiskAssessment    `json:"riskAssessment"`
	Recommendations []string          `json:"recommendations"`
	GeneratedAt     time.Time         `json:"generatedAt"`
}

type RiskAssessment struct {
	Level          string   `json:"level"`
	CriticalItems  []string `json:"criticalItems"`
	WarningItems   []string `json:"warningItems"`
	InfoItems      []string `json:"infoItems"`
}

type ComplianceRequest struct {
	TargetResource   string            `json:"targetResource"`
	ComplianceType   string            `json:"complianceType"`
	Parameters       map[string]string `json:"parameters"`
	RequesterID      string            `json:"requesterId"`
	Priority         string            `json:"priority"`
}

type ComplianceResult struct {
	Report      ComplianceReport `json:"report"`
	Approved    bool             `json:"approved"`
	CompletedAt time.Time        `json:"completedAt"`
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
	TaskID      string    `json:"taskId"`
	Approved    bool      `json:"approved"`
	Decision    string    `json:"decision"`
	CompletedAt time.Time `json:"completedAt"`
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
}
