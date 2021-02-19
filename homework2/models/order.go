package models

import "time"

type Order struct {
	ID      int32   `json:"id"`
	Phone   string  `json:"phone"`
	Name	string	`json:"name"`
	ItemIDs []int32 `json:"item_ids"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// filtered:=make([]*Item,0,len(order.Items))
// for_,item:=range items{
	// item.Id=3232
// }