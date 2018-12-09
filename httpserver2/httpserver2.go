package httpserver2

import (
	"io"
	"net/http"
	"time"
)

// Hello : hi from the webserver at 127.0.0.1:1337
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

// HTTPExample : Run the example, import into a main using package name
func HTTPExample() {
	http.HandleFunc("/", hello)
	server := &http.Server{Addr: ":1337"}

	if err := server.ListenAndServe(); err != nil {
		// handle err
	}
	// handle err

	//srv := http.ListenAndServe(":1337", nil) // noob server port l0l
	time.Sleep(10 * time.Second)
	server.Shutdown(nil)
}
