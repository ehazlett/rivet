package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/ehazlett/rivet/auth"
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

func (a *Api) doPluginHook(action string, args []string, env []string) ([]byte, error) {
	log.Debugf("running hook: name=%s args=%v", action, args)

	cmdPath, err := exec.LookPath("pluginhook")
	if err != nil {
		return nil, err
	}

	cmdArgs := []string{
		action,
	}

	cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command(cmdPath, cmdArgs...)
	// load current environment to get custom settings for plugins
	cmd.Env = os.Environ()
	// set PLUGIN_PATH env var
	cmd.Env = append(cmd.Env, fmt.Sprintf("PLUGIN_PATH=%s", a.config.HooksPath))
	cmd.Env = append(cmd.Env, env...)

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
	// enable auth if token is present
	if a.config.AuthToken != "" {
		am := auth.NewAuthMiddleware(a.config.AuthToken)
		globalMux.Handle("/", negroni.New(
			negroni.HandlerFunc(am.Handler),
			negroni.Wrap(http.Handler(router)),
		))
	} else {

		globalMux.Handle("/", router)
	}

	log.Infof("rivet version %s", version.FULL_VERSION)
	log.Infof("listening: addr=%s", a.config.ListenAddr)
	n := negroni.New()
	n.UseHandler(globalMux)
	return http.ListenAndServe(a.config.ListenAddr, n)
}

func (a *Api) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("rivet %s\n", version.FULL_VERSION)))
}

func (a *Api) create(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	cpu := r.URL.Query().Get("cpu")
	memory := r.URL.Query().Get("memory")
	storage := r.URL.Query().Get("storage")
	image := r.URL.Query().Get("image")
	env := r.URL.Query()["env"]

	if name == "" || cpu == "" || memory == "" || storage == "" {
		apiResponse(http.StatusBadRequest, []byte("you must specify name, key, cpu, memory and storage params"), w)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiResponse(http.StatusInternalServerError, []byte(err.Error()), w)
		return
	}

	key := bytes.NewBuffer(data)

	log.Infof("create: name=%s cpu=%s memory=%s storage=%s env=%v",
		name, cpu, memory, storage, env)
	log.Debugf("ssh key: %s", key.String())

	args := []string{
		name,
		key.String(),
		cpu,
		memory,
		storage,
		image,
	}

	res, err := a.doPluginHook("create", args, env)
	if err != nil {
		log.Error(err)
		apiResponse(http.StatusInternalServerError, []byte(err.Error()), w)
		return
	}

	// post-create; append the response
	postRes, err := a.doPluginHook("post_create", args, env)
	if err != nil {
		log.Errorf("post-create error: %s", err)
	}

	if postRes != nil {
		res = append([]byte("\n"))
		res = append(postRes)
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) getIP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		apiResponse(http.StatusBadRequest, []byte("you must specify a machine name"), w)
		return
	}

	args := []string{
		name,
	}

	env := []string{}

	res, err := a.doPluginHook("get_ip", args, env)
	if err != nil {
		log.Error(err)
		apiResponse(http.StatusInternalServerError, []byte(err.Error()), w)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) getState(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		apiResponse(http.StatusBadRequest, []byte("you must specify a machine name"), w)
		return
	}

	args := []string{
		name,
	}

	env := []string{}

	res, err := a.doPluginHook("get_state", args, env)
	if err != nil {
		apiResponse(http.StatusInternalServerError, []byte(err.Error()), w)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) start(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		apiResponse(http.StatusBadRequest, []byte("you must specify a machine name"), w)
		return
	}

	args := []string{
		name,
	}

	env := []string{}

	res, err := a.doPluginHook("start", args, env)
	if err != nil {
		apiResponse(http.StatusInternalServerError, []byte(err.Error()), w)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) kill(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		apiResponse(http.StatusBadRequest, []byte("you must specify a machine name"), w)
		return
	}

	args := []string{
		name,
	}

	env := []string{}

	res, err := a.doPluginHook("kill", args, env)
	if err != nil {
		apiResponse(http.StatusInternalServerError, []byte(err.Error()), w)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) remove(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		apiResponse(http.StatusBadRequest, []byte("you must specify a machine name"), w)
		return
	}

	args := []string{
		name,
	}

	env := []string{}

	res, err := a.doPluginHook("remove", args, env)
	if err != nil {
		apiResponse(http.StatusInternalServerError, []byte(err.Error()), w)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) restart(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		apiResponse(http.StatusBadRequest, []byte("you must specify a machine name"), w)
		return
	}

	args := []string{
		name,
	}

	env := []string{}

	res, err := a.doPluginHook("restart", args, env)
	if err != nil {
		apiResponse(http.StatusInternalServerError, []byte(err.Error()), w)
		return
	}

	apiResponse(http.StatusOK, res, w)
}

func (a *Api) stop(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		apiResponse(http.StatusBadRequest, []byte("you must specify a machine name"), w)
		return
	}

	args := []string{
		name,
	}

	env := []string{}

	res, err := a.doPluginHook("stop", args, env)
	if err != nil {
		apiResponse(http.StatusInternalServerError, []byte(err.Error()), w)
		return
	}

	apiResponse(http.StatusOK, res, w)
}
