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
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status" : "healty",
			"service": "clothify-go-api",
		})
})


server := &http.Server{
	Addr: ":8080",
	Handler: mux,
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
	IdleTimeout:  60 * time.Second,	
}


shutDown := make(chan os.Signal, 1)
signal.Notify(shutDown, os.Interrupt, syscall.SIGTERM)

go func(){
	log.Printf("Server Starting at http://localhost%s", server.Addr)
	if err := server.ListenAndServe();err != nil && !errors.Is(err, http.ErrServerClosed){
		log.Fatalf("Erorr starting server %v", err)
	}
}()

<- shutDown
log.Printf("Shutting down server ")


ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
defer cancel()

if err := server.Shutdown(ctx); err != nil{
	log.Fatalf("Error shutting down server: %v", err)
}



}