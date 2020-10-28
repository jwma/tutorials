package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type EchoParameters struct {
	Content string `json:"content"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method not allowed")
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error")
		return
	}

	if strings.ToLower(r.Header.Get("Content-Type")) != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "unsupported Content-Type, only support application/json")
		return
	}

	p := &EchoParameters{}
	_ = json.Unmarshal(body, p)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"echo": "%s"}`, p.Content)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"version": "%s"}`, os.Getenv("API_VERSION"))
}

func main() {
	http.HandleFunc("/", echoHandler)
	http.HandleFunc("/version", versionHandler)
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
