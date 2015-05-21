package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func getTestApi() (*Api, error) {
	log.SetLevel(log.ErrorLevel)

	cfg := &ApiConfig{
		ListenAddr: ":8080",
		HooksPath:  ".hooks",
	}

	return NewApi(cfg), nil
}

func TestApiGetIndex(t *testing.T) {
	api, err := getTestApi()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(api.index))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK, "expected response code 200")
}

func TestApiGetIP(t *testing.T) {
	api, err := getTestApi()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(api.getIP))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?name=foo")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK, "expected response code 200")
}

func TestApiGetState(t *testing.T) {
	api, err := getTestApi()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(api.getState))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?name=foo")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK, "expected response code 200")
}

func TestApiKill(t *testing.T) {
	api, err := getTestApi()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(api.kill))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?name=foo")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK, "expected response code 200")
}

func TestApiRemove(t *testing.T) {
	api, err := getTestApi()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(api.remove))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?name=foo")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK, "expected response code 200")
}

func TestApiRestart(t *testing.T) {
	api, err := getTestApi()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(api.restart))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?name=foo")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK, "expected response code 200")
}

func TestApiStart(t *testing.T) {
	api, err := getTestApi()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(api.start))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?name=foo")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK, "expected response code 200")
}

func TestApiStop(t *testing.T) {
	api, err := getTestApi()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(api.stop))
	defer ts.Close()

	res, err := http.Get(ts.URL + "?name=foo")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK, "expected response code 200")
}
