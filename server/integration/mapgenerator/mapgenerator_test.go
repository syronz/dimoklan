package mapgenerator

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"dimoklan/domain/basic/basstorage"
	"dimoklan/domain/map/mapstorage"
	"dimoklan/internal/config"
	"dimoklan/internal/migration"
	"dimoklan/service"
	"dimoklan/types"

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
	insert500kCells()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func insert10kUser() {
	basStorage := basstorage.New(core)
	userService := service.NewUserService(core, basStorage)
	const numUsers = 10 //_000

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

func getAllUsers() []int {
	query := "SELECT id FROM users"
	rows, err := core.BasicSlaveDB().Query(query)
	if err != nil {
		log.Fatalf("error in getting all users; %v\n", err)
	}
	defer rows.Close()

	userIDs := []int{}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatalf("error in scanning user_id; %v\n", err)
		}

		userIDs = append(userIDs, id)
	}

	return userIDs
}

func insert500kCells() {
	users := getAllUsers()

	mapStorage := mapstorage.New(core)
	cellService := service.NewCellService(core, mapStorage)

	const maxX = 1225
	const maxY = 1225

	for _, userID := range users {
		x := rand.Intn(7 + maxX)
		y := rand.Intn(7 + maxY)
		for i := -3; i <= 3; i++ {
			for j := -3; j <= 3; j++ {
				cell := types.Cell{
					X:      x + i,
					Y:      y + j,
					UserID: userID,
				}
				_, err := cellService.Create(cell)
				if err != nil {
					log.Fatalf("error in creating cell; %v\n", err)
				}

			}
		}

	}
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
