package main

import (
	"fmt"
	"net/http"
	"os"
)

func userGet(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Path
	fmt.Println(string(req))
	w.Write([]byte("[/] : get all user"))
}
func todoGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("[/todo] : get all user's todo"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/user", userGet)
	mux.HandleFunc("/api/todo", todoGet)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Error in starting http server: %s\n", err)
		os.Exit(1)
	}
}
