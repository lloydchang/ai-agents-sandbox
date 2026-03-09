package humanloop

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// HumanLoopManager interface for human-in-the-loop operations
type HumanLoopManager interface {
	Start(ctx context.Context)
	Stop()
	CreateTask(request TaskRequest) (*HumanTask, error)
	GetTask(taskID string) (*HumanTask, error)
	UpdateTask(taskID string, update TaskUpdate) (*HumanTask, error)
	AssignTask(taskID string, assignee string) error
	CompleteTask(taskID string, result TaskResult) error
	EscalateTask(taskID string, reason string) error
	GetTasks(filter TaskFilter) ([]*HumanTask, error)
	GetTaskHistory(taskID string) ([]*TaskHistoryEntry, error)
	NotifyAssignee(taskID string) error
}

// Enhanced human loop manager implementation
type EnhancedHumanLoopManager struct {
	ctx            context.Context
	cancel         context.CancelFunc
	isRunning      bool
	mu             sync.RWMutex
	
	// Task management
	tasks          map[string]*HumanTask
	taskHistory    map[string][]*TaskHistoryEntry
	taskQueue      map[string][]string // assignee -> taskIDs
	
	// Configuration
	config         HumanLoopConfig
	
	// Notification system
	notificationMgr NotificationManager
	
	// Escalation management
	escalationMgr  EscalationManager
	
	// Approval workflows
	approvalMgr    ApprovalManager
	
	// Analytics
	analytics       TaskAnalytics
	
	// Integration
	externalSystems map[string]ExternalSystem
}

type HumanLoopConfig struct {
	Enabled                bool          `json:"enabled"`
	DefaultTaskTimeout     time.Duration `json:"defaultTaskTimeout"`
	MaxConcurrentTasks     int           `json:"maxConcurrentTasks"`
	EscalationEnabled      bool          `json:"escalationEnabled"`
	EscalationTimeout      time.Duration `json:"escalationTimeout"`
	NotificationEnabled    bool          `json:"notificationEnabled"`
	ApprovalRequired       bool          `json:"approvalRequired"`
	AutoAssignmentEnabled  bool          `json:"autoAssignmentEnabled"`
	TaskHistoryRetention   time.Duration `json:"taskHistoryRetention"`
	AnalyticsEnabled       bool          `json:"analyticsEnabled"`
	ExternalIntegration    bool          `json:"externalIntegration"`
}

type HumanTask struct {
	ID                  string                 `json:"id"`
	Type                string                 `json:"type"`
	Title               string                 `json:"title"`
	Description         string                 `json:"description"`
	Priority            string                 `json:"priority"`
	Status              TaskStatus             `json:"status"`
	Assignee            string                 `json:"assignee"`
	Requester           string                 `json:"requester"`
	CreatedAt           time.Time              `json:"createdAt"`
	UpdatedAt           time.Time              `json:"updatedAt"`
	DueAt               *time.Time             `json:"dueAt,omitempty"`
	CompletedAt         *time.Time             `json:"completedAt,omitempty"`
	EscalatedAt         *time.Time             `json:"escalatedAt,omitempty"`
	
	// Task content
	Data                map[string]interface{} `json:"data"`
	FormData            map[string]interface{} `json:"formData,omitempty"`
	Attachments         []TaskAttachment       `json:"attachments,omitempty"`
	
	// Workflow context
	WorkflowID          string                 `json:"workflowId"`
	ActivityID          string                 `json:"activityId"`
	ExecutionID         string                 `json:"executionId"`
	
	// Approval and review
	ApprovalRequired    bool                   `json:"approvalRequired"`
	ApprovedBy          string                 `json:"approvedBy,omitempty"`
	ApprovedAt          *time.Time             `json:"approvedAt,omitempty"`
	ReviewComments      string                 `json:"reviewComments,omitempty"`
	ReviewScore         int                    `json:"reviewScore,omitempty"`
	
	// Escalation
	EscalationLevel     int                    `json:"escalationLevel"`
	EscalationReason    string                 `json:"escalationReason,omitempty"`
	EscalatedTo         string                 `json:"escalatedTo,omitempty"`
	
	// Metadata
	Tags                []string               `json:"tags"`
	Category            string                 `json:"category"`
	Complexity          string                 `json:"complexity"`
	EstimatedEffort     string                 `json:"estimatedEffort"`
	ActualEffort        string                 `json:"actualEffort,omitempty"`
	
	// Notifications
	NotificationSent     bool                   `json:"notificationSent"`
	ReminderSent        bool                   `json:"reminderSent"`
	LastReminderAt      *time.Time             `json:"lastReminderAt,omitempty"`
	
	// External integration
	ExternalID          string                 `json:"externalId,omitempty"`
	ExternalSystem      string                 `json:"externalSystem,omitempty"`
	SyncStatus          string                 `json:"syncStatus"`
	LastSyncAt          *time.Time             `json:"lastSyncAt,omitempty"`
}

type TaskStatus string

const (
	TaskStatusPending     TaskStatus = "pending"
	TaskStatusAssigned    TaskStatus = "assigned"
	TaskStatusInProgress  TaskStatus = "in_progress"
	TaskStatusReview      TaskStatus = "review"
	TaskStatusApproved    TaskStatus = "approved"
	TaskStatusRejected    TaskStatus = "rejected"
	TaskStatusCompleted   TaskStatus = "completed"
	TaskStatusEscalated   TaskStatus = "escalated"
	TaskStatusCancelled   TaskStatus = "cancelled"
	TaskStatusExpired     TaskStatus = "expired"
)

type TaskRequest struct {
	Type                string                 `json:"type"`
	Title               string                 `json:"title"`
	Description         string                 `json:"description"`
	Priority            string                 `json:"priority"`
	Data                map[string]interface{} `json:"data"`
	Assignee            string                 `json:"assignee,omitempty"`
	Requester           string                 `json:"requester"`
	DueAt               *time.Time             `json:"dueAt,omitempty"`
	WorkflowID          string                 `json:"workflowId,omitempty"`
	ActivityID          string                 `json:"activityId,omitempty"`
	Tags                []string               `json:"tags,omitempty"`
	Category            string                 `json:"category,omitempty"`
	Complexity          string                 `json:"complexity,omitempty"`
	EstimatedEffort     string                 `json:"estimatedEffort,omitempty"`
	ApprovalRequired    bool                   `json:"approvalRequired,omitempty"`
	Attachments         []TaskAttachment       `json:"attachments,omitempty"`
}

type TaskUpdate struct {
	Status              *TaskStatus            `json:"status,omitempty"`
	Assignee            *string                `json:"assignee,omitempty"`
	Priority            *string                `json:"priority,omitempty"`
	DueAt               *time.Time             `json:"dueAt,omitempty"`
	Data                map[string]interface{} `json:"data,omitempty"`
	FormData            map[string]interface{} `json:"formData,omitempty"`
	Tags                *[]string              `json:"tags,omitempty"`
	ReviewComments      *string                `json:"reviewComments,omitempty"`
	ReviewScore         *int                   `json:"reviewScore,omitempty"`
	ActualEffort        *string                `json:"actualEffort,omitempty"`
}

type TaskResult struct {
	Status              TaskStatus             `json:"status"`
	Decision            string                 `json:"decision"`
	Comments            string                 `json:"comments,omitempty"`
	FormData            map[string]interface{} `json:"formData,omitempty"`
	Attachments         []TaskAttachment       `json:"attachments,omitempty"`
	ApprovedBy          string                 `json:"approvedBy"`
	ApprovedAt          time.Time              `json:"approvedAt"`
	ReviewScore         int                    `json:"reviewScore,omitempty"`
	ActualEffort        string                 `json:"actualEffort,omitempty"`
	ExternalID          string                 `json:"externalId,omitempty"`
}

type TaskFilter struct {
	Type            string      `json:"type,omitempty"`
	Status          []TaskStatus `json:"status,omitempty"`
	Assignee        string      `json:"assignee,omitempty"`
	Requester       string      `json:"requester,omitempty"`
	Priority        string      `json:"priority,omitempty"`
	Category        string      `json:"category,omitempty"`
	CreatedAfter    *time.Time  `json:"createdAfter,omitempty"`
	CreatedBefore   *time.Time  `json:"createdBefore,omitempty"`
	DueAfter        *time.Time  `json:"dueAfter,omitempty"`
	DueBefore       *time.Time  `json:"dueBefore,omitempty"`
	Tags            []string    `json:"tags,omitempty"`
	WorkflowID      string      `json:"workflowId,omitempty"`
	Limit           int         `json:"limit,omitempty"`
	Offset          int         `json:"offset,omitempty"`
	SortBy          string      `json:"sortBy,omitempty"`
	SortOrder       string      `json:"sortOrder,omitempty"`
}

type TaskHistoryEntry struct {
	ID          string                 `json:"id"`
	TaskID      string                 `json:"taskId"`
	Action      string                 `json:"action"`
	OldValue    interface{}            `json:"oldValue,omitempty"`
	NewValue    interface{}            `json:"newValue,omitempty"`
	ChangedBy   string                 `json:"changedBy"`
	ChangedAt   time.Time              `json:"changedAt"`
	Comments    string                 `json:"comments,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type TaskAttachment struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Size        int64     `json:"size"`
	URL         string    `json:"url"`
	UploadedBy  string    `json:"uploadedBy"`
	UploadedAt  time.Time `json:"uploadedAt"`
	Description string    `json:"description,omitempty"`
}

type TaskAnalytics struct {
	TotalTasks          int                    `json:"totalTasks"`
	TasksByStatus       map[string]int          `json:"tasksByStatus"`
	TasksByPriority     map[string]int          `json:"tasksByPriority"`
	TasksByAssignee     map[string]int          `json:"tasksByAssignee"`
	TasksByCategory     map[string]int          `json:"tasksByCategory"`
	AverageCompletionTime time.Duration         `json:"averageCompletionTime"`
	AverageApprovalTime  time.Duration         `json:"averageApprovalTime"`
	EscalationRate      float64                `json:"escalationRate"`
	OverdueTasks       int                    `json:"overdueTasks"`
	PeakHours          map[string]int          `json:"peakHours"`
	ProductivityMetrics map[string]float64     `json:"productivityMetrics"`
}

// NotificationManager interface
type NotificationManager interface {
	SendNotification(notification Notification) error
	SendReminder(taskID string) error
	SendEscalationNotification(taskID string) error
	GetNotificationHistory(taskID string) ([]Notification, error)
}

type Notification struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Recipient   string                 `json:"recipient"`
	Subject     string                 `json:"subject"`
	Message     string                 `json:"message"`
	Channel     string                 `json:"channel"`
	Priority    string                 `json:"priority"`
	TaskID      string                 `json:"taskId,omitempty"`
	SentAt      time.Time              `json:"sentAt"`
	Status      string                 `json:"status"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// EscalationManager interface
type EscalationManager interface {
	EscalateTask(task *HumanTask, reason string) error
	CheckEscalations() error
	GetEscalationPolicy(taskType string) (*EscalationPolicy, error)
	UpdateEscalationPolicy(taskType string, policy EscalationPolicy) error
}

type EscalationPolicy struct {
	TaskType            string        `json:"taskType"`
	Levels              []EscalationLevel `json:"levels"`
	Timeout             time.Duration `json:"timeout"`
	AutoEscalate        bool          `json:"autoEscalate"`
	NotificationEnabled bool          `json:"notificationEnabled"`
}

type EscalationLevel struct {
	Level        int      `json:"level"`
	Name         string   `json:"name"`
	Assignee     string   `json:"assignee"`
	Timeout      time.Duration `json:"timeout"`
	Notify       bool     `json:"notify"`
	Actions      []string `json:"actions"`
}

// ApprovalManager interface
type ApprovalManager interface {
	RequestApproval(task *HumanTask) error
	ApproveTask(taskID string, approver string, comments string) error
	RejectTask(taskID string, approver string, comments string) error
	GetApprovalStatus(taskID string) (*ApprovalStatus, error)
}

type ApprovalStatus struct {
	TaskID       string    `json:"taskId"`
	Status       string    `json:"status"`
	RequestedBy  string    `json:"requestedBy"`
	RequestedAt  time.Time `json:"requestedAt"`
	ApprovedBy   string    `json:"approvedBy,omitempty"`
	ApprovedAt   *time.Time `json:"approvedAt,omitempty"`
	RejectedBy   string    `json:"rejectedBy,omitempty"`
	RejectedAt   *time.Time `json:"rejectedAt,omitempty"`
	Comments     string    `json:"comments,omitempty"`
	Required     bool      `json:"required"`
}

// ExternalSystem interface
type ExternalSystem interface {
	SyncTask(task *HumanTask) error
	CreateTask(task *HumanTask) error
	UpdateTask(task *HumanTask) error
	DeleteTask(taskID string) error
	GetTask(externalID string) (*HumanTask, error)
}

// NewEnhancedHumanLoopManager creates a new enhanced human loop manager
func NewEnhancedHumanLoopManager(config HumanLoopConfig) *EnhancedHumanLoopManager {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &EnhancedHumanLoopManager{
		ctx:             ctx,
		cancel:          cancel,
		tasks:           make(map[string]*HumanTask),
		taskHistory:     make(map[string][]*TaskHistoryEntry),
		taskQueue:       make(map[string][]string),
		config:          config,
		notificationMgr: NewNotificationManager(),
		escalationMgr:  NewEscalationManager(),
		approvalMgr:    NewApprovalManager(),
		analytics:       TaskAnalytics{},
		externalSystems: make(map[string]ExternalSystem),
	}
}

// Start begins the human loop manager operations
func (ehlm *EnhancedHumanLoopManager) Start(ctx context.Context) {
	ehlm.mu.Lock()
	defer ehlm.mu.Unlock()
	
	if ehlm.isRunning {
		return
	}
	
	ehlm.isRunning = true
	ehlm.ctx = ctx
	
	// Start background processes
	go ehlm.monitorTaskTimeouts()
	go ehlm.processEscalations()
	go ehlm.sendReminders()
	go ehlm.updateAnalytics()
	
	// Start external sync if enabled
	if ehlm.config.ExternalIntegration {
		go ehlm.syncExternalSystems()
	}
	
	log.Printf("Enhanced human loop manager started with config: %+v", ehlm.config)
}

// Stop stops the human loop manager
func (ehlm *EnhancedHumanLoopManager) Stop() {
	ehlm.mu.Lock()
	defer ehlm.mu.Unlock()
	
	if !ehlm.isRunning {
		return
	}
	
	ehlm.cancel()
	ehlm.isRunning = false
	
	log.Println("Enhanced human loop manager stopped")
}

// CreateTask creates a new human task
func (ehlm *EnhancedHumanLoopManager) CreateTask(request TaskRequest) (*HumanTask, error) {
	ehlm.mu.Lock()
	defer ehlm.mu.Unlock()
	
	// Validate request
	if err := ehlm.validateTaskRequest(&request); err != nil {
		return nil, fmt.Errorf("invalid task request: %w", err)
	}
	
	// Create task
	task := &HumanTask{
		ID:               ehlm.generateTaskID(),
		Type:             request.Type,
		Title:            request.Title,
		Description:      request.Description,
		Priority:         request.Priority,
		Status:           TaskStatusPending,
		Requester:        request.Requester,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		Data:             request.Data,
		WorkflowID:       request.WorkflowID,
		ActivityID:       request.ActivityID,
		Tags:             request.Tags,
		Category:         request.Category,
		Complexity:       request.Complexity,
		EstimatedEffort:  request.EstimatedEffort,
		ApprovalRequired: request.ApprovalRequired,
		Attachments:      request.Attachments,
	}
	
	// Set due date if not provided
	if request.DueAt != nil {
		task.DueAt = request.DueAt
	} else {
		dueAt := time.Now().Add(ehlm.config.DefaultTaskTimeout)
		task.DueAt = &dueAt
	}
	
	// Auto-assign if enabled and no assignee specified
	if ehlm.config.AutoAssignmentEnabled && request.Assignee == "" {
		assignee := ehlm.findBestAssignee(task)
		task.Assignee = assignee
		task.Status = TaskStatusAssigned
	} else if request.Assignee != "" {
		task.Assignee = request.Assignee
		task.Status = TaskStatusAssigned
	}
	
	// Add to tasks map
	ehlm.tasks[task.ID] = task
	
	// Add to task queue
	if task.Assignee != "" {
		ehlm.taskQueue[task.Assignee] = append(ehlm.taskQueue[task.Assignee], task.ID)
	}
	
	// Create history entry
	ehlm.addHistoryEntry(task.ID, "created", nil, task, "system", "Task created")
	
	// Send notification if enabled
	if ehlm.config.NotificationEnabled && task.Assignee != "" {
		go ehlm.NotifyAssignee(task.ID)
	}
	
	// Sync with external systems if enabled
	if ehlm.config.ExternalIntegration {
		go ehlm.syncTaskExternal(task)
	}
	
	log.Printf("Created task %s: %s", task.ID, task.Title)
	
	return task, nil
}

// GetTask retrieves a task by ID
func (ehlm *EnhancedHumanLoopManager) GetTask(taskID string) (*HumanTask, error) {
	ehlm.mu.RLock()
	defer ehlm.mu.RUnlock()
	
	task, exists := ehlm.tasks[taskID]
	if !exists {
		return nil, fmt.Errorf("task not found: %s", taskID)
	}
	
	return task, nil
}

// UpdateTask updates an existing task
func (ehlm *EnhancedHumanLoopManager) UpdateTask(taskID string, update TaskUpdate) (*HumanTask, error) {
	ehlm.mu.Lock()
	defer ehlm.mu.Unlock()
	
	task, exists := ehlm.tasks[taskID]
	if !exists {
		return nil, fmt.Errorf("task not found: %s", taskID)
	}
	
	// Create copy of original task for history
	originalTask := *task
	
	// Apply updates
	if update.Status != nil {
		task.Status = *update.Status
		ehlm.addHistoryEntry(taskID, "status_changed", originalTask.Status, *update.Status, "system", "Status updated")
	}
	
	if update.Assignee != nil {
		oldAssignee := task.Assignee
		task.Assignee = *update.Assignee
		task.Status = TaskStatusAssigned
		
		// Update task queue
		if oldAssignee != "" {
			ehlm.removeFromQueue(oldAssignee, taskID)
		}
		ehlm.taskQueue[*update.Assignee] = append(ehlm.taskQueue[*update.Assignee], taskID)
		
		ehlm.addHistoryEntry(taskID, "assignee_changed", oldAssignee, *update.Assignee, "system", "Assignee updated")
	}
	
	if update.Priority != nil {
		task.Priority = *update.Priority
		ehlm.addHistoryEntry(taskID, "priority_changed", originalTask.Priority, *update.Priority, "system", "Priority updated")
	}
	
	if update.DueAt != nil {
		task.DueAt = update.DueAt
		ehlm.addHistoryEntry(taskID, "due_date_changed", originalTask.DueAt, *update.DueAt, "system", "Due date updated")
	}
	
	if update.Data != nil {
		task.Data = update.Data
		ehlm.addHistoryEntry(taskID, "data_updated", originalTask.Data, update.Data, "system", "Task data updated")
	}
	
	if update.FormData != nil {
		task.FormData = update.FormData
		ehlm.addHistoryEntry(taskID, "form_data_updated", originalTask.FormData, update.FormData, "system", "Form data updated")
	}
	
	if update.Tags != nil {
		task.Tags = *update.Tags
		ehlm.addHistoryEntry(taskID, "tags_updated", originalTask.Tags, *update.Tags, "system", "Tags updated")
	}
	
	if update.ReviewComments != nil {
		task.ReviewComments = *update.ReviewComments
		ehlm.addHistoryEntry(taskID, "review_comments_updated", originalTask.ReviewComments, *update.ReviewComments, "system", "Review comments updated")
	}
	
	if update.ReviewScore != nil {
		task.ReviewScore = *update.ReviewScore
		ehlm.addHistoryEntry(taskID, "review_score_updated", originalTask.ReviewScore, *update.ReviewScore, "system", "Review score updated")
	}
	
	if update.ActualEffort != nil {
		task.ActualEffort = *update.ActualEffort
		ehlm.addHistoryEntry(taskID, "actual_effort_updated", originalTask.ActualEffort, *update.ActualEffort, "system", "Actual effort updated")
	}
	
	task.UpdatedAt = time.Now()
	
	// Sync with external systems if enabled
	if ehlm.config.ExternalIntegration {
		go ehlm.syncTaskExternal(task)
	}
	
	log.Printf("Updated task %s", taskID)
	
	return task, nil
}

// AssignTask assigns a task to a specific person
func (ehlm *EnhancedHumanLoopManager) AssignTask(taskID string, assignee string) error {
	ehlm.mu.Lock()
	defer ehlm.mu.Unlock()
	
	task, exists := ehlm.tasks[taskID]
	if !exists {
		return fmt.Errorf("task not found: %s", taskID)
	}
	
	oldAssignee := task.Assignee
	task.Assignee = assignee
	task.Status = TaskStatusAssigned
	task.UpdatedAt = time.Now()
	
	// Update task queue
	if oldAssignee != "" {
		ehlm.removeFromQueue(oldAssignee, taskID)
	}
	ehlm.taskQueue[assignee] = append(ehlm.taskQueue[assignee], taskID)
	
	// Create history entry
	ehlm.addHistoryEntry(taskID, "assigned", oldAssignee, assignee, "system", "Task assigned")
	
	// Send notification
	if ehlm.config.NotificationEnabled {
		go ehlm.NotifyAssignee(taskID)
	}
	
	// Sync with external systems if enabled
	if ehlm.config.ExternalIntegration {
		go ehlm.syncTaskExternal(task)
	}
	
	log.Printf("Assigned task %s to %s", taskID, assignee)
	
	return nil
}

// CompleteTask completes a task with the given result
func (ehlm *EnhancedHumanLoopManager) CompleteTask(taskID string, result TaskResult) error {
	ehlm.mu.Lock()
	defer ehlm.mu.Unlock()
	
	task, exists := ehlm.tasks[taskID]
	if !exists {
		return fmt.Errorf("task not found: %s", taskID)
	}
	
	// Update task with result
	task.Status = result.Status
	task.FormData = result.FormData
	task.ApprovedBy = result.ApprovedBy
	task.ApprovedAt = &result.ApprovedAt
	task.ReviewComments = result.Comments
	task.ReviewScore = result.ReviewScore
	task.ActualEffort = result.ActualEffort
	task.ExternalID = result.ExternalID
	task.CompletedAt = &result.ApprovedAt
	task.UpdatedAt = time.Now()
	
	// Add attachments if provided
	if len(result.Attachments) > 0 {
		task.Attachments = append(task.Attachages, result.Attachments...)
	}
	
	// Remove from task queue
	if task.Assignee != "" {
		ehlm.removeFromQueue(task.Assignee, taskID)
	}
	
	// Create history entry
	ehlm.addHistoryEntry(taskID, "completed", nil, result, result.ApprovedBy, "Task completed")
	
	// Send completion notification
	if ehlm.config.NotificationEnabled {
		go ehlm.sendCompletionNotification(task)
	}
	
	// Sync with external systems if enabled
	if ehlm.config.ExternalIntegration {
		go ehlm.syncTaskExternal(task)
	}
	
	log.Printf("Completed task %s with status %s", taskID, result.Status)
	
	return nil
}

// EscalateTask escalates a task
func (ehlm *EnhancedHumanLoopManager) EscalateTask(taskID string, reason string) error {
	ehlm.mu.Lock()
	defer ehlm.mu.Unlock()
	
	task, exists := ehlm.tasks[taskID]
	if !exists {
		return fmt.Errorf("task not found: %s", taskID)
	}
	
	if !ehlm.config.EscalationEnabled {
		return fmt.Errorf("escalation not enabled")
	}
	
	// Update task
	task.Status = TaskStatusEscalated
	task.EscalationLevel++
	task.EscalationReason = reason
	task.EscalatedAt = &[]time.Time{time.Now()}[0]
	task.UpdatedAt = time.Now()
	
	// Get escalation policy
	policy, err := ehlm.escalationMgr.GetEscalationPolicy(task.Type)
	if err != nil {
		log.Printf("Warning: Failed to get escalation policy for task type %s: %v", task.Type, err)
		return err
	}
	
	// Find next escalation level
	if task.EscalationLevel < len(policy.Levels) {
		level := policy.Levels[task.EscalationLevel]
		task.EscalatedTo = level.Assignee
		
		// Reassign task
		if task.Assignee != "" {
			ehlm.removeFromQueue(task.Assignee, taskID)
		}
		task.Assignee = level.Assignee
		ehlm.taskQueue[level.Assignee] = append(ehlm.taskQueue[level.Assignee], taskID)
		
		// Send escalation notification
		if ehlm.config.NotificationEnabled {
			go ehlm.escalationMgr.SendEscalationNotification(taskID)
		}
	}
	
	// Create history entry
	ehlm.addHistoryEntry(taskID, "escalated", task.EscalationLevel-1, task.EscalationLevel, "system", reason)
	
	// Sync with external systems if enabled
	if ehlm.config.ExternalIntegration {
		go ehlm.syncTaskExternal(task)
	}
	
	log.Printf("Escalated task %s to level %d: %s", taskID, task.EscalationLevel, reason)
	
	return nil
}

// GetTasks retrieves tasks based on filter
func (ehlm *EnhancedHumanLoopManager) GetTasks(filter TaskFilter) ([]*HumanTask, error) {
	ehlm.mu.RLock()
	defer ehlm.mu.RUnlock()
	
	var tasks []*HumanTask
	
	for _, task := range ehlm.tasks {
		if ehlm.matchesFilter(task, &filter) {
			tasks = append(tasks, task)
		}
	}
	
	// Sort tasks
	ehlm.sortTasks(tasks, filter.SortBy, filter.SortOrder)
	
	// Apply pagination
	if filter.Offset > 0 || filter.Limit > 0 {
		start := filter.Offset
		if start > len(tasks) {
			return []*HumanTask{}, nil
		}
		
		end := start + filter.Limit
		if filter.Limit == 0 || end > len(tasks) {
			end = len(tasks)
		}
		
		tasks = tasks[start:end]
	}
	
	return tasks, nil
}

// GetTaskHistory retrieves the history of a task
func (ehlm *EnhancedHumanLoopManager) GetTaskHistory(taskID string) ([]*TaskHistoryEntry, error) {
	ehlm.mu.RLock()
	defer ehlm.mu.RUnlock()
	
	history, exists := ehlm.taskHistory[taskID]
	if !exists {
		return []*TaskHistoryEntry{}, nil
	}
	
	return history, nil
}

// NotifyAssignee sends a notification to the task assignee
func (ehlm *EnhancedHumanLoopManager) NotifyAssignee(taskID string) error {
	ehlm.mu.RLock()
	task, exists := ehlm.tasks[taskID]
	ehlm.mu.RUnlock()
	
	if !exists {
		return fmt.Errorf("task not found: %s", taskID)
	}
	
	if task.Assignee == "" {
		return fmt.Errorf("task has no assignee: %s", taskID)
	}
	
	notification := Notification{
		ID:        ehlm.generateNotificationID(),
		Type:      "task_assigned",
		Recipient: task.Assignee,
		Subject:   fmt.Sprintf("New Task Assigned: %s", task.Title),
		Message:   fmt.Sprintf("You have been assigned a new task: %s\n\nDescription: %s\nPriority: %s\nDue: %s", 
			task.Title, task.Description, task.Priority, task.DueAt.Format("2006-01-02 15:04")),
		Channel:  "email",
		Priority: task.Priority,
		TaskID:   taskID,
		SentAt:   time.Now(),
		Status:   "pending",
	}
	
	return ehlm.notificationMgr.SendNotification(notification)
}

// Private helper methods

func (ehlm *EnhancedHumanLoopManager) generateTaskID() string {
	return fmt.Sprintf("task-%d", time.Now().UnixNano())
}

func (ehlm *EnhancedHumanLoopManager) generateNotificationID() string {
	return fmt.Sprintf("notif-%d", time.Now().UnixNano())
}

func (ehlm *EnhancedHumanLoopManager) validateTaskRequest(request *TaskRequest) error {
	if request.Title == "" {
		return fmt.Errorf("title is required")
	}
	if request.Type == "" {
		return fmt.Errorf("type is required")
	}
	if request.Requester == "" {
		return fmt.Errorf("requester is required")
	}
	if request.Priority == "" {
		request.Priority = "normal"
	}
	return nil
}

func (ehlm *EnhancedHumanLoopManager) findBestAssignee(task *HumanTask) string {
	// This is a simplified implementation
	// In production, you'd use sophisticated assignment logic
	// based on workload, expertise, availability, etc.
	
	// For now, just return a default assignee
	return "default-assignee"
}

func (ehlm *EnhancedHumanLoopManager) addHistoryEntry(taskID, action string, oldValue, newValue interface{}, changedBy, comments string) {
	entry := &TaskHistoryEntry{
		ID:        fmt.Sprintf("history-%d", time.Now().UnixNano()),
		TaskID:    taskID,
		Action:    action,
		OldValue:  oldValue,
		NewValue:  newValue,
		ChangedBy: changedBy,
		ChangedAt: time.Now(),
		Comments:  comments,
	}
	
	ehlm.taskHistory[taskID] = append(ehlm.taskHistory[taskID], entry)
}

func (ehlm *EnhancedHumanLoopManager) removeFromQueue(assignee, taskID string) {
	tasks := ehlm.taskQueue[assignee]
	for i, id := range tasks {
		if id == taskID {
			ehlm.taskQueue[assignee] = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
}

func (ehlm *EnhancedHumanLoopManager) matchesFilter(task *HumanTask, filter *TaskFilter) bool {
	// Type filter
	if filter.Type != "" && task.Type != filter.Type {
		return false
	}
	
	// Status filter
	if len(filter.Status) > 0 {
		found := false
		for _, status := range filter.Status {
			if task.Status == status {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	
	// Assignee filter
	if filter.Assignee != "" && task.Assignee != filter.Assignee {
		return false
	}
	
	// Requester filter
	if filter.Requester != "" && task.Requester != filter.Requester {
		return false
	}
	
	// Priority filter
	if filter.Priority != "" && task.Priority != filter.Priority {
		return false
	}
	
	// Category filter
	if filter.Category != "" && task.Category != filter.Category {
		return false
	}
	
	// Workflow ID filter
	if filter.WorkflowID != "" && task.WorkflowID != filter.WorkflowID {
		return false
	}
	
	// Created date filters
	if filter.CreatedAfter != nil && task.CreatedAt.Before(*filter.CreatedAfter) {
		return false
	}
	if filter.CreatedBefore != nil && task.CreatedAt.After(*filter.CreatedBefore) {
		return false
	}
	
	// Due date filters
	if filter.DueAfter != nil && task.DueAt != nil && task.DueAt.Before(*filter.DueAfter) {
		return false
	}
	if filter.DueBefore != nil && task.DueAt != nil && task.DueAt.After(*filter.DueBefore) {
		return false
	}
	
	// Tags filter
	if len(filter.Tags) > 0 {
		found := false
		for _, filterTag := range filter.Tags {
			for _, taskTag := range task.Tags {
				if taskTag == filterTag {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return false
		}
	}
	
	return true
}

func (ehlm *EnhancedHumanLoopManager) sortTasks(tasks []*HumanTask, sortBy, sortOrder string) {
	// This is a simplified sorting implementation
	// In production, you'd implement more sophisticated sorting
	
	if sortBy == "" {
		sortBy = "created_at"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}
	
	// Sort by creation time (newest first)
	if sortBy == "created_at" && sortOrder == "desc" {
		// Already in creation order (newest first due to map iteration)
		return
	}
	
	// For other sorting options, implement as needed
}

// Background processes

func (ehlm *EnhancedHumanLoopManager) monitorTaskTimeouts() {
	ticker := time.NewTicker(time.Minute * 5)
	defer ticker.Stop()
	
	for {
		select {
		case <-ehlm.ctx.Done():
			return
		case <-ticker.C:
			ehlm.checkTaskTimeouts()
		}
	}
}

func (ehlm *EnhancedHumanLoopManager) checkTaskTimeouts() {
	ehlm.mu.Lock()
	defer ehlm.mu.Unlock()
	
	now := time.Now()
	
	for _, task := range ehlm.tasks {
		if task.DueAt != nil && now.After(*task.DueAt) {
			if task.Status != TaskStatusCompleted && task.Status != TaskStatusCancelled {
				// Mark as expired
				task.Status = TaskStatusExpired
				task.UpdatedAt = now
				
				ehlm.addHistoryEntry(task.ID, "expired", nil, now, "system", "Task expired due to timeout")
				
				// Send notification
				if ehlm.config.NotificationEnabled {
					go ehlm.sendExpirationNotification(task)
				}
			}
		}
	}
}

func (ehlm *EnhancedHumanLoopManager) processEscalations() {
	if !ehlm.config.EscalationEnabled {
		return
	}
	
	ticker := time.NewTicker(time.Minute * 10)
	defer ticker.Stop()
	
	for {
		select {
		case <-ehlm.ctx.Done():
			return
		case <-ticker.C:
			ehlm.escalationMgr.CheckEscalations()
		}
	}
}

func (ehlm *EnhancedHumanLoopManager) sendReminders() {
	if !ehlm.config.NotificationEnabled {
		return
	}
	
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	
	for {
		select {
		case <-ehlm.ctx.Done():
			return
		case <-ticker.C:
			ehlm.sendTaskReminders()
		}
	}
}

func (ehlm *EnhancedHumanLoopManager) sendTaskReminders() {
	ehlm.mu.RLock()
	defer ehlm.mu.RUnlock()
	
	now := time.Now()
	
	for _, task := range ehlm.tasks {
		if task.Status == TaskStatusAssigned || task.Status == TaskStatusInProgress {
			if task.DueAt != nil {
				// Send reminder if due date is approaching (within 24 hours)
				if task.DueAt.Sub(now) < 24*time.Hour && task.DueAt.Sub(now) > 0 {
					if !task.ReminderSent || (task.LastReminderAt != nil && now.Sub(*task.LastReminderAt) > 12*time.Hour) {
						go ehlm.notificationMgr.SendReminder(task.ID)
						task.ReminderSent = true
						now := time.Now()
						task.LastReminderAt = &now
					}
				}
			}
		}
	}
}

func (ehlm *EnhancedHumanLoopManager) updateAnalytics() {
	if !ehlm.config.AnalyticsEnabled {
		return
	}
	
	ticker := time.NewTicker(time.Minute * 15)
	defer ticker.Stop()
	
	for {
		select {
		case <-ehlm.ctx.Done():
			return
		case <-ticker.C:
			ehlm.calculateAnalytics()
		}
	}
}

func (ehlm *EnhancedHumanLoopManager) calculateAnalytics() {
	ehlm.mu.Lock()
	defer ehlm.mu.Unlock()
	
	analytics := TaskAnalytics{
		TotalTasks:          len(ehlm.tasks),
		TasksByStatus:       make(map[string]int),
		TasksByPriority:     make(map[string]int),
		TasksByAssignee:     make(map[string]int),
		TasksByCategory:     make(map[string]int),
		PeakHours:          make(map[string]int),
		ProductivityMetrics: make(map[string]float64),
	}
	
	var totalCompletionTime time.Duration
	var totalApprovalTime time.Duration
	var completedTasks int
	var escalatedTasks int
	var overdueTasks int
	
	for _, task := range ehlm.tasks {
		// Status analytics
		analytics.TasksByStatus[string(task.Status)]++
		
		// Priority analytics
		analytics.TasksByPriority[task.Priority]++
		
		// Assignee analytics
		if task.Assignee != "" {
			analytics.TasksByAssignee[task.Assignee]++
		}
		
		// Category analytics
		if task.Category != "" {
			analytics.TasksByCategory[task.Category]++
		}
		
		// Peak hours analytics
		hour := task.CreatedAt.Hour()
		analytics.PeakHours[fmt.Sprintf("%02d:00", hour)]++
		
		// Completion time analytics
		if task.CompletedAt != nil {
			completionTime := task.CompletedAt.Sub(task.CreatedAt)
			totalCompletionTime += completionTime
			completedTasks++
		}
		
		// Approval time analytics
		if task.ApprovedAt != nil {
			approvalTime := task.ApprovedAt.Sub(task.CreatedAt)
			totalApprovalTime += approvalTime
		}
		
		// Escalation analytics
		if task.EscalationLevel > 0 {
			escalatedTasks++
		}
		
		// Overdue tasks
		if task.DueAt != nil && time.Now().After(*task.DueAt) && 
		   task.Status != TaskStatusCompleted && task.Status != TaskStatusCancelled {
			overdueTasks++
		}
	}
	
	// Calculate averages
	if completedTasks > 0 {
		analytics.AverageCompletionTime = totalCompletionTime / time.Duration(completedTasks)
	}
	
	if completedTasks > 0 {
		analytics.AverageApprovalTime = totalApprovalTime / time.Duration(completedTasks)
	}
	
	// Calculate escalation rate
	if len(ehlm.tasks) > 0 {
		analytics.EscalationRate = float64(escalatedTasks) / float64(len(ehlm.tasks))
	}
	
	analytics.OverdueTasks = overdueTasks
	
	ehlm.analytics = analytics
}

func (ehlm *EnhancedHumanLoopManager) syncExternalSystems() {
	ticker := time.NewTicker(time.Minute * 5)
	defer ticker.Stop()
	
	for {
		select {
		case <-ehlm.ctx.Done():
			return
		case <-ticker.C:
			ehlm.syncAllTasks()
		}
	}
}

func (ehlm *EnhancedHumanLoopManager) syncAllTasks() {
	ehlm.mu.RLock()
	tasks := make([]*HumanTask, 0, len(ehlm.tasks))
	for _, task := range ehlm.tasks {
		tasks = append(tasks, task)
	}
	ehlm.mu.RUnlock()
	
	for _, task := range tasks {
		ehlm.syncTaskExternal(task)
	}
}

func (ehlm *EnhancedHumanLoopManager) syncTaskExternal(task *HumanTask) {
	if task.ExternalSystem == "" {
		return
	}
	
	system, exists := ehlm.externalSystems[task.ExternalSystem]
	if !exists {
		return
	}
	
	if err := system.SyncTask(task); err != nil {
		log.Printf("Failed to sync task %s with external system %s: %v", task.ID, task.ExternalSystem, err)
	}
}

func (ehlm *EnhancedHumanLoopManager) sendCompletionNotification(task *HumanTask) {
	notification := Notification{
		ID:        ehlm.generateNotificationID(),
		Type:      "task_completed",
		Recipient: task.Requester,
		Subject:   fmt.Sprintf("Task Completed: %s", task.Title),
		Message:   fmt.Sprintf("Task '%s' has been completed by %s\n\nStatus: %s\nComments: %s", 
			task.Title, task.ApprovedBy, task.Status, task.ReviewComments),
		Channel:  "email",
		Priority: task.Priority,
		TaskID:   task.ID,
		SentAt:   time.Now(),
		Status:   "pending",
	}
	
	ehlm.notificationMgr.SendNotification(notification)
}

func (ehlm *EnhancedHumanLoopManager) sendExpirationNotification(task *HumanTask) {
	notification := Notification{
		ID:        ehlm.generateNotificationID(),
		Type:      "task_expired",
		Recipient: task.Requester,
		Subject:   fmt.Sprintf("Task Expired: %s", task.Title),
		Message:   fmt.Sprintf("Task '%s' assigned to %s has expired.\n\nDue date: %s\nStatus: %s", 
			task.Title, task.Assignee, task.DueAt.Format("2006-01-02 15:04"), task.Status),
		Channel:  "email",
		Priority: "high",
		TaskID:   task.ID,
		SentAt:   time.Now(),
		Status:   "pending",
	}
	
	ehlm.notificationMgr.SendNotification(notification)
}

// Placeholder implementations for interfaces

func NewNotificationManager() NotificationManager {
	return &placeholderNotificationManager{}
}

func NewEscalationManager() EscalationManager {
	return &placeholderEscalationManager{}
}

func NewApprovalManager() ApprovalManager {
	return &placeholderApprovalManager{}
}

type placeholderNotificationManager struct{}

func (p *placeholderNotificationManager) SendNotification(notification Notification) error {
	log.Printf("Notification sent: %+v", notification)
	return nil
}

func (p *placeholderNotificationManager) SendReminder(taskID string) error {
	log.Printf("Reminder sent for task: %s", taskID)
	return nil
}

func (p *placeholderNotificationManager) SendEscalationNotification(taskID string) error {
	log.Printf("Escalation notification sent for task: %s", taskID)
	return nil
}

func (p *placeholderNotificationManager) GetNotificationHistory(taskID string) ([]Notification, error) {
	return []Notification{}, nil
}

type placeholderEscalationManager struct{}

func (p *placeholderEscalationManager) EscalateTask(task *HumanTask, reason string) error {
	log.Printf("Task escalated: %s - %s", task.ID, reason)
	return nil
}

func (p *placeholderEscalationManager) CheckEscalations() error {
	log.Printf("Checking escalations")
	return nil
}

func (p *placeholderEscalationManager) GetEscalationPolicy(taskType string) (*EscalationPolicy, error) {
	return &EscalationPolicy{
		TaskType: taskType,
		Levels: []EscalationLevel{
			{Level: 1, Name: "Manager", Assignee: "manager", Timeout: time.Hour * 24, Notify: true, Actions: []string{"notify"}},
			{Level: 2, Name: "Director", Assignee: "director", Timeout: time.Hour * 48, Notify: true, Actions: []string{"notify", "reassign"}},
		},
		Timeout:             time.Hour * 24,
		AutoEscalate:        true,
		NotificationEnabled: true,
	}, nil
}

func (p *placeholderEscalationManager) UpdateEscalationPolicy(taskType string, policy EscalationPolicy) error {
	log.Printf("Escalation policy updated for %s", taskType)
	return nil
}

type placeholderApprovalManager struct{}

func (p *placeholderApprovalManager) RequestApproval(task *HumanTask) error {
	log.Printf("Approval requested for task: %s", task.ID)
	return nil
}

func (p *placeholderApprovalManager) ApproveTask(taskID string, approver string, comments string) error {
	log.Printf("Task approved: %s by %s", taskID, approver)
	return nil
}

func (p *placeholderApprovalManager) RejectTask(taskID string, approver string, comments string) error {
	log.Printf("Task rejected: %s by %s", taskID, approver)
	return nil
}

func (p *placeholderApprovalManager) GetApprovalStatus(taskID string) (*ApprovalStatus, error) {
	return &ApprovalStatus{
		TaskID:      taskID,
		Status:      "pending",
		RequestedBy: "system",
		RequestedAt: time.Now(),
		Required:    true,
	}, nil
}
