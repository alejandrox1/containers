package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func initLog() error {
	f, err := os.Create("file.log")
	if err != nil {
		return err
	}
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	initLog()

	http.HandleFunc("/", home)

	log.Println("Hi Sid :-), im running your server")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
