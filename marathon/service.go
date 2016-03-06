package marathon

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Service struct {
	BaseURL string
}

type jsonResponse struct {
	App *App
}

func NewService(host string, port int) (*Service, error) {
	var url = fmt.Sprintf("http://%v:%v", host, port)
	var ms = &Service{BaseURL: url}

	return ms, nil
}

func (service *Service) HttpGet(path string) ([]byte, error) {
	response, err := http.Get(service.BaseURL + path)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	output, err := ioutil.ReadAll(response.Body)
	return output, err
}

func (service *Service) HttpPost(path string, body io.Reader) ([]byte, error) {
	response, err := http.Post(service.BaseURL+path, "application/json", body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	output, err := ioutil.ReadAll(response.Body)
	return output, err
}

func (service *Service) GetApp(path string) (*App, error) {
	jsonBlob, err := service.HttpGet("/v2/apps" + path)
	if err != nil {
		return nil, err
	}

	var v jsonResponse
	err = json.Unmarshal(jsonBlob, &v)
	return v.App, err
}
