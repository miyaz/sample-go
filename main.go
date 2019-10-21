package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	revision string
	buildAt  string
)

func main() {
	log.Printf("revision: %s, buildAt: %s", revision, buildAt)
	http.HandleFunc("/", handler)
	srv := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	queryKeys := r.URL.Query()
	if queryKeys == nil {
		return
	}
	if urls, ok := queryKeys["url"]; ok {
		url := urls[0]
		err := validation.Validate(url,
			validation.Required,
			validation.Length(5, 256),
			is.URL,
		)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		resp, err := proxy(url)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		fmt.Fprintf(w, "%s", string(resp))
		return
	}
	for key, vals := range queryKeys {
		for _, val := range vals {
			fmt.Fprintf(w, "%s=%v\n", key, val)
		}
	}
}

func proxy(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}
