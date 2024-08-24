package util

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

// Config holds the application configuration
type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DATABASE_HOST"`
	DBDatabase        string `mapstructure:"POSTGRES_DB"`
	DBUsername        string `mapstructure:"POSTGRES_USER"`
	DBPassword        string `mapstructure:"POSTGRES_PASSWORD"`
	DBPort            string `mapstructure:"DATABASE_PORT"`
	DBSource          string `mapstructure:"-"`
	DBSslMode         string `mapstructure:"DB_SSL_MODE"`
	DBMigrationPath   string `mapstructure:"DB_MIGRATION_PATH"`
	DBMaxOpenConns    int    `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
	DBMaxIdleConns    int    `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	DBConnMaxLifetime int    `mapstructure:"DB_CONNECTION_MAX_LIFETIME"`
}

// LoadConfig reads the configuration from file and/or environment variables
func LoadConfig(path string) (config Config, err error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	// Set reasonable defeault for specific configurations
	v.SetDefault("DB_MAX_OPEN_CONNECTIONS", 10)
	v.SetDefault("DB_MAX_IDLE_CONNECTIONS", 5)
	v.SetDefault("DB_CONNECTION_MAX_LIFETIME", 3600)

	err = v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, fmt.Errorf("Error reading config file: %w", err)
		}
		// config file not found; ignore error
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("Unable to decode config into struct: %w", err)
	}

	// Validate required fields
	if err := config.validate(); err != nil {
		return config, fmt.Errorf("Config validation error: %w", err)
	}

	// Construct the DBSource
	config.buildDBSource()

	return config, nil
}

// validate checks that all required fields are set
func (c *Config) validate() error {
	missingFields := []string{}

	if c.DBDriver == "" {
		missingFields = append(missingFields, "DB_DRIVER")
	}

	if c.DBHost == "" {
		missingFields = append(missingFields, "DATABASE_HOST")
	}

	if c.DBDatabase == "" {
		missingFields = append(missingFields, "POSTGRES_DB")
	}

	if c.DBUsername == "" {
		missingFields = append(missingFields, "POSTGRES_USER")
	}

	if c.DBPassword == "" {
		missingFields = append(missingFields, "POSTGRES_PASSWORD")
	}

	if c.DBPort == "" {
		missingFields = append(missingFields, "DATABASE_PORT")
	}

	if c.DBSslMode == "" {
		missingFields = append(missingFields, "DB_SSL_MODE")
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("Missing required config fields: %s", strings.Join(missingFields, ", "))
	}

	return nil
}

func (c *Config) buildDBSource() {
	c.DBSource = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBDatabase, c.DBSslMode)
}
