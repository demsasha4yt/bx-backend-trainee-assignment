package avitomock

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/demsasha4yt/bx-backend-trainee-assignment/internal/app/models"

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

func (s *server) configureCron() {
	var TaskID int = 1
	s.cron.AddFunc("@every 10m", func() {
		start := time.Now()
		curID := TaskID
		TaskID++
		s.logger.Infof("UpdateAds start [ID:%d]", curID)
		s.store.AvitoMock().SetDeleted(context.Background())
		s.store.AvitoMock().UpdatePrices(context.Background())
		s.logger.Infof("UpdateAds end (Completed in %s) [ID:%d]", time.Since(start), curID)
	})
	s.logger.Info("Cron configured")
	s.cron.Start()
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
		avitoID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		ad, err := s.store.AvitoMock().FindByAvitoID(r.Context(), avitoID)
		if err != nil {
			ad = models.NewAvitoMockFromID(avitoID)
			if err != store.ErrRecordNotFound {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
			if err := s.store.AvitoMock().Create(r.Context(), ad); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
		}
		if ad.Deleted {
			s.respond(w, r, http.StatusOK, map[string]string{"status": "not-found", "result": ""})
			return
		}
		s.respond(w, r, http.StatusOK, map[string]interface{}{
			"status": "ok",
			"result": map[string]interface{}{
				"banners": map[string]interface{}{
					"somedata":  "somedata",
					"somedata2": "somedata",
					"somedata3": "somedata",
					"somedata4": "somedata",
					"somedata5": "somedata",
				},
				"dfpTargetings": map[string]interface{}{
					"somedata":   "somedata",
					"somedata2":  "somedata",
					"somedata3":  "somedata",
					"somedata4":  "somedata",
					"somedata5":  "somedata",
					"par_price":  ad.Price,
					"somedata6":  "somedata",
					"somedata7":  "somedata",
					"somedata8":  "somedata",
					"somedata9":  "somedata",
					"somedata10": "somedata",
				},
				"enableEventSampling": false,
				"wbPixelEnabled":      true,
			},
		})
	}
}

func (s *server) handleUpdatePrice() http.HandlerFunc {
	type request struct {
		Price int `json:"price"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		vars := mux.Vars(r)

		avitoID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.AvitoMock{
			AvitoID: avitoID,
			Price:   req.Price,
		}

		if err := s.store.AvitoMock().UpdatePrice(r.Context(), avitoID, req.Price); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) handleSetDeleted() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		avitoID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := s.store.AvitoMock().SetDeletedOne(r.Context(), avitoID); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, "ok")
	}
}
