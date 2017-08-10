package controllers

import (
	"time"
	"strconv"
	// Standard library packages
	"fmt"
	"encoding/json"
    "net/http"

	// Third party packages
	"gopkg.in/mgo.v2/bson"
    "gopkg.in/mgo.v2"
    "gopkg.in/validator.v2"
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

    if err:=json.NewDecoder(r.Body).Decode(&coupon); err!=nil { // Unmarshalling JSON
        w.WriteHeader(400)
        return
    }  

    // Validate JSON Key-Value
    if err := validator.Validate(coupon); err != nil {
        w.Header().Set("Content-Type","application/json")
        w.WriteHeader(400)
        errJSON, _ := json.Marshal(err)
        fmt.Fprintf(w,"%s",errJSON)
        return
    }

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

    if !bson.IsObjectIdHex(id) { // Verify id is ObjectId, otherwise bail
        w.WriteHeader(404)
        return
    }

    objectId := bson.ObjectIdHex(id) // Grab id

    coupon := models.Coupon{} // Stub coupon

    if err := this.dbCollection.FindId(objectId).One(&coupon); err != nil { // Fetch coupon from coupon collection
        w.WriteHeader(404)
        return
    }   

    couponJSON, _:= json.Marshal(coupon) // Marshalling coupon to couponJSON

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", couponJSON)
}

func (this CouponController) CheckCouponValidity(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id") // Grab id

    if !bson.IsObjectIdHex(id) { // Verify id is ObjectId, otherwise bail
        w.WriteHeader(404)
        return
    }

    objectId := bson.ObjectIdHex(id) // Convert id to object ID
    
    coupon := models.Coupon{} // Stub coupon
    if err := this.dbCollection.FindId(objectId).One(&coupon); err != nil { // Fetch coupon from coupon collection
        w.WriteHeader(404)
        return
    }

    nowTime := time.Now()
    validity := nowTime.After(coupon.ValidStartDate) && nowTime.Before(coupon.ValidEndDate)

    response := (struct{
        coupon models.Coupon
        validity bool
    }{
        coupon,
        validity,
    })

    responseJSON, _:= json.Marshal(response) // Marshalling response to responseJSON
    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", responseJSON)
}

// FindCoupon retrieves a coupon based on couponId
func (this CouponController) ReduceCouponQuantity(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Grab id
    decQuantityStr := p.ByName("quantity") // Grab reducedQuantity in string
    decQuantity, err := strconv.Atoi(decQuantityStr); if err !=nil {
        w.WriteHeader(400)
        return
    }

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    objectId := bson.ObjectIdHex(id) // Grab id

    coupon := models.Coupon{} // Stub coupon

    // Find coupon with objectId and decrease stock in the amount of decQuantity if stock is sufficient
    change := mgo.Change{
        Update: bson.M{"$inc": bson.M{"quantity": -decQuantity}},
        ReturnNew: true,
    }
    if _, err := this.dbCollection.Find(bson.M{"_id": objectId, "quantity": bson.M{"$gte": decQuantity}}).Apply(change, &coupon); err != nil {
        w.WriteHeader(400)
        return
    }

    couponJSON, _:= json.Marshal(coupon) // Marshalling coupon to couponJSON

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
