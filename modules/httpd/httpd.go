package httpd

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type HttpD struct {
	server *http.Server
	router *chi.Mux
	hcheck healthcheck
}

func newHttpd(addr string) *HttpD {
	var (
		router = chi.NewRouter()
		server = &http.Server{Addr: addr, Handler: router}
		hcheck = newHealthcheck()
		daemon = HttpD{server: server, router: router, hcheck: hcheck}
	)
	router.Use(middleware.Logger)
	router.Get("/health", daemon.healthcheckHandler)
	return &daemon
}

func (d *HttpD) AddHealthCheck(name string, callback HealthCheck) {
	d.hcheck.AddCheck(name, callback)
}

func (d *HttpD) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	outcome, errors := d.hcheck.GetResults()
	if errors > 0 {
		w.WriteHeader(400)
	}
	w.Write([]byte(outcome))
}
