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

// ProductController structure
type (
	ProductController struct{
		dbCollection *mgo.Collection // Product collection
	}
)

// ProductController 'constructor'
func NewProductController(dbCollection *mgo.Collection) *ProductController {
	 return &ProductController{dbCollection}
}

// CreateProduct creates a new Product based on Product model
func (this ProductController) CreateProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	product := models.Product{}  // Stub a product to be populated from the body
    json.NewDecoder(r.Body).Decode(&product)  // Populate the product data
	product.Id = bson.NewObjectId()  // Add a productId

    this.dbCollection.Insert(product) // Write the new product to database specifically to Product collection
    
    productJSON, _ := json.Marshal(product)  // Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", productJSON)
}

// FindProduct retrieves a product based on productId
func (this ProductController) FindProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Grab id

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    objectId := bson.ObjectIdHex(id) // Grab id

    product := models.Product{} // Stub product

    // Fetch product from product collection
    if err := this.dbCollection.FindId(objectId).One(&product); err != nil {
        w.WriteHeader(404)
        return
    }   

    productJSON, _ := json.Marshal(product) // Marshal provided interface into JSON structure

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", productJSON)
}

// DeleteProduct deletes an existing product based on productId
func (this ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    objectId := bson.ObjectIdHex(id)

    // Remove product
    if err := this.dbCollection.RemoveId(objectId); err != nil {
        w.WriteHeader(404)
        return
    }

    // Write status
    w.WriteHeader(200)
}
