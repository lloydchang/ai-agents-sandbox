package workflows

import (
	"fmt"
	"math"
	"sort"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"github.com/lloydchang/backstage-temporal/backend/activities"
	"github.com/lloydchang/backstage-temporal/backend/types"
)

// EnhancedAIOrchestrationWorkflowV3 - Third iteration with advanced orchestration patterns
func EnhancedAIOrchestrationWorkflowV3(ctx workflow.Context, request types.ComplianceRequest) (*types.ComplianceResult, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting Enhanced AI Orchestration Workflow V3", "request", request)

	// Advanced activity options with dynamic timeouts
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 20,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 3,
			BackoffCoefficient: 2.5,
			MaximumInterval:    time.Minute * 3,
			MaximumAttempts:    4,
			NonRetryableErrorTypes: []string{"ValidationError", "AuthenticationError", "ResourceNotFound"},
		},
		HeartbeatTimeout: time.Minute * 2,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// Advanced workflow metrics with predictive analytics
	metrics := &types.WorkflowMetricsV3{
		WorkflowID:   workflow.GetInfo(ctx).WorkflowExecution.ID,
		WorkflowType: "EnhancedAIOrchestrationV3",
		StartTime:    workflow.Now(ctx),
		Status:       "running",
		Stages:       make(map[string]*types.StageMetricsV3),
		ResourceUsage: types.ResourceUsageV3{
			AgentCount:     3,
			MaxConcurrency: 3,
		},
		PerformanceMetrics: types.PerformanceMetrics{
			EstimatedDuration: time.Minute * 15,
			ConfidenceTarget:  85.0,
		},
	}

	// Initialize circuit breaker and adaptive rate limiter
	circuitBreaker := NewAdaptiveCircuitBreaker("ai-orchestration-v3", 5, time.Minute*15)
	rateLimiter := NewAdaptiveRateLimiter(10, time.Minute)

	// Stage 1: Intelligent Infrastructure Discovery with predictive validation
	logger.Info("Stage 1: Intelligent Infrastructure Discovery")
	metrics.Stages["discovery"] = &types.StageMetricsV3{
		Name:      "Intelligent Infrastructure Discovery",
		StartTime: workflow.Now(ctx),
		Status:    "running",
	}

	infraResult, err := intelligentInfrastructureDiscovery(ctx, request, metrics, circuitBreaker)
	if err != nil {
		metrics.Stages["discovery"].Status = "failed"
		metrics.Stages["discovery"].EndTime = workflow.Now(ctx)
		metrics.ErrorCount++
		return nil, fmt.Errorf("intelligent infrastructure discovery failed: %w", err)
	}

	metrics.Stages["discovery"].Status = "completed"
	metrics.Stages["discovery"].EndTime = workflow.Now(ctx)

	// Stage 2: Dynamic Agent Orchestration with dependency graphs
	logger.Info("Stage 2: Dynamic Agent Orchestration")
	metrics.Stages["agent-orchestration"] = &types.StageMetricsV3{
		Name:      "Dynamic Agent Orchestration",
		StartTime: workflow.Now(ctx),
		Status:    "running",
	}

	agentResults, orchestrationMetrics, err := dynamicAgentOrchestrationV3(ctx, infraResult, request, metrics, circuitBreaker, rateLimiter)
	if err != nil {
		metrics.Stages["agent-orchestration"].Status = "failed"
		metrics.Stages["agent-orchestration"].EndTime = workflow.Now(ctx)
		metrics.ErrorCount++
		return nil, fmt.Errorf("dynamic agent orchestration failed: %w", err)
	}

	metrics.Stages["agent-orchestration"].Status = "completed"
	metrics.Stages["agent-orchestration"].EndTime = workflow.Now(ctx)
	metrics.ResourceUsage = orchestrationMetrics.ResourceUsage

	// Stage 3: ML-enhanced Result Aggregation with Bayesian confidence
	logger.Info("Stage 3: ML-enhanced Result Aggregation")
	metrics.Stages["ml-aggregation"] = &types.StageMetricsV3{
		Name:      "ML-enhanced Aggregation",
		StartTime: workflow.Now(ctx),
		Status:    "running",
	}

	aggregatedResult, err := mlEnhancedAggregationV3(ctx, agentResults, request, metrics)
	if err != nil {
		metrics.Stages["ml-aggregation"].Status = "failed"
		metrics.Stages["ml-aggregation"].EndTime = workflow.Now(ctx)
		metrics.ErrorCount++
		return nil, fmt.Errorf("ML-enhanced aggregation failed: %w", err)
	}

	metrics.Stages["ml-aggregation"].Status = "completed"
	metrics.Stages["ml-aggregation"].EndTime = workflow.Now(ctx)

	// Stage 4: Adaptive Human Review with predictive escalation
	logger.Info("Stage 4: Adaptive Human Review")
	metrics.Stages["adaptive-review"] = &types.StageMetricsV3{
		Name:      "Adaptive Human Review",
		StartTime: workflow.Now(ctx),
		Status:    "running",
	}

	requiresReview, reviewPriority, reviewTimeout := adaptiveHumanReviewDecisionV3(aggregatedResult, agentResults, request, metrics)

	if requiresReview {
		logger.Info("Adaptive human review required", "priority", reviewPriority, "timeout", reviewTimeout)

		reviewCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: reviewTimeout,
			RetryPolicy: &temporal.RetryPolicy{
				InitialInterval:    time.Minute * 15,
				MaximumAttempts:    1,
			},
		})

		var reviewResult types.HumanReviewResultV3
		err = workflow.ExecuteActivity(reviewCtx, activities.EnhancedHumanReviewActivity, aggregatedResult).Get(ctx, &reviewResult)
		if err != nil {
			logger.Warn("Adaptive human review failed or timed out", "error", err)
			reviewResult = generateAdaptiveFallbackReviewV3(aggregatedResult, request, reviewPriority)
		}

		aggregatedResult.HumanReviewResult = &types.HumanReviewResult{
			ReviewerID:     reviewResult.ReviewerID,
			Approved:       reviewResult.Approved,
			Decision:       reviewResult.Decision,
			Comments:       reviewResult.Comments,
			ReviewedAt:     reviewResult.ReviewedAt,
			ReviewDuration: reviewResult.ReviewDuration,
			Confidence:     reviewResult.Confidence,
			EscalationLevel: reviewResult.EscalationLevel,
		}
		metrics.Stages["adaptive-review"].Status = "completed"
	} else {
		metrics.Stages["adaptive-review"].Status = "skipped"
		logger.Info("Adaptive review not required - auto-approving")
	}

	metrics.Stages["adaptive-review"].EndTime = workflow.Now(ctx)

	// Stage 5: Intelligent Report Generation with insights
	logger.Info("Stage 5: Intelligent Report Generation")
	metrics.Stages["intelligent-reporting"] = &types.StageMetricsV3{
		Name:      "Intelligent Report Generation",
		StartTime: workflow.Now(ctx),
		Status:    "running",
	}

	finalReport, err := intelligentReportGenerationV3(ctx, aggregatedResult, request, metrics, orchestrationMetrics)
	if err != nil {
		metrics.Stages["intelligent-reporting"].Status = "failed"
		metrics.Stages["intelligent-reporting"].EndTime = workflow.Now(ctx)
		metrics.ErrorCount++
		return nil, fmt.Errorf("intelligent report generation failed: %w", err)
	}

	metrics.Stages["intelligent-reporting"].Status = "completed"
	metrics.Stages["intelligent-reporting"].EndTime = workflow.Now(ctx)

	// Final result with comprehensive metadata
	result := &types.ComplianceResult{
		Report:      finalReport,
		Approved:    determineAdaptiveApprovalStatusV3(aggregatedResult, request),
		CompletedAt: workflow.Now(ctx),
		Metadata: map[string]interface{}{
			"workflowVersion":     "v3.0",
			"totalAgents":         len(agentResults),
			"humanReviewRequired": requiresReview,
			"reviewPriority":      reviewPriority,
			"autoApproved":        !requiresReview,
			"orchestrationMetrics": orchestrationMetrics,
			"performanceMetrics": metrics.PerformanceMetrics,
			"confidenceScore":    aggregatedResult.ConfidenceScore,
			"adaptiveFeatures": []string{
				"dynamic-agent-orchestration",
				"ml-enhanced-aggregation",
				"adaptive-human-review",
				"intelligent-reporting",
				"predictive-escalation",
				"circuit-breaker-resilience",
			},
		},
		ProcessingTime: workflow.Now(ctx).Sub(metrics.StartTime),
		AutoApproved:   !requiresReview,
	}

	// Complete metrics
	metrics.EndTime = workflow.Now(ctx)
	metrics.Duration = metrics.EndTime.Sub(metrics.StartTime)
	metrics.Status = "completed"
	metrics.PerformanceMetrics.ActualDuration = metrics.Duration

	logger.Info("Enhanced AI Orchestration Workflow V3 completed successfully",
		"duration", metrics.Duration,
		"approved", result.Approved,
		"autoApproved", result.AutoApproved,
		"confidenceScore", aggregatedResult.ConfidenceScore)

	return result, nil
}

// intelligentInfrastructureDiscovery performs predictive validation and enrichment
func intelligentInfrastructureDiscovery(ctx workflow.Context, request types.ComplianceRequest, metrics *types.WorkflowMetricsV3, circuitBreaker *AdaptiveCircuitBreaker) (types.InfrastructureResult, error) {
	logger := workflow.GetLogger(ctx)

	// Predictive validation before execution
	if !circuitBreaker.ShouldAttempt() {
		return types.InfrastructureResult{}, fmt.Errorf("circuit breaker open - skipping infrastructure discovery")
	}

	var infraResult types.InfrastructureResult
	err := workflow.ExecuteActivity(ctx, activities.EnhancedDiscoverInfrastructureActivity, request.TargetResource).Get(ctx, &infraResult)
	if err != nil {
		circuitBreaker.RecordFailure()
		return types.InfrastructureResult{}, err
	}

	// Intelligent validation with enrichment
	if err := intelligentInfrastructureValidationV3(infraResult); err != nil {
		logger.Warn("Infrastructure validation failed", "error", err)
		// Continue but flag the issue
		infraResult.ValidationStatus = "warning"
		infraResult.ValidationMessage = err.Error()
	} else {
		infraResult.ValidationStatus = "success"
	}

	// Enrich with predictive metadata
	infraResult.DiscoveryMethod = "intelligent-v3"
	infraResult.DiscoveredAt = workflow.Now(ctx)
	infraResult.ConfidenceScore = 95.0 // High confidence for successful discovery

	circuitBreaker.RecordSuccess()
	return infraResult, nil
}

// dynamicAgentOrchestrationV3 implements advanced dependency graphs and adaptive execution
func dynamicAgentOrchestrationV3(ctx workflow.Context, infra types.InfrastructureResult, request types.ComplianceRequest, metrics *types.WorkflowMetricsV3, circuitBreaker *AdaptiveCircuitBreaker, rateLimiter *AdaptiveRateLimiter) ([]types.AgentResult, *types.OrchestrationMetrics, error) {
	logger := workflow.GetLogger(ctx)

	// Build dynamic agent configuration based on infrastructure analysis
	agentConfigs := buildDynamicAgentConfigsV3(infra, request)

	// Create dependency graph
	dependencyGraph := buildDependencyGraph(agentConfigs)

	// Execute with dynamic scheduling
	results, orchestrationMetrics, err := executeWithDynamicSchedulingV3(ctx, agentConfigs, dependencyGraph, infra, circuitBreaker, rateLimiter)
	if err != nil {
		return nil, nil, err
	}

	// Update metrics with orchestration data
	metrics.ResourceUsage.AgentCount = len(results)
	metrics.ResourceUsage.ParallelTasks = orchestrationMetrics.MaxConcurrencyUsed
	metrics.PerformanceMetrics.ActualAgentCount = len(results)

	return results, orchestrationMetrics, nil
}

// mlEnhancedAggregationV3 uses Bayesian confidence and ensemble methods
func mlEnhancedAggregationV3(ctx workflow.Context, agentResults []types.AgentResult, request types.ComplianceRequest, metrics *types.WorkflowMetricsV3) (types.AggregatedResult, error) {
	// Calculate Bayesian confidence intervals
	confidenceInterval := calculateBayesianConfidence(agentResults)

	// Ensemble aggregation with weighted scoring
	ensembleResult := performEnsembleAggregation(agentResults, request)

	// Apply ML-based risk assessment
	riskAssessment := applyMLRiskAssessment(ensembleResult, agentResults, confidenceInterval)

	// Combine results
	finalResult := types.AggregatedResult{
		AgentResults:     agentResults,
		OverallScore:     ensembleResult.OverallScore,
		RiskLevel:        riskAssessment.RiskLevel,
		ConfidenceScore:  confidenceInterval.ConfidenceScore,
		AggregationMethod: "ml-enhanced-ensemble-v3",
		ProcessedAt:      workflow.Now(ctx),
		ConsensusLevel:   calculateConsensusLevel(agentResults),
		RiskFactors:      riskAssessment.RiskFactors,
		Recommendations:  riskAssessment.Recommendations,
	}

	return finalResult, nil
}

// adaptiveHumanReviewDecisionV3 uses predictive analytics for review decisions
func adaptiveHumanReviewDecisionV3(aggregated types.AggregatedResult, agentResults []types.AgentResult, request types.ComplianceRequest, metrics *types.WorkflowMetricsV3) (bool, string, time.Duration) {
	// Predictive risk modeling
	predictiveRisk := calculatePredictiveRisk(aggregated, agentResults, request)

	// Adaptive threshold based on historical performance
	adaptiveThreshold := calculateAdaptiveThreshold(metrics.PerformanceMetrics)

	// Consensus analysis with temporal weighting
	temporalConsensus := analyzeTemporalConsensus(agentResults)

	// Decision matrix
	if predictiveRisk.Score >= 90.0 && temporalConsensus >= 85.0 && aggregated.ConfidenceScore >= adaptiveThreshold {
		return false, "", 0 // No review needed
	}

	// Determine priority and timeout
	priority := determineReviewPriorityV3(predictiveRisk, aggregated)
	timeout := calculateAdaptiveTimeout(priority, predictiveRisk.Urgency)

	return true, priority, timeout
}

// intelligentReportGenerationV3 creates actionable insights with predictive recommendations
func intelligentReportGenerationV3(ctx workflow.Context, aggregated types.AggregatedResult, request types.ComplianceRequest, metrics *types.WorkflowMetricsV3, orchestrationMetrics *types.OrchestrationMetrics) (types.ComplianceReport, error) {
	var finalReport types.ComplianceReport
	err := workflow.ExecuteActivity(ctx, activities.IntelligentComplianceReportActivity, types.ReportRequest{
		AggregatedResult: aggregated,
		Request:         request,
		Metrics:         metrics,
		OrchestrationMetrics: orchestrationMetrics,
	}).Get(ctx, &finalReport)
	if err != nil {
		return types.ComplianceReport{}, err
	}

	// Enhance with predictive insights
	finalReport.ReportVersion = "v3.0-intelligent"
	finalReport.Confidence = aggregated.ConfidenceScore
	finalReport.ComplianceFramework = request.ComplianceType

	// Add predictive recommendations
	finalReport.PredictiveInsights = generatePredictiveInsights(aggregated, metrics, orchestrationMetrics)
	finalReport.ActionableRecommendations = generateActionableRecommendations(aggregated, request)

	return finalReport, nil
}

// Helper functions for advanced orchestration

func buildDynamicAgentConfigsV3(infra types.InfrastructureResult, request types.ComplianceRequest) []AgentConfigV3 {
	baseConfigs := []AgentConfigV3{
		{
			Type:         "security",
			Priority:     1,
			Timeout:      time.Minute * 6,
			MaxRetries:   3,
			Required:     true,
			Dependencies: []string{},
			Weight:       1.0,
		},
		{
			Type:         "compliance",
			Priority:     1,
			Timeout:      time.Minute * 8,
			MaxRetries:   3,
			Required:     true,
			Dependencies: []string{},
			Weight:       1.0,
		},
	}

	// Add cost optimization based on infrastructure analysis
	if shouldIncludeCostAgent(infra) {
		baseConfigs = append(baseConfigs, AgentConfigV3{
			Type:         "cost-optimization",
			Priority:     2,
			Timeout:      time.Minute * 5,
			MaxRetries:   2,
			Required:     false,
			Dependencies: []string{"security", "compliance"},
			Weight:       0.8,
		})
	}

	// Add specialized agents based on compliance type
	switch request.ComplianceType {
	case "SOC2":
		baseConfigs = append(baseConfigs, AgentConfigV3{
			Type:         "soc2-audit",
			Priority:     2,
			Timeout:      time.Minute * 7,
			MaxRetries:   2,
			Required:     false,
			Dependencies: []string{"security", "compliance"},
			Weight:       0.9,
		})
	case "GDPR":
		baseConfigs = append(baseConfigs, AgentConfigV3{
			Type:         "gdpr-privacy",
			Priority:     2,
			Timeout:      time.Minute * 6,
			MaxRetries:   2,
			Required:     false,
			Dependencies: []string{"security", "compliance"},
			Weight:       0.95,
		})
	}

	return baseConfigs
}

func executeWithDynamicSchedulingV3(ctx workflow.Context, configs []AgentConfigV3, graph DependencyGraph, infra types.InfrastructureResult, circuitBreaker *AdaptiveCircuitBreaker, rateLimiter *AdaptiveRateLimiter) ([]types.AgentResult, *types.OrchestrationMetrics, error) {
	logger := workflow.GetLogger(ctx)

	metrics := &types.OrchestrationMetrics{
		MaxConcurrencyUsed: 0,
		TotalExecutionTime: 0,
		AgentExecutionOrder: []string{},
		ResourceUsage: types.ResourceUsageV3{},
	}

	// Execute in topological order respecting dependencies
	executed := make(map[string]bool)
	results := make(map[string]types.AgentResult)

	for len(executed) < len(configs) {
		// Find ready agents (dependencies satisfied)
		readyAgents := []AgentConfigV3{}
		for _, config := range configs {
			if executed[config.Type] {
				continue
			}

			// Check if all dependencies are satisfied
			depsSatisfied := true
			for _, dep := range config.Dependencies {
				if !executed[dep] {
					depsSatisfied = false
					break
				}
			}

			if depsSatisfied {
				readyAgents = append(readyAgents, config)
			}
		}

		if len(readyAgents) == 0 {
			break // No more agents can be executed
		}

		// Sort by priority for optimal execution order
		sort.Slice(readyAgents, func(i, j int) bool {
			return readyAgents[i].Priority < readyAgents[j].Priority
		})

		// Execute ready agents with adaptive concurrency
		concurrentLimit := calculateAdaptiveConcurrencyLimit(len(readyAgents), rateLimiter)
		if concurrentLimit > metrics.MaxConcurrencyUsed {
			metrics.MaxConcurrencyUsed = concurrentLimit
		}

		futures := make([]workflow.Future, 0, concurrentLimit)
		executingConfigs := make([]AgentConfigV3, 0, concurrentLimit)

		for i, config := range readyAgents {
			if i >= concurrentLimit {
				break
			}

			if !rateLimiter.Allow() {
				logger.Warn("Rate limit exceeded, reducing concurrency")
				break
			}

			future := executeAgentWithAdaptiveConfigV3(ctx, config, infra, circuitBreaker)
			futures = append(futures, future)
			executingConfigs = append(executingConfigs, config)
			metrics.AgentExecutionOrder = append(metrics.AgentExecutionOrder, config.Type)
		}

		// Collect results
		for i, future := range futures {
			config := executingConfigs[i]
			var result types.AgentResult
			err := future.Get(ctx, &result)

			if err != nil {
				if config.Required {
					return nil, nil, fmt.Errorf("required agent %s failed: %w", config.Type, err)
				}
				logger.Warn("Optional agent failed", "agent", config.Type, "error", err)
				continue
			}

			results[config.Type] = result
			executed[config.Type] = true
		}
	}

	// Convert map to slice
	var finalResults []types.AgentResult
	for _, result := range results {
		finalResults = append(finalResults, result)
	}

	return finalResults, metrics, nil
}

// Advanced calculation functions

func calculateBayesianConfidence(agentResults []types.AgentResult) types.ConfidenceInterval {
	if len(agentResults) == 0 {
		return types.ConfidenceInterval{ConfidenceScore: 0, LowerBound: 0, UpperBound: 0}
	}

	scores := make([]float64, len(agentResults))
	for i, result := range agentResults {
		scores[i] = result.Score
	}

	mean, variance := calculateMeanAndVariance(scores)
	stdDev := math.Sqrt(variance)

	// Bayesian confidence interval
	confidenceLevel := 0.95
	zScore := 1.96 // For 95% confidence
	marginOfError := zScore * (stdDev / math.Sqrt(float64(len(scores))))

	lowerBound := mean - marginOfError
	upperBound := mean + marginOfError

	// Confidence score based on interval width and consensus
	intervalWidth := upperBound - lowerBound
	confidenceScore := 100.0 - (intervalWidth * 2.5) // Penalize wide intervals

	if confidenceScore < 0 {
		confidenceScore = 0
	}
	if confidenceScore > 100 {
		confidenceScore = 100
	}

	return types.ConfidenceInterval{
		ConfidenceScore: confidenceScore,
		LowerBound:      math.Max(0, lowerBound),
		UpperBound:      math.Min(100, upperBound),
	}
}

func calculateAdaptiveConcurrencyLimit(readyCount int, rateLimiter *AdaptiveRateLimiter) int {
	baseLimit := 3 // Default concurrency

	// Adjust based on rate limiter feedback
	if rateLimiter.IsThrottled() {
		baseLimit = 1
	} else if rateLimiter.GetSuccessRate() > 0.9 {
		baseLimit = 5
	}

	// Don't exceed ready agents
	if baseLimit > readyCount {
		baseLimit = readyCount
	}

	return baseLimit
}

func determineAdaptiveApprovalStatusV3(aggregated types.AggregatedResult, request types.ComplianceRequest) bool {
	if aggregated.HumanReviewResult != nil {
		return aggregated.HumanReviewResult.Approved
	}

	// Adaptive auto-approval logic
	autoApproveThreshold := request.RequiredScore
	if aggregated.ConfidenceScore > 90.0 {
		autoApproveThreshold -= 5.0 // Lower threshold for high confidence
	}

	return aggregated.OverallScore >= autoApproveThreshold &&
		aggregated.RiskLevel != "Critical" &&
		aggregated.ConfidenceScore >= 75.0
}

// Adaptive Circuit Breaker implementation
type AdaptiveCircuitBreaker struct {
	Name           string
	FailureCount   int
	SuccessCount   int
	State          string // "closed", "open", "half-open"
	NextAttemptAt  time.Time
	FailureThreshold int
	Timeout         time.Duration
}

func NewAdaptiveCircuitBreaker(name string, failureThreshold int, timeout time.Duration) *AdaptiveCircuitBreaker {
	return &AdaptiveCircuitBreaker{
		Name:             name,
		State:            "closed",
		FailureThreshold: failureThreshold,
		Timeout:          timeout,
	}
}

func (cb *AdaptiveCircuitBreaker) ShouldAttempt() bool {
	now := time.Now()

	switch cb.State {
	case "closed":
		return true
	case "open":
		if now.After(cb.NextAttemptAt) {
			cb.State = "half-open"
			return true
		}
		return false
	case "half-open":
		return true
	default:
		return false
	}
}

func (cb *AdaptiveCircuitBreaker) RecordSuccess() {
	cb.SuccessCount++
	cb.FailureCount = 0

	if cb.State == "half-open" {
		cb.State = "closed"
	}
}

func (cb *AdaptiveCircuitBreaker) RecordFailure() {
	cb.FailureCount++

	if cb.FailureCount >= cb.FailureThreshold {
		cb.State = "open"
		cb.NextAttemptAt = time.Now().Add(cb.Timeout)
	}
}

// Adaptive Rate Limiter implementation
type AdaptiveRateLimiter struct {
	RequestsPerMinute float64
	WindowDuration    time.Duration
	RequestCount      int
	WindowStart       time.Time
	SuccessCount      int
	TotalCount        int
}

func NewAdaptiveRateLimiter(requestsPerMinute float64, windowDuration time.Duration) *AdaptiveRateLimiter {
	return &AdaptiveRateLimiter{
		RequestsPerMinute: requestsPerMinute,
		WindowDuration:    windowDuration,
		WindowStart:       time.Now(),
	}
}

func (rl *AdaptiveRateLimiter) Allow() bool {
	now := time.Now()

	// Reset window if expired
	if now.Sub(rl.WindowStart) >= rl.WindowDuration {
		rl.WindowStart = now
		rl.RequestCount = 0
		rl.SuccessCount = 0
		rl.TotalCount = 0
	}

	// Check rate limit
	currentRate := float64(rl.RequestCount) / rl.WindowDuration.Minutes() * 60
	if currentRate >= rl.RequestsPerMinute {
		return false
	}

	rl.RequestCount++
	rl.TotalCount++
	return true
}

func (rl *AdaptiveRateLimiter) IsThrottled() bool {
	currentRate := float64(rl.RequestCount) / time.Since(rl.WindowStart).Minutes() * 60
	return currentRate >= rl.RequestsPerMinute*0.8 // 80% of limit
}

func (rl *AdaptiveRateLimiter) GetSuccessRate() float64 {
	if rl.TotalCount == 0 {
		return 1.0
	}
	return float64(rl.SuccessCount) / float64(rl.TotalCount)
}

func (rl *AdaptiveRateLimiter) RecordSuccess() {
	rl.SuccessCount++
}

func (rl *AdaptiveRateLimiter) RecordFailure() {
	// Failures don't increment success count
}

// Dependency Graph implementation
type DependencyGraph struct {
	Nodes map[string][]string // node -> dependencies
}

func buildDependencyGraph(configs []AgentConfigV3) DependencyGraph {
	graph := DependencyGraph{
		Nodes: make(map[string][]string),
	}

	for _, config := range configs {
		graph.Nodes[config.Type] = config.Dependencies
	}

	return graph
}

// Enhanced agent configuration
type AgentConfigV3 struct {
	Type         string
	Priority     int
	Timeout      time.Duration
	MaxRetries   int
	Required     bool
	Dependencies []string
	Weight       float64 // For weighted aggregation
}

// Enhanced activity function
func executeAgentWithAdaptiveConfigV3(ctx workflow.Context, config AgentConfigV3, infra types.InfrastructureResult, circuitBreaker *AdaptiveCircuitBreaker) workflow.Future {
	// Adaptive timeout based on circuit breaker state
	timeout := config.Timeout
	if circuitBreaker.State == "half-open" {
		timeout = timeout * 2 // Double timeout when recovering
	}

	agentOptions := workflow.ActivityOptions{
		StartToCloseTimeout: timeout,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 5,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    int32(config.MaxRetries),
		},
		HeartbeatTimeout: time.Minute,
	}

	agentCtx := workflow.WithActivityOptions(ctx, agentOptions)

	// Enhanced agent execution with adaptive parameters
	switch config.Type {
	case "security":
		return workflow.ExecuteActivity(agentCtx, activities.AdaptiveSecurityAgentActivity, types.AdaptiveAgentRequest{
			Infrastructure: infra,
			Config:         config,
		})
	case "compliance":
		return workflow.ExecuteActivity(agentCtx, activities.AdaptiveComplianceAgentActivity, types.AdaptiveAgentRequest{
			Infrastructure: infra,
			Config:         config,
		})
	case "cost-optimization":
		return workflow.ExecuteActivity(agentCtx, activities.AdaptiveCostOptimizationAgentActivity, types.AdaptiveAgentRequest{
			Infrastructure: infra,
			Config:         config,
		})
	case "soc2-audit":
		return workflow.ExecuteActivity(agentCtx, activities.SOC2AuditAgentActivity, types.AdaptiveAgentRequest{
			Infrastructure: infra,
			Config:         config,
		})
	case "gdpr-privacy":
		return workflow.ExecuteActivity(agentCtx, activities.GDPRPrivacyAgentActivity, types.AdaptiveAgentRequest{
			Infrastructure: infra,
			Config:         config,
		})
	default:
		panic(fmt.Sprintf("Unknown agent type: %s", config.Type))
	}
}

// Utility functions
func calculateMeanAndVariance(scores []float64) (float64, float64) {
	if len(scores) == 0 {
		return 0, 0
	}

	sum := 0.0
	for _, score := range scores {
		sum += score
	}
	mean := sum / float64(len(scores))

	variance := 0.0
	for _, score := range scores {
		variance += math.Pow(score-mean, 2)
	}
	variance /= float64(len(scores))

	return mean, variance
}

func shouldIncludeCostAgent(infra types.InfrastructureResult) bool {
	// Include cost agent for cloud resources with significant usage
	if infra.Properties != nil {
		if resourceType, ok := infra.Properties["resourceType"].(string); ok {
			return resourceType == "EC2" || resourceType == "VM" || resourceType == "AKS" || resourceType == "EKS"
		}
	}
	return false
}

func calculateConsensusLevel(agentResults []types.AgentResult) float64 {
	if len(agentResults) < 2 {
		return 100.0
	}

	scores := make([]float64, len(agentResults))
	for i, result := range agentResults {
		scores[i] = result.Score
	}

	_, variance := calculateMeanAndVariance(scores)
	consensus := 100.0 - variance*0.5 // Lower variance = higher consensus

	if consensus < 0 {
		consensus = 0
	}
	if consensus > 100 {
		consensus = 100
	}

	return consensus
}

func performEnsembleAggregation(agentResults []types.AgentResult, request types.ComplianceRequest) types.AggregatedResult {
	if len(agentResults) == 0 {
		return types.AggregatedResult{}
	}

	// Weighted ensemble aggregation
	totalWeight := 0.0
	weightedSum := 0.0

	for _, result := range agentResults {
		weight := 1.0 // Default weight
		switch result.AgentType {
		case "security":
			weight = 1.2
		case "compliance":
			weight = 1.1
		case "cost-optimization":
			weight = 0.9
		}

		weightedSum += result.Score * weight
		totalWeight += weight
	}

	overallScore := weightedSum / totalWeight

	// Determine risk level based on weighted score and individual agent findings
	riskLevel := "Low"
	if overallScore < 70.0 {
		riskLevel = "High"
	} else if overallScore < 85.0 {
		riskLevel = "Medium"
	}

	return types.AggregatedResult{
		OverallScore: overallScore,
		RiskLevel:    riskLevel,
	}
}

func applyMLRiskAssessment(ensemble types.AggregatedResult, agentResults []types.AgentResult, confidence types.ConfidenceInterval) types.MLRiskAssessment {
	riskFactors := []string{}
	recommendations := []string{}

	// Analyze score variance
	scoreVariance := calculateScoreVariance(agentResults)
	if scoreVariance > 20.0 {
		riskFactors = append(riskFactors, "High agent disagreement")
		recommendations = append(recommendations, "Consider human review due to agent consensus issues")
	}

	// Analyze confidence interval
	intervalWidth := confidence.UpperBound - confidence.LowerBound
	if intervalWidth > 30.0 {
		riskFactors = append(riskFactors, "Wide confidence interval")
		recommendations = append(recommendations, "Gather additional evidence to narrow confidence range")
	}

	// Risk level adjustment based on ML analysis
	adjustedRiskLevel := ensemble.RiskLevel
	if confidence.ConfidenceScore < 70.0 && ensemble.RiskLevel == "Low" {
		adjustedRiskLevel = "Medium"
		riskFactors = append(riskFactors, "Low confidence despite good score")
		recommendations = append(recommendations, "Verify results with additional validation")
	}

	return types.MLRiskAssessment{
		RiskLevel:       adjustedRiskLevel,
		RiskFactors:     riskFactors,
		Recommendations: recommendations,
	}
}

func calculateScoreVariance(agentResults []types.AgentResult) float64 {
	if len(agentResults) < 2 {
		return 0.0
	}

	scores := make([]float64, len(agentResults))
	for i, result := range agentResults {
		scores[i] = result.Score
	}

	_, variance := calculateMeanAndVariance(scores)
	return variance
}

func calculatePredictiveRisk(aggregated types.AggregatedResult, agentResults []types.AgentResult, request types.ComplianceRequest) types.PredictiveRisk {
	// Simple predictive model based on historical patterns
	baseRisk := 50.0

	// Adjust based on score
	if aggregated.OverallScore < 60.0 {
		baseRisk += 30.0
	} else if aggregated.OverallScore > 90.0 {
		baseRisk -= 20.0
	}

	// Adjust based on variance
	variance := calculateScoreVariance(agentResults)
	baseRisk += variance * 0.3

	// Adjust based on risk level
	switch aggregated.RiskLevel {
	case "Critical":
		baseRisk += 40.0
	case "High":
		baseRisk += 25.0
	case "Medium":
		baseRisk += 10.0
	case "Low":
		baseRisk -= 10.0
	}

	// Cap risk between 0 and 100
	if baseRisk < 0 {
		baseRisk = 0
	}
	if baseRisk > 100 {
		baseRisk = 100
	}

	urgency := "low"
	if baseRisk > 80.0 {
		urgency = "critical"
	} else if baseRisk > 60.0 {
		urgency = "high"
	} else if baseRisk > 40.0 {
		urgency = "medium"
	}

	return types.PredictiveRisk{
		Score:   baseRisk,
		Level:   aggregated.RiskLevel,
		Urgency: urgency,
	}
}

func calculateAdaptiveThreshold(metrics types.PerformanceMetrics) float64 {
	// Start with default threshold
	threshold := 80.0

	// Adjust based on historical performance
	if metrics.ActualDuration > metrics.EstimatedDuration*1.5 {
		threshold += 5.0 // Increase threshold if consistently taking longer
	} else if metrics.ActualDuration < metrics.EstimatedDuration*0.8 {
		threshold -= 5.0 // Decrease threshold if consistently faster
	}

	return threshold
}

func analyzeTemporalConsensus(agentResults []types.AgentResult) float64 {
	// For now, use simple consensus calculation
	// In a real implementation, this would consider temporal patterns
	return calculateConsensusLevel(agentResults)
}

func determineReviewPriorityV3(risk types.PredictiveRisk, aggregated types.AggregatedResult) string {
	if risk.Urgency == "critical" || aggregated.RiskLevel == "Critical" {
		return "critical"
	}
	if risk.Urgency == "high" || aggregated.RiskLevel == "High" {
		return "high"
	}
	if risk.Score > 60.0 || aggregated.RiskLevel == "Medium" {
		return "medium"
	}
	return "low"
}

func calculateAdaptiveTimeout(priority string, urgency string) time.Duration {
	baseTimeout := time.Hour * 24 // Default

	switch priority {
	case "critical":
		baseTimeout = time.Hour * 4
	case "high":
		baseTimeout = time.Hour * 12
	case "medium":
		baseTimeout = time.Hour * 24
	case "low":
		baseTimeout = time.Hour * 48
	}

	// Adjust based on urgency
	switch urgency {
	case "critical":
		baseTimeout = baseTimeout / 2
	case "high":
		baseTimeout = baseTimeout * 3 / 4
	}

	return baseTimeout
}

func intelligentInfrastructureValidationV3(result types.InfrastructureResult) error {
	// Enhanced validation logic
	if result.ResourceID == "" {
		return fmt.Errorf("resource ID is required")
	}
	if result.ResourceType == "" {
		return fmt.Errorf("resource type is required")
	}
	if result.Region == "" {
		return fmt.Errorf("region is required")
	}

	// Additional business logic validation
	if result.Properties != nil {
		if env, ok := result.Properties["environment"].(string); ok && env == "production" {
			// Production resources require stricter validation
			if result.Properties["owner"] == nil {
				return fmt.Errorf("production resources must have an owner specified")
			}
			if result.Properties["backupEnabled"] == nil {
				return fmt.Errorf("production resources must have backup configuration")
			}
		}
	}

	return nil
}

func generatePredictiveInsights(aggregated types.AggregatedResult, metrics *types.WorkflowMetricsV3, orchestrationMetrics *types.OrchestrationMetrics) []types.PredictiveInsight {
	insights := []types.PredictiveInsight{}

	// Performance insights
	if metrics.Duration > metrics.PerformanceMetrics.EstimatedDuration*1.2 {
		insights = append(insights, types.PredictiveInsight{
			Type:        "performance",
			Severity:    "medium",
			Description: "Workflow execution took longer than estimated",
			Recommendation: "Consider optimizing agent concurrency or reducing timeouts",
		})
	}

	// Reliability insights
	if metrics.ErrorCount > 0 {
		insights = append(insights, types.PredictiveInsight{
			Type:        "reliability",
			Severity:    "high",
			Description: fmt.Sprintf("Workflow encountered %d errors", metrics.ErrorCount),
			Recommendation: "Review error patterns and consider circuit breaker adjustments",
		})
	}

	// Concurrency insights
	if orchestrationMetrics.MaxConcurrencyUsed < orchestrationMetrics.ResourceUsage.MaxConcurrency {
		insights = append(insights, types.PredictiveInsight{
			Type:        "optimization",
			Severity:    "low",
			Description: "Underutilized concurrency capacity detected",
			Recommendation: "Consider increasing parallel agent execution",
		})
	}

	// Confidence insights
	if aggregated.ConfidenceScore < 80.0 {
		insights = append(insights, types.PredictiveInsight{
			Type:        "confidence",
			Severity:    "medium",
			Description: "Low confidence score in aggregated results",
			Recommendation: "Consider additional agent validation or human review",
		})
	}

	return insights
}

func generateActionableRecommendations(aggregated types.AggregatedResult, request types.ComplianceRequest) []types.ActionableRecommendation {
	recommendations := []types.ActionableRecommendation{}

	// Generate recommendations based on findings
	if aggregated.OverallScore < 80.0 {
		recommendations = append(recommendations, types.ActionableRecommendation{
			Priority:      "high",
			Category:      "compliance",
			Description:   "Address identified compliance gaps",
			EstimatedEffort: "2-4 weeks",
			BusinessImpact: "high",
		})
	}

	if aggregated.RiskLevel == "High" || aggregated.RiskLevel == "Critical" {
		recommendations = append(recommendations, types.ActionableRecommendation{
			Priority:      "critical",
			Category:      "security",
			Description:   "Implement immediate security remediation",
			EstimatedEffort: "1-2 weeks",
			BusinessImpact: "critical",
		})
	}

	// Cost optimization recommendations
	if aggregated.OverallScore > 85.0 {
		recommendations = append(recommendations, types.ActionableRecommendation{
			Priority:      "medium",
			Category:      "optimization",
			Description:   "Review cost optimization opportunities",
			EstimatedEffort: "1 week",
			BusinessImpact: "medium",
		})
	}

	return recommendations
}

func generateAdaptiveFallbackReviewV3(aggregated types.AggregatedResult, request types.ComplianceRequest, priority string) types.HumanReviewResultV3 {
	autoApprove := false

	switch priority {
	case "critical":
		autoApprove = false // Never auto-approve critical reviews
	case "high":
		autoApprove = aggregated.OverallScore >= 95.0 && aggregated.ConfidenceScore >= 90.0
	case "medium":
		autoApprove = aggregated.OverallScore >= 90.0 && aggregated.ConfidenceScore >= 85.0
	case "low":
		autoApprove = aggregated.OverallScore >= 85.0 && aggregated.ConfidenceScore >= 80.0
	}

	decision := "System fallback: Auto-rejected due to timeout"
	if autoApprove {
		decision = "System fallback: Auto-approved based on high confidence and score"
	}

	return types.HumanReviewResultV3{
		ReviewerID:     "system-adaptive-fallback-v3",
		Approved:       autoApprove,
		Decision:       decision,
		Comments:       fmt.Sprintf("Adaptive fallback - Score: %.1f, Confidence: %.1f, Priority: %s", aggregated.OverallScore, aggregated.ConfidenceScore, priority),
		ReviewedAt:     workflow.Now(ctx),
		ReviewDuration: 0,
		Confidence:     aggregated.ConfidenceScore,
		EscalationLevel: func() int {
			switch priority {
			case "critical": return 3
			case "high": return 2
			case "medium": return 1
			default: return 0
			}
		}(),
	}
}
