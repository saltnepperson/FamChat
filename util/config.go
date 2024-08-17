package util

import (
	"fmt"

  "github.com/spf13/viper"
)

type Config struct {
  DBDriver  string `mapstructure:"DB_DRIVER"`
  DBHost  string `mapstructure:"DATABASE_HOST"`
  DBDatabase string `mapstructure:"POSTGRES_DB"`
  DBUsername  string `mapstructure:"POSTGRES_USER"`
  DBPassword  string `mapstructure:"POSTGRES_PASSWORD"`
  DBPort  string `mapstructure:"DATABASE_PORT"`
	DBSslMode string `mapstructure:"DB_SSL_MODE"`
}

func LoadConfig(path string) (config Config, err error) {
  viper.AddConfigPath(path)
  viper.SetConfigFile(".env")
  viper.AutomaticEnv()

  err = viper.ReadInConfig()
  if err != nil {
    return
  }	

  err = viper.Unmarshal(&config)
  return
}

func (c *Config) BuildDBSource() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBDatabase, c.DBSslMode)
}
