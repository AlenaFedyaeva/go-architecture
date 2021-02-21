# go-architecture

# HW2
postman:

1) create item
  http://localhost:8085/items  POST body type JSON 
{
     "name": "newItem333",
      "price": 10   
}


2) create order 
 http://localhost:8085/orders  POST body type JSON 
{
    "customer_name": "newItem333",
    "customer_phone": "+222", 
     "item_ids": [1,2]  
}
