package config

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"

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
	logger        *zap.Logger
}

// Core make accessible to private config variables
type Core struct {
	config config
}

const jwtLocalSecret = "81027ac7103d791abacd19ac9f1e8722c19ad6c9"

// GetConf load config file for game
func GetConf(configPath string) (cfg Core, err error) {
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

	// set up zap logger
	logFile, err := os.OpenFile(cfg.GetLogPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error in opening log file: %v\n", err)
	}
	// defer logFile.Close()

	cfg.config.logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(logFile),
		cfg.GetLogLevel(),
	))

	return
}

// configs
func (c Core) GetEnvironment() string {
	return c.config.Environment
}

func (c Core) GetAppName() string {
	return c.config.AppName
}

func (c Core) GetPort() string {
	return c.config.Port
}

func (c Core) GetSalt() string {
	return c.config.Salt
}

func (c Core) GetJwtSecret() string {
	return c.config.JwtSecret
}

func (c Core) GetLocalJwtSecret() string {
	return jwtLocalSecret
}

func (c Core) GetLogPath() string {
	return c.config.LogPath
}

func (c Core) GetLogLevel() zapcore.Level {
	switch c.config.LogLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	}

	log.Fatalf("log level is not valid: %v\n", c.config.LogLevel)

	return 0
}

func (c Core) GetDefaultLang() string {
	return "en"
}

func (c Core) GetDatabaseMasterDNS() string {
	return c.config.BasicMasterDatabaseDSN
}

func (c Core) GetDatabaseSlaveDNS() string {
	return c.config.BasicSlaveDatabaseDSN
}

func (c Core) GetDatabaseActivityDNS() string {
	return c.config.ActivityDatabaseDSN
}

func (c Core) ShowOriginalError() bool {
	return c.config.OriginalError
}

// generated
func (c Core) BasicMasterDB() *sql.DB {
	return c.config.basicMasterDB
}

func (c Core) BasicSlaveDB() *sql.DB {
	return c.config.basicSlaveDB
}

func (c Core) ActivityDB() *sql.DB {
	return c.config.activityDB
}

func (c Core) Log() *zap.Logger {
	return c.config.logger
}
