package struct_models

import (
	"encoding/json"
)

func UnmarshalCreateCart(data []byte) (CreateCart, error) {
	var r CreateCart
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateCart) MarshalCreateCart() ([]byte, error) {
	return json.Marshal(r)
}

type CreateCart struct {
	Created bool   `json:"created"`
	CartID  string `json:"cartId"`
}
