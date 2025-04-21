package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	tokenConfig "github.com/leandro-andrade-candido/api-go/src/config/token"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/context"
	"github.com/leandro-andrade-candido/api-go/src/libs/application/errs"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// RequireJWTAuth is a middleware function that enforces JWT authentication for protected routes.
// It performs the following:
// 1. Extracts the JWT token from the Authorization header
// 2. Validates the token and extracts user claims
// 3. Creates an application context with user information
// 4. Sets OpenTelemetry span attributes for authentication status and user ID
// Returns an echo.MiddlewareFunc that:
// - Returns 401 Unauthorized if token is missing or invalid
// - Calls next handler with ApplicationContext if authentication succeeds
func RequireJWTAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// trace span
			span := trace.SpanFromContext(c.Request().Context())

			bearer := c.Request().Header.Get("Authorization")
			if bearer == "" {
				span.SetAttributes(attribute.Bool("x.user.auth", false))
				return errs.NewApiError("user token is missing", nil, http.StatusUnauthorized)
			}

			token := getJwtFromBearer(bearer)
			claims, err := ParseJwtToken(token, tokenConfig.GetTokenConfiguration)
			if err != nil {
				span.SetAttributes(attribute.Bool("x.user.auth", false))
				return errs.NewApiError("user is not authorized", err, http.StatusUnauthorized)
			}

			applicationContext := &context.ApplicationContext{
				Context: c,
				User:    *claims,
			}

			// sets user id as span attribute
			span.SetAttributes(
				attribute.Bool("x.user.auth", true),
				attribute.String("x.user.id", claims.Id.String()),
			)
			return next(applicationContext)
		}
	}
}

func getJwtFromBearer(bearer string) string {
	value, _ := strings.CutPrefix(bearer, "Bearer ")
	return value
}

type ConfigProvider func() (*tokenConfig.TokenConfig, error)

// parses JWT token and get user claims.
// Returns error if token is not valid or if parsing fails
func ParseJwtToken(token string, getTokenConfig ConfigProvider) (*context.UserClaims, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		tConfig, err := getTokenConfig()
		if err != nil {
			return nil, err
		}

		return []byte(tConfig.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	userId, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return nil, err
	}

	return &context.UserClaims{
		Id:    userId,
		Email: claims["email"].(string),
	}, nil
}
