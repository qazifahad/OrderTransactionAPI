package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
)

func main() {  
    // Instantiate a new router
	r := httprouter.New()

	// Setup DB Session
	dbSession := GetSession()

	defer dbSession.Close()

	SetupRoutes(r,dbSession)

	// Fire up the server
    http.ListenAndServe("localhost:8080", r)
}

func GetSession() *mgo.Session {  
    // Connect to our local mongo
    s, err := mgo.Dial("localhost")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
	}

    return s
}