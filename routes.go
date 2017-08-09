package main

import (
	// Third party packages
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"

	// Internal packages
	CustomerController "./CustomerService/controllers"
	ProductController "./ProductService/controllers"
)

func SetupRoutes(r *httprouter.Router,s *mgo.Session) {
	db := s.DB("order_transaction") // Assign database

	// Assign collections
	customerCollection := db.C("customer") 
	productCollection := db.C("product")

	// Get controller instances
	customerController := CustomerController.NewCustomerController(customerCollection)
	productController := ProductController.NewProductController(productCollection)

	// Customer routes
	r.POST("/customer", customerController.CreateCustomer)
	r.GET("/customer/:id", customerController.FindCustomer)
	r.DELETE("/customer/:id", customerController.DeleteCustomer)

	// Product routes
	r.POST("/product", productController.CreateProduct)
	r.GET("/product/:id", productController.FindProduct)
	r.DELETE("/product/:id", productController.DeleteProduct)

	
}