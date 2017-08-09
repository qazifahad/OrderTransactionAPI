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

// CouponController structure
type (
	CouponController struct{
		dbCollection *mgo.Collection // Coupon collection
	}
)

// CouponController 'constructor'
func NewCouponController(dbCollection *mgo.Collection) *CouponController {
	 return &CouponController{dbCollection}
}

// CreateCoupon creates a new Coupon based on Coupon model
func (this CouponController) CreateCoupon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	coupon := models.Coupon{}  // Stub a coupon to be populated from the body
    json.NewDecoder(r.Body).Decode(&coupon)  // Populate the coupon data
	coupon.Id = bson.NewObjectId()  // Add a couponId

    this.dbCollection.Insert(coupon) // Write the new coupon to database specifically to Coupon collection
    
    couponJSON, _ := json.Marshal(coupon)  // Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", couponJSON)
}

// FindCoupon retrieves a coupon based on couponId
func (this CouponController) FindCoupon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Grab id

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    objectId := bson.ObjectIdHex(id) // Grab id

    coupon := models.Coupon{} // Stub coupon

    // Fetch coupon from coupon collection
    if err := this.dbCollection.FindId(objectId).One(&coupon); err != nil {
        w.WriteHeader(404)
        return
    }   

    couponJSON, _ := json.Marshal(coupon) // Marshal provided interface into JSON structure

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", couponJSON)
}

// DeleteCoupon deletes an existing coupon based on couponId
func (this CouponController) DeleteCoupon(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    objectId := bson.ObjectIdHex(id)

    // Remove coupon
    if err := this.dbCollection.RemoveId(objectId); err != nil {
        w.WriteHeader(404)
        return
    }

    // Write status
    w.WriteHeader(200)
}
