package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/smartystreets/cproxy"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	filter := &LoggingFilter{}
	handler := cproxy.Configure(cproxy.WithFilter(filter))

	log.Println("Listening on:", addr)
	http.ListenAndServe(addr, handler)
}

type LoggingFilter struct{}

func (f *LoggingFilter) IsAuthorized(req *http.Request) bool {
	fmt.Println("PROXY REQUEST: ", req.Method)
	return true
}
