package publicapi

import "net/http"

func (s *server) transferHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("transfer ok"))
}

func (s *server) balanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("balance ok"))
}