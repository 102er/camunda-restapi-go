package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blossom102er/camunda-restapi-go/entitys"
	"github.com/go-resty/resty/v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type BaseRestApiService struct {
	httpClient  *http.Client
	endpointUrl string
	resty       *resty.Client
}

func NewBaseRestApiService(endpointUrl string, timeoutSec time.Duration) *BaseRestApiService {
	baseCamundaRestApiService := &BaseRestApiService{
		httpClient: &http.Client{
			Timeout: time.Second * timeoutSec,
		},
		endpointUrl: endpointUrl,
		resty:       resty.New(), //暂未改造成resty方式
	}
	return baseCamundaRestApiService
}

func (c *BaseRestApiService) doPostJson(path string, query map[string]string, v interface{}) (res *http.Response, err error) {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(v); err != nil {
		return nil, err
	}
	res, err = c.do(http.MethodPost, path, query, body, "application/json")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *BaseRestApiService) doPutJson(path string, query map[string]string, v interface{}) error {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(v); err != nil {
		return err
	}

	_, err := c.do(http.MethodPut, path, query, body, "application/json")
	return err
}

func (c *BaseRestApiService) doDelete(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodDelete, path, query, nil, "")
}

func (c *BaseRestApiService) doPost(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodPost, path, query, nil, "")
}

func (c *BaseRestApiService) doPut(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodPut, path, query, nil, "")
}

func (c *BaseRestApiService) do(method, path string, query map[string]string, body io.Reader, contentType string) (res *http.Response, err error) {
	useUrl, err := c.buildUrl(path, query)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, useUrl, body)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.SetBasicAuth("demo", "123")
	res, err = c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err := c.checkResponse(res); err != nil {
		return nil, err
	}

	return
}

func (c *BaseRestApiService) doGet(path string, query map[string]string) (res *http.Response, err error) {
	return c.do(http.MethodGet, path, query, nil, "")
}

func (c *BaseRestApiService) checkResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	defer res.Body.Close()

	if res.Header.Get("Content-Type") == "application/json" {
		if res.StatusCode == 404 {
			return fmt.Errorf("not found")
		}

		jsonErr := &Error{}
		err := json.NewDecoder(res.Body).Decode(jsonErr)
		if err != nil {
			return fmt.Errorf("response error with status code %d: failed unmarshal error response: %w", res.StatusCode, err)
		}

		return jsonErr
	}

	errText, err := ioutil.ReadAll(res.Body)
	if err == nil {
		return fmt.Errorf("response error with status code %d: %s", res.StatusCode, string(errText))
	}

	return fmt.Errorf("response error with status code %d", res.StatusCode)
}

func (c *BaseRestApiService) readJsonResponse(res *http.Response, v interface{}) error {
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (c *BaseRestApiService) buildUrl(path string, query map[string]string) (string, error) {
	if len(query) == 0 {
		return c.endpointUrl + path, nil
	}
	useUrl, err := url.Parse(c.endpointUrl + path)
	if err != nil {
		return "", err
	}

	q := useUrl.Query()
	for k, v := range query {
		q.Set(k, v)
	}

	useUrl.RawQuery = q.Encode()
	return useUrl.String(), nil
}

// Error a custom error type
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Error error message
func (e *Error) Error() string {
	return e.Message
}

// toCamundaTime return time formatted for camunda
func toCamundaTime(dt time.Time) string {
	if dt.IsZero() {
		return ""
	}

	return dt.Format(entitys.DefaultDateTimeFormat)
}
