// This file is where we are going to create the API server or handlers
package main

type APIserver struct {
	listenAddr string
}

func newAPIserver(listenAddr string) "APIserver"{
	return &APIserver{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run(){

}

func (s *APIServer) handleAccount(w http.ResponseWriter,r *http.Request) error {
	return nil
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