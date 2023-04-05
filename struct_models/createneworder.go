package struct_models

import (
	"encoding/json"
)

func UnmarshalCreateNewOrder(data []byte) (CreateNewOrder, error) {
	var r CreateNewOrder
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateUser) MarshalCreateNewOrder() ([]byte, error) {
	return json.Marshal(r)
}

type CreateNewOrder struct {
	CartID       string `json:"cartId"`
	CustomerName string `json:"customerName"`
	OrderID      string `json:"orderId"`
	Created      bool   `json:"created"`
	Error        string `json:"error"`
}
