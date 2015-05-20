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

	ts := httptest.NewServer(http.HandlerFunc(api.apiIndex))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, http.StatusOK, "expected response code 200")
}
