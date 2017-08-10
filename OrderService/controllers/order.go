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

// OrderController structure
type (
	OrderController struct{
		dbCollection *mgo.Collection // Order collection
	}
)

// OrderController 'constructor'
func NewOrderController(dbCollection *mgo.Collection) *OrderController {
	 return &OrderController{dbCollection}
}

// CreateOrder creates a new Order based on Order model
func (this OrderController) CreateOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	order := models.Order{}  // Stub a order to be populated from the body
    json.NewDecoder(r.Body).Decode(&order)  // Populate the order data
	order.Id = bson.NewObjectId()  // Add a orderId

    this.dbCollection.Insert(order) // Write the new order to database specifically to order collection
    
    orderJSON, _ := json.Marshal(order)  // Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", orderJSON)
}

// FindOrder retrieves a order based on orderId
func (this OrderController) FindOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Grab id

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    objectId := bson.ObjectIdHex(id) // Grab id
    order := models.Order{} // Stub order

    // Fetch order from order collection
    if err := this.dbCollection.FindId(objectId).One(&order); err!=nil {
        w.WriteHeader(400)
        return
    } 

    orderJSON, _ := json.Marshal(order) // Marshal provided interface into JSON structure

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", orderJSON)
}

// DeleteOrder deletes an existing order based on orderId
func (this OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Convert id to objectId
    objectId := bson.ObjectIdHex(id)

    // Remove order
    if err := this.dbCollection.RemoveId(objectId); err != nil {
        w.WriteHeader(404)
        return
    }

    // Write status
    w.WriteHeader(200)
}

func (this OrderController) AddOrderLine(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // Stub data to be populated from the body
    data := (struct{
        orderId bson.ObjectId
        orderLine models.OrderLine 
    }{})
    json.NewDecoder(r.Body).Decode(&data)  // Populate the order data

    orderLine := models.OrderLine{
        ProductId: data.orderLine.ProductId,
        Quantity: data.orderLine.Quantity,
    }


    order := models.Order{} // Stub coupon

    // Find product with objectId and decrease stock in the amount of decQuantity if stock is sufficient
    change := mgo.Change{
        Update: bson.M{"$push": bson.M{"orderLines":orderLine}},
        ReturnNew: true,
    }
    if _, err := this.dbCollection.FindId(data.orderId).Apply(change, &order); err != nil {
        w.WriteHeader(400)
        return
    }
}