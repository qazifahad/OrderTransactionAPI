package routes

import (
	// Third party packages
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"

	// Internal packages
	CustomerController "../CustomerService/controllers"
	ProductController "../ProductService/controllers"
	CouponController "../CouponService/controllers"
	ShipmentController "../ShipmentService/controllers"
	OrderController "../OrderService/controllers"
)

func SetupRoutes(r *httprouter.Router,s *mgo.Session) {
	db := s.DB("order_transaction") // Assign database

	// Assign collections
	customerCollection := db.C("customer") 
	productCollection := db.C("product")
	couponCollection := db.C("coupon")
	shipmentCollection := db.C("shipment")
	orderCollection := db.C("order")

	// Get controller instances
	customerController := CustomerController.NewCustomerController(customerCollection)
	productController := ProductController.NewProductController(productCollection)
	couponController := CouponController.NewCouponController(couponCollection)
	shipmentController := ShipmentController.NewShipmentController(shipmentCollection)
	orderController := OrderController.NewOrderController(orderCollection)

	// Customer routes
	r.POST("/customer", customerController.CreateCustomer)
	r.GET("/customer/:id", customerController.FindCustomer)
	r.DELETE("/customer/:id", customerController.DeleteCustomer)

	// Product routes
	r.POST("/product", productController.CreateProduct)
	r.GET("/product/:id", productController.FindProduct)
	r.POST("/product/:id/reduce/:quantity",productController.ReduceProductQuantity)
	r.DELETE("/product/:id", productController.DeleteProduct)

	// Coupon routes
	r.POST("/coupon", couponController.CreateCoupon)
	r.GET("/coupon/:id", couponController.FindCoupon)
	r.GET("/coupon/:id/valid", couponController.CheckCouponValidity)
	r.POST("/coupon/:id/reduce/:quantity", couponController.ReduceCouponQuantity)
	r.DELETE("/coupon/:id", couponController.DeleteCoupon)

	// Shipment routes
	r.POST("/shipment", shipmentController.CreateShipment)
	r.GET("/shipment/:id", shipmentController.FindShipment)
	r.DELETE("/shipment/:id", shipmentController.DeleteShipment)

	// Order routes
	r.POST("/order", orderController.CreateOrder)
	r.GET("/order/:id", orderController.FindOrder)
	r.PUT("/order/",orderController.AddOrderLine)
	r.DELETE("/order/:id", orderController.DeleteOrder)
}