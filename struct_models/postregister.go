package struct_models

import "encoding/json"

func UnmarshalRegister(data []byte) (Register, error) {
	var r Register
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Register) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Register struct {
	ClientName  string `json:"clientName"`
	ClientEmail string `json:"clientEmail"`
	AccessToken string `json:"accessToken"`
}
