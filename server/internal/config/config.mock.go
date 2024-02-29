package config

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetCoreForTest() (core Core, err error) {
	core.config = config{
		Environment:   "testing",
		AppName:       "dimoklan",
		Port:          ":3091",
		Salt:          "salt for test",
		JwtSecret:     "secret for jwt",
		LogPath:       "/home/diako/projects/dimoklan/server/logs/dimoklan_test.log",
		LogLevel:      "debug",
		DefaultLang:   "en",
		OriginalError: true,
	}

	if err != nil {
		err = fmt.Errorf("error in opening basic master database; %w", err)
		return
	}

	// set up zap logger
	logFile, err := os.OpenFile(core.GetLogPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error in opening log file: %v\n", err)
	}

	core.logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(logFile),
		core.GetLogLevel(),
	))

	return
}
