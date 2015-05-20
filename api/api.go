package api

import (
	"fmt"
	"net/http"

	"github.com/ehazlett/rivet/version"
	"github.com/gorilla/mux"
)

type Api struct {
	config *ApiConfig
}

func NewApi(cfg *ApiConfig) *Api {
	return &Api{
		config: cfg,
	}
}

func (a *Api) Run() error {
	globalMux := http.NewServeMux()

	router := mux.NewRouter()
	router.HandleFunc("/", a.apiIndex).Methods("GET")
	globalMux.Handle("/", router)

	return http.ListenAndServe(a.config.ListenAddr, globalMux)
}

func (a *Api) apiIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("rivet %s\n", version.FULL_VERSION)))
}
