package monitoring

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

// MetricsCollector interface for collecting and storing metrics
type MetricsCollector interface {
	Start(ctx context.Context)
	Stop()
	GetSummary() map[string]interface{}
	GetAll() map[string]interface{}
	RecordWorkflowMetric(metric WorkflowMetric)
	RecordActivityMetric(metric ActivityMetric)
	RecordSystemMetric(metric SystemMetric)
	GetWorkflowMetrics(workflowID string) []WorkflowMetric
	GetActivityMetrics(activityID string) []ActivityMetric
}

// Enhanced metrics implementation
type EnhancedMetricsCollector struct {
	ctx          context.Context
	cancel       context.CancelFunc
	isRunning    bool
	mu           sync.RWMutex
	
	// Metrics storage
	workflowMetrics map[string][]WorkflowMetric
	activityMetrics map[string][]ActivityMetric
	systemMetrics   []SystemMetric
	
	// Aggregated metrics
	workflowStats map[string]*WorkflowStats
	activityStats map[string]*ActivityStats
	systemStats    *SystemStats
	
	// Configuration
	config MetricsConfig
	
	// Channels for real-time processing
	workflowChan chan WorkflowMetric
	activityChan chan ActivityMetric
	systemChan   chan SystemMetric
}

type MetricsConfig struct {
	CollectionInterval    time.Duration `json:"collectionInterval"`
	RetentionPeriod      time.Duration `json:"retentionPeriod"`
	MaxMetricsPerEntity  int           `json:"maxMetricsPerEntity"`
	EnableRealTime       bool          `json:"enableRealTime"`
	EnableAggregation    bool          `json:"enableAggregation"`
	ExportFormat         string        `json:"exportFormat"`
	ExportInterval       time.Duration `json:"exportInterval"`
}

type WorkflowMetric struct {
	ID           string                 `json:"id"`
	WorkflowID   string                 `json:"workflowId"`
	WorkflowType string                 `json:"workflowType"`
	Timestamp    time.Time              `json:"timestamp"`
	MetricType   string                 `json:"metricType"`
	Value        float64                `json:"value"`
	Unit         string                 `json:"unit"`
	Tags         map[string]string      `json:"tags"`
	Metadata     map[string]interface{} `json:"metadata"`
}

type ActivityMetric struct {
	ID           string                 `json:"id"`
	ActivityID   string                 `json:"activityId"`
	ActivityType string                 `json:"activityType"`
	WorkflowID   string                 `json:"workflowId"`
	Timestamp    time.Time              `json:"timestamp"`
	MetricType   string                 `json:"metricType"`
	Value        float64                `json:"value"`
	Unit         string                 `json:"unit"`
	Status       string                 `json:"status"`
	Duration     time.Duration          `json:"duration"`
	Tags         map[string]string      `json:"tags"`
	Metadata     map[string]interface{} `json:"metadata"`
}

type SystemMetric struct {
	ID        string                 `json:"id"`
	Timestamp time.Time              `json:"timestamp"`
	MetricType string                 `json:"metricType"`
	Value     float64                `json:"value"`
	Unit      string                 `json:"unit"`
	Source    string                 `json:"source"`
	Tags      map[string]string      `json:"tags"`
	Metadata  map[string]interface{} `json:"metadata"`
}

type WorkflowStats struct {
	WorkflowID         string                 `json:"workflowId"`
	WorkflowType       string                 `json:"workflowType"`
	TotalExecutions    int64                  `json:"totalExecutions"`
	SuccessfulRuns     int64                  `json:"successfulRuns"`
	FailedRuns         int64                  `json:"failedRuns"`
	AverageDuration    time.Duration          `json:"averageDuration"`
	MinDuration        time.Duration          `json:"minDuration"`
	MaxDuration        time.Duration          `json:"maxDuration"`
	TotalCost          float64                `json:"totalCost"`
	AverageCost        float64                `json:"averageCost"`
	SuccessRate        float64                `json:"successRate"`
	ErrorRate          float64                `json:"errorRate"`
	LastExecution      time.Time              `json:"lastExecution"`
	FirstExecution     time.Time              `json:"firstExecution"`
	CustomMetrics      map[string]interface{} `json:"customMetrics"`
}

type ActivityStats struct {
	ActivityType       string                 `json:"activityType"`
	TotalExecutions    int64                  `json:"totalExecutions"`
	SuccessfulRuns     int64                  `json:"successfulRuns"`
	FailedRuns         int64                  `json:"failedRuns"`
	AverageDuration    time.Duration          `json:"averageDuration"`
	MinDuration        time.Duration          `json:"minDuration"`
	MaxDuration        time.Duration          `json:"maxDuration"`
	RetryCount         int64                  `json:"retryCount"`
	SuccessRate        float64                `json:"successRate"`
	ErrorRate          float64                `json:"errorRate"`
	LastExecution      time.Time              `json:"lastExecution"`
	FirstExecution     time.Time              `json:"firstExecution"`
	CustomMetrics      map[string]interface{} `json:"customMetrics"`
}

type SystemStats struct {
	Goroutines         int                    `json:"goroutines"`
	MemoryAlloc        uint64                 `json:"memoryAlloc"`
	MemoryTotalAlloc   uint64                 `json:"memoryTotalAlloc"`
	MemorySys          uint64                 `json:"memorySys"`
	NumGC              uint32                 `json:"numGC"`
	GCPauseTotal       time.Duration          `json:"gcPauseTotal"`
	CPUUsage           float64                `json:"cpuUsage"`
	LastUpdated        time.Time              `json:"lastUpdated"`
	HostInfo           map[string]interface{} `json:"hostInfo"`
	CustomMetrics      map[string]interface{} `json:"customMetrics"`
}

// NewEnhancedMetricsCollector creates a new enhanced metrics collector
func NewEnhancedMetricsCollector(config MetricsConfig) *EnhancedMetricsCollector {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &EnhancedMetricsCollector{
		ctx:             ctx,
		cancel:          cancel,
		workflowMetrics: make(map[string][]WorkflowMetric),
		activityMetrics: make(map[string][]ActivityMetric),
		systemMetrics:   make([]SystemMetric, 0),
		workflowStats:   make(map[string]*WorkflowStats),
		activityStats:   make(map[string]*ActivityStats),
		systemStats:     &SystemStats{},
		config:          config,
		workflowChan:    make(chan WorkflowMetric, 1000),
		activityChan:    make(chan ActivityMetric, 1000),
		systemChan:      make(chan SystemMetric, 1000),
	}
}

// Start begins the metrics collection process
func (emc *EnhancedMetricsCollector) Start(ctx context.Context) {
	emc.mu.Lock()
	defer emc.mu.Unlock()
	
	if emc.isRunning {
		return
	}
	
	emc.isRunning = true
	emc.ctx = ctx
	
	// Start real-time metrics processing if enabled
	if emc.config.EnableRealTime {
		go emc.processWorkflowMetrics()
		go emc.processActivityMetrics()
		go emc.processSystemMetrics()
	}
	
	// Start periodic system metrics collection
	go emc.collectSystemMetrics()
	
	// Start aggregation if enabled
	if emc.config.EnableAggregation {
		go emc.aggregateMetrics()
	}
	
	// Start cleanup goroutine
	go emc.cleanupOldMetrics()
	
	log.Printf("Enhanced metrics collector started with config: %+v", emc.config)
}

// Stop stops the metrics collection
func (emc *EnhancedMetricsCollector) Stop() {
	emc.mu.Lock()
	defer emc.mu.Unlock()
	
	if !emc.isRunning {
		return
	}
	
	emc.cancel()
	emc.isRunning = false
	
	// Close channels
	close(emc.workflowChan)
	close(emc.activityChan)
	close(emc.systemChan)
	
	log.Println("Enhanced metrics collector stopped")
}

// RecordWorkflowMetric records a workflow metric
func (emc *EnhancedMetricsCollector) RecordWorkflowMetric(metric WorkflowMetric) {
	if emc.config.EnableRealTime {
		select {
		case emc.workflowChan <- metric:
		default:
			log.Printf("Workflow metrics channel full, dropping metric: %s", metric.ID)
		}
	} else {
		emc.storeWorkflowMetric(metric)
	}
}

// RecordActivityMetric records an activity metric
func (emc *EnhancedMetricsCollector) RecordActivityMetric(metric ActivityMetric) {
	if emc.config.EnableRealTime {
		select {
		case emc.activityChan <- metric:
		default:
			log.Printf("Activity metrics channel full, dropping metric: %s", metric.ID)
		}
	} else {
		emc.storeActivityMetric(metric)
	}
}

// RecordSystemMetric records a system metric
func (emc *EnhancedMetricsCollector) RecordSystemMetric(metric SystemMetric) {
	if emc.config.EnableRealTime {
		select {
		case emc.systemChan <- metric:
		default:
			log.Printf("System metrics channel full, dropping metric: %s", metric.ID)
		}
	} else {
		emc.storeSystemMetric(metric)
	}
}

// GetSummary returns a summary of all metrics
func (emc *EnhancedMetricsCollector) GetSummary() map[string]interface{} {
	emc.mu.RLock()
	defer emc.mu.RUnlock()
	
	summary := make(map[string]interface{})
	
	// Workflow summary
	workflowSummary := map[string]interface{}{
		"totalWorkflows":       len(emc.workflowMetrics),
		"totalExecutions":      emc.calculateTotalWorkflowExecutions(),
		"averageSuccessRate":   emc.calculateAverageSuccessRate(),
		"averageDuration":      emc.calculateAverageDuration(),
		"activeWorkflows":      emc.getActiveWorkflowCount(),
	}
	summary["workflows"] = workflowSummary
	
	// Activity summary
	activitySummary := map[string]interface{}{
		"totalActivities":      len(emc.activityMetrics),
		"totalExecutions":      emc.calculateTotalActivityExecutions(),
		"averageSuccessRate":   emc.calculateAverageActivitySuccessRate(),
		"averageDuration":      emc.calculateAverageActivityDuration(),
	}
	summary["activities"] = activitySummary
	
	// System summary
	summary["system"] = emc.systemStats
	
	// Collection info
	summary["collection"] = map[string]interface{}{
		"isRunning":       emc.isRunning,
		"collectionInterval": emc.config.CollectionInterval,
		"retentionPeriod": emc.config.RetentionPeriod,
		"lastUpdated":     time.Now(),
	}
	
	return summary
}

// GetAll returns all collected metrics
func (emc *EnhancedMetricsCollector) GetAll() map[string]interface{} {
	emc.mu.RLock()
	defer emc.mu.RUnlock()
	
	all := make(map[string]interface{})
	all["workflowMetrics"] = emc.workflowMetrics
	all["activityMetrics"] = emc.activityMetrics
	all["systemMetrics"] = emc.systemMetrics
	all["workflowStats"] = emc.workflowStats
	all["activityStats"] = emc.activityStats
	all["systemStats"] = emc.systemStats
	
	return all
}

// GetWorkflowMetrics returns metrics for a specific workflow
func (emc *EnhancedMetricsCollector) GetWorkflowMetrics(workflowID string) []WorkflowMetric {
	emc.mu.RLock()
	defer emc.mu.RUnlock()
	
	return emc.workflowMetrics[workflowID]
}

// GetActivityMetrics returns metrics for a specific activity
func (emc *EnhancedMetricsCollector) GetActivityMetrics(activityID string) []ActivityMetric {
	emc.mu.RLock()
	defer emc.mu.RUnlock()
	
	return emc.activityMetrics[activityID]
}

// Private methods

func (emc *EnhancedMetricsCollector) processWorkflowMetrics() {
	for {
		select {
		case <-emc.ctx.Done():
			return
		case metric, ok := <-emc.workflowChan:
			if !ok {
				return
			}
			emc.storeWorkflowMetric(metric)
		}
	}
}

func (emc *EnhancedMetricsCollector) processActivityMetrics() {
	for {
		select {
		case <-emc.ctx.Done():
			return
		case metric, ok := <-emc.activityChan:
			if !ok {
				return
			}
			emc.storeActivityMetric(metric)
		}
	}
}

func (emc *EnhancedMetricsCollector) processSystemMetrics() {
	for {
		select {
		case <-emc.ctx.Done():
			return
		case metric, ok := <-emc.systemChan:
			if !ok {
				return
			}
			emc.storeSystemMetric(metric)
		}
	}
}

func (emc *EnhancedMetricsCollector) storeWorkflowMetric(metric WorkflowMetric) {
	emc.mu.Lock()
	defer emc.mu.Unlock()
	
	emc.workflowMetrics[metric.WorkflowID] = append(emc.workflowMetrics[metric.WorkflowID], metric)
	
	// Limit metrics per entity
	if len(emc.workflowMetrics[metric.WorkflowID]) > emc.config.MaxMetricsPerEntity {
		emc.workflowMetrics[metric.WorkflowID] = emc.workflowMetrics[metric.WorkflowID][1:]
	}
	
	// Update stats
	emc.updateWorkflowStats(metric)
}

func (emc *EnhancedMetricsCollector) storeActivityMetric(metric ActivityMetric) {
	emc.mu.Lock()
	defer emc.mu.Unlock()
	
	emc.activityMetrics[metric.ActivityID] = append(emc.activityMetrics[metric.ActivityID], metric)
	
	// Limit metrics per entity
	if len(emc.activityMetrics[metric.ActivityID]) > emc.config.MaxMetricsPerEntity {
		emc.activityMetrics[metric.ActivityID] = emc.activityMetrics[metric.ActivityID][1:]
	}
	
	// Update stats
	emc.updateActivityStats(metric)
}

func (emc *EnhancedMetricsCollector) storeSystemMetric(metric SystemMetric) {
	emc.mu.Lock()
	defer emc.mu.Unlock()
	
	emc.systemMetrics = append(emc.systemMetrics, metric)
	
	// Limit total system metrics
	if len(emc.systemMetrics) > emc.config.MaxMetricsPerEntity*10 {
		emc.systemMetrics = emc.systemMetrics[1:]
	}
}

func (emc *EnhancedMetricsCollector) collectSystemMetrics() {
	ticker := time.NewTicker(emc.config.CollectionInterval)
	defer ticker.Stop()
	
	for {
		select {
		case <-emc.ctx.Done():
			return
		case <-ticker.C:
			emc.collectCurrentSystemMetrics()
		}
	}
}

func (emc *EnhancedMetricsCollector) collectCurrentSystemMetrics() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	// Memory metrics
	emc.RecordSystemMetric(SystemMetric{
		ID:        fmt.Sprintf("mem-alloc-%d", time.Now().UnixNano()),
		Timestamp: time.Now(),
		MetricType: "memory_alloc",
		Value:     float64(m.Alloc),
		Unit:      "bytes",
		Source:    "runtime",
		Tags:      map[string]string{"type": "memory"},
	})
	
	emc.RecordSystemMetric(SystemMetric{
		ID:        fmt.Sprintf("mem-total-alloc-%d", time.Now().UnixNano()),
		Timestamp: time.Now(),
		MetricType: "memory_total_alloc",
		Value:     float64(m.TotalAlloc),
		Unit:      "bytes",
		Source:    "runtime",
		Tags:      map[string]string{"type": "memory"},
	})
	
	// Goroutine metrics
	emc.RecordSystemMetric(SystemMetric{
		ID:        fmt.Sprintf("goroutines-%d", time.Now().UnixNano()),
		Timestamp: time.Now(),
		MetricType: "goroutines",
		Value:     float64(runtime.NumGoroutine()),
		Unit:      "count",
		Source:    "runtime",
		Tags:      map[string]string{"type": "runtime"},
	})
	
	// GC metrics
	emc.RecordSystemMetric(SystemMetric{
		ID:        fmt.Sprintf("gc-pause-%d", time.Now().UnixNano()),
		Timestamp: time.Now(),
		MetricType: "gc_pause_total",
		Value:     float64(m.PauseTotalNs) / 1e9, // Convert to seconds
		Unit:      "seconds",
		Source:    "runtime",
		Tags:      map[string]string{"type": "gc"},
	})
	
	// Update system stats
	emc.updateSystemStats()
}

func (emc *EnhancedMetricsCollector) aggregateMetrics() {
	ticker := time.NewTicker(emc.config.ExportInterval)
	defer ticker.Stop()
	
	for {
		select {
		case <-emc.ctx.Done():
			return
		case <-ticker.C:
			emc.performAggregation()
		}
	}
}

func (emc *EnhancedMetricsCollector) performAggregation() {
	emc.mu.Lock()
	defer emc.mu.Unlock()
	
	// Aggregate workflow stats
	for workflowID, metrics := range emc.workflowMetrics {
		if len(metrics) == 0 {
			continue
		}
		
		stats := emc.workflowStats[workflowID]
		if stats == nil {
			stats = &WorkflowStats{
				WorkflowID:   workflowID,
				WorkflowType: metrics[0].WorkflowType,
				CustomMetrics: make(map[string]interface{}),
			}
			emc.workflowStats[workflowID] = stats
		}
		
		// Update aggregation logic here
		// This is a simplified version - in production, you'd want more sophisticated aggregation
		stats.LastExecution = time.Now()
		if stats.FirstExecution.IsZero() {
			stats.FirstExecution = metrics[0].Timestamp
		}
	}
	
	// Aggregate activity stats
	for activityID, metrics := range emc.activityMetrics {
		if len(metrics) == 0 {
			continue
		}
		
		stats := emc.activityStats[activityID]
		if stats == nil {
			stats = &ActivityStats{
				ActivityType: metrics[0].ActivityType,
				CustomMetrics: make(map[string]interface{}),
			}
			emc.activityStats[activityID] = stats
		}
		
		stats.LastExecution = time.Now()
		if stats.FirstExecution.IsZero() {
			stats.FirstExecution = metrics[0].Timestamp
		}
	}
}

func (emc *EnhancedMetricsCollector) cleanupOldMetrics() {
	ticker := time.NewTicker(time.Hour) // Run cleanup every hour
	defer ticker.Stop()
	
	for {
		select {
		case <-emc.ctx.Done():
			return
		case <-ticker.C:
			emc.performCleanup()
		}
	}
}

func (emc *EnhancedMetricsCollector) performCleanup() {
	emc.mu.Lock()
	defer emc.mu.Unlock()
	
	cutoff := time.Now().Add(-emc.config.RetentionPeriod)
	
	// Clean workflow metrics
	for workflowID, metrics := range emc.workflowMetrics {
		filtered := make([]WorkflowMetric, 0)
		for _, metric := range metrics {
			if metric.Timestamp.After(cutoff) {
				filtered = append(filtered, metric)
			}
		}
		emc.workflowMetrics[workflowID] = filtered
	}
	
	// Clean activity metrics
	for activityID, metrics := range emc.activityMetrics {
		filtered := make([]ActivityMetric, 0)
		for _, metric := range metrics {
			if metric.Timestamp.After(cutoff) {
				filtered = append(filtered, metric)
			}
		}
		emc.activityMetrics[activityID] = filtered
	}
	
	// Clean system metrics
	filtered := make([]SystemMetric, 0)
	for _, metric := range emc.systemMetrics {
		if metric.Timestamp.After(cutoff) {
			filtered = append(filtered, metric)
		}
	}
	emc.systemMetrics = filtered
	
	log.Printf("Cleaned up metrics older than %v", cutoff)
}

// Helper methods for calculations
func (emc *EnhancedMetricsCollector) calculateTotalWorkflowExecutions() int64 {
	var total int64
	for _, metrics := range emc.workflowMetrics {
		total += int64(len(metrics))
	}
	return total
}

func (emc *EnhancedMetricsCollector) calculateAverageSuccessRate() float64 {
	if len(emc.workflowStats) == 0 {
		return 0
	}
	
	var totalRate float64
	count := 0
	
	for _, stats := range emc.workflowStats {
		if stats.SuccessRate > 0 {
			totalRate += stats.SuccessRate
			count++
		}
	}
	
	if count == 0 {
		return 0
	}
	
	return totalRate / float64(count)
}

func (emc *EnhancedMetricsCollector) calculateAverageDuration() time.Duration {
	if len(emc.workflowStats) == 0 {
		return 0
	}
	
	var totalDuration time.Duration
	count := 0
	
	for _, stats := range emc.workflowStats {
		if stats.AverageDuration > 0 {
			totalDuration += stats.AverageDuration
			count++
		}
	}
	
	if count == 0 {
		return 0
	}
	
	return totalDuration / time.Duration(count)
}

func (emc *EnhancedMetricsCollector) getActiveWorkflowCount() int {
	// This would track currently running workflows
	// For now, return a placeholder
	return 0
}

func (emc *EnhancedMetricsCollector) calculateTotalActivityExecutions() int64 {
	var total int64
	for _, metrics := range emc.activityMetrics {
		total += int64(len(metrics))
	}
	return total
}

func (emc *EnhancedMetricsCollector) calculateAverageActivitySuccessRate() float64 {
	if len(emc.activityStats) == 0 {
		return 0
	}
	
	var totalRate float64
	count := 0
	
	for _, stats := range emc.activityStats {
		if stats.SuccessRate > 0 {
			totalRate += stats.SuccessRate
			count++
		}
	}
	
	if count == 0 {
		return 0
	}
	
	return totalRate / float64(count)
}

func (emc *EnhancedMetricsCollector) calculateAverageActivityDuration() time.Duration {
	if len(emc.activityStats) == 0 {
		return 0
	}
	
	var totalDuration time.Duration
	count := 0
	
	for _, stats := range emc.activityStats {
		if stats.AverageDuration > 0 {
			totalDuration += stats.AverageDuration
			count++
		}
	}
	
	if count == 0 {
		return 0
	}
	
	return totalDuration / time.Duration(count)
}

func (emc *EnhancedMetricsCollector) updateWorkflowStats(metric WorkflowMetric) {
	stats := emc.workflowStats[metric.WorkflowID]
	if stats == nil {
		stats = &WorkflowStats{
			WorkflowID:   metric.WorkflowID,
			WorkflowType: metric.WorkflowType,
			CustomMetrics: make(map[string]interface{}),
		}
		emc.workflowStats[metric.WorkflowID] = stats
	}
	
	// Update stats based on metric type
	switch metric.MetricType {
	case "execution":
		stats.TotalExecutions++
	case "success":
		stats.SuccessfulRuns++
	case "failure":
		stats.FailedRuns++
	case "duration":
		// Update duration statistics
		if stats.AverageDuration == 0 {
			stats.AverageDuration = time.Duration(metric.Value)
			stats.MinDuration = time.Duration(metric.Value)
			stats.MaxDuration = time.Duration(metric.Value)
		} else {
			avg := time.Duration(metric.Value)
			stats.AverageDuration = (stats.AverageDuration + avg) / 2
			if avg < stats.MinDuration {
				stats.MinDuration = avg
			}
			if avg > stats.MaxDuration {
				stats.MaxDuration = avg
			}
		}
	case "cost":
		stats.TotalCost += metric.Value
		if stats.TotalExecutions > 0 {
			stats.AverageCost = stats.TotalCost / float64(stats.TotalExecutions)
		}
	}
	
	// Calculate rates
	if stats.TotalExecutions > 0 {
		stats.SuccessRate = float64(stats.SuccessfulRuns) / float64(stats.TotalExecutions)
		stats.ErrorRate = float64(stats.FailedRuns) / float64(stats.TotalExecutions)
	}
	
	stats.LastExecution = metric.Timestamp
	if stats.FirstExecution.IsZero() {
		stats.FirstExecution = metric.Timestamp
	}
}

func (emc *EnhancedMetricsCollector) updateActivityStats(metric ActivityMetric) {
	stats := emc.activityStats[metric.ActivityType]
	if stats == nil {
		stats = &ActivityStats{
			ActivityType: metric.ActivityType,
			CustomMetrics: make(map[string]interface{}),
		}
		emc.activityStats[metric.ActivityType] = stats
	}
	
	// Update stats based on metric type
	switch metric.MetricType {
	case "execution":
		stats.TotalExecutions++
	case "success":
		stats.SuccessfulRuns++
	case "failure":
		stats.FailedRuns++
	case "duration":
		// Update duration statistics
		if stats.AverageDuration == 0 {
			stats.AverageDuration = metric.Duration
			stats.MinDuration = metric.Duration
			stats.MaxDuration = metric.Duration
		} else {
			avg := metric.Duration
			stats.AverageDuration = (stats.AverageDuration + avg) / 2
			if avg < stats.MinDuration {
				stats.MinDuration = avg
			}
			if avg > stats.MaxDuration {
				stats.MaxDuration = avg
			}
		}
	case "retry":
		stats.RetryCount++
	}
	
	// Calculate rates
	if stats.TotalExecutions > 0 {
		stats.SuccessRate = float64(stats.SuccessfulRuns) / float64(stats.TotalExecutions)
		stats.ErrorRate = float64(stats.FailedRuns) / float64(stats.TotalExecutions)
	}
	
	stats.LastExecution = metric.Timestamp
	if stats.FirstExecution.IsZero() {
		stats.FirstExecution = metric.Timestamp
	}
}

func (emc *EnhancedMetricsCollector) updateSystemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	emc.systemStats.Goroutines = runtime.NumGoroutine()
	emc.systemStats.MemoryAlloc = m.Alloc
	emc.systemStats.MemoryTotalAlloc = m.TotalAlloc
	emc.systemStats.MemorySys = m.Sys
	emc.systemStats.NumGC = m.NumGC
	emc.systemStats.GCPauseTotal = time.Duration(m.PauseTotalNs) * time.Nanosecond
	emc.systemStats.LastUpdated = time.Now()
}
