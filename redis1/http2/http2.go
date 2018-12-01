package http2

import (
	"io"
	"net/http"
)

// Hello : hi!
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

// HTTPExample : yep.
func HTTPExample() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":1337", nil) // noob server port l0l
}
