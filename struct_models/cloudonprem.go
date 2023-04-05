package struct_models

import (
	"encoding/json"
)

func UnmarshalCloudPrem(data []byte) (CloudPrem, error) {
	var r CloudPrem
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloudPrem) MarshalCloudPrem() ([]byte, error) {
	return json.Marshal(r)
}

type CloudPrem struct {
	Body string `json:"body"`
	//Bucket       string `json:"bucket"`
	FileName     string `json:"fileName"`
	Filelocation string `json:"filelocation"`
	Deploymentid string `json:"deploymentid"`
	Error        string `json:"error"`
}
