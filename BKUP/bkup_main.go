// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// 	"github.com/gorilla/mux"
// )

// func main() {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/", handler)

// 	srv := &http.Server{
// 		Handler:      router,
// 		Addr:         ":8080",
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 	}

// 	go func() {
// 		log.Println("Starting Server now...")
// 		if err := srv.ListenAndServe(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	waitForShutdown(srv)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Received request")
// 	w.Write([]byte(greet()))
// }

// func greet() string {
// 	return "Hello from Form-App"
// }

// func waitForShutdown(srv *http.Server) {
// 	interruptChan := make(chan os.Signal, 1)
// 	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

// 	// Block until we receive our signal.
// 	<-interruptChan

// 	// Create a deadline to wait for.
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()
// 	srv.Shutdown(ctx)

// 	log.Println("Shutting down")
// 	os.Exit(0)
// }
