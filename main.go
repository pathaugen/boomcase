
package main

import (
	"fmt"
    "net/http"
)

// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========
// GAE config handlers for URIs
func init() {
	http.HandleFunc("/savecasedriver",	handlerSaveCaseDriver)
	http.HandleFunc("/saveimage",		handlerSaveImage)
	http.HandleFunc("/serve/",			handlerServe)
	http.HandleFunc("/login",			handlerLogin)
	
	// API Versions
	http.HandleFunc("/api/1.0/",		handlerAPI10) // API version 1.0
	
    http.HandleFunc("/",				handlerRoot)
}

// if true, running in production environment
// os.getenv('SERVER_SOFTWARE', '').startswith('Google App Engine/')
// ========== ========== ========== ========== ========== ========== ========== ========== ========== ==========

// main() used for executable when running after build, and also required for Travis CI
func main() {
	fmt.Println("BoomCase Google App Engine Application. Deploy code to Google App Engine (GAE) to utilize. No command line functions available.")
	
	// TODO: Creation of offline BoomCase customization web server application for POS kiosks utilizing same code
}
