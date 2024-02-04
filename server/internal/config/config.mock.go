package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetCoreForTest() (core Core, err error) {
	core.config = config{
		Environment:            "testing",
		AppName:                "dimoklan",
		Port:                   ":3091",
		Salt:                   "salt for test",
		JwtSecret:              "secret for jwt",
		BasicMasterDatabaseDSN: "root:root@tcp(127.001:3306)/dimotest_basic",
		BasicSlaveDatabaseDSN:  "root:root@tcp(127.001:3306)/dimotest_basic",
		ActivityDatabaseDSN:    "root:root@tcp(127.001:3306)/dimotest_activity",
		LogPath:                "/home/diako/projects/dimoklan/server/logs/dimoklan_test.log",
		LogLevel:               "debug",
		DefaultLang:            "en",
		OriginalError:          true,
	}

	core.basicMasterDB, err = sql.Open("mysql", core.GetDatabaseMasterDNS())
	if err != nil {
		err = fmt.Errorf("error in opening basic master database; %w", err)
		return
	}

	core.basicSlaveDB, err = sql.Open("mysql", core.GetDatabaseSlaveDNS())
	if err != nil {
		err = fmt.Errorf("error in opening basic slave database; %w", err)
		return
	}

	core.activityDB, err = sql.Open("mysql", core.GetDatabaseActivityDNS())
	if err != nil {
		err = fmt.Errorf("error in opening activity database; %w", err)
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
