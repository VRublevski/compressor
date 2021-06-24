package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vrubleuskii/image-compression/cache"
	"github.com/vrubleuskii/image-compression/service/compression"
	"github.com/vrubleuskii/image-compression/web"
)

func main() {
	const cacheSize = 64
	ch := cache.New(cacheSize)
	compressService := compression.NewService(ch)
	r := mux.NewRouter()
	web.Register(compressService, r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
