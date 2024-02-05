package mapgenerator

import (
	"dimoklan/domain/basic/basstorage"
	"dimoklan/internal/config"
	"dimoklan/internal/migration"
	"dimoklan/service"
	"dimoklan/types"
	"log"
	"os"
	"testing"

	"github.com/go-faker/faker/v4"
)

var core config.Core

func TestMain(m *testing.M) {
	var err error
	if core, err = config.SetCoreForTest(); err != nil {
		log.Fatalf("error in setting config in test environment; %v\n", err)
	}

	migration.MigrateDB(core.GetDatabaseMasterDNS(), "up", 2)

	insert10kUser()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func insert10kUser() {
	basStorage := basstorage.New(core)
	userService := service.NewUserService(core, basStorage)
	const numUsers = 10//_000

	for i := 0; i < numUsers; i++ {
		user := types.User{}
		err := faker.FakeData(&user)
		if err != nil {
			log.Fatalf("error in faking user; %v\n", err)
		}

		// Set additional fields or modify generated data if needed
		user.Status = "active"
		user.Language = "en"

		_, err = userService.Create(user)
		if err != nil {
			log.Fatalf("error in creating user; %v\n", err)
		}
	}
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
