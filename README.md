# go-architecture

# hw 1
 
create request with Postman

1) create item 
In postman -  body (json)
{
     "name": "newItem",
      "price": 10   
}
send POST request http://localhost:8085/item

2) list items

send GET (example with offset=1) request http://localhost:8085/item?offset=1 