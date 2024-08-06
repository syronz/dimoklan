package integration

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"

	"dimoklan/api"
	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/repo"
	"dimoklan/service"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterAPI Integration Tests", func() {
	var (
		e           *echo.Echo
		core        config.Core
		registerAPI *api.RegisterAPI
	)

	BeforeEach(func() {
		testConfig := "./test_config.yaml"
		testConfigPath, err := filepath.Abs(testConfig)
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}

		// Setup configuration and services before each test.
		core, err = config.GetCore(testConfigPath)
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}

		storage := repo.NewRepo(core)

		cellService := service.NewCellService(core, storage)

		registerService := service.NewRegisterService(core, storage, cellService)
		registerAPI = api.NewRegisterAPI(core, registerService)

		// Create a new Echo instance.
		e = echo.New()

		e.POST("/register", registerAPI.Create)
		e.GET("/register", registerAPI.Confirm)
	})

	It("successfull registration", func() {
		randomEmail := fmt.Sprintf("a:User%v@gmail.com", rand.Intn(10000000))
		registerPayload := `{
			"email": "` + randomEmail + `",
			"password": "StrongPassword2000",
			"kingdom":"Eldoria",
			"cell": {
				"x": 7,
				"y": 8
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(registerPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		err := registerAPI.Create(c)
		Expect(err).NotTo(HaveOccurred())
		Expect(rec.Code).To(Equal(http.StatusCreated))

		responseRegister := model.Register{}
		err = json.Unmarshal(rec.Body.Bytes(), &responseRegister)
		Expect(err).NotTo(HaveOccurred())
		Expect(responseRegister.Email).To(Equal(randomEmail))

		// activation code: a1d7f2c0da2ab6e559ccac5390fe4905c967270e143fd8afd0a7f3331bfe41f1
	})
})
