package main

import (
	"context"
	"log"
	"net"
	pb "server/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Item struct {
	ID    int32
	Name  string
	Price int64
}
//Interface ItemReposiotory
type ItemRepositoryService struct {
	maxID int32
	db    map[int32]*Item
	pb.UnimplementedItemRepositoryServer
}

func (s *ItemRepositoryService) BlockUser(ctx context.Context, req *pb.BlockUserRequest) (*pb.BlockUserResponse, error) {

	return &pb.BlockUserResponse{}, nil
}
func (s *ItemRepositoryService) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.Item, error) {
	
	
	
	//****************
	newItem := &Item{
		Name:  req.Name,
		Price: req.Price,
	}
	s.maxID++
	newItem.ID = s.maxID

	s.db[s.maxID] = newItem

	return &pb.Item{
		Id:    s.maxID,
		Name:  newItem.Name,
		Price: newItem.Price,
	}, nil
}

func (s *ItemRepositoryService) ListItems(ctx context.Context, req *pb.ListItemRequest) (*pb.ListItemsResponse, error) {
	item, ok := s.db[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}

	item.Name = req.Name
	item.Price = req.Price

	return &pb.ListItemsResponse{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}, nil
}

func NewItemRepositoryServerStart(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := &ItemRepositoryService{
		db: make(map[int32]*Item),
	}
	serv := grpc.NewServer()

	log.Println("starting grpc server at", addr)

	pb.RegisterItemRepositoryServer(serv, s)
	if err = serv.Serve(lis); err != nil {
		return err
	}
	return nil
}

func main() {
	NewItemRepositoryServerStart("localhost:9094")
}
