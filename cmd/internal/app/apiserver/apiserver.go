package apiserver

import (
	"e_katalog_killer/cmd/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(cnf *Config) *ApiServer {
	return &ApiServer{
		config: cnf,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *ApiServer) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}

	a.configureRouter()

	if err := a.configureStore(); err != nil {
		return err
	}

	a.logger.Info("starting api server")

	return http.ListenAndServe(a.config.BindAddr, a.router)
}

func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	s.configureRouter()

	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())

}

func (s *ApiServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *ApiServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
