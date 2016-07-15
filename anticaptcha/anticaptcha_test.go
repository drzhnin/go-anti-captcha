package anticaptcha

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// github client configured to use test server
	client = NewClient("123123")
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	c := NewClient("123123")
	apiKey := "123123"

	assert.Equal(t, c.BaseURL.String(), defaultBaseURL)
	assert.Equal(t, c.ApiKey, apiKey)
}

func TestNewRequest(t *testing.T) {
	c := NewClient("123123")

	inURL, outURL := "/foo", defaultBaseURL+"/foo"
	inBody, outBody := &Account{Balance: 1.0}, `{"Balance":1}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	assert.Equal(t, req.URL.String(), outURL)

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	assert.Equal(t, string(body), outBody)

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
}

func testMethod(t *testing.T, r *http.Request, want string) {
	assert.Equal(t, r.Method, want)
}
