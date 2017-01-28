package main

import (
	"net/http"
	"time"
	"github.com/uber-go/zap"
)

var logger = zap.New(zap.NewJSONEncoder())

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		logger.Info("request received",
			zap.String("method", r.Method),
			zap.String("uri", r.RequestURI),
			zap.String("handler", name),
			zap.Int64("duration", time.Since(start).Nanoseconds()))
	})
}
