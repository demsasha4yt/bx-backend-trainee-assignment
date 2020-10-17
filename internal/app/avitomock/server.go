package avitomock

import (
	"fmt"
	"net/http"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/crontab"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/store"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)

// Server structure
type server struct {
	logger *logrus.Logger
	router *mux.Router
	cron   crontab.Cron
	store  store.Store
	conf   *Config
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(store store.Store, conf *Config) *server {
	s := &server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		cron:   crontab.New(),
		store:  store,
		conf:   conf,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Use(s.accessLogMiddleware)
	s.router.Use(s.panicMiddleware)
	s.router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, map[string]bool{"ok": true})
	})
	s.router.HandleFunc("/api/1/rmp/show/{id:[0-9]+}", s.handleAdInfo())
	s.router.HandleFunc("/api/1/rmp/update/{id:[0-9]+}", s.handleUpdatePrice())
	s.router.HandleFunc("/api/1/rmp/delete/{id:[0-9]+}", s.handleSetDeleted())
}

func (s *server) handleAdInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "SHOW: %v\n", vars["id"])
	}
}

func (s *server) handleUpdatePrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "UPPDATE ID: %v\n", vars["id"])
	}
}

func (s *server) handleSetDeleted() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "DELETE ID: %v\n", vars["id"])
	}
}
