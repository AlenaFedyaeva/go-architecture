package main

import (
	pb "client/api/proto"
	"context"
	"log"

	"google.golang.org/grpc"
)

func NewItemRepisitoryClient(addr string) (pb.itemRepositoryClient, error)  {
	//Заблокироваться до тех пор, пока подключение к серверу не будет установлено
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return pb.NewItemRepisitoryClient(conn), nil
}


func main() {
	itemRepository,err:= NewItemRepisitoryClient("localhost:9094")
	if err != nil{
		log.Println(err)
	}


	createItemReq := &pb.CreateItemsRequest{
		Name: "test_item1",
		Price: 1400,
	}

	item,err :=itemRepository.CreateItem(context.Background(),createItemReq)
	if err != nil{
		log.Println(err)
	}
	log.Println("item created", *item)
	
	listItemReq:=&pb.ListItemRequest{//}  ItemRequest{
		Limit: 1,
	}
	resp, err:= itemRepository.ListItems(context.Background(), listItemReq)
	if err!=nil {
		log.Println(err)
	}
	log.Println("item updated", *resp)
}
