package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	serverList := []*http.Server{
		shortTimeout(),
		withoutTimeout(),
		longTimeout(),
	}

	for _, v := range serverList {
		go func(server *http.Server) {
			server.Handler = getHandler(cancel)
			log.Fatal(server.ListenAndServe())
		}(v)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Finish by context done")
	}
}
//withoutTimeout func show behavior what if server without any timeouts handle request
//Return server with 0 sec(infinity) timeout settings
func withoutTimeout() *http.Server  {
	server := &http.Server{
		Addr:              "localhost:9871",
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
	}

	return server
}
//shortTimeout func show behavior what if server with short timeouts handle request
//Return server with 5 sec timeout settings
func shortTimeout() *http.Server  {
	server := &http.Server{
		Addr:              "localhost:9872",
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	return server
}

//longTimeout func show behavior what if server with long timeouts handle request
//Return server with 100 sec timeout settings
func longTimeout() *http.Server {
	server := &http.Server{
		Addr:              "localhost:9873",
		ReadTimeout:       100 * time.Second,
		ReadHeaderTimeout: 100 * time.Second,
		WriteTimeout:      100 * time.Second,
		IdleTimeout:       100 * time.Second,
	}

	return server
}

// getHandler get handlers for all servers
func getHandler(cancel context.CancelFunc) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)

		w.Write([]byte("short"))
	})
	r.HandleFunc("/long", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)

		w.Write([]byte("long"))
	})

	r.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		cancel()
	})

	r.HandleFunc("/longest", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1000 * time.Second)

		w.Write([]byte("longest"))
	})

	return r
}

