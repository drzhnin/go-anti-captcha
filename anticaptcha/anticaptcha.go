package anticaptcha

import (
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "http://anti-captcha.com"
)

//SystemStat contains statistics from anti-captcha system
type SystemStat struct {
	Waiting                  int     `xml:"waiting"`
	WaitingRU                int     `xml:"waitingRU"`
	Load                     float64 `xml:"load"`
	Minbid                   float64 `xml:"minbid"`
	MinbidRU                 float64 `xml:"minbidRU"`
	AverageRecognitionTime   float64 `xml:"averageRecognitionTime"`
	AverageRecognitionTimeRU float64 `xml:"averageRecognitionTimeRU"`
}

// A Client manages communication with the API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	//Api key
	APIKey string

	// Services used for talking to different parts of the API.
	Account                 *AccountService
	Captcha                 *CaptchaService
	CaptchaAdditionalParams *AdditionalParams
}

type service struct {
	client *Client
}

// NewClient returns a new API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(apiKey string) *Client {
	httpClient := http.DefaultClient

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, APIKey: apiKey}
	c.Account = &AccountService{client: c}
	c.Captcha = &CaptchaService{client: c}
	c.CaptchaAdditionalParams = &AdditionalParams{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if string(data) == "ERROR_KEY_DOES_NOT_EXIST" || string(data) == "ERROR_WRONG_USER_KEY" {
		return nil, errors.New("Api key does not exist, plaese set correct api key from http://anti-captcha.com")
	}

	return data, err
}

//GetSystemStat represents load real time load statistics
func (c *Client) GetSystemStat() (*SystemStat, error) {
	var stats *SystemStat
	req, err := c.NewRequest("GET", "load.php", nil)
	if err != nil {
		return stats, err
	}
	data, err := c.Do(req)
	if err != nil {
		return stats, err
	}
	err = xml.Unmarshal(data, &stats)
	if err != nil {
		return stats, err
	}
	return stats, err
}
