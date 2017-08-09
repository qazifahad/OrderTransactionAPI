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

// ShipmentController structure
type (
	ShipmentController struct{
		dbCollection *mgo.Collection // Shipment collection
	}
)

// ShipmentController 'constructor'
func NewShipmentController(dbCollection *mgo.Collection) *ShipmentController {
	 return &ShipmentController{dbCollection}
}

// CreateShipment creates a new Shipment based on Shipment model
func (this ShipmentController) CreateShipment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	shipment := models.Shipment{}  // Stub a shipment to be populated from the body
    json.NewDecoder(r.Body).Decode(&shipment)  // Populate the shipment data
	shipment.Id = bson.NewObjectId()  // Add a shipment

    this.dbCollection.Insert(shipment) // Write the new shipment to database specifically to Shipment collection
    
    shipmentJSON, _ := json.Marshal(shipment)  // Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", shipmentJSON)
}

// FindShipment retrieves a shipment based on shipmentId
func (this ShipmentController) FindShipment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Grab id

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    objectId := bson.ObjectIdHex(id) // Grab id

    shipment := models.Shipment{} // Stub shipment

    // Fetch shipment from shipment collection
    if err := this.dbCollection.FindId(objectId).One(&shipment); err != nil {
        w.WriteHeader(404)
        return
    }   

    shipmentJSON, _ := json.Marshal(shipment) // Marshal provided interface into JSON structure

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", shipmentJSON)
}

// DeleteShipment deletes an existing shipment based on shipmentId
func (this ShipmentController) DeleteShipment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    objectId := bson.ObjectIdHex(id)

    // Remove shipment
    if err := this.dbCollection.RemoveId(objectId); err != nil {
        w.WriteHeader(404)
        return
    }

    // Write status
    w.WriteHeader(200)
}
