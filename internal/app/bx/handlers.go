package bx

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (s *server) handleSubscription() http.HandlerFunc {
	type request struct {
		Email string `json:"email"`
		Link  string `json:"link"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := validation.ValidateStruct(
			req,
			validation.Field(&req.Email, validation.Required, is.Email),
			validation.Field(
				&req.Link,
				validation.Required,
				is.URL,
				validation.Match(
					regexp.MustCompile(`^((http?|https?)://)?(www.)?(m.)?(avito.ru/)..*_[0-9]{1,}$`),
				),
			),
		); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		result, err := s.service.Subscription().Subscribe(r.Context(), req.Email, req.Link)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, result)
	}
}

func parseInt(value string) (int64, error) {
	if result, err := strconv.ParseInt(value, 10, 64); err == nil {
		return result, nil
	}
	return 0, errors.New("Not an integer")
}

func (s *server) handleConfirmSubscribe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adID, emailID, token, err := getConfirmUnsubscribeQueryParams(r)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		result, err := s.service.Subscription().ConfirmSubscribe(r.Context(), emailID, adID, token)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, result)
	}
}

func (s *server) handleUnubscribe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		adID, emailID, token, err := getConfirmUnsubscribeQueryParams(r)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		result, err := s.service.Subscription().Unsubscribe(r.Context(), emailID, adID, token)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, result)
	}
}

func getConfirmUnsubscribeQueryParams(req *http.Request) (int64, int64, string, error) {
	adID, err := parseInt(req.URL.Query().Get("adID"))
	if err != nil {
		return 0, 0, "", errors.New(`adID is not and integer`)
	}

	emailID, err := parseInt(req.URL.Query().Get("emailID"))
	if err != nil {
		return 0, 0, "", errors.New(`emailID is not and integer`)
	}
	token := req.URL.Query().Get("token")
	if token == "" {
		return 0, 0, "", errors.New("Unknown token")
	}
	return adID, emailID, token, nil
}
