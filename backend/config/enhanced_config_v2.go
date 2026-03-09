package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Config represents the application configuration
type Config struct {
	Server           ServerConfig           `json:"server"`
	Temporal         TemporalConfig         `json:"temporal"`
	Database         DatabaseConfig         `json:"database"`
	Security         SecurityConfig         `json:"security"`
	Monitoring       MonitoringConfig       `json:"monitoring"`
	Workflow         WorkflowConfig         `json:"workflow"`
	AI               AIConfig               `json:"ai"`
	Infrastructure   InfrastructureConfig   `json:"infrastructure"`
	Notification     NotificationConfig     `json:"notification"`
	Features         FeatureFlags           `json:"features"`
}

// Server configuration
type ServerConfig struct {
	Host                 string        `json:"host"`
	Port                 string        `json:"port"`
	ReadTimeout          time.Duration `json:"readTimeout"`
	WriteTimeout         time.Duration `json:"writeTimeout"`
	IdleTimeout          time.Duration `json:"idleTimeout"`
	EnableTLS            bool          `json:"enableTLS"`
	TLSCertFile          string        `json:"tlsCertFile"`
	TLSKeyFile           string        `json:"tlsKeyFile"`
	EnableCORS          bool          `json:"enableCORS"`
	EnableMetrics        bool          `json:"enableMetrics"`
	EnableProfiling     bool          `json:"enableProfiling"`
	MaxRequestSize       int64         `json:"maxRequestSize"`
	RateLimiting         RateLimitConfig `json:"rateLimiting"`
}

type RateLimitConfig struct {
	Enabled              bool          `json:"enabled"`
	RequestsPerMinute    int           `json:"requestsPerMinute"`
	BurstSize            int           `json:"burstSize"`
	WhitelistIPs         []string      `json:"whitelistIPs"`
	BlacklistIPs         []string      `json:"blacklistIPs"`
}

// Temporal configuration
type TemporalConfig struct {
	Host                 string        `json:"host"`
	Port                 string        `json:"port"`
	TaskQueue            string        `json:"taskQueue"`
	Namespace            string        `json:"namespace"`
	MaxConcurrentWorkers int           `json:"maxConcurrentWorkers"`
	WorkerOptions        WorkerOptions `json:"workerOptions"`
	ClientOptions        ClientOptions `json:"clientOptions"`
}

type WorkerOptions struct {
	MaxConcurrentActivityExecutionSize int           `json:"maxConcurrentActivityExecutionSize"`
	MaxConcurrentWorkflowTaskExecutionSize int        `json:"maxConcurrentWorkflowTaskExecutionSize"`
	WorkflowPollThrottle               time.Duration `json:"workflowPollThrottle"`
	ActivityPollThrottle               time.Duration `json:"activityPollThrottle"`
	StickyWorkflowToScheduleTimeout    time.Duration `json:"stickyWorkflowToScheduleTimeout"`
	WorkerActivitiesPerSecond          float64       `json:"workerActivitiesPerSecond"`
	WorkerWorkflowTasksPerSecond       float64       `json:"workerWorkflowTasksPerSecond"`
}

type ClientOptions struct {
	HostPort                      string        `json:"hostPort"`
	Namespace                     string        `json:"namespace"`
	Identity                      string        `json:"identity"`
	DataConverter                 string        `json:"dataConverter"`
	Headers                       map[string]string `json:"headers"`
	RPCTimeout                    time.Duration `json:"rpcTimeout"`
	WorkflowExecutionTimeout      time.Duration `json:"workflowExecutionTimeout"`
	WorkflowTaskTimeout           time.Duration `json:"workflowTaskTimeout"`
	ActivityTimeout               time.Duration `json:"activityTimeout"`
	ContextTimeout                time.Duration `json:"contextTimeout"`
	GRPCDialOptions               GRPCDialOptions `json:"grpcDialOptions"`
}

type GRPCDialOptions struct {
	MaxMessageSize             int           `json:"maxMessageSize"`
	KeepAliveTime              time.Duration `json:"keepAliveTime"`
	KeepAliveTimeout           time.Duration `json:"keepAliveTimeout"`
	PerRPCCredentials          []string      `json:"perRPCCredentials"`
	TransportCredentials       string        `json:"transportCredentials"`
	DisableTLS                 bool          `json:"disableTLS"`
}

// Database configuration
type DatabaseConfig struct {
	Type                 string        `json:"type"`
	Host                 string        `json:"host"`
	Port                 string        `json:"port"`
	Database             string        `json:"database"`
	Username             string        `json:"username"`
	Password             string        `json:"password"`
	SSLMode              string        `json:"sslMode"`
	MaxOpenConnections   int           `json:"maxOpenConnections"`
	MaxIdleConnections   int           `json:"maxIdleConnections"`
	ConnectionMaxLifetime time.Duration `json:"connectionMaxLifetime"`
	ConnectionMaxIdleTime time.Duration  `json:"connectionMaxIdleTime"`
	EnableMetrics        bool          `json:"enableMetrics"`
	EnableLogging        bool          `json:"enableLogging"`
}

// Security configuration
type SecurityConfig struct {
	Enabled              bool          `json:"enabled"`
	TokenSecret          string        `json:"tokenSecret"`
	TokenExpiration      time.Duration `json:"tokenExpiration"`
	RefreshTokenExpiration time.Duration `json:"refreshTokenExpiration"`
	EnableRateLimiting   bool          `json:"enableRateLimiting"`
	RateLimitWindow      time.Duration `json:"rateLimitWindow"`
	MaxRequestsPerWindow int           `json:"maxRequestsPerWindow"`
	EnableAuditLogging   bool          `json:"enableAuditLogging"`
	AuditRetentionPeriod time.Duration `json:"auditRetentionPeriod"`
	EncryptionKey        string        `json:"encryptionKey"`
	EnableEncryption     bool          `json:"enableEncryption"`
	AllowedOrigins       []string      `json:"allowedOrigins"`
	AllowedMethods       []string      `json:"allowedMethods"`
	AllowedHeaders       []string      `json:"allowedHeaders"`
	ExposeHeaders        []string      `json:"exposeHeaders"`
}

// Monitoring configuration
type MonitoringConfig struct {
	Enabled              bool          `json:"enabled"`
	MetricsInterval      time.Duration `json:"metricsInterval"`
	RetentionPeriod      time.Duration `json:"retentionPeriod"`
	MaxMetricsPerEntity  int           `json:"maxMetricsPerEntity"`
	EnableRealTime       bool          `json:"enableRealTime"`
	EnableAggregation    bool          `json:"enableAggregation"`
	ExportFormat         string        `json:"exportFormat"`
	ExportInterval       time.Duration `json:"exportInterval"`
	ExternalEndpoint     string        `json:"externalEndpoint"`
	EnableTracing        bool          `json:"enableTracing"`
	TracingEndpoint      string        `json:"tracingEndpoint"`
	TracingSampleRate    float64       `json:"tracingSampleRate"`
	EnableHealthCheck    bool          `json:"enableHealthCheck"`
	HealthCheckInterval  time.Duration `json:"healthCheckInterval"`
}

// Workflow configuration
type WorkflowConfig struct {
	DefaultTimeout       time.Duration `json:"defaultTimeout"`
	DefaultRetryPolicy   RetryPolicy   `json:"defaultRetryPolicy"`
	MaxWorkflowDuration  time.Duration `json:"maxWorkflowDuration"`
	EnableWorkflowCaching bool         `json:"enableWorkflowCaching"`
	WorkflowCacheSize    int           `json:"workflowCacheSize"`
	EnableActivityCaching bool         `json:"enableActivityCaching"`
	ActivityCacheSize    int           `json:"activityCacheSize"`
	EnableDeadlockDetection bool       `json:"enableDeadlockDetection"`
	DeadlockCheckInterval time.Duration `json:"deadlockCheckInterval"`
}

type RetryPolicy struct {
	InitialInterval    time.Duration `json:"initialInterval"`
	BackoffCoefficient float64       `json:"backoffCoefficient"`
	MaximumInterval    time.Duration `json:"maximumInterval"`
	MaximumAttempts    int           `json:"maximumAttempts"`
	NonRetryableErrorTypes []string   `json:"nonRetryableErrorTypes"`
}

// AI configuration
type AIConfig struct {
	Enabled              bool          `json:"enabled"`
	Provider             string        `json:"provider"`
	APIKey               string        `json:"apiKey"`
	APISecret            string        `json:"apiSecret"`
	BaseURL              string        `json:"baseUrl"`
	Model                string        `json:"model"`
	MaxTokens            int           `json:"maxTokens"`
	Temperature          float64       `json:"temperature"`
	Timeout              time.Duration `json:"timeout"`
	MaxRetries           int           `json:"maxRetries"`
	RetryInterval        time.Duration `json:"retryInterval"`
	EnableCaching        bool          `json:"enableCaching"`
	CacheSize            int           `json:"cacheSize"`
	CacheTTL             time.Duration `json:"cacheTTL"`
	EnableRateLimiting   bool          `json:"enableRateLimiting"`
	RateLimitPerMinute   int           `json:"rateLimitPerMinute"`
	EnableLogging        bool          `json:"enableLogging"`
	LogLevel             string        `json:"logLevel"`
}

// Infrastructure configuration
type InfrastructureConfig struct {
	Enabled              bool          `json:"enabled"`
	Provider             string        `json:"provider"`
	Region               string        `json:"region"`
	AccessKey            string        `json:"accessKey"`
	SecretKey            string        `json:"secretKey"`
	SessionToken         string        `json:"sessionToken"`
	EnableEmulator       bool          `json:"enableEmulator"`
	EmulatorConfig       EmulatorConfig `json:"emulatorConfig"`
	EnableAutoDiscovery  bool          `json:"enableAutoDiscovery"`
	DiscoveryInterval    time.Duration `json:"discoveryInterval"`
	EnableResourceCache  bool          `json:"enableResourceCache"`
	ResourceCacheSize    int           `json:"resourceCacheSize"`
	ResourceCacheTTL     time.Duration `json:"resourceCacheTTL"`
}

type EmulatorConfig struct {
	Enabled              bool          `json:"enabled"`
	DataDirectory        string        `json:"dataDirectory"`
	EnablePersistence    bool          `json:"enablePersistence"`
	SyncInterval         time.Duration `json:"syncInterval"`
	MaxResources         int           `json:"maxResources"`
	EnableMetrics        bool          `json:"enableMetrics"`
	EnableLogging        bool          `json:"enableLogging"`
	LogLevel             string        `json:"logLevel"`
}

// Notification configuration
type NotificationConfig struct {
	Enabled              bool          `json:"enabled"`
	Providers            []NotificationProvider `json:"providers"`
	DefaultChannel       string        `json:"defaultChannel"`
	EnableBatching       bool          `json:"enableBatching"`
	BatchSize            int           `json:"batchSize"`
	BatchInterval        time.Duration `json:"batchInterval"`
	EnableRetry          bool          `json:"enableRetry"`
	MaxRetries           int           `json:"maxRetries"`
	RetryInterval        time.Duration `json:"retryInterval"`
	EnableQueue          bool          `json:"enableQueue"`
	QueueSize            int           `json:"queueSize"`
	QueueTimeout         time.Duration `json:"queueTimeout"`
}

type NotificationProvider struct {
	Name                 string            `json:"name"`
	Type                 string            `json:"type"`
	Config               map[string]interface{} `json:"config"`
	Enabled              bool              `json:"enabled"`
	Priority             int               `json:"priority"`
}

// Feature flags
type FeatureFlags struct {
	EnableAdvancedWorkflows bool `json:"enableAdvancedWorkflows"`
	EnableMultiTenant      bool `json:"enableMultiTenant"`
	EnableGraphQL          bool `json:"enableGraphQL"`
	EnableWebSocket       bool `json:"enableWebSocket"`
	EnableFileUpload       bool `json:"enableFileUpload"`
	EnableBackupRestore    bool `json:"enableBackupRestore"`
	EnableAPIKeyAuth       bool `json:"enableAPIKeyAuth"`
	EnableOAuth2Auth       bool `json:"enableOAuth2Auth"`
	EnableLDAPAuth         bool `json:"enableLDAPAuth"`
	EnableSAMLAuth         bool `json:"enableSAMLAuth"`
	EnableCustomThemes     bool `json:"enableCustomThemes"`
	EnablePluginSystem     bool `json:"enablePluginSystem"`
	EnableVersioning       bool `json:"enableVersioning"`
}

// ConfigManager manages configuration loading and validation
type ConfigManager struct {
	config *Config
	path   string
}

// NewConfigManager creates a new configuration manager
func NewConfigManager(configPath string) (*ConfigManager, error) {
	cm := &ConfigManager{
		path: configPath,
	}
	
	if err := cm.loadConfig(); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	
	if err := cm.validateConfig(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	
	return cm, nil
}

// loadConfig loads configuration from file or environment variables
func (cm *ConfigManager) loadConfig() error {
	// Start with default configuration
	cm.config = cm.getDefaultConfig()
	
	// Try to load from file
	if cm.path != "" {
		if _, err := os.Stat(cm.path); err == nil {
			file, err := os.Open(cm.path)
			if err != nil {
				return fmt.Errorf("failed to open config file: %w", err)
			}
			defer file.Close()
			
			decoder := json.NewDecoder(file)
			if err := decoder.Decode(cm.config); err != nil {
				return fmt.Errorf("failed to decode config file: %w", err)
			}
		}
	}
	
	// Override with environment variables
	cm.overrideWithEnv()
	
	return nil
}

// getDefaultConfig returns the default configuration
func (cm *ConfigManager) getDefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host:            "0.0.0.0",
			Port:            "8081",
			ReadTimeout:     15 * time.Second,
			WriteTimeout:    15 * time.Second,
			IdleTimeout:     60 * time.Second,
			EnableTLS:       false,
			EnableCORS:      true,
			EnableMetrics:   true,
			EnableProfiling: false,
			MaxRequestSize:  10 * 1024 * 1024, // 10MB
			RateLimiting: RateLimitConfig{
				Enabled:           true,
				RequestsPerMinute: 100,
				BurstSize:         20,
			},
		},
		Temporal: TemporalConfig{
			Host:                 "localhost",
			Port:                 "7233",
			TaskQueue:            "ai-agent-task-queue-v2",
			Namespace:            "default",
			MaxConcurrentWorkers: 10,
			WorkerOptions: WorkerOptions{
				MaxConcurrentActivityExecutionSize:     10,
				MaxConcurrentWorkflowTaskExecutionSize: 10,
				WorkflowPollThrottle:                  100 * time.Millisecond,
				ActivityPollThrottle:                  100 * time.Millisecond,
				WorkerActivitiesPerSecond:             100.0,
				WorkerWorkflowTasksPerSecond:          100.0,
			},
			ClientOptions: ClientOptions{
				HostPort:                 "localhost:7233",
				Namespace:                "default",
				Identity:                 "ai-agent-worker",
				RPCTimeout:               30 * time.Second,
				WorkflowExecutionTimeout: 24 * time.Hour,
				WorkflowTaskTimeout:      10 * time.Second,
				ActivityTimeout:          30 * time.Second,
				ContextTimeout:           10 * time.Second,
				GRPCDialOptions: GRPCDialOptions{
					MaxMessageSize: 1024 * 1024 * 4, // 4MB
					KeepAliveTime:  30 * time.Second,
					DisableTLS:     true,
				},
			},
		},
		Database: DatabaseConfig{
			Type:                     "sqlite",
			Host:                     "localhost",
			Port:                     "5432",
			Database:                 "ai_agent_db",
			MaxOpenConnections:       25,
			MaxIdleConnections:       5,
			ConnectionMaxLifetime:    5 * time.Minute,
			ConnectionMaxIdleTime:    5 * time.Minute,
			EnableMetrics:            true,
			EnableLogging:            true,
		},
		Security: SecurityConfig{
			Enabled:               true,
			TokenExpiration:        24 * time.Hour,
			RefreshTokenExpiration: 7 * 24 * time.Hour,
			EnableRateLimiting:     true,
			RateLimitWindow:        time.Minute,
			MaxRequestsPerWindow:   100,
			EnableAuditLogging:     true,
			AuditRetentionPeriod:   30 * 24 * time.Hour,
			EnableEncryption:       true,
			AllowedOrigins:         []string{"*"},
			AllowedMethods:         []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:         []string{"Content-Type", "Authorization", "X-Requested-With"},
		},
		Monitoring: MonitoringConfig{
			Enabled:             true,
			MetricsInterval:     30 * time.Second,
			RetentionPeriod:     24 * time.Hour,
			MaxMetricsPerEntity: 1000,
			EnableRealTime:      true,
			EnableAggregation:   true,
			ExportFormat:        "json",
			ExportInterval:      time.Minute * 5,
			EnableTracing:       false,
			TracingSampleRate:   0.1,
			EnableHealthCheck:   true,
			HealthCheckInterval: 30 * time.Second,
		},
		Workflow: WorkflowConfig{
			DefaultTimeout:            30 * time.Minute,
			MaxWorkflowDuration:       24 * time.Hour,
			EnableWorkflowCaching:     true,
			WorkflowCacheSize:         1000,
			EnableActivityCaching:     true,
			ActivityCacheSize:         1000,
			EnableDeadlockDetection:   true,
			DeadlockCheckInterval:     time.Minute * 5,
			DefaultRetryPolicy: RetryPolicy{
				InitialInterval:    time.Second,
				BackoffCoefficient: 2.0,
				MaximumInterval:    time.Minute,
				MaximumAttempts:    3,
				NonRetryableErrorTypes: []string{"ValidationError", "AuthenticationError"},
			},
		},
		AI: AIConfig{
			Enabled:            true,
			Provider:           "openai",
			Model:              "gpt-3.5-turbo",
			MaxTokens:          2048,
			Temperature:        0.7,
			Timeout:            30 * time.Second,
			MaxRetries:         3,
			RetryInterval:      time.Second * 2,
			EnableCaching:      true,
			CacheSize:          1000,
			CacheTTL:           time.Hour,
			EnableRateLimiting: true,
			RateLimitPerMinute: 60,
			EnableLogging:      true,
			LogLevel:           "info",
		},
		Infrastructure: InfrastructureConfig{
			Enabled:             true,
			Provider:            "aws",
			Region:              "us-west-2",
			EnableEmulator:      true,
			EmulatorConfig: EmulatorConfig{
				Enabled:           true,
				DataDirectory:     "./emulator_data",
				EnablePersistence: true,
				SyncInterval:      time.Minute * 5,
				MaxResources:      1000,
				EnableMetrics:      true,
				EnableLogging:      true,
				LogLevel:          "info",
			},
			EnableAutoDiscovery: true,
			DiscoveryInterval:   time.Minute * 10,
			EnableResourceCache: true,
			ResourceCacheSize:   1000,
			ResourceCacheTTL:    time.Hour,
		},
		Notification: NotificationConfig{
			Enabled:       true,
			DefaultChannel: "email",
			EnableBatching: true,
			BatchSize:     100,
			BatchInterval: time.Minute * 5,
			EnableRetry:   true,
			MaxRetries:    3,
			RetryInterval: time.Second * 5,
			EnableQueue:   true,
			QueueSize:     1000,
			QueueTimeout:  time.Minute * 5,
		},
		Features: FeatureFlags{
			EnableAdvancedWorkflows: true,
			EnableMultiTenant:       false,
			EnableGraphQL:           false,
			EnableWebSocket:        false,
			EnableFileUpload:        false,
			EnableBackupRestore:     false,
			EnableAPIKeyAuth:        false,
			EnableOAuth2Auth:        false,
			EnableLDAPAuth:          false,
			EnableSAMLAuth:          false,
			EnableCustomThemes:      false,
			EnablePluginSystem:      false,
			EnableVersioning:        false,
		},
	}
}

// overrideWithEnv overrides configuration with environment variables
func (cm *ConfigManager) overrideWithEnv() {
	// Server configuration
	if host := os.Getenv("SERVER_HOST"); host != "" {
		cm.config.Server.Host = host
	}
	if port := os.Getenv("SERVER_PORT"); port != "" {
		cm.config.Server.Port = port
	}
	
	// Temporal configuration
	if host := os.Getenv("TEMPORAL_HOST"); host != "" {
		cm.config.Temporal.Host = host
	}
	if port := os.Getenv("TEMPORAL_PORT"); port != "" {
		cm.config.Temporal.Port = port
	}
	if namespace := os.Getenv("TEMPORAL_NAMESPACE"); namespace != "" {
		cm.config.Temporal.Namespace = namespace
	}
	
	// Database configuration
	if dbType := os.Getenv("DB_TYPE"); dbType != "" {
		cm.config.Database.Type = dbType
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		cm.config.Database.Host = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		cm.config.Database.Port = dbPort
	}
	if dbDatabase := os.Getenv("DB_DATABASE"); dbDatabase != "" {
		cm.config.Database.Database = dbDatabase
	}
	if dbUsername := os.Getenv("DB_USERNAME"); dbUsername != "" {
		cm.config.Database.Username = dbUsername
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		cm.config.Database.Password = dbPassword
	}
	
	// Security configuration
	if tokenSecret := os.Getenv("TOKEN_SECRET"); tokenSecret != "" {
		cm.config.Security.TokenSecret = tokenSecret
	}
	if encryptionKey := os.Getenv("ENCRYPTION_KEY"); encryptionKey != "" {
		cm.config.Security.EncryptionKey = encryptionKey
	}
	
	// AI configuration
	if aiProvider := os.Getenv("AI_PROVIDER"); aiProvider != "" {
		cm.config.AI.Provider = aiProvider
	}
	if apiKey := os.Getenv("AI_API_KEY"); apiKey != "" {
		cm.config.AI.APIKey = apiKey
	}
	if apiSecret := os.Getenv("AI_API_SECRET"); apiSecret != "" {
		cm.config.AI.APISecret = apiSecret
	}
	if baseURL := os.Getenv("AI_BASE_URL"); baseURL != "" {
		cm.config.AI.BaseURL = baseURL
	}
	if model := os.Getenv("AI_MODEL"); model != "" {
		cm.config.AI.Model = model
	}
	
	// Infrastructure configuration
	if provider := os.Getenv("INFRA_PROVIDER"); provider != "" {
		cm.config.Infrastructure.Provider = provider
	}
	if region := os.Getenv("INFRA_REGION"); region != "" {
		cm.config.Infrastructure.Region = region
	}
	if accessKey := os.Getenv("INFRA_ACCESS_KEY"); accessKey != "" {
		cm.config.Infrastructure.AccessKey = accessKey
	}
	if secretKey := os.Getenv("INFRA_SECRET_KEY"); secretKey != "" {
		cm.config.Infrastructure.SecretKey = secretKey
	}
}

// validateConfig validates the configuration
func (cm *ConfigManager) validateConfig() error {
	// Validate server configuration
	if cm.config.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	
	// Validate temporal configuration
	if cm.config.Temporal.Host == "" {
		return fmt.Errorf("temporal host is required")
	}
	if cm.config.Temporal.Port == "" {
		return fmt.Errorf("temporal port is required")
	}
	if cm.config.Temporal.TaskQueue == "" {
		return fmt.Errorf("temporal task queue is required")
	}
	
	// Validate security configuration
	if cm.config.Security.Enabled && cm.config.Security.TokenSecret == "" {
		return fmt.Errorf("security token secret is required when security is enabled")
	}
	
	// Validate AI configuration
	if cm.config.AI.Enabled {
		if cm.config.AI.Provider == "" {
			return fmt.Errorf("AI provider is required when AI is enabled")
		}
		if cm.config.AI.APIKey == "" {
			return fmt.Errorf("AI API key is required when AI is enabled")
		}
	}
	
	// Validate infrastructure configuration
	if cm.config.Infrastructure.Enabled {
		if cm.config.Infrastructure.Provider == "" {
			return fmt.Errorf("infrastructure provider is required when infrastructure is enabled")
		}
	}
	
	return nil
}

// GetConfig returns the configuration
func (cm *ConfigManager) GetConfig() *Config {
	return cm.config
}

// SaveConfig saves the configuration to file
func (cm *ConfigManager) SaveConfig() error {
	if cm.path == "" {
		return fmt.Errorf("no config path specified")
	}
	
	file, err := os.Create(cm.path)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	
	if err := encoder.Encode(cm.config); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}
	
	return nil
}

// ReloadConfig reloads the configuration
func (cm *ConfigManager) ReloadConfig() error {
	return cm.loadConfig()
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	cm := &ConfigManager{}
	return cm.getDefaultConfig()
}
