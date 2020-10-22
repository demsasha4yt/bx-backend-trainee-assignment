package avitoapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const (
	infoPath       string = "/api/1/rmp/show/"
	statusNotFound        = "not-found"
	statusOK              = "ok"
)

var (
	//ErrNotFound returns while ad not found
	ErrNotFound = errors.New("Not found")
	//ErrNotOK return while respone from AVITO is not OK
	ErrNotOK = errors.New("Can't get ad info from avito")
)

// AvitoAPI interface
type AvitoAPI interface {
	GetInfo(avitoID int64) (*Response, error)
}

type avitoAPI struct {
	conf *Config
}

type infoResponse struct {
	Status string `json:"status"`
	Result struct {
		DFPTargetings struct {
			Price int `json:"price"`
		} `json:"dfpTargetings,omitempty"`
	} `json:"result,omitempty"`
}

// Response ...
type Response struct {
	Price int
}

// New returns new AvitoAPI
func New(conf *Config) AvitoAPI {
	return &avitoAPI{
		conf: conf,
	}
}

func (s *avitoAPI) generateInfoURL(avitoID int64) string {
	return s.conf.URL + infoPath + strconv.FormatInt(avitoID, 10) + "?key=" + s.conf.Key
}

// GetInfo returns info about ad
func (s *avitoAPI) GetInfo(avitoID int64) (*Response, error) {
	info := &infoResponse{}

	resp, err := http.Get(s.generateInfoURL(avitoID))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, ErrNotOK
	}

	if strings.Compare(info.Status, statusNotFound) == 0 {
		return nil, ErrNotFound
	}

	if strings.Compare(info.Status, statusOK) != 0 {
		return nil, ErrNotOK
	}

	return &Response{
		Price: info.Result.DFPTargetings.Price,
	}, nil
}
