package main

import "net/http"

/*
In order to set up a web application in go, Its pretty simple
What we need to do is
1. Register a function that going to respond to our web request.
2. Then start up the server.
*/
func main() {
	// Handle all request and hence forward slash "/" - request handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	// nil - defualtServeMux
	http.ListenAndServe(":8000", nil)

}
