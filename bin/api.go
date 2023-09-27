// This file is where we are going to create the API server or handlers
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct{
	error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		if err = f(r, w); err !=nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}


type APIserver struct {
	listenAddr string
}

func newAPIserver(listenAddr string) "APIserver"{
	return &APIserver{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run(){
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc( s.handleAccount))

	log.Println("Json API server running on port: ", s.listenAddr)

	http.listenAndServer(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter,r *http.Request) error {
	if r.Method == "GET" {
		s.handleGetAccount(w,r)
	}
	if r.Method == "POST" {
		s.handleCreateAccount(w,r)
	}
	if r.Method == "DELETE" {
		s.handleDeleteAccount(w,r)
	}
	
	return fmt.error("Method not allowed %s", r.Method)

}

func (s *APIServer) handleGetAccount(w http.ResponseWriter,r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter,r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter,r *http.Request) error {
	return nil
}