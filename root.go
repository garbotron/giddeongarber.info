package giddeongarber

import (
	"net/http"
	"os"
)

func Init() error {
	http.Handle("/gg/", http.StripPrefix("/gg/", http.FileServer(http.Dir(os.ExpandEnv("$GOPATH/src/giddeongarber.info")))))
	return nil
}
