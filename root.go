package giddeongarber

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Init(r *mux.Router) error {
	fileServer := http.FileServer(http.Dir(os.ExpandEnv("$GOPATH/src/github.com/garbotron/giddeongarber.info")))
	r.Handle("/{path:.*}", fileServer).Host("{subdomain:(www\\.)?}giddeongarber.info")
	return nil
}
