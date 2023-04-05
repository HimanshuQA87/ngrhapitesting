package components

import (
	"io/ioutil"
	"log"

	"github.com/HimanshuQA87/ngrhapitesting/struct_models"
	"github.com/HimanshuQA87/ngrhapitesting/utilities"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var body interface{}

//var downloadbody interface{}

var _ = Describe("Post Request - Cloud to On-Prem", Label("Smoke"), func() {

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
			log.Println(Fail.FailureNodeLocation)
			log.Println(Fail.Message)

		}

		log.Println("----Test Ended----")
		log.Println("")
		log.Println("ending")
	})

	It("Can send a file to S3 bucket using load balancer API", Label("Smoke"), func() {

		/* 		filename, _ := utilities.RandomNameandEmail()
		   		endpoint := viperenvcloud
		   		headers := utilities.Get_common_headers()
		   		encodedstring := utilities.ConvertFiletoBase64("Test.pdf")
		   		body := struct_models.CloudPrem{FileName: filename + ".pdf", Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

		   		resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)
		   		data, err := struct_models.UnmarshalCloudPrem(resp.Body())
		   		if err != nil {
		   			log.Println(err)
		   		} else {
		   			log.Println("Response Deserialized successfully")
		   		} */
		resp, data, filename, _, _ := utilities.SendFileToBucket("baseCloudOnPrem", "Test.pdf", "deploymentidsqe")
		Expect(resp.StatusCode()).To(Equal(200), "Status code does not match")
		Expect(resp.Time().Seconds()).To(BeNumerically("<", 5), "Response was not within 5 seconds")
		Expect(filename).To(Equal(data.FileName), "File name does not match")

	})

	It("Can send multiple pdf files to S3 bucket using load balancer API ", Label("Smoke"), func() {

		filename, _ := ioutil.ReadDir("../../data/files")

		for _, file := range filename {

			/* filename1, _ := utilities.RandomNameandEmail()
			endpoint := viperenvcloud
			headers := utilities.Get_common_headers()
			encodedstring := utilities.ConvertFiletoBase64(file.Name())
			body := struct_models.CloudPrem{FileName: filename1 + ".pdf", Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}

			resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

			data, err := struct_models.UnmarshalCloudPrem(resp.Body())
			if err != nil {
				log.Println(err)
			} */

			resp, data, filename1, _, _ := utilities.SendFileToBucket("baseCloudOnPrem", file.Name(), "deploymentidsqe")

			Expect(resp.StatusCode()).To(Equal(200), "Status code is not 200")
			Expect(filename1).To(Equal(data.FileName), "File name does not match")

		}

	})

	XIt("Can send files to different S3 buckets", func() {

		filename, _ := ioutil.ReadDir("../../data/files")

		for _, file := range filename {

			filename1, _ := utilities.RandomNameandEmail()
			endpoint := viperenvcloud
			headers := utilities.Get_common_headers()
			encodedstring := utilities.ConvertFiletoBase64(file.Name())
			if file.Name() == "Test2.pdf" || file.Name() == "Testin.pdf" {
				body = struct_models.CloudPrem{FileName: filename1 + ".pdf", Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe2")}

			} else {
				body = struct_models.CloudPrem{FileName: filename1 + ".pdf", Body: encodedstring, Deploymentid: utilities.ViperEnvVariable("deploymentidsqe")}
			}

			resp := utilities.Fire_post_request_bodystruct(headers, body, endpoint)

			data, err := struct_models.UnmarshalCloudPrem(resp.Body())
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Response Deserialized successfully")
			}
			Expect(resp.StatusCode()).To(Equal(200))
			Expect(filename1 + ".pdf").To(Equal(data.FileName))

		}

	})

})
