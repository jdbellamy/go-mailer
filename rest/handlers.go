package rest

import (
	"fmt"
	"net/http"
	"github.com/uber-go/zap"
	"github.com/google/jsonapi"
	"github.com/jdbellamy/go-mailer/mail"
	. "github.com/jdbellamy/go-mailer/middleware"
)

var emails = []*mail.Email{}

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
	var m = new(mail.Email)
	if err := jsonapiRuntime.UnmarshalPayload(r.Body, m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Z.Error("Unexpected error while decoding request body",
			zap.Error(err),
			zap.Int("status", http.StatusInternalServerError))
		return
	}
	var mailer = mail.SmtpClient{
		Server: "localhost",
		Port: 26,
	}
	if err := mailer.Send(m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		Z.Error("Unexpected error while sending email",
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
