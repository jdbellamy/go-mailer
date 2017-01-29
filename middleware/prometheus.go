package middleware



import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Prometheus(inner http.Handler, name string) http.Handler {
	return promhttp.Handler()
}