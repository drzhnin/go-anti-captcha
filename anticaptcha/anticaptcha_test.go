package anticaptcha

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	//client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	//client configured to use test server
	client = NewClient("F629EBDA-D89A-4A0E-AAA5-069761578237")
	client.CaptchaAdditionalParams.ResetToDefault()
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	c := NewClient("F629EBDA-D89A-4A0E-AAA5-069761578237")
	apiKey := "F629EBDA-D89A-4A0E-AAA5-069761578237"

	assert.Equal(t, c.BaseURL.String(), defaultBaseURL)
	assert.Equal(t, c.APIKey, apiKey)
}

func TestNewRequest(t *testing.T) {
	c := NewClient("F629EBDA-D89A-4A0E-AAA5-069761578237")

	inURL, outURL := "/foo", defaultBaseURL+"/foo"
	inBody, outBody := strings.NewReader(`{"Balance":1}`+"\n"), `{"Balance":1}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	assert.Equal(t, req.URL.String(), outURL)

	body, _ := ioutil.ReadAll(req.Body)
	assert.Equal(t, string(body), outBody)

}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient("")
	_, err := c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, 1.1)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := 1.0
	client.Do(req)

	want := 1.0
	assert.Equal(t, body, want)

	httpmock.RegisterResponder("GET", "http://anti-captcha.com/res.php?&action=getbalance",
		httpmock.NewStringResponder(200, `ERROR_KEY_DOES_NOT_EXIST`))
	reqTestKeyNotExst, _ := client.NewRequest("GET", "http://anti-captcha.com/res.php?key=1&action=getbalance", nil)
	_, err := client.Do(reqTestKeyNotExst)
	assert.Equal(t, err.Error(), "Api key does not exist, plaese set correct api key from http://anti-captcha.com")

	httpmock.RegisterResponder("GET", "http://anti-captcha.com/res.php?&action=getbalance",
		httpmock.NewStringResponder(200, `ERROR_WRONG_USER_KEY`))
	reqTestWrongKey, _ := client.NewRequest("GET", "http://anti-captcha.com/res.php?key=1&action=getbalance", nil)
	_, err = client.Do(reqTestWrongKey)
	assert.Equal(t, err.Error(), "Api key does not exist, plaese set correct api key from http://anti-captcha.com")

}

func TestAccountService_GetSystemStat(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "http://anti-captcha.com/load.php", httpmock.NewStringResponder(200, `<RESPONSE>
	<waiting>50</waiting>
	<waitingRU>0</waitingRU>
	<load>75.4</load>
	<minbid>0.0009072645</minbid>
	<minbidRU>0.0010887339</minbidRU>
	<averageRecognitionTime>14.853999734384</averageRecognitionTime>
	<averageRecognitionTimeRU>13.725848563969</averageRecognitionTimeRU>
	</RESPONSE>`))
	sysStat, err := client.GetSystemStat()
	if err != nil {
		t.Errorf("Account.GetSystemStat returned error: %v", err)
	}
	assert.Equal(t, sysStat.Waiting, 50)
	assert.Equal(t, sysStat.WaitingRU, 0)
	assert.Equal(t, sysStat.Load, 75.4)
	assert.Equal(t, sysStat.Minbid, 0.0009072645)
	assert.Equal(t, sysStat.MinbidRU, 0.0010887339)
	assert.Equal(t, sysStat.AverageRecognitionTime, 14.853999734384)
	assert.Equal(t, sysStat.AverageRecognitionTimeRU, 13.725848563969)
}

func testMethod(t *testing.T, r *http.Request, want string) {
	assert.Equal(t, r.Method, want)
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}
