package api

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HueApi struct {
	IP    string
	Token string
}

func (api *HueApi) HttpPut(url string, body map[string]interface{}) (bool, error) {
	jsonBytes, _ := json.Marshal(body)
	req, reqErr := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBytes))
	if reqErr != nil {
		log.Error(reqErr.Error())
		return false, reqErr
	}
	client := &http.Client{}
	response, resErr := client.Do(req)
	if resErr != nil {
		log.Error(resErr.Error())
		return false, resErr
	}
	return response.Status == "200", nil
}

func New(ip string, token string) (HueApi, error) {
	if ip == "" {
		return HueApi{}, errors.New("no IP provided")
	}
	if token == "" {
		log.Error("no token provided")
		return HueApi{}, errors.New("no token provided")
	}
	return HueApi{IP: ip, Token: token}, nil
}
