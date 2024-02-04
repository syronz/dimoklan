package mapgenerator

import (
	"log"
	"os"
	"testing"

	"dimoklan/internal/config"
)

var core config.Core

func TestMain(m *testing.M) {
	var err error
	if core, err = config.SetCoreForTest(); err != nil {
		log.Fatalf("error in setting config in test environment; %v\n", err)
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}

func insert100kUser() {

}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
