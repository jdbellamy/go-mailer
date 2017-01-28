package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/uber-go/zap"

	"time"
)



func Index(w http.ResponseWriter, r *http.Request) {
	logger := zap.New(
		zap.NewJSONEncoder(zap.NoTime()), // drop timestamps in tests
	)

	logger.Warn("Log without structured data...")
	logger.Warn(
		"Or use strongly-typed wrappers to add structured context.",
		zap.String("library", "zap"),
		zap.Duration("latency", time.Nanosecond),
	)

	logger.Info("Failed to fetch URL.",
		zap.String("url", r.URL.String()),
	)
	fmt.Fprintln(w, "/")
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	type MessageRequest struct {
		Subject string
		Body	string
	}
	var messageRequest MessageRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&messageRequest)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	msg := NewMessage()
	msg.From("admin@example.com")
	msg.To("bellamy.john.d@gmail.com", "john.devon.bellamy@gmail.com")
	msg.Subject(messageRequest.Subject)
	msg.Body(messageRequest.Body)
	msg.SendSmtp(&SMTPConfig{
		Server: "smtp",
		Port: 25,
		Username: "user",
		Password: "123456",
	})
}