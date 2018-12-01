package httpserver2

import (
	"io"
	"net/http"
)

// Hello : hi from the webserver at 127.0.0.1:1337
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

// HTTPExample : Run the example, import into a main using package name
func HTTPExample() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":1337", nil) // noob server port l0l
}
