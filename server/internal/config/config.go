package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"

	"github.com/aws/aws-sdk-go-v2/aws"
	dconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// config is a struct for saving game env variables
type config struct {
	Environment   string `yaml:"environment"`
	AppName       string `yaml:"app_name"`
	Port          string `yaml:"port"`
	Salt          string `yaml:"salt"`
	JwtSecret     string `yaml:"jwt_secret"`
	LogPath       string `yaml:"log_path"`
	LogLevel      string `yaml:"log_level"`
	DefaultLang   string `yaml:"default_lang"`
	OriginalError bool   `yaml:"original_error"`

	// Map domain
	MapDynamoDBRegion   string `yaml:"map_dynamodb_region"`
	MapDynamoDBEndpoint string `yaml:"map_dynamodb_endpoint"`

	// Global
	LoginPage string `yaml:"login_page"`
}

// Core make accessible to private config variables
type Core struct {
	config config

	logger   *zap.Logger
	dynamodb *dynamodb.Client
}

const jwtLocalSecret = "81027ac7103d791abacd19ac9f1e8722c19ad6c9"

// GetCore load config file for game
func GetCore(configPath string) (cfg Core, err error) {
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

	// set up zap logger
	logFile, err := os.OpenFile(cfg.GetLogPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error in opening log file: %v\n", err)
	}
	// defer logFile.Close()

	cfg.logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(logFile),
		cfg.GetLogLevel(),
	))

	// Create a DynamoDB client
	dynamodbCfg, err := dconfig.LoadDefaultConfig(context.TODO(),
		dconfig.WithRegion(cfg.config.MapDynamoDBRegion),
		dconfig.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: cfg.config.MapDynamoDBEndpoint}, nil
			})),
		dconfig.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		log.Fatal("error in connecting to dynamodb;", err)
	}
	cfg.dynamodb = dynamodb.NewFromConfig(dynamodbCfg)

	return
}

func validateConfig(cfg config) error {
	if cfg.MapDynamoDBRegion == "" {
		return errors.New("environment is required in config file")
	}

	switch {
	case cfg.Environment == "":
		return errors.New("environment is required in config file")
	case cfg.AppName == "":
		return errors.New("app_name is required in config file")
	case cfg.Port == "":
		return errors.New("port is required in config file")
	case cfg.Salt == "":
		return errors.New("salt is required in config file")
	case cfg.JwtSecret == "":
		return errors.New("jwt_secret is required in config file")
	case cfg.DefaultLang == "":
		return errors.New("default_lang is required in config file")
	case cfg.LogPath == "":
		return errors.New("log_path is required in config file")
	case cfg.LogLevel == "":
		return errors.New("log_level is required in config file")
	case cfg.MapDynamoDBRegion == "":
		return errors.New("map_dynamodb_region is required in config file")
	case cfg.MapDynamoDBEndpoint == "":
		return errors.New("map_dynamodb_endpoint is required in config file")
	case cfg.LoginPage == "":
		return errors.New("login_page is required in config file")
	}

	return nil
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

func (c Core) GetLoginPage() string {
	return c.config.LoginPage
}

func (c Core) ShowOriginalError() bool {
	return c.config.OriginalError
}

func (c Core) Debug(msg string, fields ...zap.Field) {
	c.logger.Debug(msg, fields...)
}

func (c Core) Info(msg string, fields ...zap.Field) {
	c.logger.Info(msg, fields...)
}

func (c Core) Warn(msg string, fields ...zap.Field) {
	c.logger.Warn(msg, fields...)
}

func (c Core) Error(msg string, fields ...zap.Field) {
	c.logger.Error(msg, fields...)
}

func (c Core) DPanic(msg string, fields ...zap.Field) {
	c.logger.DPanic(msg, fields...)
}

func (c Core) Panic(msg string, fields ...zap.Field) {
	c.logger.Panic(msg, fields...)
}

func (c Core) Fatal(msg string, fields ...zap.Field) {
	c.logger.Fatal(msg, fields...)
}

func (c Core) DynamoDB() *dynamodb.Client {
	return c.dynamodb
}
