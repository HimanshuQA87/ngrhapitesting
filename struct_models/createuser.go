package struct_models

import (
	"encoding/json"
	"time"
)

func UnmarshalCreateUser(data []byte) (CreateUser, error) {
	var r CreateUser
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreateUser) MarshalCreateUser() ([]byte, error) {
	return json.Marshal(r)
}

type CreateUser struct {
	Name      string    `json:"name"`
	Job       string    `json:"job"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
