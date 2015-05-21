package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/ehazlett/rivet/version"
	"github.com/gorilla/mux"
)

type Api struct {
	config *ApiConfig
}

func apiResponse(status int, response []byte, w http.ResponseWriter) error {
	w.WriteHeader(status)

	b := bytes.NewBuffer(response)

	resp := &ApiResponse{
		StatusCode: status,
		Response:   b.String(),
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return err
	}

	return nil
}

func NewApi(cfg *ApiConfig) *Api {
	return &Api{
		config: cfg,
	}
}

func (a *Api) doPluginHook(action string, args []string) ([]byte, error) {
	cmdPath, err := exec.LookPath("pluginhook")
	if err != nil {
		return nil, err
	}

	cmdArgs := []string{
		action,
	}

	cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command(cmdPath, cmdArgs...)
	// set PLUGIN_PATH env var
	cmd.Env = append(cmd.Env, fmt.Sprintf("PLUGIN_PATH=%s", a.config.HooksPath))

	return cmd.CombinedOutput()
}

func (a *Api) Run() error {
	globalMux := http.NewServeMux()

	router := mux.NewRouter()
	router.HandleFunc("/", a.index).Methods("GET")
	router.HandleFunc("/create", a.create).Methods("POST")
	router.HandleFunc("/ip", a.getIP).Methods("GET")
	router.HandleFunc("/state", a.getState).Methods("GET")
	router.HandleFunc("/start", a.start).Methods("GET")
	router.HandleFunc("/kill", a.kill).Methods("GET")
	router.HandleFunc("/remove", a.remove).Methods("GET")
	router.HandleFunc("/restart", a.restart).Methods("GET")
	router.HandleFunc("/stop", a.stop).Methods("GET")
	globalMux.Handle("/", router)

	log.Infof("listening: addr=%s", a.config.ListenAddr)
	return http.ListenAndServe(a.config.ListenAddr, globalMux)
}

func (a *Api) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("rivet %s\n", version.FULL_VERSION)))
}

func (a *Api) create(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	cpu := r.URL.Query().Get("cpu")
	memory := r.URL.Query().Get("memory")
	storage := r.URL.Query().Get("storage")

	if name == "" || cpu == "" || memory == "" || storage == "" {
		http.Error(w, "you must specify name, key, cpu, memory and storage params", http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key := bytes.NewBuffer(data)

	log.Infof("create: name=%s cpu=%s memory=%s storage=%s",
		name, cpu, memory, storage)
	log.Debugf("ssh key: %s", key.String())

	args := []string{
		name,
		key.String(),
		cpu,
		memory,
		storage,
	}

	res, err := a.doPluginHook("create", args)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) getIP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "you must specify a machine name", http.StatusBadRequest)
		return
	}

	args := []string{
		name,
	}

	res, err := a.doPluginHook("get_ip", args)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) getState(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "you must specify a machine name", http.StatusBadRequest)
		return
	}

	args := []string{
		name,
	}

	res, err := a.doPluginHook("get_state", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) start(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "you must specify a machine name", http.StatusBadRequest)
		return
	}

	args := []string{
		name,
	}

	res, err := a.doPluginHook("start", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) kill(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "you must specify a machine name", http.StatusBadRequest)
		return
	}

	args := []string{
		name,
	}

	res, err := a.doPluginHook("kill", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) remove(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "you must specify a machine name", http.StatusBadRequest)
		return
	}

	args := []string{
		name,
	}

	res, err := a.doPluginHook("remove", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) restart(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "you must specify a machine name", http.StatusBadRequest)
		return
	}

	args := []string{
		name,
	}

	res, err := a.doPluginHook("restart", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) stop(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "you must specify a machine name", http.StatusBadRequest)
		return
	}

	args := []string{
		name,
	}

	res, err := a.doPluginHook("stop", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiResponse(http.StatusOK, res, w)
}
