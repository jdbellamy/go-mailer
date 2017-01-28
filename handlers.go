package main

import (
	"fmt"
	"net/http"
	"github.com/uber-go/zap"
	"github.com/google/jsonapi"
)

var emails = []*Email{}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go-mailer Service")
}

func ListMessages(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("msgs.list")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapiRuntime.MarshalManyPayload(w, emails); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("msgs.create")
	var m = new(Email)
	if err := jsonapiRuntime.UnmarshalPayload(r.Body, m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error("Unexpected error while decoding request body",
			zap.Error(err),
			zap.Int("status", http.StatusInternalServerError))
		return
	}
	var mailer = SmtpClient{Server: "smtp", Port: 25}
	if err := mailer.Send(m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error("Unexpected error while sending email",
			zap.Error(err),
			zap.Int("status", http.StatusInternalServerError))
		return
	}
	emails = append(emails, m)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", jsonapi.MediaType)
	if err := jsonapi.MarshalOnePayload(w, m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
