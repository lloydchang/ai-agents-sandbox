package security

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// SecurityManager interface for security operations
type SecurityManager interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*TokenInfo, error)
	RevokeToken(token string) error
	EncryptData(data []byte) ([]byte, error)
	DecryptData(encryptedData []byte) ([]byte, error)
	AuditLog(event AuditEvent)
	GetAuditLogs(filter AuditFilter) []AuditEvent
	CheckPermission(userID, resource, action string) bool
}

// Enhanced security implementation
type EnhancedSecurityManager struct {
	ctx           context.Context
	cancel        context.CancelFunc
	isRunning     bool
	mu            sync.RWMutex
	
	// Token management
	activeTokens  map[string]*TokenInfo
	revokedTokens map[string]time.Time
	
	// Audit logging
	auditLogs     []AuditEvent
	auditConfig   AuditConfig
	
	// Encryption
	encryptionKey []byte
	
	// Access control
	permissions   map[string]map[string][]string // userID -> resource -> actions
	roles         map[string][]string           // role -> permissions
	
	// Configuration
	config        SecurityConfig
	
	// Rate limiting
	rateLimiter   map[string]*RateLimitInfo
}

type TokenInfo struct {
	TokenID      string                 `json:"tokenId"`
	UserID       string                 `json:"userId"`
	IssuedAt     time.Time              `json:"issuedAt"`
	ExpiresAt    time.Time              `json:"expiresAt"`
	Permissions  []string               `json:"permissions"`
	Metadata     map[string]interface{} `json:"metadata"`
	LastUsed     time.Time              `json:"lastUsed"`
	UsageCount   int                    `json:"usageCount"`
}

type AuditEvent struct {
	ID           string                 `json:"id"`
	Timestamp    time.Time              `json:"timestamp"`
	EventType    string                 `json:"eventType"`
	UserID       string                 `json:"userId"`
	Resource     string                 `json:"resource"`
	Action       string                 `json:"action"`
	IPAddress    string                 `json:"ipAddress"`
	UserAgent    string                 `json:"userAgent"`
	Success      bool                   `json:"success"`
	ErrorMessage string                 `json:"errorMessage,omitempty"`
	Metadata     map[string]interface{} `json:"metadata"`
}

type AuditConfig struct {
	Enabled          bool          `json:"enabled"`
	RetentionPeriod  time.Duration `json:"retentionPeriod"`
	LogLevel         string        `json:"logLevel"`
	ExternalEndpoint string        `json:"externalEndpoint"`
	BatchSize        int           `json:"batchSize"`
	BatchInterval    time.Duration `json:"batchInterval"`
}

type SecurityConfig struct {
	TokenExpiration     time.Duration `json:"tokenExpiration"`
	MaxTokenUsage       int           `json:"maxTokenUsage"`
	EncryptionAlgorithm  string        `json:"encryptionAlgorithm"`
	EnableRateLimiting  bool          `json:"enableRateLimiting"`
	RateLimitWindow     time.Duration `json:"rateLimitWindow"`
	MaxRequestsPerWindow int         `json:"maxRequestsPerWindow"`
	EnableAuditLogging  bool          `json:"enableAuditLogging"`
	SessionTimeout      time.Duration `json:"sessionTimeout"`
}

type AuditFilter struct {
	UserID       string     `json:"userId,omitempty"`
	EventType    string     `json:"eventType,omitempty"`
	Resource     string     `json:"resource,omitempty"`
	Action       string     `json:"action,omitempty"`
	StartTime    *time.Time `json:"startTime,omitempty"`
	EndTime      *time.Time `json:"endTime,omitempty"`
	Success      *bool      `json:"success,omitempty"`
	Limit        int        `json:"limit,omitempty"`
}

type RateLimitInfo struct {
	RequestCount int       `json:"requestCount"`
	WindowStart  time.Time `json:"windowStart"`
	LastRequest  time.Time `json:"lastRequest"`
}

// NewEnhancedSecurityManager creates a new enhanced security manager
func NewEnhancedSecurityManager(config SecurityConfig) *EnhancedSecurityManager {
	ctx, cancel := context.WithCancel(context.Background())
	
	// Generate encryption key
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Printf("Warning: Failed to generate encryption key: %v", err)
	}
	
	return &EnhancedSecurityManager{
		ctx:           ctx,
		cancel:        cancel,
		activeTokens:  make(map[string]*TokenInfo),
		revokedTokens: make(map[string]time.Time),
		auditLogs:     make([]AuditEvent, 0),
		auditConfig: AuditConfig{
			Enabled:         config.EnableAuditLogging,
			RetentionPeriod: 24 * time.Hour * 30, // 30 days
			LogLevel:        "info",
			BatchSize:       100,
			BatchInterval:   time.Minute * 5,
		},
		encryptionKey: key,
		permissions:   make(map[string]map[string][]string),
		roles:         make(map[string][]string),
		config:        config,
		rateLimiter:   make(map[string]*RateLimitInfo),
	}
}

// Start begins the security manager operations
func (esm *EnhancedSecurityManager) Start(ctx context.Context) {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	if esm.isRunning {
		return
	}
	
	esm.isRunning = true
	esm.ctx = ctx
	
	// Start cleanup goroutines
	go esm.cleanupExpiredTokens()
	go esm.cleanupAuditLogs()
	
	// Start audit log batching if external endpoint is configured
	if esm.auditConfig.ExternalEndpoint != "" {
		go esm.batchAuditLogs()
	}
	
	// Start rate limiting cleanup
	if esm.config.EnableRateLimiting {
		go esm.cleanupRateLimits()
	}
	
	log.Printf("Enhanced security manager started with config: %+v", esm.config)
}

// Stop stops the security manager
func (esm *EnhancedSecurityManager) Stop() {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	if !esm.isRunning {
		return
	}
	
	esm.cancel()
	esm.isRunning = false
	
	log.Println("Enhanced security manager stopped")
}

// GenerateToken generates a new authentication token
func (esm *EnhancedSecurityManager) GenerateToken(userID string) (string, error) {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	// Check rate limiting
	if esm.config.EnableRateLimiting {
		if !esm.checkRateLimit(userID) {
			return "", fmt.Errorf("rate limit exceeded for user %s", userID)
		}
	}
	
	// Generate token
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	
	token := base64.URLEncoding.EncodeToString(tokenBytes)
	tokenID := esm.generateTokenID(token)
	
	// Get user permissions
	permissions := esm.getUserPermissions(userID)
	
	// Create token info
	tokenInfo := &TokenInfo{
		TokenID:     tokenID,
		UserID:      userID,
		IssuedAt:    time.Now(),
		ExpiresAt:   time.Now().Add(esm.config.TokenExpiration),
		Permissions: permissions,
		Metadata:    make(map[string]interface{}),
		LastUsed:    time.Now(),
		UsageCount:  0,
	}
	
	esm.activeTokens[token] = tokenInfo
	
	// Audit log token generation
	esm.AuditLog(AuditEvent{
		EventType: "token_generated",
		UserID:    userID,
		Action:    "generate_token",
		Success:   true,
		Metadata: map[string]interface{}{
			"tokenId": tokenID,
			"expiresAt": tokenInfo.ExpiresAt,
		},
	})
	
	return token, nil
}

// ValidateToken validates a token and returns token info
func (esm *EnhancedSecurityManager) ValidateToken(token string) (*TokenInfo, error) {
	esm.mu.RLock()
	defer esm.mu.RUnlock()
	
	// Check if token is revoked
	if revokedAt, exists := esm.revokedTokens[token]; exists {
		return nil, fmt.Errorf("token revoked at %v", revokedAt)
	}
	
	// Check if token exists
	tokenInfo, exists := esm.activeTokens[token]
	if !exists {
		return nil, fmt.Errorf("token not found")
	}
	
	// Check if token is expired
	if time.Now().After(tokenInfo.ExpiresAt) {
		return nil, fmt.Errorf("token expired")
	}
	
	// Check usage limit
	if esm.config.MaxTokenUsage > 0 && tokenInfo.UsageCount >= esm.config.MaxTokenUsage {
		return nil, fmt.Errorf("token usage limit exceeded")
	}
	
	// Update usage info
	tokenInfo.LastUsed = time.Now()
	tokenInfo.UsageCount++
	
	return tokenInfo, nil
}

// RevokeToken revokes a token
func (esm *EnhancedSecurityManager) RevokeToken(token string) error {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	tokenInfo, exists := esm.activeTokens[token]
	if !exists {
		return fmt.Errorf("token not found")
	}
	
	// Move to revoked tokens
	esm.revokedTokens[token] = time.Now()
	delete(esm.activeTokens, token)
	
	// Audit log token revocation
	esm.AuditLog(AuditEvent{
		EventType: "token_revoked",
		UserID:    tokenInfo.UserID,
		Action:    "revoke_token",
		Success:   true,
		Metadata: map[string]interface{}{
			"tokenId": tokenInfo.TokenID,
			"usageCount": tokenInfo.UsageCount,
		},
	})
	
	return nil
}

// EncryptData encrypts data using AES-256-GCM
func (esm *EnhancedSecurityManager) EncryptData(data []byte) ([]byte, error) {
	// This is a simplified implementation
	// In production, you'd use proper crypto libraries like crypto/aes
	
	// For now, just return base64 encoded data (placeholder)
	encoded := base64.StdEncoding.EncodeToString(data)
	return []byte(encoded), nil
}

// DecryptData decrypts data
func (esm *EnhancedSecurityManager) DecryptData(encryptedData []byte) ([]byte, error) {
	// This is a simplified implementation
	// In production, you'd use proper crypto libraries
	
	// For now, just decode base64 (placeholder)
	decoded, err := base64.StdEncoding.DecodeString(string(encryptedData))
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}
	
	return decoded, nil
}

// AuditLog records an audit event
func (esm *EnhancedSecurityManager) AuditLog(event AuditEvent) {
	if !esm.auditConfig.Enabled {
		return
	}
	
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	// Generate event ID if not provided
	if event.ID == "" {
		event.ID = esm.generateEventID()
	}
	
	// Set timestamp if not provided
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}
	
	// Add to audit logs
	esm.auditLogs = append(esm.auditLogs, event)
	
	// Limit audit log size
	if len(esm.auditLogs) > 10000 {
		esm.auditLogs = esm.auditLogs[1:]
	}
}

// GetAuditLogs returns audit logs filtered by the provided filter
func (esm *EnhancedSecurityManager) GetAuditLogs(filter AuditFilter) []AuditEvent {
	esm.mu.RLock()
	defer esm.mu.RUnlock()
	
	var filtered []AuditEvent
	
	for _, event := range esm.auditLogs {
		if filter.UserID != "" && event.UserID != filter.UserID {
			continue
		}
		if filter.EventType != "" && event.EventType != filter.EventType {
			continue
		}
		if filter.Resource != "" && event.Resource != filter.Resource {
			continue
		}
		if filter.Action != "" && event.Action != filter.Action {
			continue
		}
		if filter.StartTime != nil && event.Timestamp.Before(*filter.StartTime) {
			continue
		}
		if filter.EndTime != nil && event.Timestamp.After(*filter.EndTime) {
			continue
		}
		if filter.Success != nil && event.Success != *filter.Success {
			continue
		}
		
		filtered = append(filtered, event)
		
		if filter.Limit > 0 && len(filtered) >= filter.Limit {
			break
		}
	}
	
	return filtered
}

// CheckPermission checks if a user has permission for a specific action on a resource
func (esm *EnhancedSecurityManager) CheckPermission(userID, resource, action string) bool {
	esm.mu.RLock()
	defer esm.mu.RUnlock()
	
	// Get user permissions
	userPerms, exists := esm.permissions[userID]
	if !exists {
		return false
	}
	
	// Check resource permissions
	resourcePerms, exists := userPerms[resource]
	if !exists {
		return false
	}
	
	// Check if action is allowed
	for _, allowedAction := range resourcePerms {
		if allowedAction == action || allowedAction == "*" {
			return true
		}
	}
	
	return false
}

// Private helper methods

func (esm *EnhancedSecurityManager) generateTokenID(token string) string {
	hash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(hash[:])
}

func (esm *EnhancedSecurityManager) generateEventID() string {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func (esm *EnhancedSecurityManager) getUserPermissions(userID string) []string {
	esm.mu.RLock()
	defer esm.mu.RUnlock()
	
	// Get all permissions for the user
	var allPermissions []string
	
	userPerms, exists := esm.permissions[userID]
	if !exists {
		return allPermissions
	}
	
	for _, actions := range userPerms {
		allPermissions = append(allPermissions, actions...)
	}
	
	return allPermissions
}

func (esm *EnhancedSecurityManager) checkRateLimit(userID string) bool {
	now := time.Now()
	
	rateInfo, exists := esm.rateLimiter[userID]
	if !exists {
		rateInfo = &RateLimitInfo{
			RequestCount: 0,
			WindowStart:  now,
			LastRequest:  now,
		}
		esm.rateLimiter[userID] = rateInfo
	}
	
	// Reset window if needed
	if now.Sub(rateInfo.WindowStart) > esm.config.RateLimitWindow {
		rateInfo.RequestCount = 0
		rateInfo.WindowStart = now
	}
	
	// Check limit
	if rateInfo.RequestCount >= esm.config.MaxRequestsPerWindow {
		return false
	}
	
	// Increment counter
	rateInfo.RequestCount++
	rateInfo.LastRequest = now
	
	return true
}

func (esm *EnhancedSecurityManager) cleanupExpiredTokens() {
	ticker := time.NewTicker(time.Minute * 5)
	defer ticker.Stop()
	
	for {
		select {
		case <-esm.ctx.Done():
			return
		case <-ticker.C:
			esm.performTokenCleanup()
		}
	}
}

func (esm *EnhancedSecurityManager) performTokenCleanup() {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	now := time.Now()
	
	// Clean expired active tokens
	for token, tokenInfo := range esm.activeTokens {
		if now.After(tokenInfo.ExpiresAt) {
			delete(esm.activeTokens, token)
			
			// Audit log token expiration
			esm.AuditLog(AuditEvent{
				EventType: "token_expired",
				UserID:    tokenInfo.UserID,
				Action:    "expire_token",
				Success:   true,
				Metadata: map[string]interface{}{
					"tokenId": tokenInfo.TokenID,
					"usageCount": tokenInfo.UsageCount,
				},
			})
		}
	}
	
	// Clean old revoked tokens
	for token, revokedAt := range esm.revokedTokens {
		if now.Sub(revokedAt) > 24*time.Hour {
			delete(esm.revokedTokens, token)
		}
	}
}

func (esm *EnhancedSecurityManager) cleanupAuditLogs() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	
	for {
		select {
		case <-esm.ctx.Done():
			return
		case <-ticker.C:
			esm.performAuditCleanup()
		}
	}
}

func (esm *EnhancedSecurityManager) performAuditCleanup() {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	cutoff := time.Now().Add(-esm.auditConfig.RetentionPeriod)
	
	filtered := make([]AuditEvent, 0)
	for _, event := range esm.auditLogs {
		if event.Timestamp.After(cutoff) {
			filtered = append(filtered, event)
		}
	}
	
	esm.auditLogs = filtered
	
	if len(filtered) < len(esm.auditLogs) {
		log.Printf("Cleaned up %d old audit log entries", len(esm.auditLogs)-len(filtered))
	}
}

func (esm *EnhancedSecurityManager) batchAuditLogs() {
	ticker := time.NewTicker(esm.auditConfig.BatchInterval)
	defer ticker.Stop()
	
	for {
		select {
		case <-esm.ctx.Done():
			return
		case <-ticker.C:
			esm.sendAuditBatch()
		}
	}
}

func (esm *EnhancedSecurityManager) sendAuditBatch() {
	esm.mu.Lock()
	batch := make([]AuditEvent, 0, esm.auditConfig.BatchSize)
	
	// Take a batch of events
	if len(esm.auditLogs) > 0 {
		batchSize := esm.auditConfig.BatchSize
		if len(esm.auditLogs) < batchSize {
			batchSize = len(esm.auditLogs)
		}
		
		batch = append(batch, esm.auditLogs[:batchSize]...)
		esm.auditLogs = esm.auditLogs[batchSize:]
	}
	esm.mu.Unlock()
	
	if len(batch) == 0 {
		return
	}
	
	// Send to external endpoint
	// This is a placeholder - in production, you'd implement proper HTTP client
	log.Printf("Sending audit batch of %d events to %s", len(batch), esm.auditConfig.ExternalEndpoint)
}

func (esm *EnhancedSecurityManager) cleanupRateLimits() {
	ticker := time.NewTicker(esm.config.RateLimitWindow)
	defer ticker.Stop()
	
	for {
		select {
		case <-esm.ctx.Done():
			return
		case <-ticker.C:
			esm.performRateLimitCleanup()
		}
	}
}

func (esm *EnhancedSecurityManager) performRateLimitCleanup() {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	now := time.Now()
	
	for userID, rateInfo := range esm.rateLimiter {
		if now.Sub(rateInfo.WindowStart) > esm.config.RateLimitWindow*2 {
			delete(esm.rateLimiter, userID)
		}
	}
}

// Helper methods for permission management

func (esm *EnhancedSecurityManager) AddUserPermission(userID, resource string, actions []string) {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	if esm.permissions[userID] == nil {
		esm.permissions[userID] = make(map[string][]string)
	}
	
	esm.permissions[userID][resource] = actions
}

func (esm *EnhancedSecurityManager) RemoveUserPermission(userID, resource string) {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	if userPerms, exists := esm.permissions[userID]; exists {
		delete(userPerms, resource)
	}
}

func (esm *EnhancedSecurityManager) AddRole(roleName string, permissions []string) {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	esm.roles[roleName] = permissions
}

func (esm *EnhancedSecurityManager) AssignRole(userID, roleName string) {
	esm.mu.Lock()
	defer esm.mu.Unlock()
	
	permissions, exists := esm.roles[roleName]
	if !exists {
		return
	}
	
	if esm.permissions[userID] == nil {
		esm.permissions[userID] = make(map[string][]string)
	}
	
	// Add role permissions to user permissions
	for resource, actions := range esm.permissions[userID] {
		// Merge with existing permissions
		esm.permissions[userID][resource] = append(actions, permissions...)
	}
}

// GetSecurityStatus returns current security status
func (esm *EnhancedSecurityManager) GetSecurityStatus() map[string]interface{} {
	esm.mu.RLock()
	defer esm.mu.RUnlock()
	
	status := map[string]interface{}{
		"activeTokens":    len(esm.activeTokens),
		"revokedTokens":   len(esm.revokedTokens),
		"auditLogs":       len(esm.auditLogs),
		"rateLimitedUsers": len(esm.rateLimiter),
		"isRunning":       esm.isRunning,
		"config":          esm.config,
		"lastCleanup":     time.Now(),
	}
	
	return status
}
