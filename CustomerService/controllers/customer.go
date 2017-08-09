package controllers

import (
	// Standard library packages
	"fmt"
	"encoding/json"
	"net/http"

	// Third party packages
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"

	// Internal packages
	"../models"
)

type (
	CustomerController struct{
		dbCollection *mgo.Collection
	}
)

func NewCustomerController(dbCollection *mgo.Collection) *CustomerController {
	 return &CustomerController{dbCollection}
}

// CreateCustomer creates a new customer based on customer model
func (this CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	customer := models.Customer{}  // Stub an user to be populated from the body
    json.NewDecoder(r.Body).Decode(&customer)  // Populate the user data
	customer.Id = bson.NewObjectId()  // Add a customer Id

    this.dbCollection.Insert(customer)  // Write the new customer to database in dbCollection

    customerJSON, _ := json.Marshal(customer)  // Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", customerJSON)
}

// FindCustomer retrieves a customer based on customerId
func (this CustomerController) FindCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Grab id

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    objectId := bson.ObjectIdHex(id) // Grab id

    customer := models.Customer{} // Stub customer

    // Fetch user from collection
    if err := this.dbCollection.FindId(objectId).One(&customer); err != nil {
        w.WriteHeader(404)
        return
    }   

    customerJSON, _ := json.Marshal(customer) // Marshal provided interface into JSON structure

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", customerJSON)
}

// DeleteCustomer deletes an existing customer based on customerId
func (this CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    objectId := bson.ObjectIdHex(id)

    // Remove user
    if err := this.dbCollection.RemoveId(objectId); err != nil {
        w.WriteHeader(404)
        return
    }

    // Write status
    w.WriteHeader(200)
}
