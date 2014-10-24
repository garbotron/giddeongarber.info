package main

import "net/http"

func main() {
	http.Handle("/gg/", http.StripPrefix("/gg/", http.FileServer(http.Dir("/www/giddeongarber.info"))))
	panic(http.ListenAndServe(":80", nil))
}
