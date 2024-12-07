package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"samb/endpoint"
)

func CORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/penerimaan-barang", endpoint.PenerimaanBarangEndpointWithoutId).Methods("POST", "GET", "OPTIONS")
	r.HandleFunc("/pengeluaran-barang", endpoint.PengeluaranBarangEndpointWithoutId).Methods("POST", "GET", "OPTIONS")

	r.Use(CORSHandler)

	return r
}
