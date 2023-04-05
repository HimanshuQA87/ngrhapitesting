package struct_models

import (
	"encoding/json"
)

func UnmarshalGetCartItems(data []byte) (GetCartItems, error) {
	var r GetCartItems
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GetCartItems) MarshalGetCartItems() ([]byte, error) {
	return json.Marshal(r)
}

type GetCartItems []struct {
	Id        int64 `json:"id"`
	ProductId int   `json:"productId"`
	Quantity  int   `json:"quantity"`
}
