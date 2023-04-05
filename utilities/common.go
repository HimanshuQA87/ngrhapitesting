package utilities

import (
	"bufio"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"time"

	"github.com/HimanshuQA87/ngrhapitesting/struct_models"

	"github.com/go-resty/resty/v2"
	"github.com/goombaio/namegenerator"
	"github.com/spf13/viper"
)

var viperenv string = ViperEnvVariable("baseURL")

// Function to return value of a key mentioned in .env file
func ViperEnvVariable(key string) string {
	viper.SetConfigFile("../../config.env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

// Function to generate random name and email
func RandomNameandEmail() (string, string) {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	return name, name + "@example.com"
}

// Function to get common headers
func Get_common_headers() map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	return headers
}

// Function to send GET request
func Fire_get_request(headers map[string]string, qparams string, endpoint string) (response *resty.Response) {

	client := resty.New()
	resp, err := client.R().
		SetQueryString(qparams).
		SetHeaders(headers).
		Get(endpoint)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Get Request successful")
	}
	log.Println(resp)
	return resp
}

// Function to send POST request using payload as a string
func Fire_post_request(headers map[string]string, payload string, endpoint string) (response *resty.Response) {

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(payload).
		Post(endpoint)

	if err != nil {
		log.Println("error:", err)
	} else {
		log.Println("Post Request successful")
	}
	log.Println(resp)
	return resp
}

// Function to send POST request using payload as a struct
func Fire_post_request_bodystruct(headers map[string]string, payload interface{}, endpoint string) (response *resty.Response) {

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	resp, err := client.R().
		SetHeaders(headers).
		SetBody(payload).
		Post(endpoint)

	if err != nil {
		log.Println("error:", err)
	} else {
		log.Println("Post Request successful")
	}
	//log.Println(resp)
	return resp
}

// Function to send PUT request using payload as a struct
func Fire_put_request_bodystruct(headers map[string]string, payload interface{}, endpoint string) (response *resty.Response) {

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(payload).
		Put(endpoint)

	if err != nil {
		log.Println("error:", err)
	} else {
		log.Println("Put Request successful")
	}
	return resp
}

// Function to send PUT request using payload as a string
func Fire_put_request(headers map[string]string, payload string, endpoint string) (response *resty.Response) {

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(payload).
		Put(endpoint)

	if err != nil {
		log.Println("error:", err)
	} else {
		log.Println("Put Request successful")
	}
	//log.Println(resp)
	return resp
}

// Function to send DELETE request
func Fire_delete_request(headers map[string]string, endpoint string) (response *resty.Response) {

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		Delete(endpoint)

	if err != nil {
		log.Println("error:", err)
	} else {
		log.Println("Delete Request successful")
	}
	//log.Println(resp)
	return resp
}

// Function to send Patch request
func Fire_patch_request(headers map[string]string, payload interface{}, endpoint string) (response *resty.Response) {

	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(payload).
		Put(endpoint)

	if err != nil {
		log.Println("error:", err)
	} else {
		log.Println("Patch Request successful")
	}
	return resp
}

// Function to compare two json objects
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// Function to convert any file to encoded string
func ConvertFiletoBase64(filename string) string {

	f, _ := os.Open("../../data/files/" + filename)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}

// Function to convert any file to encoded string
func ConvertAllFilestoBase64(filename string) string {

	f, _ := os.Open("../../data/allformatfiles/" + filename)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}

func GetAuthorizationToken() string {

	cName, cEmail := RandomNameandEmail()
	endpoint := viperenv + "/api-clients"

	headers := Get_common_headers()
	body := struct_models.Register{ClientName: cName, ClientEmail: cEmail}
	resp := Fire_post_request_bodystruct(headers, body, endpoint)

	data, err := struct_models.UnmarshalRegister(resp.Body())

	if err != nil {
		fmt.Println(err)
	}
	return data.AccessToken
}

func GetCartId() string {
	endpoint := viperenv + "/carts"
	headers := Get_common_headers()
	resp := Fire_post_request(headers, "", endpoint)

	data, err := struct_models.UnmarshalCreateCart(resp.Body())
	if err != nil {
		fmt.Println(err)
	}
	return data.CartID
}

func ReadFile(partialfilepath string) (string, error) {
	data, err := os.ReadFile("../../data/" + partialfilepath)
	if err != nil {
		log.Println(" File not found")
		return "", err
	}
	temp := string(data)
	re := regexp.MustCompile(` +\r?\n +`)
	temp1 := re.ReplaceAllString(temp, "")
	//log.Println(temp1)
	return temp1, nil
}

func GetFileCount() int {
	files, _ := ioutil.ReadDir("../../data/files")
	//fmt.Println(len(files))
	return len(files)
}

func GetCurrentDateTime() string {
	currentTime := time.Now()
	datetime := currentTime.Format("January 2, 2006 3:04 PM")
	return datetime
}

func SendFileToBucket(baseEndPoint string, fileName string, deployment string) (*resty.Response, struct_models.CloudPrem, string, string, map[string]string) {
	filename, _ := RandomNameandEmail()
	endpoint := ViperEnvVariable(baseEndPoint)
	headers := Get_common_headers()
	encodedstring := ConvertFiletoBase64(fileName)
	body := struct_models.CloudPrem{FileName: filename + ".pdf", Body: encodedstring, Deploymentid: ViperEnvVariable(deployment)}

	resp := Fire_post_request_bodystruct(headers, body, endpoint)
	data, err := struct_models.UnmarshalCloudPrem(resp.Body())
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Response Deserialized successfully")
	}
	return resp, data, data.FileName, deployment, headers
}

func DownloadFileOnPrem(baseEndPoint string, fileName string, deployment string, headers map[string]string) (*resty.Response, struct_models.CloudPrem, string) {
	downloadendpoint := ViperEnvVariable(baseEndPoint) + "/download"
	downloadbody := struct_models.CloudPrem{FileName: fileName, Deploymentid: ViperEnvVariable(deployment)}

	downloadresp := Fire_post_request_bodystruct(headers, downloadbody, downloadendpoint)

	downloaddata, downloaderr := struct_models.UnmarshalCloudPrem(downloadresp.Body())

	if downloaderr != nil {
		log.Println(downloaderr)
	} else {
		log.Println("Response Deserialized successfully")
	}

	return downloadresp, downloaddata, fileName
}
