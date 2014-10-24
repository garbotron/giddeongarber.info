package main

import (
	"net/http"
	"os"
)

func main() {
	http.Handle("/gg/", http.StripPrefix("/gg/", http.FileServer(http.Dir(os.ExpandEnv("$GOPATH/src/giddeongarber.info")))))
	panic(http.ListenAndServe(":80", nil))
}
