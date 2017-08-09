package main

import (
	// Third party packages
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"

	// Internal packages
	CustomerController "./CustomerService/controllers"
	ProductController "./ProductService/controllers"
	CouponController "./CouponService/controllers"
)

func SetupRoutes(r *httprouter.Router,s *mgo.Session) {
	db := s.DB("order_transaction") // Assign database

	// Assign collections
	customerCollection := db.C("customer") 
	productCollection := db.C("product")
	couponCollection := db.C("coupon")

	// Get controller instances
	customerController := CustomerController.NewCustomerController(customerCollection)
	productController := ProductController.NewProductController(productCollection)
	couponController := CouponController.NewCouponController(couponCollection)

	// Customer routes
	r.POST("/customer", customerController.CreateCustomer)
	r.GET("/customer/:id", customerController.FindCustomer)
	r.DELETE("/customer/:id", customerController.DeleteCustomer)

	// Product routes
	r.POST("/product", productController.CreateProduct)
	r.GET("/product/:id", productController.FindProduct)
	r.DELETE("/product/:id", productController.DeleteProduct)

	// Coupon routes
	r.POST("/coupon", couponController.CreateCoupon)
	r.GET("/coupon/:id", couponController.FindCoupon)
	r.DELETE("/coupon/:id", couponController.DeleteCoupon)
}