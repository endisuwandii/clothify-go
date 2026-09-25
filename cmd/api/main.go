package main


import (

	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)



func main (){
	mux := newServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWritter, r *http.Request){
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.statusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status" : "healty",
			"service": "clothify-go-api"
		})


	})


}