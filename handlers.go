package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/uber-go/zap"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go-mailer Service")
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	type MessageRequest struct {
		Subject string
		Body	string
	}
	var messageRequest MessageRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		logger.Error("Missing request body",  zap.Int("status", 400))
		return
	}
	err := json.NewDecoder(r.Body).Decode(&messageRequest)
	if err != nil {
		http.Error(w, err.Error(), 500)
		logger.Error("Unexpected error while decoding request body",
			zap.Error(err),
			zap.Int("status", 500))
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