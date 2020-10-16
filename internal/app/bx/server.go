package bx

import (
	"net/http"
	"time"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/avitoapi"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/crontab"
	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/emailsender"
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
	logger      *logrus.Logger
	router      *mux.Router
	cron        crontab.Cron
	store       store.Store
	service     service.Service
	emailSender emailsender.EmailSender
	avitoAPI    avitoapi.AvitoAPI
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(store store.Store, service service.Service,
	emailsender emailsender.EmailSender, avitoAPI avitoapi.AvitoAPI) *server {
	s := &server{
		logger:      logrus.New(),
		router:      mux.NewRouter(),
		cron:        crontab.New(),
		store:       store,
		service:     service,
		emailSender: emailsender,
		avitoAPI:    avitoAPI,
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
	s.router.HandleFunc("/confirm", s.handleConfirmSubscribe())
	s.router.HandleFunc("/unsubscribe", s.handleUnubscribe())
}

func (s *server) configureCron() {
	var TaskID int = 1
	s.cron.AddFunc("@every 10s", func() {
		start := time.Now()
		curID := TaskID
		TaskID++
		s.logger.Infof("PriceTracking start [ID:%d]", curID)
		if err := s.service.PriceTracker().CheckAdsTask(); err != nil {
			s.logger.Error(err)
			s.logger.Errorf("PriceTracking failed (in %s) [ID:%d]", time.Since(start), curID)
			return
		}
		s.logger.Infof("PriceTracking end (Completed in %s) [ID:%d]", time.Since(start), curID)
	})
	s.logger.Info("Cron configured")
	s.cron.Start()
}
