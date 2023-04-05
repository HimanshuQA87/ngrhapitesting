package struct_models

import (
	"encoding/json"
)

func UnmarshalUpdateItemsToCart(data []byte) (UpdateItemsToCart, error) {
	var r UpdateItemsToCart
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AddItemToCart) MarshalUpdateItemsToCart() ([]byte, error) {
	return json.Marshal(r)
}

type UpdateItemsToCart struct {
	CartID   string `json:"cartId"`
	Quantity int    `json:"quantity"`
	ItemId   int    `json:"itemId"`
}
