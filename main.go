package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
)

// https://github.com/tunapanda/h5p-standalone

func main() {
	StartHTTPServer(8080)
}

// 200331223000002
func index(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovering from panic", r)
		}
	}()

	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("parse files failed", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	t.Execute(w, nil)
	w.WriteHeader(http.StatusOK)
}

func StartHTTPServer(port int) error {
	httpMux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./"))
	httpMux.Handle("/", fs)

	server := &http.Server{
		Handler: httpMux,
	}

	// creating a listener for server
	httpLis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("listen port failed", err)
		return err
	}

	fmt.Println("ready to listen on port", port)

	err = server.Serve(httpLis)
	if err != nil {
		fmt.Println("serve failed", err)
		return err
	}

	return nil
}
