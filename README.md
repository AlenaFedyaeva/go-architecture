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

3) GOOGLE. To send mail need to !Less secure app access! (settings in account)

# Links 
 
1) Сложность алгоритмов https://habr.com/ru/post/263765/
2) Бинарный поиск https://ru.wikipedia.org/wiki/Двоичный_поиск
3) Шпаргалка структуры данных https://habr.com/ru/post/188010/


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

# hw7 

1) Install proto 

2) In my case I also had to add 
   sudo apt install protobuf-compiler
    sudo apt install golang-goprotobuf-dev

3) gen proto
protoc -I . fname_spec.proto --go_out=plugins=grpc:.

