package middleware

import (
	"context"
	"time"

	"github.com/FauzanAr/go-init/pkg/logger"
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
