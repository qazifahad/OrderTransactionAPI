# OrderTransactionAPI
REST API to manage order transaction.

#Added this comment just for fun.

## Requirements
* Go 1.8.x+
* MongoDB 3.4.x+


### 3. Install Third Party Packages
Install packages by issuing in cmd:
```
go get gopkg.in/validator.v2 
go get github.com/julienschmidt/httprouter
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
```


## API Route
### Customer
1. GET /customer/:id - Find customer info by id
2. POST/customer - Create new customer
3. DELETE /customer/:id - Remove customer by id

### Product
1. GET /product/:id - Find product by id
2. POST /product - Create new product
3. POST /product/:id/reduce/:quantity - Reduce product quantity
4. DELETE /product/:id - Remove product by id

### Coupon
1. GET /coupon/:id - Find coupon by id 
2. GET /coupon/:id/valid - Check if coupon is valid by its starting date and expired date
3. POST /coupon - Create new coupon
4. POST /coupon/:id/reduce/:quantity - Reduce coupon quantity
5. DELETE /coupon/:id - Remove coupon by id

### Order
1. GET /order/:id - Find order by id
2. POST /order - Create new order
3. PUT /order/ - Adding orderLine which contains productId and quantity (when ordering)
4. DELETE /order/:id - Remove order by id

### Shipment
1. GET /shipment/:id - Find shipment status by id
2. POST /shipment - Create new shipment
3. DELETE /shipment/:id - Remove shipment by id



## License
MIT
