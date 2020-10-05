package bx

import (
	"encoding/json"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (s *server) handleSubscription(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	type request struct {
		email string
		link  string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := validation.ValidateStruct(
			req,
			validation.Field(&req.email, validation.Required, is.Email),
			validation.Field(&req.link, validation.Required, is.URL, validation.Match(regexp.MustCompile("^((http?|https?)://)?(www.)?(m.)?avito.ru([/w .-]*)*/?_[0-9]{1,}$"))),
		); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := s.service.Subscription().Subscribe(r.Context(), req.email, req.link); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, map[string]string{"message": "ok"})
	}
}
