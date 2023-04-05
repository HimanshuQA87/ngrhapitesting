package components

import (
	"fmt"
	"log"

	"github.com/HimanshuQA87/ngrhapitesting/struct_models"
	"github.com/HimanshuQA87/ngrhapitesting/utilities" // Give modulename of go.mod file along with the file from where to call the function

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var viperenv string = utilities.ViperEnvVariable("baseURL")
var viperenvcloud string = utilities.ViperEnvVariable("baseCloudOnPrem")

//var viperenv1 string = utilities.ViperEnvVariable("baseURL1")

/* var _ = BeforeEach(func() {
	fmt.Println("Before Each1")

})
var _ = AfterEach(func() {
	fmt.Println("After Each1")

}) */

var _ = XDescribe("Get Request", Label("Regression"), func() {

	BeforeEach(func() {
		log.Println("----Test Started---- ")
		log.Printf("Running test : %s\n", CurrentSpecReport().FullText())
	})

	AfterEach(func() {
		Fail := CurrentSpecReport().Failure
		if Fail.Message == "" {
			log.Println("Test Passed")
		} else {
			log.Println("Test Failed")
			log.Println(Fail.Message)

		}
		log.Println("----Test Ended----")
		log.Println("")
	})

	XIt("Get Request - Status 200", func() {

		endpoint := viperenv + "/status"
		headers := utilities.Get_common_headers()
		//headers["Authorization"] = "123141"
		//queryString:= "page=2"

		resp := utilities.Fire_get_request(headers, "", endpoint)

		data, err := struct_models.UnmarshalStatus(resp.Body())
		if err != nil {
			fmt.Println(err)
		}
		Expect(data.Status).To(Equal("UP"), "Status not Equal to UP")
		Expect(resp.StatusCode()).To(Equal(200), "Status code not equal to 200")
	})

})

var _ = Describe("Post Request", func() {

	XIt("Post Request - Register a new client 201 ", func() {

		cName, cEmail := utilities.RandomNameandEmail()
		endpoint := viperenv + "/api-clients"

		headers := utilities.Get_common_headers()
		//headers["Authorization"] = "123141"
		body := struct_models.Register{ClientName: cName, ClientEmail: cEmail}
		resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

		data, err := struct_models.UnmarshalRegister(resp.Body())

		if err != nil {
			fmt.Println(err)
		}
		Expect(resp.StatusCode()).To(Equal(201))
		Expect(data.AccessToken).ToNot(Equal("Null"))

	})

})
