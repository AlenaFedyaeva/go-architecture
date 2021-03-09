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

# hw7 

1) Install proto 

2) In my case I also had to add 
   sudo apt install protobuf-compiler
    sudo apt install golang-goprotobuf-dev

3) gen proto
protoc -I . fname_spec.proto --go_out=plugins=grpc:.