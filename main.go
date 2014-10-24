package giddeongarber

import (
	"net/http"
	"os"
)

func init() {
	http.Handle("/gg/", http.StripPrefix("/gg/", http.FileServer(http.Dir(os.ExpandEnv("$GOPATH/src/giddeongarber.info")))))
}
