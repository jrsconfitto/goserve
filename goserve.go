package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var port *int

func init() {
	port = flag.Int("port", 4000, "The port from which static files will be served.")
	flag.Parse()
}

func main() {
	currentDirectory, err := os.Getwd()
	if err == nil {
		thePort := strings.TrimSpace(strconv.Itoa(*port))
		fmt.Printf("Serving from http://localhost:%v\n", thePort)
		http.ListenAndServe(":"+thePort, http.FileServer(http.Dir(currentDirectory)))
	} else {
		fmt.Println(err.Error())
	}
}
