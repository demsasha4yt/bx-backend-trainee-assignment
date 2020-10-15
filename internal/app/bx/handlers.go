package bx

import (
	"encoding/json"
	"net/http"
	"regexp"

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
