package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "server/proto"
	"sort"
	"sync"
	"time"

	"google.golang.org/grpc"
)

//Interface ItemReposiotory
type ItemRepositoryService struct {
	maxID       int32
	mu          *sync.RWMutex
	itemsTable  *itemsTable
	// ordersTable *ordersTable

	pb.UnimplementedItemRepositoryServer
}
type itemsTable struct {
	items map[int32]*Item
	maxID int32
}
type Item struct {
	ID        int32
	Name      string
	Price     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}


// type Order struct {
// 	ID      int32   `json:"id"`
// 	CustomerPhone   string  `json:"customer_phone"`
// 	CustomerName	string	`json:"customer_name"`
// 	ItemIDs []int32 `json:"item_ids"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// // }


// type ordersTable struct {
// 	orders map[int32]*models.Order
// 	maxID  int32
// }

func (s *ItemRepositoryService) BlockUser(ctx context.Context, req *pb.BlockUserRequest) (*pb.BlockUserResponse, error) {

	return &pb.BlockUserResponse{}, nil
}
func (s *ItemRepositoryService) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.Item, error) {
	s.itemsTable.maxID++

	timeNow := time.Now().UTC()

	newItem := &Item{
		ID:        s.maxID,
		Price:     req.Price,
		Name:      req.Name,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	s.mu.Lock()
	s.itemsTable.items[newItem.ID] = newItem
	s.mu.Unlock()

	return &pb.Item{
		ID:        newItem.ID,
		Name:      newItem.Name,
		Price:     newItem.Price,
	}, nil
}

func (s *ItemRepositoryService) ListItems(ctx context.Context, req *pb.ListItemRequest) (*pb.ListItemsResponse, error) {

	var res []*pb.ListItemsResponse
	fmt.Println(res)
	//sort items
	s.mu.RLock()
	itemSlice := make([]*pb.Item, 0, len(s.itemsTable.items))
	for _, item := range s.itemsTable.items {
		i:=&pb.Item{
			Id: item.ID,
			Name: item.Name,
			Price: item.Price,
		}
		itemSlice = append(itemSlice, i)
	}
	s.mu.RUnlock()

	sort.Slice(itemSlice, func(i, j int) bool {
		return itemSlice[i].Id < itemSlice[j].Id	})

	for _, item := range itemSlice {
		//надо еще обработать: передан ли 0 или ничего не передано
		if req.PriceLeft == 0 && req.PriceRight == 0 {
			res = itemSlice
			break
		}
		fmt.Println(*req.PriceLeft, item.Price, *req.PriceRight)
		if (*req.PriceLeft <= item.Price) && (item.Price <= *req.PriceRight) {

			res = append(res, item)
		}
	}
	//If filter is off
	if req.Limit == 0 && req.Offset == 0 && req.PriceLeft == nil && req.PriceRight == nil {
		return res, nil
	}

	resFiltered := make([]*pb.Item, 0, len(res))
	for idx, item := range res {
		if int32(len(resFiltered)) == req.Limit {
			break
		}
		if int32(idx) < req.Offset {
			continue
		}
		itemPb:=&pb.Item{
			Id: item.ID,
			Name: item.Name,
			Price: item.Price,
		}
		resFiltered = append(resFiltered, itemPb)
	}
	return &pb.ListItemsResponse{
		Items: resFiltered,
	}, nil

}

func NewItemRepositoryServerStart(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := NewMapDB()
	serv := grpc.NewServer()

	log.Println("starting grpc server at", addr)

	pb.RegisterItemRepositoryServer(serv, s)
	if err = serv.Serve(lis); err != nil {
		return err
	}
	return nil
}

func NewMapDB() pb.ItemRepositoryServer {
	return &pb.ItemRepositoryServer{
		mu: &sync.RWMutex{},
		itemsTable: &itemsTable{
			// items: map[int32]*models.Item{
				// 1: &models.Item{
				// 	ID:        1,
				// 	Name:      "item1",
				// 	Price:     50,
				// 	CreatedAt: time.Now(),
				// 	UpdatedAt: time.Now(),
				// },
				// 2: &models.Item{
				// 	ID:        2,
				// 	Name:      "item2",
				// 	Price:     60,
				// 	CreatedAt: time.Now(),
				// 	UpdatedAt: time.Now(),
				// },
				// 3: &models.Item{
				// 	ID:        3,
				// 	Name:      "item3",
				// 	Price:     70,
				// 	CreatedAt: time.Now(),
				// 	UpdatedAt: time.Now(),
				// },
				// 4: &models.Item{
				// 	ID:        4,
				// 	Name:      "item4",
				// 	Price:     80,
				// 	CreatedAt: time.Now(),
				// 	UpdatedAt: time.Now(),
				// },
			// },
			maxID: 4,
		},
		// ordersTable: &ordersTable{
		// 	orders: make(map[int32]*models.Order),
		// 	maxID:  0,
		// },
	}
}

func main() {

	


	NewItemRepositoryServerStart("localhost:9094")
}
