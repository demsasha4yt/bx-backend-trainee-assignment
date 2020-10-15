package avitoapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	key = "?key=af0deccbgcgidddjgnvljitntccdduijhdinfgjgfjir"
	url = "https://m.avito.ru/api/1/rmp/show/"
)

// Response ...
type Response struct {
	price int
}

type avitoResponse struct {
	DFPTargetings struct {
		Price int `json:"par_price"`
	} `json:"dfpTargetings"`
}

func generateQuery(avitoID int64) string {
	return url + strconv.FormatInt(avitoID, 10) + key
}

func getAvitoResponse(b []byte) (*avitoResponse, error) {
	r := &avitoResponse{}
	fmt.Println(string(b))
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetInfo returns info about ad
func GetInfo(avitoID int64) (*Response, error) {
	client := &http.Client{
		Timeout: 20 * time.Second,
		// Transport: transport,
	}

	request, err := http.NewRequest("GET", "https://avito.ru", nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	getAvitoResponse(body)
	return nil, nil
}
