package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"dimoklan/api"
	"dimoklan/internal/config"
	"dimoklan/model"
	"dimoklan/repo"
	"dimoklan/service"
)

func TestAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Suite")
}

var _ = Describe("AuthAPI Integration Tests", func() {
	var (
		e       *echo.Echo
		core    config.Core
		authAPI *api.AuthAPI
	)

	BeforeEach(func() {
		testConfig := "./test_config.yaml"
		testConfigPath, err := filepath.Abs(testConfig)
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}

		// Setup your configuration and services before each test.
		core, err = config.GetCore(testConfigPath)
		if err != nil {
			Expect(err).NotTo(HaveOccurred())
		}

		storage := repo.NewRepo(core)
		authService := service.NewAuthService(core, storage)
		authAPI = api.NewAuthAPI(core, authService)

		// Create a new Echo instance.
		e = echo.New()

		e.POST("/login", authAPI.Login)
	})

	It("should successfully handle a login request", func() {
		// Create a test user for the login request.
		authPayload := `{"email": "sabina.diako@gmail.com", "password": "StrongPassword2000"}`

		// Create a request with the test user JSON.
		// req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonUser))
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(authPayload))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		// Serve the request using the Echo instance.
		c := e.NewContext(req, rec)
		// err = e.ServeHTTP(rec, req)
		err := authAPI.Login(c)

		// Assert that there is no error.
		Expect(err).NotTo(HaveOccurred())

		// Assert the HTTP status code is OK.
		Expect(rec.Code).To(Equal(http.StatusOK))

		// Parse the response JSON into a model.Auth object.
		var responseAuth model.Auth
		err = json.Unmarshal(rec.Body.Bytes(), &responseAuth)
		Expect(err).NotTo(HaveOccurred())
		Expect(responseAuth.Email).To(Equal("sabina.diako@gmail.com"))
	})

	Context("when payload is not valid", func() {
		It("should successfully handle a login request", func() {
			authPayload := `{"email": "invalid-email", "password": "StrongPassword2000"}`

			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(authPayload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			// Serve the request using the Echo instance.
			c := e.NewContext(req, rec)
			err := authAPI.Login(c)
			Expect(err).NotTo(HaveOccurred())

			// Assert the HTTP status code is OK.
			Expect(rec.Code).To(Equal(http.StatusOK))

			// Parse the response JSON into a model.Auth object.
			var responseAuth model.Auth
			err = json.Unmarshal(rec.Body.Bytes(), &responseAuth)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseAuth.Email).To(Equal("sabina.diako@gmail.com"))
		})
	})
})
