package config

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// config is a struct for saving game env variables
type config struct {
	Environment            string `yaml:"environment"`
	AppName                string `yaml:"app_name"`
	Port                   string `yaml:"port"`
	Salt                   string `yaml:"salt"`
	JwtSecret              string `yaml:"jwt_secret"`
	BasicMasterDatabaseDSN string `yaml:"basic_master_database_dsn"`
	BasicSlaveDatabaseDSN  string `yaml:"basic_slave_database_dsn"`
	ActivityDatabaseDSN    string `yaml:"activity_database_dsn"`
	LogPath                string `yaml:"log_path"`
	LogLevel               string `yaml:"log_level"`
	DefaultLang            string `yaml:"default_lang"`
	OriginalError          bool   `yaml:"original_error"`

	basicMasterDB *sql.DB
	basicSlaveDB  *sql.DB
	activityDB    *sql.DB
}

// Config make accessible to private config variables
type Config struct {
	config config
}

const jwtLocalSecret = "81027ac7103d791abacd19ac9f1e8722c19ad6c9"

// GetConf load config file for game
func GetConf(configPath string) (cfg Config, err error) {

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		err = fmt.Errorf("error in reading game config file; %w", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &cfg.config)
	if err != nil {
		err = fmt.Errorf("error in unmarshal game config file; %w", err)
		return
	}

	// set up database connections
	cfg.config.basicMasterDB, err = sql.Open("mysql", cfg.GetDatabaseMasterDNS())
	if err != nil {
		err = fmt.Errorf("error in opening basic master database; %w", err)
		return
	}

	cfg.config.basicSlaveDB, err = sql.Open("mysql", cfg.GetDatabaseSlaveDNS())
	if err != nil {
		err = fmt.Errorf("error in opening basic slave database; %w", err)
		return
	}

	cfg.config.activityDB, err = sql.Open("mysql", cfg.GetDatabaseActivityDNS())
	if err != nil {
		err = fmt.Errorf("error in opening activity database; %w", err)
		return
	}
	rand.Seed(time.Now().UnixNano())

	return
}

// configs
func (c Config) GetEnvironment() string {
	return c.config.Environment
}

func (c Config) GetAppName() string {
	return c.config.AppName
}

func (c Config) GetPort() string {
	return c.config.Port
}

func (c Config) GetSalt() string {
	return c.config.Salt
}

func (c Config) GetJwtSecret() string {
	return c.config.JwtSecret
}

func (c Config) GetLocalJwtSecret() string {
	return jwtLocalSecret
}

func (c Config) GetLogPath() string {
	return c.config.LogPath
}

func (c Config) GetLogLevel() string {
	return c.config.LogLevel
}

func (c Config) GetDefaultLang() string {
	return "en"
}

func (c Config) GetDatabaseMasterDNS() string {
	return c.config.BasicMasterDatabaseDSN
}

func (c Config) GetDatabaseSlaveDNS() string {
	return c.config.BasicSlaveDatabaseDSN
}

func (c Config) GetDatabaseActivityDNS() string {
	return c.config.ActivityDatabaseDSN
}

func (c Config) ShowOriginalError() bool {
	return c.config.OriginalError
}

// generated
func (c Config) BasicMasterDB() *sql.DB {
	return c.config.basicMasterDB
}

func (c Config) BasicSlaveDB() *sql.DB {
	return c.config.basicSlaveDB
}

func (c Config) ActivityDB() *sql.DB {
	return c.config.activityDB
}
