package main

import (
	"context"
	"hello-world/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "hello-world", log.LstdFlags|log.Lshortfile)
	helloHandler := handlers.NewHelloHandler(logger)

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", helloHandler.GetRoot)
	serverMux.HandleFunc("/hello", helloHandler.GetHello)
	serverMux.HandleFunc("/echo", helloHandler.ServeRequestAsResponse)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	timeOutContext, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancelFunc()

	server.Shutdown(timeOutContext)
}
