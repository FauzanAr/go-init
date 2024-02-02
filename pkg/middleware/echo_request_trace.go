package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/FauzanAr/go-init/pkg/helper"
	"github.com/FauzanAr/go-init/pkg/logger"
	"github.com/FauzanAr/go-init/pkg/wrapper"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func EchoRequestTrace(log logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			traceID := uuid.New().String()
			ctx := context.WithValue(c.Request().Context(), "trace_id", traceID)
			ctx = context.WithValue(ctx, "url", c.Request().RequestURI)
			ctx = context.WithValue(ctx, "method", c.Request().Method)
			ctx = context.WithValue(ctx, "remote_ip", c.RealIP())
			c.SetRequest(c.Request().WithContext(ctx))

			log.Info(ctx, "Request started", nil)

			err := next(c)

			log.Info(ctx, "Request completed", logger.MetaData{
				"status":       c.Response().Status,
				"elapsed_time": time.Since(start).Seconds(),
			})

			return err
		}
	}
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("auth header missing"), nil, http.StatusUnauthorized)
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := helper.VerifyToken(c.Request().Context(), tokenString)
		if err != nil {
			return wrapper.SendErrorResponse(c, wrapper.UnauthorizedError("invalid token"), nil, http.StatusUnauthorized)
		}

		ctx := context.WithValue(c.Request().Context(), "user", claims)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
