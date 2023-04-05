package struct_models

import (
	"encoding/json"
)

func UnmarshalAddItemToCart(data []byte) (AddItemToCart, error) {
	var r AddItemToCart
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AddItemToCart) MarshalAddItemToCart() ([]byte, error) {
	return json.Marshal(r)
}

type AddItemToCart struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
	Created   bool  `json:"created"`
	ItemId    int64 `json:"itemId"`
}
