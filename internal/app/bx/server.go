package bx

import (
	"net/http"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/crontab"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/service"
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
	logger  *logrus.Logger
	router  *mux.Router
	cron    crontab.Cron
	store   store.Store
	service service.Interface
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(store store.Store, service service.Interface) *server {
	s := &server{
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		cron:    crontab.New(),
		store:   store,
		service: service,
	}
	s.configureCron()
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
	s.router.HandleFunc("/api/subscribe", s.handleSubscription())
}

func (s *server) configureCron() {
	s.cron.AddFunc("@every 10s", func() {
		s.logger.Info("CRONTAB every 1s")
	})
	s.logger.Info("Cron configured")
	s.cron.Start()
}
