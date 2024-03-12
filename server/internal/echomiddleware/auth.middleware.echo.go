package echomiddleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"dimoklan/consts"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// AuthMiddleware is the Echo middleware for JWT authentication
func (m *Middleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the Authorization header
		authHeader := c.Request().Header.Get("Authorization")

		// Check if the header is present
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		// Split the header value into parts
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing Bearer token")
		}

		// Extract the token from the second part
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(consts.HashSalt + m.core.GetSalt()), nil
		})
		if err != nil {
			m.core.Warn("signature_is_not_valid", zap.String("error", err.Error()), zap.String("ip", c.Get("ip").(string)))
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token 1")
		}

		var claims jwt.MapClaims
		ok := false
		if claims, ok = token.Claims.(jwt.MapClaims); !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token 2")
		}

		expInterface, ok2 := claims["exp"]
		if !ok2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token 3")
		}

		expString, ok3 := expInterface.(float64)
		if !ok3 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token 4")
		}

		exp := int64(expString)

		if exp < time.Now().Unix() {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token Expired")
		}

		userID := claims["user_id"]
		if userID == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token 5")
		}

		c.Set("user_id", claims["user_id"])

		return next(c)

	}
}
