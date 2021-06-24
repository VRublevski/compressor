// Package web provides functions for configuring routers.
package web

import (
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vrubleuskii/image-compression/service/compression"
)

func Register(compressService *compression.Service, r *mux.Router) {
	r.HandleFunc("/compressed/{name}", compress(compressService)).Methods(http.MethodGet)
}

func compress(service *compression.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		vars := mux.Vars(r)
		name := vars["name"]

		quality, err := parseQuality(r.URL.Query().Get("quality"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		compressed, err := service.Compress(name, quality)
		if err == compression.ErrNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Contet-Type", "image/json")

		jpeg.Encode(w, compressed, nil)
	}
}

func parseQuality(param string) (int, error) {
	if param == "" {
		return 50, nil
	}

	return strconv.Atoi(param)
}
